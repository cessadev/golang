package internal

import (
	"go-api-plus/internal/handler"
	authmiddleware "go-api-plus/internal/middleware"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)    // Log de cada request
	r.Use(middleware.Recoverer) // Recupera de panics
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	// Rutas de usuario
	r.Route("/users", func(r chi.Router) {
		r.Use(authmiddleware.AuthMiddleware)
		r.Post("/", handler.CreateUser)
		r.Get("/{id}", handler.GetUser)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the advanced Go API!"))
	})

	return r
}
