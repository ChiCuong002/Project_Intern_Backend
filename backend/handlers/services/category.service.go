package services

import (
	"errors"
	"fmt"
	storage "main/database"
	"main/helper/scope"
	helper "main/helper/struct"
	"main/schema"

	"gorm.io/gorm"
)

func SearchCategories(query *gorm.DB, search string) *gorm.DB {
	if search != "" {
		query = query.Where("category_name like ?", "%"+search+"%")
	}
	return query

}
func GetAllCategories(pagination helper.Pagination) (*helper.Pagination, error) {
	db := storage.GetDB()
	var categories []helper.Categories
	query := db.Model(&categories)
	query = SearchCategories(query, pagination.Search)
	query = query.Scopes(scope.Paginate(query, &pagination))
	query.Find(&categories)
	pagination.Rows = categories
	return &pagination, nil
}
func GetDetailCategory(id uint) (schema.Category, error) {
	var db *gorm.DB = storage.GetDB()
	category := schema.Category{}
	err := db.First(&category, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return category, err
	}
	return category, nil
}
func CategoriesDropDown() ([]helper.CategoriesDropDown, error) {
	db := storage.GetDB()
	category := []helper.CategoriesDropDown{}
	result := db.Select("category_id, category_name").Find(&category, schema.Category{IsActive: true})
	if result.Error != nil {
		return category, result.Error
	}
	if result.RowsAffected == 0 {
		return category, fmt.Errorf("Can't found any categories")
	}
	return category, nil
}
