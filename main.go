package main

import (
	"log"
	"practica/api"
	_ "practica/docs" // Importa los documentos generados por swag
	"practica/internal/auth"
	"practica/internal/database"
	"practica/internal/models"
	"practica/pkg/config"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Practica API
// @version 1.0
// @description Esta es una API de ejemplo para Practica.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

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

	// Realizar la migración de la tabla fuera de la transacción
	if !database.DB.Migrator().HasTable(&models.Usuario{}) {
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciar el servidor
	router.Run(":8080")
}
