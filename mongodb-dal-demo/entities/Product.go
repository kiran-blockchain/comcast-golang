package entities


// 👈 Product struct
type Product struct {
	Name               string    `json:"name" bson:"name" binding:"required"`
	Price              uint       `json:"price" bson:"price" binding: "required"`
	ImageUrl           string     `json:"imageurl" bson:"imageurl" binding: "required"`
	Description           string     `json:"description" bson:"description" binding: "required"`
}
