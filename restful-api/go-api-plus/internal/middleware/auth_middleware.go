package authmiddleware

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")

		if apiKey != "my-secret-key" { // Puedes reemplazar "my-secret-key" por tu clave real
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r) // Llama al siguiente handler si la autenticación es correcta
	})
}
