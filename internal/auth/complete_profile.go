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

// SuccessResponse representa la estructura para respuestas exitosas
type SuccessResponse struct {
	Message string `json:"message"`
}

// CompleteProfileHandler permite a los usuarios completar o actualizar su perfil
// @Summary Completar o actualizar perfil de usuario
// @Description Permite a los usuarios autenticados completar o actualizar su perfil
// @Tags profile
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param profile body ProfileUpdateRequest true "Datos para actualizar el perfil"
// @Success 200 {object} SuccessResponse "Perfil actualizado correctamente"
// @Failure 400 {object} ErrorResponse "Datos inválidos"
// @Failure 401 {object} ErrorResponse "Usuario no autenticado"
// @Failure 500 {object} ErrorResponse "Error al actualizar el perfil"
// @Router /complete-profile [post]
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
		PerfilCompletado: true,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el perfil"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Perfil actualizado correctamente"})
}
