package service

import (
	"errors"
	"fmt"
	storage "main/database"
	"main/helper/scope"
	helper "main/helper/struct"
	"main/models"
	"main/schema"

	"gorm.io/gorm"
)

func UpdateUser(userData *helper.UpdateData) error {
	fmt.Println("service")
	db := storage.GetDB()
	result := db.Model(&schema.User{}).Where("user_id = ?", userData.UserID).Updates(
		map[string]interface{}{"first_name": userData.FirstName, "last_name": userData.LastName,
			"address": userData.Address, "email": userData.Email})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
const (
	ADMIN = 1
	USER  = 2
)

func Search(scope *gorm.DB, search string) *gorm.DB {
	if search != "" {
		scope = scope.Where("first_name LIKE ? OR last_name LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	return scope
}
func GetAllUserPagination(pagination helper.Pagination) (*helper.Pagination, error) {
	var db *gorm.DB = storage.GetDB()
	var users []models.User
	query := db.Model(&users)
	query = Search(query, pagination.Search)
	query = query.Scopes(scope.Paginate(query, &pagination))
	query.Find(&users)
	if query.RowsAffected == 0 {
		pagination.TotalPages = 1
	}
	pagination.Rows = users
	return &pagination, nil
}
func UserDetail(id uint) (helper.UserResponse, error) {
	var db *gorm.DB = storage.GetDB()
	user := helper.UserResponse{}
	err := db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//SELECT * FROM users WHERE id = 10;
		return user, err
	}
	return user, nil
}
func BlockUser(id uint) (helper.UserResponse, error) {
	var db *gorm.DB = storage.GetDB()
	user, err := UserDetail(id)
	if err != nil {
		return user, err
	}
	db.Model(&user).Where("user_id = ?", id).Update("is_active", !user.IsActive)
	return user, nil
}