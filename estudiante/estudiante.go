package estudiante

import (
	"errors"
	"time"

	"crud_postgres/connection"
)

// Estudiante Estructura del estudiante
type Estudiante struct {
	ID       int
	Name     string
	Age      int16
	Active   bool
	CreateAt time.Time
	UpdateAt time.Time
}

// Crear Registra un estudiante en la BD
func EstudianteCrear(e Estudiante) error {
	q := `INSERT INTO estudiantes (name, age, active) 
	      VALUES ($1, $2, $3)`

	db := connection.GetConecction()

	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(e.Name, e.Age, e.Active)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error: Se esperaba una fila afectada.")
	}

	return nil
}
