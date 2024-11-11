package service

import (
	"go-api/internal/domain"
	"go-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

// Constructor
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Función que retorna un usuario a través de su ID.
func (s *UserService) GetUser(id int) (domain.User, error) {
	return s.repo.GetUserByID(id)
}

// Función que guarda un usuario en el mapa (base de datos)
func (s *UserService) CreateUser(user domain.User) error {
	return s.repo.SaveUser(user)
}
