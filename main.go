package main

import (
	"log"
	"practica/api"
	"practica/internal/auth"
	"practica/internal/database"
	"practica/internal/models"
	"practica/internal/storage"
	"practica/pkg/config"

	_ "practica/docs" // Importar la documentación generada

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// Aquí solo asignamos con `=`, ya que la variable `err` ya fue declarada antes
	err = storage.InitStorage()
	if err != nil {
		log.Fatalf("Error inicializando Supabase Storage: %v", err)
	}

	// Registrar rutas
	router := api.SetupRoutes()

	// Agregar la ruta de Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciar el servidor
	router.Run(":8080")
}
