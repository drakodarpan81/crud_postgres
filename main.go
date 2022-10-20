package main

import (
	"crud_postgres/estudiante"
	"fmt"
	"log"
)

func main() {
	e := estudiante.Estudiante{
		Name:   "David",
		Age:    40,
		Active: true,
	}

	err := estudiante.EstudianteCrear(e)
	if err != nil {
		log.Fatal(err)
	}

	//connection.Conexion()

	fmt.Println("Creado exitosamente!!!")

}
