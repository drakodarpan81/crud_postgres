package connection

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// GetConecction Obtiene una conexión a la BD
func GetConecction() *sql.DB {
	//dsn := "postgres://golang:Dar113017@192.168.1.232:5432/testgo?sslmode=disable"
	dsn := "postgres://golang:Dar113017@127.0.0.1:5432/testgo?sslmode=disable"

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
