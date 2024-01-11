package services

import (
	"errors"
	storage "main/database"
	"main/helper/scope"
	helper "main/helper/struct"
	"main/models"

	"gorm.io/gorm"
	//"github.com/labstack/echo/v4"
)

const (
	ADMIN = 1
	USER  = 2
)

func GetAllUser() ([]models.User, error) {
	var db *gorm.DB = storage.GetDB()
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		// Default behavior with QueryFields set to true
		//SELECT * FROM users
		return nil, err
	}

	return users, nil
}
func GetAllUserPagination(pagination helper.Pagination) (*helper.Pagination, error) {
	var db *gorm.DB = storage.GetDB()
	var users []models.User
	query := db.Scopes(scope.Paginate(users, &pagination, db)) //.Where("role_id <> ?", ADMIN).Find(&users)
	if pagination.Search != "" {
		query.Where("first_name LIKE ? OR last_name LIKE ?", "%"+pagination.Search+"%", "%"+pagination.Search+"%")
	}
	query.Find(&users)
	pagination.Rows = users
	return &pagination, nil
}
func UserDetail(id uint) (models.User, error) {
	var db *gorm.DB = storage.GetDB()
	user := models.User{}
	err := db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//SELECT * FROM users WHERE id = 10;
		return user, err
	}
	return user, nil
}
