package image

import (
	storage "main/database"
	helper "main/helper/struct/product"

	"gorm.io/gorm"
)

func GetImageByID(id uint) (helper.ImageInsert, error) {
	db := storage.GetDB()
	image := helper.ImageInsert{}
	result := db.First(&image, id)
	if result.Error != nil {
		return image, result.Error
	}
	return image, nil
}
func InsertImage(tx *gorm.DB, image *helper.ImageInsert) error {
	result := tx.Create(image)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func UpdateImage(tx *gorm.DB, image *helper.ImageInsert) error {
	result := tx.Save(image)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
