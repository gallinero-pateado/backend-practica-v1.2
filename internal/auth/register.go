package auth

import (
	"context"
	"net/http"
	"practica/internal/database"
	"practica/internal/models"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"time"
)

type RegisterRequest struct {
	Email     string `json:"email" binding:"required"`
	Nombres   string `json:"nombres" binding:"required"`
	Apellidos string `json:"apellidos" binding:"required"`
}

// RegisterHandler maneja el registro del usuario
func RegisterHandler(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear el usuario en Firebase
	params := (&auth.UserToCreate{}).
		Email(req.Email).
		Password("password123") // Asegúrate de que la contraseña se gestiona solo en Firebase

	user, err := authClient.CreateUser(context.Background(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear usuario en Firebase: " + err.Error()})
		return
	}

	// Crear el usuario en la base de datos usando GORM
	usuario := models.Usuario{
		Correo:            req.Email,
		Nombres:           req.Nombres,
		Apellidos:         req.Apellidos,
		Fecha_creacion:    time.Now(),
		Rol:               "estudiante",
		Firebase_usuario:  user.UID,
	}

	// Guardar el usuario en la base de datos
	result := database.DB.Create(&usuario)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el usuario en la base de datos: " + result.Error.Error()})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "Usuario creado correctamente", "firebase_uid": user.UID})
}
