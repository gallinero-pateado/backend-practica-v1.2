package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Estructura de solicitud para el correo de recuperación
type PasswordResetRequest struct {
	Email string `json:"email" binding:"required"`
}

// @Summary Enviar correo de recuperación de contraseña
// @Description Enviar un correo de recuperación de contraseña usando la API de Firebase
// @Tags Autenticación
// @Accept json
// @Produce json
// @Param email body string true "Correo electrónico"
// @Success 200 {object} map[string]string{"message": "Correo de recuperación enviado con éxito"}
// @Failure 500 {object} map[string]string{"error": "Error al enviar el correo de recuperación"}
// @Failure 400 {object} map[string]string{"error": "Error desde Firebase al enviar el correo de recuperación"}
// @Router /api/v1/password-recovery [post]

// Handler para enviar correo de recuperación de contraseña usando la API REST de Firebase
func SendPasswordResetEmailHandler(c *gin.Context) {
	var req PasswordResetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email requerido"})
		return
	}

	// Construir la solicitud para la API REST de Firebase
	apiKey := os.Getenv("FIREBASE_API_KEY")
	url := "https://identitytoolkit.googleapis.com/v1/accounts:sendOobCode?key=" + apiKey

	// Payload para la solicitud
	resetPayload := map[string]string{
		"requestType": "PASSWORD_RESET",
		"email":       req.Email,
	}

	jsonPayload, _ := json.Marshal(resetPayload)

	// Enviar la solicitud HTTP POST a la API de Firebase
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar el correo de recuperación"})
		return
	}
	defer resp.Body.Close()

	// Verificar la respuesta
	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Error desde Firebase al enviar el correo de recuperación"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Correo de recuperación enviado con éxito"})
}
