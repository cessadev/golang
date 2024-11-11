package handler

import (
	"encoding/json"
	"go-api/internal/domain"
	"go-api/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type UserHandler struct {
	service *service.UserService
}

// Constructor
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Función (controlador) para obtener un usuario por su ID
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Función (controlador) para guardar un usuario en el mapa (base de datos)
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
