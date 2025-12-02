package utils

import (
	"src/golang_mssql/config"
	"src/golang_mssql/dto"
	"src/golang_mssql/models"
)

func GetByUsername(username string) ([]models.User, error) {
	var users []models.User

	db := config.Connection()
	result := db.Where("username = ?", username).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetByUserId(id string) ([]dto.Users, error) {
	var users []dto.Users

	db := config.Connection()
	result := db.Where("id = ?", id).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
