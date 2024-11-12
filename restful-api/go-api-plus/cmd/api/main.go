package main

import (
	dbconfig "go-api-plus/config"
	"go-api-plus/internal"
	"log"
	"net/http"
)

func main() {
	// Inicializa la configuración (p. ej., conexión a la base de datos)
	dbconfig.InitDB()

	// Configura el router
	r := internal.SetupRouter()

	// Inicia el servidor
	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
