package main

import (
	"crud_postgres/estudiantes"
	"fmt"
	"log"
)

func main() {
	e := estudiantes.Estudiante{
		Name:   "David",
		Age:    40,
		Active: true,
	}
	
	err := estudiantes.EstudianteCrear(e)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creado exitosamente!!!")

}
