package api

import (
	"practica/internal/auth"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()


	// Configurar CORS
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"}, // Cambia el puerto si es necesario
        AllowMethods:     []string{ }, // Métodos permitidos
        AllowHeaders:     []string{"Content-Type"},
        AllowCredentials: true,
    }))


	router.POST("/register", auth.RegisterHandler)
	router.POST("/login", auth.LoginHandler)
	router.GET("/verify-email", auth.VerifyEmailHandler)
	router.POST("/password-reset", auth.SendPasswordResetEmailHandler)
	router.POST("/resend-verification", auth.ResendVerificationEmailHandler)

	// Rutas protegidas
	protected := router.Group("/").Use(auth.AuthMiddleware) // Agrupar las rutas protegidas con el middleware
	{
		protected.POST("/complete-profile", auth.CompleteProfileHandler) // Ruta para completar perfil
		// Otras rutas protegidas aquí
	}

	return router
}
