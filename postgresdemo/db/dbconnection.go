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
		connStr := "host=SG-storm-bath-6462-5504-pgsql-master.servers.mongodirector.com user=ecom2 password=Ecom2#2024 dbname=ecom sslmode=disable"
		//connStr:="postgresql://ecom2:Ecom2#2024@SG-storm-bath-6462-5504-pgsql-master.servers.mongodirector.com:5432/ecom"
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


//postgresql://SG-storm-bath-6462-5504-pgsql-master.servers.mongodirector.com:5432/ecom