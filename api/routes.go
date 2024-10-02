package api

import (
	"practica/internal/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/register", auth.RegisterHandler)

	return router
}
