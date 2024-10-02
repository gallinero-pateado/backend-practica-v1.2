package database

import (
    "context"
    "fmt"
    "os"
    "github.com/jackc/pgx/v4"
    "log"
)

var Conn *pgx.Conn

func Connect() error {
    // Obtener las variables de entorno
    host := os.Getenv("SUPABASE_HOST")
    dbname := os.Getenv("SUPABASE_DB")
    user := os.Getenv("SUPABASE_USER")
    password := os.Getenv("SUPABASE_PASSWORD")
    port := os.Getenv("SUPABASE_PORT")

    // Construir la URL de conexión
    url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbname)

    // Conectar a la base de datos
    conn, err := pgx.Connect(context.Background(), url)
    if err != nil {
        return fmt.Errorf("No se pudo conectar a la base de datos: %v", err)
    }
    Conn = conn
    log.Println("Conexión a Supabase exitosa")
    return nil
}
