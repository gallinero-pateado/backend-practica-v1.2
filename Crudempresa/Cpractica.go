package Crudempresa

import (
	"net/http"
	"practica/internal/database"
	"practica/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

// RegisterRequest estructura de los datos recibidos
type practicasRequest struct {
	Titulo             string    `json:"Titulo"`
	Descripcion        string    `json:"Descripcion"`
	Id_empresa         int       `json:"Id_Empresa"`
	Ubicacion          string    `json:"Ubicacion"`
	Fecha_inicio       time.Time `json:"Fecha_inicio"`
	Fecha_fin          time.Time `json:"Fecha_fin"`
	Requisitos         string    `json:"Requisitos"`
	Fecha_expiracion   time.Time `json:"Fecha_expiracion"`
	Id_estado_practica int       `json:"Id_estado_practica"`
	Fecha_publicacion  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func Createpractica(c *gin.Context) {
	var req practicasRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//tomar hora de creacion
	localTime := time.Now().Local()

	// Crear practica
	practica := models.Practica{
		Titulo:             req.Titulo,
		Descripcion:        req.Descripcion,
		Id_empresa:         req.Id_empresa,
		Ubicacion:          req.Ubicacion,
		Fecha_inicio:       req.Fecha_inicio,
		Fecha_fin:          req.Fecha_fin,
		Requisitos:         req.Requisitos,
		Fecha_expiracion:   req.Fecha_expiracion,
		Fecha_publicacion:  localTime,
		Id_estado_practica: 1,
	}

	result := database.DB.Create(&practica)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la practica en la base de datos"})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "La Oferta de practica fue creada exitosamente", "id_practica": practica.Id})
}
