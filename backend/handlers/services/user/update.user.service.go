package service

import (
	"fmt"
	storage "main/database"
	helper "main/helper/struct"
	"main/schema"
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
