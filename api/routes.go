package api

import (
	"practica/internal/auth"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Ruta para autenticación con Firebase
	router.HandleFunc("/verify", auth.VerifyHandler).Methods("POST")

	return router
}
