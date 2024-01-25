package service

import (
	//"fmt"
	//storage "main/database"

	"fmt"
	storage "main/database"
	helper "main/helper/struct/product"
	productHelper "main/helper/struct/product"
	"main/schema"

	"gorm.io/gorm"
)

func InsertProduct(tx *gorm.DB, product *helper.ProductInsert) error {
	result := tx.Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func InsertImage(tx *gorm.DB, image *helper.ImageInsert) error {
	result := tx.Create(image)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func InsertProductImage(tx *gorm.DB, productImg *helper.ProductImageInsert) error {
	result := tx.Create(productImg)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func UpdateProduct(product *helper.UpdateProduct) error {
	db := storage.GetDB()
	result := db.First(&schema.Product{}, product.ProductID)
	if result.Error != nil {
		return fmt.Errorf("Invalid product id")
	}
	result = db.First(&schema.Category{}, product.CategoryID)
	if result.Error != nil {
		return fmt.Errorf("Invalid category id")
	}
	result = db.Model(&schema.Product{}).Where("product_id = ?", product.ProductID).Updates(map[string]interface{}{
		"category_id":  product.CategoryID,
		"product_name": product.ProductName,
		"description":  product.Description,
		"price":        product.Price,
		"quantity":     product.Quantity})
	if result.Error != nil {
		return fmt.Errorf(result.Error.Error())
	}
	return nil
}
func DetailProduct(id uint) (productHelper.DetailProductRes, error) {
	db := storage.GetDB()
	product := productHelper.DetailProductRes{}
	err := db.Model(&productHelper.DetailProductRes{}).Preload("ProductImages.Image").First(&product, id).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
