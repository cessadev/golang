package repository

import (
	dbconfig "go-api-plus/config"
	"go-api-plus/internal/domain"
)

// Método de repositorio para crear un usuario
func CreateUser(user *domain.User) error {
	return dbconfig.DB.Create(user).Error
}

// Método de repositorio para consultar un usuario a través de su ID
func GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := dbconfig.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
