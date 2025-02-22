package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func ConnectDB() *sql.DB {
	dsn := "postgres://calvario:root@localhost:15432/DBGo?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("No se pudo hacer ping a la base de datos: %v", err)
	}

	fmt.Println("Conectado a la base de datos")
	return db
}
