package models

import (
	"time"
)

type Practica struct {
	Id                 uint      `gorm:"primaryKey;autoIncrement"`
	Titulo             string    `json:"Titulo"`
	Descripcion        string    `json:"Descripcion"`
	Id_empresa         int       `json:"Id_Empresa"`
	Ubicacion          string    `json:"Ubicacion"`
	Fecha_inicio       time.Time `json:"Fecha_inicio"`
	Fecha_fin          time.Time `json:"Fecha_fin"`
	Requisitos         string    `json:"Requisitos"`
	Fecha_publicacion  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Fecha_expiracion   time.Time `json:"Fecha_expiracion"`
	Id_estado_practica int       `json:"Id_estado_practica"`
	Modalidad          string    `json:"Modalidad"`
	Area_practica      string    `json:"Area_practica"`
	Jornada            string    `json:"Jornada"`
}

// TableName establece el nombre de la tabla para GORM
func (Practica) TableName() string {
	return "practica"
}
