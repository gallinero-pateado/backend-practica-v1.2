package auth

import (
	"net/http"
	"practica/internal/database"
	"practica/internal/models"

	"github.com/gin-gonic/gin"
)

// Obtener práctica por ID
func GetUsuarioByUID(c *gin.Context) {
	var usuario models.Usuario

	// Obtener el ID de la práctica desde los parámetros de la URL
	iud := c.Param("firebase_usuario")

	// Buscar la práctica por ID en la base de datos
	if err := database.DB.First(&usuario, iud).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "usuario no encontrada"})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, usuario)
}
