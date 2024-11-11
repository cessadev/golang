package main

import (
	"go-api/internal/handler"
	"go-api/internal/repository"
	"go-api/internal/service"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	// Inicializar repositorio, servicio y handler
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Configuración del router
	r := chi.NewRouter()
	r.Get("/users/{id}", userHandler.GetUser)       // Endpoint para obtener un usuario
	r.Post("/users/create", userHandler.CreateUser) // Endpoint para guardar un usuario

	// Iniciar el servidor en el puerto 8080
	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
