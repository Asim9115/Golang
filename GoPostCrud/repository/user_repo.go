package repository

import (
		"github.com/asim9115/GoLang/config"
	"github.com/asim9115/GoLang/models"
	"errors"

)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

func GetUserById(id string) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	return &user, result.Error
}

func UpdateUser(user *models.User, id string) (error) {
	
	
	return config.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}


func DeleteUser(id string) (error) {
	result := config.DB.Delete(&models.User{}, id)
		if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil

}
