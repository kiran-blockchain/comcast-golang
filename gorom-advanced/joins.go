package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID     uint
	Name   string
	Orders []Order
}

type Order struct {
	ID        uint
	UserID    uint
	ProductID uint
	Quantity  int
	User      User
	Product   Product
}

type Product struct {
	ID     uint
	Name   string
	Price  float64
	Orders []Order
}

func main() {
	//dsn := "host=localhost user=youruser password=yourpassword dbname=gorm_demo port=5432 sslmode=disable"
	ConnStr := "host=SG-storm-bath-6462-5504-pgsql-master.servers.mongodirector.com user=ecom2 password=Ecom2#2024 dbname=ecom sslmode=disable"
	db, err := gorm.Open(postgres.Open(ConnStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &Order{}, &Product{})

	// Insert sample data
	db.Create(&User{Name: "John"})
	db.Create(&User{Name: "Jane"})
	db.Create(&Product{Name: "Laptop", Price: 1000.0})
	db.Create(&Product{Name: "Phone", Price: 500.0})
	db.Create(&Order{UserID: 1, ProductID: 1, Quantity: 2})
	db.Create(&Order{UserID: 1, ProductID: 2, Quantity: 1})
	db.Create(&Order{UserID: 2, ProductID: 1, Quantity: 1})

	// Inner Join
	var results []struct {
		UserName    string
		ProductName string
		Quantity    int
	}
	db.Table("users").Select("users.name as user_name, products.name as product_name, orders.quantity").
		Joins("inner join orders on orders.user_id = users.id").
		Joins("inner join products on products.id = orders.product_id").
		Scan(&results)

	fmt.Println("Inner Join Results:")
	for _, result := range results {
		fmt.Printf("User: %s, Product: %s, Quantity: %d\n", result.UserName, result.ProductName, result.Quantity)
	}

	// Left Join
	db.Table("users").Select("users.name as user_name, products.name as product_name, orders.quantity").
		Joins("left join orders on orders.user_id = users.id").
		Joins("left join products on products.id = orders.product_id").
		Scan(&results)

	fmt.Println("Left Join Results:")
	for _, result := range results {
		fmt.Printf("User: %s, Product: %s, Quantity: %d\n", result.UserName, result.ProductName, result.Quantity)
	}

	// Right Join
	db.Table("users").Select("users.name as user_name, products.name as product_name, orders.quantity").
		Joins("right join orders on orders.user_id = users.id").
		Joins("right join products on products.id = orders.product_id").
		Scan(&results)

	fmt.Println("Right Join Results:")
	for _, result := range results {
		fmt.Printf("User: %s, Product: %s, Quantity: %d\n", result.UserName, result.ProductName, result.Quantity)
	}

	 // Drop the table
	 if err := db.Migrator().DropTable(&User{}); err != nil {
        fmt.Println("Error dropping table:", err)
    } else {
        fmt.Println("User table dropped successfully")
    }
	if err := db.Migrator().DropTable(&Order{}); err != nil {
        fmt.Println("Error dropping table:", err)
    } else {
        fmt.Println("User table dropped successfully")
    }
	if err := db.Migrator().DropTable(&Product{}); err != nil {
        fmt.Println("Error dropping table:", err)
    } else {
        fmt.Println("User table dropped successfully")
    }

}
