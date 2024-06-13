package dataaccess

import (
	//"database/sql"
	//"log"

	"context"
	db "postgresdemo/db"

	"github.com/jackc/pgx/v4/pgxpool"
)

// DataAccessLayer represents the data access layer.
type DataAccessLayer struct {
	//db *sql.DB
	db  *pgxpool.Pool
}
// User represents a user entity.
type User struct {
	UserID   int
	Username string
	Email    string
}

// Product represents a product entity.
type Product struct {
	ProductID   int
	ProductName string
	Price       float64
}

// NewDataAccessLayer creates a new instance of the data access layer.
func NewDataAccessLayer() *DataAccessLayer {
	return &DataAccessLayer{db: db.GetDBConnectionPool()}
}

// GetUserByID retrieves a user by ID.
func (dal *DataAccessLayer) GetUserByID(userID int) (User, error) {
	var user User
	err := dal.db.QueryRow(context.Background(),"SELECT UserID, Username, Email FROM Users WHERE UserID = $1", userID).Scan(&user.UserID, &user.Username, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetProductByID retrieves a product by ID.
func (dal *DataAccessLayer) GetProductByID(productID int) (Product, error) {
	var product Product
	err := dal.db.QueryRow(context.Background(),"SELECT ProductID, ProductName, Price FROM Products WHERE ProductID = $1", productID).Scan(&product.ProductID, &product.ProductName, &product.Price)
	if err != nil {
		return product, err
	}
	return product, nil
}
