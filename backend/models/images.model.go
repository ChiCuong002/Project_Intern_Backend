package models

type Image struct {
	ImageID   uint `json:"image_id"`
	ProductID uint	`json:"product_id"`
	Path      string	`json:"path"`
}