package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Database connection string
	connStr := "host=SG-storm-bath-6462-5504-pgsql-master.servers.mongodirector.com user=ecom2 password=Ecom2#2024 dbname=ecom sslmode=disable"
    //connStr:="postgresql://ecom2:Ecom2#2024@SG-storm-bath-6462-5504-pgsql-master.servers.mongodirector.com:5432/ecom"
	// Open the connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")

	// Query users
	// query := "SELECT UserID, Username, Email FROM Users"
	// rows, err := db.Query(query)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// fmt.Println("Users:")
	// for rows.Next() {
	// 	var userID int
	// 	var username, email string
	// 	err = rows.Scan(&userID, &username, &email)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("UserID: %d, Username: %s, Email: %s\n", userID, username, email)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Query products
	// query = "SELECT ProductID, ProductName, Price FROM Products"
	// rows, err = db.Query(query)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// fmt.Println("Products:")
	// for rows.Next() {
	// 	var productID int
	// 	var productName string
	// 	var price float64
	// 	err = rows.Scan(&productID, &productName, &price)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("ProductID: %d, ProductName: %s, Price: %.2f\n", productID, productName, price)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
