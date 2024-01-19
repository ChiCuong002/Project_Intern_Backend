package service

import (
	//"fmt"
	//storage "main/database"
	"main/schema"

	"gorm.io/gorm"
)

func InsertProduct(tx *gorm.DB,product *schema.Product) error {
	result := tx.Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func InsertImage(tx *gorm.DB, image *schema.Image) error {
	result := tx.Create(image)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
