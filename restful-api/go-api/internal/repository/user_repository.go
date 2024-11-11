package repository

import (
	"errors"
	"go-api/internal/domain"
)

type UserRepository struct {
	user map[int]domain.User // Simulación de una base de datos en memoria
}

// Constructor
func NewUserRepository() *UserRepository {
	return &UserRepository{
		user: make(map[int]domain.User),
	}
}

// Función que recupera un usuario por ID. Devuelve un error si no existe.
func (r *UserRepository) GetUserByID(id int) (domain.User, error) {
	user, exists := r.user[id]
	if !exists {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

// Función que guarda un nuevo usuario en el mapa
func (r *UserRepository) SaveUser(user domain.User) error {
	r.user[user.ID] = user
	return nil
}
