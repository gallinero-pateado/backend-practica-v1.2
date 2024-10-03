package api

import (
	"practica/internal/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/register", auth.RegisterHandler)
	router.POST("/login", auth.LoginHandler)
	router.GET("/verify-email", auth.VerifyEmailHandler)
	router.POST("/password-reset", auth.SendPasswordResetEmailHandler)

	// Rutas protegidas
	protected := router.Group("/").Use(auth.AuthMiddleware)  // Agrupar las rutas protegidas con el middleware
	{
		protected.POST("/complete-profile", auth.CompleteProfileHandler)  // Ruta para completar perfil
		// Otras rutas protegidas aquí
	}


	return router
}
