package entities

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User represents a user entity.
type Profile struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
}
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
}

type ProfileService struct{
	db  *gorm.DB
}

func NewProfileService(connStr string) *ProfileService{
	db,err:= gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if(err!=nil){
		log.Fatalf("Error in connecting to database %v",err)
	}
	db.AutoMigrate(&Profile{})
	return &ProfileService{db:db}
}

func (ps *ProfileService) Create(profile *Profile) error{
	return ps.db.Create(profile).Error
}

func (ps *ProfileService) Update(id uint,newEmail string) error{
	return ps.db.Model(&Profile{}).Where("id = ?",id).Update("email",newEmail).Error
}

func (ps *ProfileService) GetOrdersByUserId() error {
	var result [] struct{
		Userid uint
		Username string
		Email string
		Orderid uint
	}
	ps.db.Table("users").Select("users.userId,users.username,orders.orderId").Joins("INNER JOIN orders on orders.userid = users.userid").Scan(&result)
	
	for _,row := range result{
		fmt.Printf("UserId : %d, Name : %s orderId: %d",row.Userid,row.Username,row.Orderid)
	}
	
	return nil
}