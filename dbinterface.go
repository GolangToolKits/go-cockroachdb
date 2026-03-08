package gocockroachdb

// DbRow database row
type DbRow struct {
	Columns []string
	Row     []string
}

// DbRows array of database rows
type DbRows struct {
	Columns []string
	Rows    [][]string
}

// Database Database
type Database interface {
	Connect() bool
	// BeginTransaction() Transaction
	// Test(query string, args ...any) *DbRow
	// Insert(query string, args ...any) (bool, int64)
	// Update(query string, args ...any) bool
	// Get(query string, args ...any) *DbRow
	// GetList(query string, args ...any) *DbRows
	// Delete(query string, args ...any) bool
	Close() bool
}

// Transaction transaction
type Transaction interface {
	Insert(query string, args ...any) (bool, int64)
	Update(query string, args ...any) bool
	Delete(query string, args ...any) bool
	Commit() bool
	Rollback() bool
}

// go mod init github.com/GolangToolKits/go-cockroachdb
