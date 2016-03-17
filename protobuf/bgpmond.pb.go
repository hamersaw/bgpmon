// Code generated by protoc-gen-go.
// source: protobuf/bgpmond.proto
// DO NOT EDIT!

/*
Package bgpmond is a generated protocol buffer package.

It is generated from these files:
	protobuf/bgpmond.proto

It has these top-level messages:
	Empty
	GoBGPLinkModule
	PrefixHijackModule
	ListModulesReply
	RunModuleRequest
	RunModuleReply
	StartModuleRequest
	StartModuleReply
	StopModuleRequest
	CassandraSession
	FileSession
	CloseSessionRequest
	ListSessionsReply
	OpenSessionRequest
	OpenSessionReply
	WriteRequest
	ASNumberLocation
	BGPUpdateMessage
	IPAddressLocation
	Location
	PrefixLocation
*/
package bgpmond

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//
// Module Messages
type ModuleType int32

const (
	ModuleType_GOBGP_LINK    ModuleType = 0
	ModuleType_PREFIX_HIJACK ModuleType = 1
)

var ModuleType_name = map[int32]string{
	0: "GOBGP_LINK",
	1: "PREFIX_HIJACK",
}
var ModuleType_value = map[string]int32{
	"GOBGP_LINK":    0,
	"PREFIX_HIJACK": 1,
}

