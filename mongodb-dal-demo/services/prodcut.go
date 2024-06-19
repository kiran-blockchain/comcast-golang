package services

import (
	"context"
	"fmt"
	// "errors"
	// "fmt"
	"mongodemo/entities"
	"mongodemo/interfaces"

	// "strings"
	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	ProductCollection *mongo.Collection // database table
	Ctx               context.Context   // database connection
}

func ProductServiceInit(productCollection *mongo.Collection, ctx context.Context) interfaces.IProduct {
	return &ProductService{ProductCollection: productCollection, Ctx: ctx}
}
func (p *ProductService) Create(user *entities.Product) (*entities.Product, error) {
if(user!=nil){
	fmt.Println("invalid user")
}
	return nil, nil
}
