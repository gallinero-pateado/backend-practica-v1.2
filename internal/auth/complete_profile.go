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

// ProfileUpdateResponse estructura de la respuesta de éxito
type ProfileUpdateResponse struct {
	Message string `json:"message"`
}

// CompleteProfileHandler permite a los usuarios completar o actualizar su perfil
// @Summary Completa o actualiza el perfil del usuario
// @Description Permite a los usuarios completar o actualizar su perfil
// @Tags auth
// @Accept json
// @Produce json
// @Param profile body ProfileUpdateRequest true "Datos del perfil"
// @Success 200 {object} ProfileUpdateResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /complete_profile [post]
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
