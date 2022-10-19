package coneccion

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

/*
//Estructura
type conector struct{
	ip 	
}
*/

// getConecction => Obtiene una conexión a la BD
func getConecction() *sql.DB {
	dsn := "postgres://golang:Dar113017@192.168.1.232:5432/testgo?sslmode=disable"

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	return db
}