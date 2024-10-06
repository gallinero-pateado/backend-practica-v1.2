package api

import (
	Cempresa "practica/Crudempresa"
	"practica/internal/auth"
	"practica/internal/upload"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Cambia el puerto si es necesario
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	router.POST("/register", auth.RegisterHandler)
	router.POST("/login", auth.LoginHandler)
	router.POST("/register_empresa", auth.RegisterHandler_empresa)
	router.GET("/verify-email", auth.VerifyEmailHandler)
	router.POST("/password-reset", auth.SendPasswordResetEmailHandler)
	router.POST("/resend-verification", auth.ResendVerificationEmailHandler)
	// rutas crud practicas
	router.POST("/Create-practicas", Cempresa.Createpractica)
	router.GET("/Get-practicas", Cempresa.GetAllPracticas)
	router.PUT("/Update-practicas/:id", Cempresa.UpdatePractica)
	router.DELETE("/Delete-practica/:id", Cempresa.DeletePractica)
	//filtros pagina
	router.GET("/filtro-practicas", Cempresa.FiltroPracticas)

	router.POST("/upload-image", upload.UploadImageHandler)

	// Rutas protegidas
	protected := router.Group("/").Use(auth.AuthMiddleware) // Agrupar las rutas protegidas con el middleware
	{
		protected.POST("/complete-profile", auth.CompleteProfileHandler) // Ruta para completar perfil
		// Otras rutas protegidas aquí
	}

	return router
}
