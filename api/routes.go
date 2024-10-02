package api

import (
	"practica/internal/auth"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	authRoutes := r.Group("/")

	// Ruta para autenticación con Firebase
	router.HandleFunc("/verify", auth.VerifyHandler).Methods("POST")

	authRoutes.POST("/register", func(c *gin.Context) {
		controllers.RegisterUser(c, firestoreClient)
	})

	return router
}
