package auth

import (
	"net/http"
	"practica/internal/database"
	"practica/internal/models"
	"github.com/gin-gonic/gin"
)

type ProfileUpdateRequest struct {
	FechaNacimiento string `json:"fecha_nacimiento"`
	AnoIngreso      string `json:"ano_ingreso"`
	IdCarrera       uint   `json:"id_carrera"`
	FotoPerfil      string `json:"foto_perfil"`
}

// CompleteProfileHandler permite a los usuarios completar o actualizar su perfil
func CompleteProfileHandler(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	var req ProfileUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Actualizar el perfil en la base de datos
	var usuario models.Usuario
	result := database.DB.Model(&usuario).Where("firebase_usuario = ?", uid).Updates(models.Usuario{
		Fecha_nacimiento: req.FechaNacimiento,
		Ano_ingreso:      req.AnoIngreso,
		Id_carrera:       req.IdCarrera,
		Foto_perfil:      req.FotoPerfil,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el perfil"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Perfil actualizado correctamente"})
}
