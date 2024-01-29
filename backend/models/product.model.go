package models

type Product struct {
	ProductID   uint `json:"product_id"`
	ProductName string	`json:"product_name"`
	Description string	`json:"description"`
	Price       float64	`json:"price"`
	Quantity    uint	`json:"quantity"`
	ImagePath   string	`json:"image_path"`
}