package service

import (
	"go-api-plus/internal/domain"
	"go-api-plus/internal/repository"
)

// Método de servicio para crear un usuario
func CreateUser(user *domain.User) error {
	return repository.CreateUser(user)
}

// Método de servicio para consultar un usuario a través de su ID
func GetUserByID(id int) (*domain.User, error) {
	return repository.GetUserByID(uint(id))
}
