package auth

import (
	"context"
	"net/http"
	"practica/internal/database"
	"practica/internal/models"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

// RegisterRequest estructura de los datos recibidos
type RegisterRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Nombres   string `json:"nombres" binding:"required"`
	Apellidos string `json:"apellidos" binding:"required"`
}

// RegisterResponse estructura de la respuesta de éxito
type RegisterResponse struct {
	Message     string `json:"message"`
	FirebaseUID string `json:"firebase_uid"`
}

// ErrorResponse estructura de la respuesta de error
type ErrorResponse struct {
	Error string `json:"error"`
}

// RegisterHandler maneja el registro del usuario
// @Summary Registra un nuevo usuario
// @Description Registra un nuevo usuario con email y contraseña
// @Tags auth
// @Accept json
// @Produce json
// @Param register body RegisterRequest true "Datos de registro"
// @Success 200 {object} RegisterResponse
// @Failure 500 {object} ErrorResponse
// @Router /register [post]
func RegisterHandler(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear el usuario en Firebase con email y password
	params := (&auth.UserToCreate{}).
		Email(req.Email).
		Password(req.Password)

	user, err := authClient.CreateUser(context.Background(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear usuario en Firebase: " + err.Error()})
		return
	}

	// Crear el usuario en la base de datos sin almacenar la contraseña
	usuario := models.Usuario{
		Correo:           req.Email,
		Nombres:          req.Nombres,
		Apellidos:        req.Apellidos,
		Firebase_usuario: user.UID,
		Rol:              "estudiante", // Rol por defecto
	}

	result := database.DB.Create(&usuario)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el usuario en la base de datos"})
		return
	}

	// Generar token de verificación de correo
	token, err := GenerateVerificationToken(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token de verificación"})
		return
	}

	// Enviar correo de verificación
	err = SendVerificationEmail(req.Email, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar el correo de verificación"})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "Usuario creado correctamente. Verifica tu correo", "firebase_uid": user.UID})
}
