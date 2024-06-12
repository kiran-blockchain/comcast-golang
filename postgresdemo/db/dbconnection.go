package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	once sync.Once
	db   *sql.DB
)

// GetDBConnection returns a singleton instance of the database connection.
func GetDBConnection() *sql.DB {
	once.Do(func() {
		// Database connection string
		connStr := "user=ecom1 password=ecom1 dbname=ecom sslmode=disable"

		// Open the connection
		var err error
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
	})

	// Verify the connection
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
