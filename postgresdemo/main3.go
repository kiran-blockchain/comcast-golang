package main

import (
	"fmt"
	"log"
	"postgresdemo/dataaccess"


)

func main() {
	// Create a new instance of the data access layer
	dal := dataaccess.NewDataAccessLayer()

	// Get user by ID
	user, err := dal.GetUserByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User: %+v\n", user)

	// Get product by ID
	product, err := dal.GetProductByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Product: %+v\n", product)
}
