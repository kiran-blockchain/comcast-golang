package entities

import "gorm.io/gorm"


type Order struct{
	gorm.Model
	OrderID uint
	UserId uint
}