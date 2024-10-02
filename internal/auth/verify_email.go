package auth

import (
	"net/http"
	"net/smtp"
	"os"
	"time"

	"practica/internal/database"
	"practica/internal/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY")) // Asegúrate de tener una clave secreta en tu .env

// Función para generar el token de verificación
func GenerateVerificationToken(email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // El token expira en 24 horas
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// Función para enviar el correo de verificación
func SendVerificationEmail(email, token string) error {
	from := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")
	to := email
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	msg := []byte("Subject: Verificación de correo\n\nPor favor verifica tu correo haciendo clic en el siguiente enlace:\n" +
		"http://localhost:8080/verify-email?token=" + token)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)

	return err
}

func VerifyEmailHandler(c *gin.Context) {
	tokenString := c.Query("token")

	// Parsear el token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Manejar el error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el token"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)

		// Actualizar el estado de validación en la base de datos
		var usuario models.Usuario
		result := database.DB.Model(&usuario).Where("correo = ?", email).Update("Id_estado_usuario", true)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el estado del usuario"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Correo verificado exitosamente. Perfil activado."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token inválido o expirado"})
	}
}
