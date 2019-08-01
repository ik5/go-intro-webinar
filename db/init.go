package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
	"time"

	// Loading postgresql driver into local package
	_ "github.com/lib/pq"
)

// Conn os the database connection
type Conn struct {
	db         *sql.DB
	ctx        context.Context
	cancelFunc context.CancelFunc
}

// PGSSLFields passing information for
type PGSSLFields struct {
	SSLMode     string `ssl:"sslmode"`
	SSLCert     string `ssl:"sslcert"`
	SSLKey      string `ssl:"sslkey"`
	SSLRootCert string `ssl:"sslrootcert"`
}

// String returns the fields settings for using them as connections
func (sslFields *PGSSLFields) String() string {
	structElements := reflect.ValueOf(sslFields).Elem()
	fieldsLen := structElements.NumField()

	// generate a tmp slice size 0, initialize it memory capable of out of
	// n bytes to caps
	tmp := make([]string, 0, fieldsLen)

	// Go over all fields of the struct
	for i := 0; i < fieldsLen; i++ {
		fieldValue := structElements.Field(i)       // get current field value
		fieldType := structElements.Type().Field(i) // get the field type
		tag := fieldType.Tag                        // get the field tag

		// cast interface{} to string can you guess what might go be wrong here?
		value := fieldValue.Interface().(string)
		if value != "" {
			tmpStr := fmt.Sprintf("%s=%s", tag.Get("ssl"), value)
			tmp = append(tmp, tmpStr) // add the content to the slice
		}
	}

	return strings.Join(tmp, " ") // return a space separated fields
}

// Open creates a new database connection, and if fails, return an error instead
func Open(address, db, username, password string, port int, sslFields *PGSSLFields) (conn *Conn, err error) {
	sslString := ""
	if sslFields != nil {
		sslString = sslFields.String()
	}
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s %s",
		username, password, address, port, db, sslString,
	)
	conn = &Conn{}
	conn.db, err = sql.Open("postgres", connStr)
	if err != nil {
		conn = nil
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	conn.ctx = ctx
	conn.cancelFunc = cancelFunc
	return
}

// Begin  starts a transaction.
//
// The provided context is used until the transaction is committed or rolled
// back. If the context is canceled, the sql package will roll back the
// transaction. Tx.Commit will return an error if the context provided to
// Begin is canceled.
//
// The provided TxOptions is optional and may be nil if defaults should be used.
// If a non-default isolation level is used that the driver doesn't support, an
// error will be returned.
func (conn *Conn) Begin(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	if ctx != nil {
		return conn.db.BeginTx(ctx, opts)
	}
	return conn.db.Begin()
}

// Close closes the database and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server
// to finish.
//
// It is rare to Close a DB, as the DB handle is meant to be long-lived and
// shared between many goroutines.
func (conn *Conn) Close() error {
	conn.cancelFunc()
	return conn.db.Close()
}

// Driver returns the database's underlying driver.
func (conn *Conn) Driver() driver.Driver {
	return conn.db.Driver()
}

// Exec executes a query without returning any rows. The args are for any
// placeholder parameters in the query.
func (conn *Conn) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if ctx != nil {
		return conn.db.ExecContext(ctx, query, args...)
	}
	return conn.db.Exec(query, args...)
}

// Ping verifies a connection to the database is still alive, establishing a
// connection if necessary.
func (conn *Conn) Ping(ctx context.Context) error {
	if ctx != nil {
		return conn.db.PingContext(ctx)
	}
	return conn.db.Ping()
}

// Prepare creates a prepared statement for later queries or executions.
// Multiple queries or executions may be run concurrently from the returned
// statement. The caller must call the statement's Close method when the
// statement is no longer needed.
//
// The provided context is used for the preparation of the statement, not for
// the execution of the statement.
func (conn *Conn) Prepare(ctx context.Context, query string) (*sql.Stmt, error) {
	if ctx != nil {
		return conn.db.PrepareContext(ctx, query)
	}
	return conn.db.Prepare(query)
}

// Query executes a query that returns rows, typically a SELECT. The args are
// for any placeholder parameters in the query
func (conn *Conn) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if ctx != nil {
		return conn.db.QueryContext(ctx, query, args...)
	}
	return conn.db.Query(query, args...)
}

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always returns a non-nil value. Errors are deferred until Row's Scan
// method is called. If the query selects no rows, the *Row's Scan will return
// ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and
// discards the rest.
func (conn *Conn) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if ctx != nil {
		return conn.db.QueryRowContext(ctx, query, args...)
	}
	return conn.db.QueryRow(query, args...)
}

// SetConnMaxLifetime sets the maximum amount of time a connection may be
// reused.
//
// Expired connections may be closed lazily before reuse.
//
// If d <= 0, connections are reused forever.
func (conn *Conn) SetConnMaxLifetime(d time.Duration) {
	conn.db.SetConnMaxLifetime(d)
}

// SetMaxIdleConns sets the maximum number of connections in the idle
// connection pool.
//
// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns, then
// the new MaxIdleConns will be reduced to match the MaxOpenConns limit.
//
// If n <= 0, no idle connections are retained.
//
// The default max idle connections is currently 2. This may change in a future
//release.
func (conn *Conn) SetMaxIdleConns(n int) {
	conn.db.SetMaxIdleConns(n)
}

// SetMaxOpenConns sets the maximum number of open connections to the database.
//
// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
// MaxIdleConns, then MaxIdleConns will be reduced to match the new MaxOpenConns
// limit.
//
// If n <= 0, then there is no limit on the number of open connections.
// The default is 0 (unlimited).
func (conn *Conn) SetMaxOpenConns(n int) {
	conn.db.SetMaxOpenConns(n)
}

// Stats returns database statistics.
func (conn *Conn) Stats() sql.DBStats {
	return conn.db.Stats()
}
