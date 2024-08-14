package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib" // Import the pgx driver
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

// Database connection pool settings
const (
	maxOpenDbConn = 10
	maxIdleDbConn = 5
	maxDbLifetime = 5 * time.Minute
)

// ConnectSQL creates a database pool for Postgres
func ConnectSQL(dsn string) (*DB, error) {
	// Create a new database connection
	d, err := NewDatabase(dsn)
	if err != nil {
		return nil, err
	}

	// Set connection pool parameters
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = d

	// Test the database connection
	err = testDB(d)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

// testDB tries to ping the database
func testDB(d *sql.DB) error {
	return d.Ping()
}

// NewDatabase creates a new database connection for the application
func NewDatabase(dsn string) (*sql.DB, error) {
	// Open a database connection using the pgx driver
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Ping the database to ensure the connection is established
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
