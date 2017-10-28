package txdb

import (
	"fmt"
	"sync"

	"github.com/lytics/flo/internal/txdb/driver"
	"github.com/lytics/flo/window"
)

var (
	driversMu = sync.Mutex{}
	drivers   = map[string]driver.Driver{}
)

// Register a database driver.
func Register(name string, d driver.Driver) {
	driversMu.Lock()
	defer driversMu.Unlock()

	_, ok := drivers[name]
	if ok {
		panic("txdb: driver registered twice: " + name)
	}
	drivers[name] = d
}

// Open a new database connection.
func Open(driverName, sourceName string) (*DB, error) {
	drvr, ok := drivers[driverName]
	if !ok {
		return nil, fmt.Errorf("txdb: unknown driver: %v", driverName)
	}
	conn, err := drvr.Open(sourceName)
	if err != nil {
		return nil, fmt.Errorf("txdb: failed to open connection: %v", err)
	}
	return &DB{
		conn: conn,
	}, nil
}

// DB handle for interaction with a database.
type DB struct {
	conn driver.Conn
}

// Apply the mutation to the graph key's row.
func (db *DB) Apply(key string, mutation func(window.State) error) error {
	return db.conn.Apply(key, mutation)
}

// Drain the keys into the sink.
func (db *DB) Drain(keys []string, sink func(span window.Span, key string, vs []interface{}) error) {
	// for key, row := range m.snapshot(keys) {
	// 	for span, vs := range row.Windows {
	// 		sink(span, key, vs)
	// 	}
	// }
}
