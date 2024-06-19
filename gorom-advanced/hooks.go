package main

import (
    "fmt"
    "time"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string
    Email     string
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate hook will be triggered before a User record is created
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    fmt.Println("BeforeCreate: Setting default values")
    if u.CreatedAt.IsZero() {
        u.CreatedAt = time.Now()
    }
    return
}

// AfterCreate hook will be triggered after a User record is created
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
    fmt.Println("AfterCreate: User has been created")
    return
}

// BeforeUpdate hook will be triggered before a User record is updated
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
    fmt.Println("BeforeUpdate: Updating UpdatedAt timestamp")
    u.UpdatedAt = time.Now()
    return
}

// AfterUpdate hook will be triggered after a User record is updated
func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
    fmt.Println("AfterUpdate: User has been updated")
    return
}

func main() {
    dsn := "host=SG-storm-bath-6462-5504-pgsql-master.servers.mongodirector.com user=ecom2 password=Ecom2#2024 dbname=ecom sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&User{})

    // Create a new user
    user := User{Name: "John", Email: "john@example.com"}
    db.Create(&user)

    // Update the user's name
    user.Name = "John Doe"
    db.Save(&user)

	if err := db.Migrator().DropTable(&User{}); err != nil {
        fmt.Println("Error dropping table:", err)
    } else {
        fmt.Println("User table dropped successfully")
    }
}
