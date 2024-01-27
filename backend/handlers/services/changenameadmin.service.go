package services

import (
	"main/schema"

	"gorm.io/gorm"
)

func ChangeNameCategory(c *schema.Category, db *gorm.DB, newName string) error {
	result := db.Model(c).Update("category_name", newName)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
