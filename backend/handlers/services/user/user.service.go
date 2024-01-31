package service

import (
	"fmt"
	storage "main/database"
	"main/helper/scope"
	helper "main/helper/struct"
	"main/models"
	"main/schema"

	"gorm.io/gorm"
)

const ZERO_VALUE_INT = 0

func UpdateUser(tx *gorm.DB, userData *helper.UserInsert) error {
	fmt.Println("service")
	updateData := map[string]interface{}{
		"first_name": userData.FirstName,
		"last_name":  userData.LastName,
		"address":    userData.Address,
		"email":      userData.Email,
		"image_id": userData.Image,
	}
	switch {
	case userData.FirstName != "":
		updateData["first_name"] = userData.FirstName
	case userData.LastName != "":
		updateData["last_name"] = userData.LastName
	case userData.Address != "":
		updateData["address"] = userData.Address
	case userData.Email != "":
		updateData["email"] = userData.Email
	case userData.Image != ZERO_VALUE_INT:
		updateData["image_id"] = userData.Image
	}
	fmt.Println("updateData: ", updateData)
	result := tx.Model(&schema.User{}).Where("user_id = ?", userData.UserID).Updates(updateData)
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
	fmt.Println("id: ", id)
	var user helper.UserResponse
	result := db.Model(&user).Preload("Image").First(&user, "user_id = ?", id)
	if result.Error != nil {
		//SELECT * FROM users WHERE id = 10;
		return user, fmt.Errorf("Failed to get user")
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
