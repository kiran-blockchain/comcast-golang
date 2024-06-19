package main

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"gorm.io/gorm/logger"

    "time"
)

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string
    Email     string
    Profile   Profile
    Orders    []Order
    Groups    []*Group `gorm:"many2many:user_groups;"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Profile struct {
    ID     uint
    UserID uint
    Bio    string
    User   *User
}

type Order struct {
    ID       uint
    UserID   uint
    Product  string
    Quantity int
    User     *User
}

type Group struct {
    ID    uint
    Name  string
    Users []*User `gorm:"many2many:user_groups;"`
}

func main() {
    dsn := "host=SG-storm-bath-6462-5504-pgsql-master.servers.mongodirector.com user=ecom2 password=Ecom2#2024 dbname=ecom sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&User{}, &Profile{}, &Order{}, &Group{})

    // Create sample data
    user := User{Name: "John", Email: "john@example.com"}
    db.Create(&user)

    profile := Profile{Bio: "Golang Developer", UserID: user.ID}
    db.Create(&profile)

    order1 := Order{Product: "Laptop", Quantity: 1, UserID: user.ID}
    order2 := Order{Product: "Phone", Quantity: 2, UserID: user.ID}
    db.Create(&order1)
    db.Create(&order2)

    group := Group{Name: "Developers"}
    db.Create(&group)
    db.Model(&group).Association("Users").Append(&user)

    // Query with associations
    var queriedUser User
    db.Preload("Profile").Preload("Orders").Preload("Groups").First(&queriedUser, user.ID)
    fmt.Printf("User: %v\n", queriedUser)
    fmt.Printf("Profile: %v\n", queriedUser.Profile)
    fmt.Printf("Orders: %v\n", queriedUser.Orders)
    fmt.Printf("Groups: %v\n", queriedUser.Groups)

	if err := db.Migrator().DropTable(&User{}); err != nil {
		fmt.Println("Error dropping table:", err)
	} else {
		fmt.Println("User table dropped successfully")
	}
	if err := db.Migrator().DropTable(&Order{}); err != nil {
		fmt.Println("Error dropping table:", err)
	} else {
		fmt.Println("Order table dropped successfully")
	}
	if err := db.Migrator().DropTable(&Group{}); err != nil {
		fmt.Println("Error dropping table:", err)
	} else {
		fmt.Println("Group table dropped successfully")
	}
	if err := db.Migrator().DropTable(&Profile{}); err != nil {
		fmt.Println("Error dropping table:", err)
	} else {
		fmt.Println("Profile table dropped successfully")
	}
}
