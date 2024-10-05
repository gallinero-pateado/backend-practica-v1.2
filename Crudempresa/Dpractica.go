package Crudempresa

import (
	"net/http"
	"practica/internal/database"
	"practica/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeletePractica elimina una práctica por su ID
func DeletePractica(c *gin.Context) {
	// Obtener el ID de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam) // Convertir el ID de string a int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Eliminar la práctica de la base de datos
	var practica models.Practica
	result := database.DB.First(&practica, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Práctica no encontrada"})
		return
	}

	if err := database.DB.Delete(&practica).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la práctica"})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "La práctica fue eliminada exitosamente"})
}
