package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"practica/internal/database"
	"practica/internal/models"

	"github.com/gin-gonic/gin"
)

// LoginRequest representa los datos de inicio de sesión del usuario
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// FirebaseLoginResponse representa la respuesta de Firebase
type FirebaseLoginResponse struct {
	IDToken string `json:"idToken"`
}

// LoginHandler maneja el inicio de sesión
// @Summary Inicia sesión en la aplicación
// @Description Permite a un usuario iniciar sesión con su email y contraseña
// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Datos de inicio de sesión"
// @Success 200 {object} FirebaseLoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Autenticar con Firebase
	token, err := SignInWithEmailAndPassword(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	// Buscar al usuario en la base de datos usando su email
	var usuario models.Usuario
	result := database.DB.Where("correo = ?", req.Email).First(&usuario)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Responder con el token JWT y el UID
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"uid":   usuario.Firebase_usuario,
	})
}

// SignInWithEmailAndPassword autentica al usuario con Firebase
func SignInWithEmailAndPassword(email, password string) (string, error) {
	apiKey := os.Getenv("FIREBASE_API_KEY") // Cargar la clave de API desde las variables de entorno
	url := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=" + apiKey

	// Datos a enviar a Firebase
	loginPayload := map[string]string{
		"email":             email,
		"password":          password,
		"returnSecureToken": "true",
	}
	jsonPayload, _ := json.Marshal(loginPayload)

	// Enviar la solicitud a Firebase
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Decodificar la respuesta de Firebase
	var firebaseResp FirebaseLoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&firebaseResp); err != nil {
		return "", err
	}

	// Retornar el token ID de Firebase
	return firebaseResp.IDToken, nil
}
