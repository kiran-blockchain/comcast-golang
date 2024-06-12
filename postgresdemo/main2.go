package main

import (

	"fmt"
	"log"
	"postgresdemo/db"
)

func main() {
	// Get database connection
	dbConn := db.GetDBConnection()
	defer dbConn.Close()

	// Query users
	rows, err := dbConn.Query("SELECT UserID, Username, Email FROM Users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var userID int
		var username, email string
		err := rows.Scan(&userID, &username, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("UserID: %d, Username: %s, Email: %s\n", userID, username, email)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Query products
	rows, err = dbConn.Query("SELECT ProductID, ProductName, Price FROM Products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Products:")
	for rows.Next() {
		var productID int
		var productName string
		var price float64
		err := rows.Scan(&productID, &productName, &price)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ProductID: %d, ProductName: %s, Price: %.2f\n", productID, productName, price)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