func (x ModuleType) String() string {
	return proto.EnumName(ModuleType_name, int32(x))
}
func (ModuleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

//
// Session Messages
type SessionType int32

const (
	SessionType_CASSANDRA SessionType = 0
	SessionType_FILE      SessionType = 1
)

var SessionType_name = map[int32]string{
	0: "CASSANDRA",
	1: "FILE",
}
var SessionType_value = map[string]int32{
	"CASSANDRA": 0,
	"FILE":      1,
}

func (x SessionType) String() string {
	return proto.EnumName(SessionType_name, int32(x))
}
func (SessionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type WriteRequest_Type int32

const (
	WriteRequest_AS_NUMBER_LOCATION  WriteRequest_Type = 0
	WriteRequest_BGP_UPDATE          WriteRequest_Type = 1
	WriteRequest_IP_ADDRESS_LOCATION WriteRequest_Type = 2
	WriteRequest_PREFIX_LOCATION     WriteRequest_Type = 3
)

var WriteRequest_Type_name = map[int32]string{
	0: "AS_NUMBER_LOCATION",
	1: "BGP_UPDATE",
	2: "IP_ADDRESS_LOCATION",
	3: "PREFIX_LOCATION",
}
var WriteRequest_Type_value = map[string]int32{
	"AS_NUMBER_LOCATION":  0,
	"BGP_UPDATE":          1,
	"IP_ADDRESS_LOCATION": 2,
	"PREFIX_LOCATION":     3,
}

func (x WriteRequest_Type) String() string {
	return proto.EnumName(WriteRequest_Type_name, int32(x))
}
func (WriteRequest_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{15, 0} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GoBGPLinkModule struct {
	Address      string   `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	OutSessionId []string `protobuf:"bytes,2,rep,name=out_session_id" json:"out_session_id,omitempty"`
}

func (m *GoBGPLinkModule) Reset()                    { *m = GoBGPLinkModule{} }
func (m *GoBGPLinkModule) String() string            { return proto.CompactTextString(m) }
func (*GoBGPLinkModule) ProtoMessage()               {}
func (*GoBGPLinkModule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PrefixHijackModule struct {
	Prefix          string   `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
	AsNumber        []uint32 `protobuf:"varint,2,rep,name=as_number" json:"as_number,omitempty"`
	PeriodicSeconds int32    `protobuf:"varint,3,opt,name=periodic_seconds" json:"periodic_seconds,omitempty"`
	TimeoutSeconds  int32    `protobuf:"varint,4,opt,name=timeout_seconds" json:"timeout_seconds,omitempty"`
	InSessionId     []string `protobuf:"bytes,5,rep,name=in_session_id" json:"in_session_id,omitempty"`
}

func (m *PrefixHijackModule) Reset()                    { *m = PrefixHijackModule{} }
func (m *PrefixHijackModule) String() string            { return proto.CompactTextString(m) }
func (*PrefixHijackModule) ProtoMessage()               {}
func (*PrefixHijackModule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

//
// Module Command Messages
type ListModulesReply struct {
	ModuleId []string `protobuf:"bytes,1,rep,name=module_id" json:"module_id,omitempty"`
}

func (m *ListModulesReply) Reset()                    { *m = ListModulesReply{} }
func (m *ListModulesReply) String() string            { return proto.CompactTextString(m) }
func (*ListModulesReply) ProtoMessage()               {}
func (*ListModulesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type RunModuleRequest struct {
	Type               ModuleType          `protobuf:"varint,1,opt,name=type,enum=bgpmond.ModuleType" json:"type,omitempty"`
	PrefixHijackModule *PrefixHijackModule `protobuf:"bytes,2,opt,name=prefix_hijack_module" json:"prefix_hijack_module,omitempty"`
}

func (m *RunModuleRequest) Reset()                    { *m = RunModuleRequest{} }
func (m *RunModuleRequest) String() string            { return proto.CompactTextString(m) }
func (*RunModuleRequest) ProtoMessage()               {}
func (*RunModuleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RunModuleRequest) GetPrefixHijackModule() *PrefixHijackModule {
	if m != nil {
		return m.PrefixHijackModule
	}
	return nil
}

type RunModuleReply struct {
	ModuleMessage string `protobuf:"bytes,3,opt,name=module_message" json:"module_message,omitempty"`
}

func (m *RunModuleReply) Reset()                    { *m = RunModuleReply{} }
func (m *RunModuleReply) String() string            { return proto.CompactTextString(m) }
func (*RunModuleReply) ProtoMessage()               {}
func (*RunModuleReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type StartModuleRequest struct {
	Type               ModuleType          `protobuf:"varint,1,opt,name=type,enum=bgpmond.ModuleType" json:"type,omitempty"`
	ModuleId           string              `protobuf:"bytes,2,opt,name=module_id" json:"module_id,omitempty"`
	GobgpLinkModule    *GoBGPLinkModule    `protobuf:"bytes,3,opt,name=gobgp_link_module" json:"gobgp_link_module,omitempty"`
	PrefixHijackModule *PrefixHijackModule `protobuf:"bytes,4,opt,name=prefix_hijack_module" json:"prefix_hijack_module,omitempty"`
}

func (m *StartModuleRequest) Reset()                    { *m = StartModuleRequest{} }
func (m *StartModuleRequest) String() string            { return proto.CompactTextString(m) }
func (*StartModuleRequest) ProtoMessage()               {}
func (*StartModuleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *StartModuleRequest) GetGobgpLinkModule() *GoBGPLinkModule {
	if m != nil {
		return m.GobgpLinkModule
	}
	return nil
}

func (m *StartModuleRequest) GetPrefixHijackModule() *PrefixHijackModule {
	if m != nil {
		return m.PrefixHijackModule
	}
	return nil
}

type StartModuleReply struct {
	ModuleId string `protobuf:"bytes,3,opt,name=module_id" json:"module_id,omitempty"`
}

func (m *StartModuleReply) Reset()                    { *m = StartModuleReply{} }
func (m *StartModuleReply) String() string            { return proto.CompactTextString(m) }
func (*StartModuleReply) ProtoMessage()               {}
func (*StartModuleReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type StopModuleRequest struct {
	ModuleId string `protobuf:"bytes,1,opt,name=module_id" json:"module_id,omitempty"`
}

func (m *StopModuleRequest) Reset()                    { *m = StopModuleRequest{} }
func (m *StopModuleRequest) String() string            { return proto.CompactTextString(m) }
func (*StopModuleRequest) ProtoMessage()               {}
func (*StopModuleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type CassandraSession struct {
	Username string   `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string   `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Hosts    []string `protobuf:"bytes,3,rep,name=hosts" json:"hosts,omitempty"`
}

func (m *CassandraSession) Reset()                    { *m = CassandraSession{} }
func (m *CassandraSession) String() string            { return proto.CompactTextString(m) }
func (*CassandraSession) ProtoMessage()               {}
func (*CassandraSession) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type FileSession struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
}

func (m *FileSession) Reset()                    { *m = FileSession{} }
func (m *FileSession) String() string            { return proto.CompactTextString(m) }
func (*FileSession) ProtoMessage()               {}
func (*FileSession) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

//
// Session Command Messages
type CloseSessionRequest struct {
	SessionId string `protobuf:"bytes,1,opt,name=session_id" json:"session_id,omitempty"`
}

func (m *CloseSessionRequest) Reset()                    { *m = CloseSessionRequest{} }
func (m *CloseSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*CloseSessionRequest) ProtoMessage()               {}
func (*CloseSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type ListSessionsReply struct {
	SessionId []string `protobuf:"bytes,1,rep,name=session_id" json:"session_id,omitempty"`
}

func (m *ListSessionsReply) Reset()                    { *m = ListSessionsReply{} }
func (m *ListSessionsReply) String() string            { return proto.CompactTextString(m) }
func (*ListSessionsReply) ProtoMessage()               {}
func (*ListSessionsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type OpenSessionRequest struct {
	Type             SessionType       `protobuf:"varint,1,opt,name=type,enum=bgpmond.SessionType" json:"type,omitempty"`
	SessionId        string            `protobuf:"bytes,2,opt,name=session_id" json:"session_id,omitempty"`
	CassandraSession *CassandraSession `protobuf:"bytes,3,opt,name=cassandra_session" json:"cassandra_session,omitempty"`
	FileSession      *FileSession      `protobuf:"bytes,4,opt,name=file_session" json:"file_session,omitempty"`
}

func (m *OpenSessionRequest) Reset()                    { *m = OpenSessionRequest{} }
func (m *OpenSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*OpenSessionRequest) ProtoMessage()               {}
func (*OpenSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *OpenSessionRequest) GetCassandraSession() *CassandraSession {
	if m != nil {
		return m.CassandraSession
	}
	return nil
}

func (m *OpenSessionRequest) GetFileSession() *FileSession {
	if m != nil {
		return m.FileSession
	}
	return nil
}

type OpenSessionReply struct {
	SessionId string `protobuf:"bytes,3,opt,name=session_id" json:"session_id,omitempty"`
}

func (m *OpenSessionReply) Reset()                    { *m = OpenSessionReply{} }
func (m *OpenSessionReply) String() string            { return proto.CompactTextString(m) }
func (*OpenSessionReply) ProtoMessage()               {}
func (*OpenSessionReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

//
// Write Messages
type WriteRequest struct {
	Type              WriteRequest_Type  `protobuf:"varint,1,opt,name=type,enum=bgpmond.WriteRequest_Type" json:"type,omitempty"`
	AsNumberLocation  *ASNumberLocation  `protobuf:"bytes,2,opt,name=as_number_location" json:"as_number_location,omitempty"`
	BgpUpdateMessage  *BGPUpdateMessage  `protobuf:"bytes,3,opt,name=bgp_update_message" json:"bgp_update_message,omitempty"`
	IpAddressLocation *IPAddressLocation `protobuf:"bytes,4,opt,name=ip_address_location" json:"ip_address_location,omitempty"`
	PrefixLocation    *PrefixLocation    `protobuf:"bytes,5,opt,name=prefix_location" json:"prefix_location,omitempty"`
}

func (m *WriteRequest) Reset()                    { *m = WriteRequest{} }
func (m *WriteRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()               {}
func (*WriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *WriteRequest) GetAsNumberLocation() *ASNumberLocation {
	if m != nil {
		return m.AsNumberLocation
	}
	return nil
}

func (m *WriteRequest) GetBgpUpdateMessage() *BGPUpdateMessage {
	if m != nil {
		return m.BgpUpdateMessage
	}
	return nil
}

func (m *WriteRequest) GetIpAddressLocation() *IPAddressLocation {
	if m != nil {
		return m.IpAddressLocation
	}
	return nil
}

func (m *WriteRequest) GetPrefixLocation() *PrefixLocation {
	if m != nil {
		return m.PrefixLocation
	}
	return nil
}

type ASNumberLocation struct {
	ASNumber    uint32    `protobuf:"varint,1,opt,name=ASNumber" json:"ASNumber,omitempty"`
	MeasureDate string    `protobuf:"bytes,2,opt,name=measure_date" json:"measure_date,omitempty"`
	Location    *Location `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
}

func (m *ASNumberLocation) Reset()                    { *m = ASNumberLocation{} }
func (m *ASNumberLocation) String() string            { return proto.CompactTextString(m) }
func (*ASNumberLocation) ProtoMessage()               {}
func (*ASNumberLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *ASNumberLocation) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type BGPUpdateMessage struct {
}

func (m *BGPUpdateMessage) Reset()                    { *m = BGPUpdateMessage{} }
func (m *BGPUpdateMessage) String() string            { return proto.CompactTextString(m) }
func (*BGPUpdateMessage) ProtoMessage()               {}
func (*BGPUpdateMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

type IPAddressLocation struct {
	IpAddress   string    `protobuf:"bytes,1,opt,name=ip_address" json:"ip_address,omitempty"`
	MeasureDate string    `protobuf:"bytes,2,opt,name=measure_date" json:"measure_date,omitempty"`
	Location    *Location `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
}

func (m *IPAddressLocation) Reset()                    { *m = IPAddressLocation{} }
func (m *IPAddressLocation) String() string            { return proto.CompactTextString(m) }
func (*IPAddressLocation) ProtoMessage()               {}
func (*IPAddressLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *IPAddressLocation) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type Location struct {
	CountryCode string  `protobuf:"bytes,1,opt,name=country_code" json:"country_code,omitempty"`
	StateCode   string  `protobuf:"bytes,2,opt,name=state_code" json:"state_code,omitempty"`
	City        string  `protobuf:"bytes,3,opt,name=city" json:"city,omitempty"`
	Latitude    float64 `protobuf:"fixed64,4,opt,name=latitude" json:"latitude,omitempty"`
	Longitude   float64 `protobuf:"fixed64,5,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

type PrefixLocation struct {
	PrefixIpAddress string    `protobuf:"bytes,1,opt,name=prefix_ip_address" json:"prefix_ip_address,omitempty"`
	PrefixMask      uint32    `protobuf:"varint,2,opt,name=prefix_mask" json:"prefix_mask,omitempty"`
	MeasureDate     string    `protobuf:"bytes,3,opt,name=measure_date" json:"measure_date,omitempty"`
	Location        *Location `protobuf:"bytes,4,opt,name=location" json:"location,omitempty"`
}

func (m *PrefixLocation) Reset()                    { *m = PrefixLocation{} }
func (m *PrefixLocation) String() string            { return proto.CompactTextString(m) }
func (*PrefixLocation) ProtoMessage()               {}
func (*PrefixLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *PrefixLocation) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "bgpmond.Empty")
	proto.RegisterType((*GoBGPLinkModule)(nil), "bgpmond.GoBGPLinkModule")
	proto.RegisterType((*PrefixHijackModule)(nil), "bgpmond.PrefixHijackModule")
	proto.RegisterType((*ListModulesReply)(nil), "bgpmond.ListModulesReply")
	proto.RegisterType((*RunModuleRequest)(nil), "bgpmond.RunModuleRequest")
	proto.RegisterType((*RunModuleReply)(nil), "bgpmond.RunModuleReply")
	proto.RegisterType((*StartModuleRequest)(nil), "bgpmond.StartModuleRequest")
	proto.RegisterType((*StartModuleReply)(nil), "bgpmond.StartModuleReply")
	proto.RegisterType((*StopModuleRequest)(nil), "bgpmond.StopModuleRequest")
	proto.RegisterType((*CassandraSession)(nil), "bgpmond.CassandraSession")
	proto.RegisterType((*FileSession)(nil), "bgpmond.FileSession")
	proto.RegisterType((*CloseSessionRequest)(nil), "bgpmond.CloseSessionRequest")
	proto.RegisterType((*ListSessionsReply)(nil), "bgpmond.ListSessionsReply")
	proto.RegisterType((*OpenSessionRequest)(nil), "bgpmond.OpenSessionRequest")
	proto.RegisterType((*OpenSessionReply)(nil), "bgpmond.OpenSessionReply")
	proto.RegisterType((*WriteRequest)(nil), "bgpmond.WriteRequest")
	proto.RegisterType((*ASNumberLocation)(nil), "bgpmond.ASNumberLocation")
	proto.RegisterType((*BGPUpdateMessage)(nil), "bgpmond.BGPUpdateMessage")
	proto.RegisterType((*IPAddressLocation)(nil), "bgpmond.IPAddressLocation")
	proto.RegisterType((*Location)(nil), "bgpmond.Location")
	proto.RegisterType((*PrefixLocation)(nil), "bgpmond.PrefixLocation")
	proto.RegisterEnum("bgpmond.ModuleType", ModuleType_name, ModuleType_value)
	proto.RegisterEnum("bgpmond.SessionType", SessionType_name, SessionType_value)
	proto.RegisterEnum("bgpmond.WriteRequest_Type", WriteRequest_Type_name, WriteRequest_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Bgpmond service

type BgpmondClient interface {
	CloseSession(ctx context.Context, in *CloseSessionRequest, opts ...grpc.CallOption) (*Empty, error)
	ListModules(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListModulesReply, error)
	ListSessions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListSessionsReply, error)
	RunModule(ctx context.Context, in *RunModuleRequest, opts ...grpc.CallOption) (*RunModuleReply, error)
	StartModule(ctx context.Context, in *StartModuleRequest, opts ...grpc.CallOption) (*StartModuleReply, error)
	StopModule(ctx context.Context, in *StopModuleRequest, opts ...grpc.CallOption) (*Empty, error)
	OpenSession(ctx context.Context, in *OpenSessionRequest, opts ...grpc.CallOption) (*OpenSessionReply, error)
}

type bgpmondClient struct {
	cc *grpc.ClientConn
}

func NewBgpmondClient(cc *grpc.ClientConn) BgpmondClient {
	return &bgpmondClient{cc}
}

func (c *bgpmondClient) CloseSession(ctx context.Context, in *CloseSessionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/bgpmond.Bgpmond/CloseSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bgpmondClient) ListModules(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListModulesReply, error) {
	out := new(ListModulesReply)
	err := grpc.Invoke(ctx, "/bgpmond.Bgpmond/ListModules", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bgpmondClient) ListSessions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListSessionsReply, error) {
	out := new(ListSessionsReply)
	err := grpc.Invoke(ctx, "/bgpmond.Bgpmond/ListSessions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bgpmondClient) RunModule(ctx context.Context, in *RunModuleRequest, opts ...grpc.CallOption) (*RunModuleReply, error) {
	out := new(RunModuleReply)
	err := grpc.Invoke(ctx, "/bgpmond.Bgpmond/RunModule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bgpmondClient) StartModule(ctx context.Context, in *StartModuleRequest, opts ...grpc.CallOption) (*StartModuleReply, error) {
	out := new(StartModuleReply)
	err := grpc.Invoke(ctx, "/bgpmond.Bgpmond/StartModule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bgpmondClient) StopModule(ctx context.Context, in *StopModuleRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/bgpmond.Bgpmond/StopModule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bgpmondClient) OpenSession(ctx context.Context, in *OpenSessionRequest, opts ...grpc.CallOption) (*OpenSessionReply, error) {
	out := new(OpenSessionReply)
	err := grpc.Invoke(ctx, "/bgpmond.Bgpmond/OpenSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Bgpmond service

type BgpmondServer interface {
	CloseSession(context.Context, *CloseSessionRequest) (*Empty, error)
	ListModules(context.Context, *Empty) (*ListModulesReply, error)
	ListSessions(context.Context, *Empty) (*ListSessionsReply, error)
	RunModule(context.Context, *RunModuleRequest) (*RunModuleReply, error)
	StartModule(context.Context, *StartModuleRequest) (*StartModuleReply, error)
	StopModule(context.Context, *StopModuleRequest) (*Empty, error)
	OpenSession(context.Context, *OpenSessionRequest) (*OpenSessionReply, error)
}

func RegisterBgpmondServer(s *grpc.Server, srv BgpmondServer) {
	s.RegisterService(&_Bgpmond_serviceDesc, srv)
}

func _Bgpmond_CloseSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CloseSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BgpmondServer).CloseSession(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Bgpmond_ListModules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BgpmondServer).ListModules(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Bgpmond_ListSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BgpmondServer).ListSessions(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Bgpmond_RunModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RunModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BgpmondServer).RunModule(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Bgpmond_StartModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(StartModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BgpmondServer).StartModule(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Bgpmond_StopModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(StopModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BgpmondServer).StopModule(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Bgpmond_OpenSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(OpenSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BgpmondServer).OpenSession(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Bgpmond_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bgpmond.Bgpmond",
	HandlerType: (*BgpmondServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CloseSession",
			Handler:    _Bgpmond_CloseSession_Handler,
		},
		{
			MethodName: "ListModules",
			Handler:    _Bgpmond_ListModules_Handler,
		},
		{
			MethodName: "ListSessions",
			Handler:    _Bgpmond_ListSessions_Handler,
		},
		{
			MethodName: "RunModule",
			Handler:    _Bgpmond_RunModule_Handler,
		},
		{
			MethodName: "StartModule",
			Handler:    _Bgpmond_StartModule_Handler,
		},
		{
			MethodName: "StopModule",
			Handler:    _Bgpmond_StopModule_Handler,
		},
		{
			MethodName: "OpenSession",
			Handler:    _Bgpmond_OpenSession_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 1168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0xdb, 0x6e, 0xe3, 0x54,
	0x17, 0xae, 0x9b, 0x64, 0x9a, 0xac, 0x9c, 0x9c, 0xdd, 0x6a, 0xda, 0x3f, 0x33, 0xbf, 0x18, 0x2c,
	0x0e, 0x9d, 0x4a, 0xb4, 0x50, 0x90, 0x46, 0x42, 0x08, 0xe1, 0x26, 0x6d, 0x27, 0x33, 0x69, 0x1b,
	0x9c, 0x56, 0x70, 0x81, 0x64, 0xdc, 0xd8, 0xed, 0x98, 0x49, 0x6c, 0xe3, 0x83, 0xa0, 0x77, 0x5c,
	0x70, 0xc5, 0x23, 0xf0, 0x12, 0xbc, 0x00, 0x37, 0xbc, 0x03, 0xd7, 0xbc, 0x02, 0xaf, 0xc0, 0x3e,
	0x7b, 0xdb, 0x1e, 0x40, 0x1a, 0x89, 0x3b, 0xaf, 0x6f, 0xaf, 0xb5, 0xbc, 0xd6, 0xb7, 0xbf, 0xb5,
	0xf7, 0x86, 0xfb, 0x51, 0x1c, 0xa6, 0xe1, 0x75, 0x76, 0x73, 0x70, 0x7d, 0x1b, 0xad, 0xc2, 0xc0,
	0xdd, 0xa7, 0x00, 0xda, 0xe0, 0xa6, 0xb1, 0x01, 0x8d, 0xe3, 0x55, 0x94, 0xde, 0x19, 0x9f, 0x43,
	0xff, 0x34, 0x3c, 0x3a, 0x9d, 0x4d, 0xfd, 0xe0, 0xe5, 0x59, 0xe8, 0x66, 0x4b, 0x0f, 0xed, 0xc0,
	0x86, 0xe3, 0xba, 0xb1, 0x97, 0x24, 0x3b, 0xda, 0x23, 0x6d, 0xb7, 0x65, 0x09, 0x13, 0xbd, 0x05,
	0xbd, 0x30, 0x4b, 0xed, 0x04, 0x7f, 0xfb, 0x61, 0x60, 0xfb, 0xee, 0xce, 0xfa, 0xa3, 0x1a, 0x76,
	0xe8, 0x60, 0x74, 0xce, 0xc0, 0x89, 0x6b, 0xfc, 0xa6, 0x01, 0x9a, 0xc5, 0xde, 0x8d, 0xff, 0xfd,
	0x53, 0xff, 0x1b, 0x67, 0x21, 0xd2, 0xde, 0x87, 0x7b, 0x11, 0x45, 0x79, 0x56, 0x6e, 0xa1, 0x07,
	0xd0, 0x72, 0x12, 0x3b, 0xc8, 0x56, 0xd7, 0x5e, 0x4c, 0xf3, 0x75, 0xad, 0xa6, 0x93, 0x9c, 0x53,
	0x1b, 0x3d, 0x06, 0x3d, 0xf2, 0x62, 0x3f, 0x74, 0xfd, 0x05, 0xfe, 0xed, 0x02, 0x97, 0x9e, 0xec,
	0xd4, 0x70, 0x78, 0xc3, 0xea, 0x0b, 0x7c, 0xce, 0x60, 0xf4, 0x2e, 0xf4, 0x53, 0x7f, 0xe5, 0xb1,
	0x02, 0x99, 0x67, 0x9d, 0x7a, 0xf6, 0x38, 0x2c, 0x1c, 0x0d, 0xe8, 0xfa, 0x81, 0xda, 0x44, 0x83,
	0x36, 0xd1, 0xf6, 0x83, 0xbc, 0x87, 0x03, 0xd0, 0xa7, 0x7e, 0x92, 0xb2, 0xd2, 0x13, 0xcb, 0x8b,
	0x96, 0x77, 0xa4, 0xd0, 0x15, 0xb5, 0x49, 0x8c, 0x46, 0x63, 0x9a, 0x0c, 0xc0, 0x01, 0x3f, 0x69,
	0xa0, 0x5b, 0x59, 0xc0, 0x02, 0x2c, 0xef, 0xdb, 0xcc, 0x4b, 0x52, 0x5c, 0x52, 0x3d, 0xbd, 0x8b,
	0x3c, 0xda, 0x70, 0xef, 0x70, 0x73, 0x5f, 0x6c, 0x06, 0xf3, 0xba, 0xc4, 0x4b, 0x16, 0x75, 0x40,
	0x67, 0xb0, 0xc5, 0xd8, 0xb0, 0x5f, 0x50, 0xca, 0x6c, 0x96, 0x17, 0xd3, 0xa1, 0xed, 0xb6, 0x0f,
	0x1f, 0xc8, 0xc0, 0x2a, 0xad, 0x16, 0x8a, 0x2a, 0x98, 0xf1, 0x04, 0x7a, 0x4a, 0x2d, 0xa4, 0xf6,
	0xb7, 0xa1, 0xc7, 0x6b, 0x5f, 0xe1, 0x1e, 0x9d, 0x5b, 0x8f, 0xb2, 0xd8, 0xb2, 0xba, 0x0c, 0x3d,
	0x63, 0xa0, 0xf1, 0x27, 0xde, 0xba, 0x79, 0xea, 0xc4, 0xe9, 0x6b, 0xf6, 0x51, 0xa0, 0x68, 0x9d,
	0xfe, 0x41, 0x52, 0x84, 0xc6, 0x30, 0xb8, 0x0d, 0x71, 0xa8, 0xbd, 0xc4, 0x5a, 0x13, 0x1d, 0xd6,
	0x68, 0x87, 0x3b, 0x32, 0x65, 0x49, 0x8c, 0x56, 0x9f, 0x86, 0x28, 0xea, 0xfc, 0x3b, 0xaa, 0xea,
	0xaf, 0x47, 0x15, 0xde, 0xe8, 0x42, 0xc3, 0x95, 0x8d, 0xae, 0x15, 0xbb, 0x30, 0xde, 0x87, 0xc1,
	0x3c, 0x0d, 0xa3, 0x22, 0x41, 0x25, 0x69, 0x14, 0x23, 0xbe, 0x06, 0x7d, 0xe4, 0x60, 0x7e, 0x03,
	0x37, 0x76, 0xb8, 0xc2, 0xd0, 0x10, 0x9a, 0x59, 0xe2, 0xc5, 0x81, 0xb3, 0xf2, 0x84, 0xbf, 0xb0,
	0xc9, 0x5a, 0x84, 0xfd, 0xbf, 0x0b, 0x63, 0xc9, 0xa1, 0xb0, 0xd1, 0x16, 0x34, 0x5e, 0x84, 0x49,
	0x4a, 0x86, 0x80, 0xe8, 0x8f, 0x19, 0xc6, 0x63, 0x68, 0x9f, 0xf8, 0x4b, 0x4f, 0x49, 0x7e, 0x83,
	0x4d, 0x35, 0xb9, 0xb0, 0x8d, 0x8f, 0x60, 0x73, 0xb4, 0x0c, 0x13, 0xe1, 0x2b, 0x1a, 0xf8, 0x3f,
	0x80, 0x32, 0x10, 0x2c, 0xa8, 0x95, 0xc8, 0x71, 0x38, 0x84, 0x01, 0x19, 0x07, 0x1e, 0xc4, 0xe7,
	0xa1, 0x1c, 0x53, 0x2b, 0xc6, 0xfc, 0x81, 0xb5, 0x74, 0x11, 0x79, 0x41, 0xe9, 0x4f, 0xbb, 0x05,
	0x2d, 0x6d, 0xc9, 0xfd, 0xe2, 0x6e, 0x8a, 0x98, 0x8a, 0xf9, 0xd7, 0x4b, 0x35, 0xa1, 0x13, 0x18,
	0x2c, 0x04, 0xad, 0x62, 0x9a, 0xb9, 0x9c, 0xfe, 0x27, 0xb3, 0x96, 0x89, 0xb7, 0xf4, 0x45, 0x79,
	0x2b, 0x9e, 0x40, 0x87, 0xb0, 0x23, 0x53, 0x30, 0x21, 0xe5, 0x85, 0x29, 0xcc, 0x5a, 0xed, 0x9b,
	0xdc, 0x30, 0x3e, 0x00, 0xbd, 0xd0, 0x5f, 0x95, 0x93, 0x5a, 0x99, 0xc7, 0x5f, 0x6b, 0xd0, 0xf9,
	0x22, 0xf6, 0x53, 0x29, 0x9c, 0xfd, 0x02, 0x1b, 0x43, 0xf9, 0x53, 0xd5, 0x69, 0x5f, 0xe1, 0xe4,
	0x14, 0x90, 0x3c, 0x2c, 0xed, 0x65, 0xb8, 0x70, 0x52, 0x52, 0xf2, 0x7a, 0xa9, 0x6b, 0x73, 0xce,
	0x8e, 0xcf, 0x29, 0x77, 0xb0, 0x74, 0x71, 0xa0, 0x0a, 0x84, 0x24, 0x22, 0xa3, 0x98, 0x45, 0xae,
	0x93, 0x16, 0x0f, 0x05, 0x35, 0x11, 0x9e, 0xc5, 0x2b, 0xea, 0xc1, 0x0f, 0x08, 0x4b, 0xc7, 0x2b,
	0x05, 0x04, 0x3d, 0x83, 0x4d, 0x3f, 0xb2, 0xf9, 0x0d, 0x91, 0x97, 0xc4, 0x58, 0xcc, 0x1b, 0x9a,
	0xcc, 0x4c, 0xe6, 0x22, 0x6b, 0x1a, 0xf8, 0x51, 0x09, 0x42, 0x9f, 0x41, 0x9f, 0xcf, 0xb6, 0xcc,
	0xd3, 0xa0, 0x79, 0xb6, 0x4b, 0x63, 0x2d, 0x93, 0xf4, 0xa2, 0x82, 0x6d, 0x7c, 0x05, 0x75, 0xc2,
	0x16, 0xbe, 0x6c, 0x90, 0x39, 0xb7, 0xcf, 0xaf, 0xce, 0x8e, 0x8e, 0x2d, 0x7b, 0x7a, 0x31, 0x32,
	0x2f, 0x27, 0x17, 0xe7, 0xfa, 0x1a, 0xea, 0x01, 0xe0, 0x9e, 0xec, 0xab, 0xd9, 0xd8, 0xbc, 0x3c,
	0xd6, 0x35, 0xb4, 0x0d, 0x9b, 0x93, 0x99, 0x6d, 0x8e, 0xc7, 0xd6, 0xf1, 0x7c, 0x9e, 0x3b, 0xae,
	0xa3, 0x4d, 0xe8, 0xcf, 0xac, 0xe3, 0x93, 0xc9, 0x97, 0x39, 0x58, 0x33, 0x7e, 0xc0, 0x87, 0x7c,
	0x99, 0x5b, 0x32, 0x6d, 0x02, 0xa3, 0xdb, 0x48, 0xae, 0x2f, 0x6e, 0xa3, 0x37, 0xa1, 0xb3, 0xf2,
	0x9c, 0x24, 0x8b, 0x3d, 0x9b, 0x70, 0xc6, 0x45, 0xdc, 0xe6, 0xd8, 0x18, 0x43, 0xe8, 0x3d, 0x68,
	0xca, 0x66, 0x19, 0xfd, 0x03, 0xd9, 0xac, 0x6c, 0x53, 0xba, 0x18, 0x08, 0xf4, 0xf2, 0xa6, 0x18,
	0x3f, 0x6a, 0x30, 0xa8, 0xf0, 0x4b, 0xa4, 0x98, 0x6f, 0x8c, 0x18, 0x69, 0xc9, 0xf9, 0x7f, 0x50,
	0xda, 0xcf, 0x1a, 0x34, 0xe5, 0xdf, 0x71, 0xfa, 0x45, 0x98, 0x05, 0x69, 0x7c, 0x67, 0x2f, 0x42,
	0x57, 0x9c, 0x43, 0x6d, 0x8e, 0x8d, 0x30, 0x44, 0x67, 0x25, 0x25, 0xea, 0xa3, 0x0e, 0x62, 0xbe,
	0x09, 0x42, 0x97, 0x11, 0xd4, 0x17, 0x7e, 0x7a, 0xc7, 0x87, 0x88, 0x7e, 0x13, 0xae, 0x97, 0x38,
	0x7f, 0x9a, 0xb9, 0xec, 0xc0, 0xd7, 0x2c, 0x69, 0xa3, 0x87, 0xd0, 0x5a, 0x86, 0xc1, 0x2d, 0x5b,
	0x6c, 0xd0, 0xc5, 0x1c, 0x30, 0x7e, 0xd1, 0xa0, 0x57, 0xd4, 0x0e, 0xda, 0x83, 0x01, 0x57, 0x5b,
	0x85, 0x27, 0x2e, 0xc3, 0x89, 0x64, 0xeb, 0x0d, 0x68, 0x73, 0xdf, 0x95, 0x93, 0xbc, 0xa4, 0xc5,
	0x76, 0x2d, 0x60, 0xd0, 0x19, 0x46, 0x2a, 0x74, 0xd6, 0xfe, 0x99, 0xce, 0xfa, 0xbf, 0xd2, 0xb9,
	0x77, 0x00, 0x90, 0xdf, 0xaf, 0x44, 0xb8, 0xa7, 0x17, 0x44, 0xba, 0xd3, 0xc9, 0xf9, 0x73, 0x2c,
	0xe4, 0x01, 0x74, 0xb9, 0x3e, 0x9f, 0x4e, 0x9e, 0x99, 0xa3, 0xe7, 0xba, 0xb6, 0xf7, 0x0e, 0xb4,
	0x95, 0x43, 0x14, 0x75, 0xa1, 0x35, 0x32, 0xe7, 0x73, 0xf3, 0x7c, 0x6c, 0x99, 0x38, 0xa0, 0x09,
	0xf5, 0x93, 0xc9, 0x14, 0x6b, 0xfe, 0xf0, 0xf7, 0x1a, 0x6c, 0x1c, 0xb1, 0xff, 0xa2, 0x4f, 0xa1,
	0xa3, 0x5e, 0x07, 0xe8, 0x61, 0x7e, 0x72, 0x56, 0x6f, 0x89, 0x61, 0x4f, 0xae, 0xb2, 0xc7, 0xe3,
	0x1a, 0xfa, 0x18, 0xda, 0xca, 0x3b, 0x09, 0x95, 0x1c, 0x86, 0xf9, 0x49, 0x52, 0x7e, 0x4d, 0xe1,
	0xd8, 0x4f, 0xa0, 0xa3, 0x5e, 0x2a, 0x95, 0xe0, 0x61, 0x21, 0xb8, 0x70, 0xf7, 0xe0, 0x68, 0x13,
	0x5a, 0xf2, 0x8d, 0x83, 0xf2, 0xff, 0x94, 0xdf, 0x60, 0xc3, 0xed, 0x57, 0x2d, 0xb1, 0x14, 0xa7,
	0x98, 0xb0, 0xfc, 0xee, 0x47, 0xf9, 0xdb, 0xa1, 0xfa, 0x04, 0x52, 0x3a, 0x29, 0x3f, 0x17, 0x28,
	0x0b, 0x90, 0xbf, 0x09, 0xd0, 0x50, 0x71, 0x2d, 0x3d, 0x14, 0x5e, 0xc1, 0x20, 0x2e, 0x42, 0xb9,
	0x45, 0x94, 0x22, 0xaa, 0x77, 0xa7, 0x52, 0x44, 0xf9, 0xe2, 0x31, 0xd6, 0xae, 0xef, 0xd1, 0x27,
	0xfe, 0x87, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x69, 0xc3, 0x3e, 0xfc, 0x0b, 0x00, 0x00,
}
