package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifica el token JWT
// AuthMiddleware verifica el token JWT
// @Summary Verifica el token JWT
// @Description Middleware que verifica el token JWT en el encabezado Authorization
// @Tags middleware
// @Success 200 {string} string "Token verificado correctamente"
// @Failure 401 {object} ErrorResponse "No se proporcionó token o token malformado"
// @Router /auth_middleware [get]
func AuthMiddleware(c *gin.Context) {
	// Obtener el token del encabezado Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se proporcionó token"})
		c.Abort()
		return
	}

	// Eliminar "Bearer " del token
	idToken := strings.TrimPrefix(authHeader, "Bearer ")
	if idToken == authHeader {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token malformado"})
		c.Abort()
		return
	}

	// Verificar el token con Firebase
	ctx := context.Background()
	token, err := authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		c.Abort()
		return
	}

	// Guardar el UID del usuario en el contexto para usarlo en otras rutas
	c.Set("uid", token.UID)
	c.Next() // Continuar la ejecución de la ruta
}
