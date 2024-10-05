package upload

import (
	"net/http"
	"practica/internal/storage"

	"github.com/gin-gonic/gin"
)

// UploadImageHandler maneja la subida de imágenes de perfil a Supabase
// @Summary Subir una imagen de perfil
// @Description Sube una imagen al bucket de Supabase y devuelve la URL pública
// @Tags upload
// @Accept mpfd
// @Produce json
// @Param file formData file true "Imagen a subir"
// @Success 200 {object} map[string]string "URL de la imagen subida"
// @Failure 400 {object} map[string]string "Error al subir la imagen"
// @Router /upload-image [post]
func UploadImageHandler(c *gin.Context) {
	// Obtener el archivo del formulario
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se ha proporcionado un archivo"})
		return
	}

	// Subir la imagen a Firebase Storage
	url, err := storage.UploadFileToFirebase(file, "ulink-sprint-1.appspot.com") // Reemplaza con tu bucket
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al subir la imagen", "details": err.Error()})
		return
	}

	// Responder con la URL pública de la imagen
	c.JSON(http.StatusOK, gin.H{"message": "Imagen subida correctamente", "url": url})
}
