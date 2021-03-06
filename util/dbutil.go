package util

import (
	"database/sql"
	"database/sql/driver"
	"net"
	"strings"

	"github.com/lib/pq"
)

// SQLExecutor is a wrapper around sql.Tx, sql.Db, and others we implement. It represents
// something that can execute queries on a database.
type SQLExecutor interface {
	// Exec is meant for queries that perform an action on a database. They may return
	// a success code, but shouldn't return any rows. This includes INSERT, CREATE, DROP,
	// etc.
	Exec(query string, args ...interface{}) (sql.Result, error)

	// Query is meant for queries that return multiple rows. The *sql.Rows returned
	// can be used to iterate over the results, and must be closed when it is done
	// being used.
	Query(query string, args ...interface{}) (*sql.Rows, error)

	// QueryRow is meant for queries that only expect one row, such as specifically
	// crafted SELECT statements.
	QueryRow(query string, args ...interface{}) *sql.Row
}

// AtomicSQLExecutor is a SQLExecutor that can be committed or rolled back. This
// wraps sql.Tx and others we implement
type AtomicSQLExecutor interface {
	SQLExecutor

	// Since this interface is meant for use with transactions and other structs
	// that wrap transactions, Commit and Rollback are within the context of that
	// underlying transaction.
	Commit() error
	Rollback() error
}

// PrefixesToPQArray handles a strange case where protobuf deserialize an array element of nil as "<nil>"
// and that kills the db insert statement cause it can't make it into a cidr. The return
// value is an inline interface which matches the return value of pq.Arrary.
func PrefixesToPQArray(n []*net.IPNet) interface {
	driver.Valuer
	sql.Scanner
} {
	if n == nil || len(n) == 0 {
		return nil // database will accept NULL on this field
	}

	ret := make([]string, len(n))
	for i := range n {
		ret[i] = n[i].String()
		if ret[i] == "" || ret[i] == "<nil>" {
			// This is a sane default value for an IPNet that also shows there was a parse error.
			ret[i] = "0.0.0.0/0"
		}
	}
	return pq.Array(ret)
}

// SanitizeDBString removes any characters from s that might be intended
// for a SQL injection attack.
func SanitizeDBString(s string) string {
	bad := []string{"'", ";", "\"", "DROP", "TRUNCATE", "DELETE", " "}
	ret := s
	problem := true

	// This outer loop checks for a nested problem strings.
	// Ex. DDROPROP
	// With only one iteration, the interior drop would be
	// removed, leaving an outer drop. This way, if the string
	// changes between iterations, it will keep going.
	for problem {
		prevRet := ret
		// This loop removes all instances of bad strings
		for _, v := range bad {
			ret = strings.Replace(ret, v, "", -1)
		}

		if ret == prevRet {
			problem = false
		}
	}

	return ret
}
