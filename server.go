// Package bgpmon provides the core interfaces for developing bgpmon client programs
// and modules.
package bgpmon

import (
	"fmt"
	"os"
	"sync"

	"github.com/CSUNetSec/bgpmon/config"
	"github.com/CSUNetSec/bgpmon/db"
	"github.com/CSUNetSec/bgpmon/util"

	pb "github.com/CSUNetSec/netsec-protobufs/bgpmon/v2"
)

var (
	coreLogger = util.NewLogger("system", "core")
)

// BgpmondServer is the interface for interacting with a server instance.
// It provides function to open and close sessions and modules, and get
// data on the state on the server.
type BgpmondServer interface {
	// Opens a session of type sType, which must come from the config file.
	// The ID of this session is sID, which is used to interact with this
	// session, and wc is the worker count, or 0, to use a default wc of 1.
	OpenSession(sType, sID string, wc int) error
	ListSessionTypes() []*pb.SessionType
	ListSessions() []SessionHandle
	CloseSession(string) error
	OpenWriteStream(string) (db.WriteStream, error)
	OpenReadStream(string, db.ReadFilter) (db.ReadStream, error)

	RunModule(string, string, map[string]string) error
	ListModuleTypes() []ModuleInfo
	ListRunningModules() []OpenModuleInfo
	CloseModule(string) error
	Close() error
}

// SessionHandle is used to return information on an open session.
type SessionHandle struct {
	Name     string
	SessType *pb.SessionType
	Session  *db.Session
}

// server is the primary server object. It contains a map of sessions
// and modules, a mutex to coordinate concurrent calls, and a configer
// to load sessions. It is the only "real" implementation of the BgpmondServer
// interface, but it does not replace it so the interface can be implemented
// in tests.
type server struct {
	sessions map[string]SessionHandle
	modules  map[string]Module
	conf     config.Configer
	mux      *sync.Mutex
}

// NewServer creates a BgpmondServer instance from a configuration. It loads
// session types from the configuration, and launches any modules specified.
// It returns an error if a module type is specified that isn't registered
// with the server.
func NewServer(conf config.Configer) (BgpmondServer, error) {
	s := &server{}
	s.sessions = make(map[string]SessionHandle)
	s.modules = make(map[string]Module)
	s.mux = &sync.Mutex{}
	s.conf = conf

	for _, mod := range conf.GetModules() {
		err := s.RunModule(mod.GetType(), mod.GetID(), mod.GetArgs())
		if err != nil {
			s.Close()
			return nil, err
		}
	}
	return s, nil
}

// NewServerFromFile does the same thing as NewServer, but loads the
// configuration from a specified file name. Returns an error if the
// file can't be parsed, or specifies an invalid module.
func NewServerFromFile(fName string) (BgpmondServer, error) {
	fd, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	bc, err := config.NewConfig(fd)
	if err != nil {
		return nil, err
	}

	return NewServer(bc)
}

// OpenSession will look for a configured session with type sType. It will
// create that session with ID sID and worker count workers. If that type
// does not exist, the ID is already taken, or the session fails to open,
// this function returns an error.
func (s *server) OpenSession(sType, sID string, workers int) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	if _, ok := s.sessions[sID]; ok {
		return coreLogger.Errorf("Session ID: %s already exists.", sID)
	}

	sc, err := s.conf.GetSessionConfigWithName(sType)
	if err != nil {
		return coreLogger.Errorf("No session type with session type name: %s found", sType)
	}

	session, err := db.NewSession(sc, sID, workers)
	if err != nil {
		return coreLogger.Errorf("Create session failed: %v", err)
	}

	hosts := fmt.Sprintf("Hosts: %v", sc.GetHostNames())

	stype := &pb.SessionType{
		Name: sc.GetName(),
		Type: sc.GetTypeName(),
		Desc: hosts,
	}
	sh := SessionHandle{
		Name:     sID,
		SessType: stype,
		Session:  session,
	}
	s.sessions[sID] = sh

	return nil
}

// CloseSession closes a single session with ID sID. It will return an error
// if sID does not exist on the server or the session returns an error in
// closing.
func (s *server) CloseSession(sID string) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	sh, ok := s.sessions[sID]
	if !ok {
		return coreLogger.Errorf("No session found with ID: %s", sID)
	}

	if err := sh.Session.Close(); err != nil {
		return err
	}

	delete(s.sessions, sID)

	return nil
}

// CloseAllSessions is a convenience method to shut down every open
// session on the server. This function could block if there are
// open streams on any of the sessions.
func (s *server) CloseAllSessions() {
	s.mux.Lock()
	defer s.mux.Unlock()

	for k, v := range s.sessions {
		v.Session.Close()
		delete(s.sessions, k)
	}

	return
}

