package main

import (
	"log"
	"practica/api"
	"practica/internal/auth"
	"practica/internal/database"
	"practica/internal/models"
	"practica/pkg/config"
)

func main() {
	// Cargar las configuraciones desde el archivo .env
	err := config.LoadConfig("config/app.env")
	if err != nil {
		log.Fatalf("Error cargando configuración: %v", err)
	}

	// Inicializar la base de datos
	err = database.InitDatabase()
	if err != nil {
		log.Fatalf("Error inicializando la base de datos: %v", err)
	}

	// Verificar si la tabla usuarios ya existe
	if !database.DB.Migrator().HasTable(&models.Usuario{}) {
		// Si no existe, migrar el modelo
		err = database.DB.AutoMigrate(&models.Usuario{})
		if err != nil {
			log.Fatalf("Error al migrar modelos: %v", err)
		}
	}

	// Inicializar Firebase
	err = auth.InitFirebase()
	if err != nil {
		log.Fatalf("Error inicializando Firebase: %v", err)
	}

	// Registrar rutas
	router := api.SetupRoutes()

	// Iniciar el servidor
	router.Run(":8080")
}
