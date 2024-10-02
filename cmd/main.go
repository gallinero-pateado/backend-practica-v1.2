package main

import (
    "fmt"
    "log"
    "net/http"
    "practica/internal/database"
    "practica/internal/auth"      // Añadir el paquete de auth
    "practica/pkg/config"
    "practica/api"
)

func main() {
    // Cargar las configuraciones desde el archivo .env
    err := config.LoadConfig("config/app.env")
    if err != nil {
        log.Fatalf("Error cargando configuración: %v", err)
    }

    // Inicializar Firebase
    err = auth.InitFirebase()
    if err != nil {
        log.Fatalf("Error inicializando Firebase: %v", err)
    }

    // Conectar a la base de datos Supabase
    err = database.Connect()
    if err != nil {
        log.Fatalf("Error conectando a Supabase: %v", err)
    }

    // Configurar rutas y arrancar el servidor
    router := api.SetupRoutes()
    fmt.Println("Servidor corriendo en http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