// ListSessionTypes will return the configured types the server
// is aware of.
func (s *server) ListSessionTypes() []*pb.SessionType {
	// This function doesn't need to lock the sessions map because
	// it is only concerned with available sessions, not open ones.
	availSessions := []*pb.SessionType{}
	for _, sc := range s.conf.GetSessionConfigs() {
		availsess := &pb.SessionType{
			Name: sc.GetName(),
			Type: sc.GetTypeName(),
			Desc: fmt.Sprintf("Hosts: %v", sc.GetHostNames()),
		}

		availSessions = append(availSessions, availsess)
	}
	return availSessions
}

// ListSessions will return the handles for all currently open sessions
// on the server.
func (s *server) ListSessions() []SessionHandle {
	s.mux.Lock()
	defer s.mux.Unlock()

	sList := []SessionHandle{}
	for _, sh := range s.sessions {
		sList = append(sList, sh)
	}
	return sList
}

// OpenWriteStream will look up the session with ID sID and create/return a WriteStream
// on that session. This function can block if the session is already saturated with
// Streams.
func (s *server) OpenWriteStream(sID string) (db.WriteStream, error) {

	// The sessions are only locked here because the OpenWriteStream function below
	// can be blocking. If it blocked while the mutex was locked, this would lock
	// most functions on the server.
	s.mux.Lock()
	sh, ok := s.sessions[sID]
	s.mux.Unlock()

	if !ok {
		return nil, coreLogger.Errorf("Can't open stream on nonexistant session: %s", sID)
	}

	stream, err := sh.Session.OpenWriteStream(db.SessionWriteCapture)
	if err != nil {
		return nil, coreLogger.Errorf("Failed to open stream on session(%s): %s", sID, err)
	}

	return stream, nil
}

// OpenReadStream will look up the session with ID sID, and create/return a ReadStream
// with the provided filter. This function can block if the session is already saturated
// with Streams.
func (s *server) OpenReadStream(sID string, rf db.ReadFilter) (db.ReadStream, error) {

	// The sessions are only locked here because the OpenReadStream function below
	// can be blocking. If it blocked while the mutex was locked, this would lock
	// most functions on the server.
	s.mux.Lock()
	sh, ok := s.sessions[sID]
	s.mux.Unlock()

	if !ok {
		return nil, coreLogger.Errorf("Can't open stream on nonexistant session: %s", sID)
	}

	stream, err := sh.Session.OpenReadStream(db.SessionReadCapture, rf)
	if err != nil {
		return nil, coreLogger.Errorf("Failed to open stream on session(%s): %s", sID, err)
	}

	return stream, nil
}

// RunModule will look up the modType in the registered modules, create the module,
// add it to it's module map, and launch the module with modargs.
func (s *server) RunModule(modType, name string, modargs map[string]string) error {
	coreLogger.Infof("Running module %s with ID %s", modType, name)
	s.mux.Lock()
	defer s.mux.Unlock()

	if _, ok := s.modules[name]; ok {
		return coreLogger.Errorf("Module with ID: %s is already running", name)
	}

	maker, ok := getModuleMaker(modType)
	if !ok {
		return coreLogger.Errorf("No module type: %s found", modType)
	}

	newMod := maker(s, getModuleLogger(modType, name))
	s.modules[name] = newMod
	go newMod.Run(modargs, s.getFinishFunc(name))
	return nil
}

// getFinishFuncn generates a function which deletes the module with
// ID from the server.
func (s *server) getFinishFunc(id string) FinishFunc {
	return func() {
		s.mux.Lock()
		delete(s.modules, id)
		s.mux.Unlock()
	}
}

// ListModuleTypes lists the modules available for the server to run.
func (s *server) ListModuleTypes() []ModuleInfo {
	return getModuleTypes()
}

// ListRunningModules lists all currently running modules on the server.
func (s *server) ListRunningModules() []OpenModuleInfo {
	s.mux.Lock()
	defer s.mux.Unlock()

	var ret []OpenModuleInfo
	for k, v := range s.modules {
		info := v.GetInfo()
		// Modules aren't aware of their own ID, so it has to be populated here
		info.ID = k
		ret = append(ret, info)
	}
	return ret
}

// CloseModule will stop a module with ID name. If name is not the ID of
// a running module, or there is an error while closing it, this function
// returns an error.
func (s *server) CloseModule(name string) error {

	// This is only locked here because the Stop() function below can block,
	// and if it blocked while the mux was locked, this could lock most functions
	// on the server.
	s.mux.Lock()
	mod, ok := s.modules[name]
	s.mux.Unlock()

	if !ok {
		return coreLogger.Errorf("No module with ID: %s found", name)
	}

	if err := mod.Stop(); err != nil {
		return err
	}

	delete(s.modules, name)
	return nil
}

// CloseAllModules is a convenience method to close all modules running
// on the server.
func (s *server) CloseAllModules() {
	s.mux.Lock()
	defer s.mux.Unlock()

	for k, v := range s.modules {
		v.Stop()
		delete(s.modules, k)
	}
}

// Close completely shuts down the server.
func (s *server) Close() error {
	s.CloseAllModules()
	s.CloseAllSessions()
	return nil
}
