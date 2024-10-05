package Crudempresa

import (
	"net/http"
	"practica/internal/database"
	"practica/internal/models"

	"github.com/gin-gonic/gin"
)

// Obtener todas las prácticas
func GetAllPracticas(c *gin.Context) {
	var practicas []models.Practica

	// Obtener todas las prácticas de la base de datos
	if err := database.DB.Find(&practicas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las prácticas"})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, practicas)
}
