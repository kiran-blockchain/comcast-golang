package db

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

var (
	once sync.Once
	db   *sql.DB
	poolInstance *pgxpool.Pool
	poolOnce sync.Once
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
func GetDBConnectionPool() *pgxpool.Pool{
	poolOnce.Do(func() {
		connStr := "host=SG-storm-bath-6462-5504-pgsql-master.servers.mongodirector.com user=ecom2 password=Ecom2#2024 dbname=ecom sslmode=disable"
		config,err := pgxpool.ParseConfig(connStr)
		if(err!=nil){
			log.Fatal("Unable to parse the connection string",err)
		}
		config.MaxConns=100
		config.MinConns=2
		config.MaxConnLifetime =30*time.Minute
		config.MaxConnIdleTime =30*time.Minute
		config.HealthCheckPeriod = 5* time.Minute
		pool,err := pgxpool.ConnectConfig(context.Background(),config)
		if(err!=nil){
			log.Fatalf("unable to connect to db%v\n",err)
		}
		poolInstance=pool
	})
	return poolInstance

}


//postgresql://SG-storm-bath-6462-5504-pgsql-master.servers.mongodirector.com:5432/ecom