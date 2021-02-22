package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablasInscripcionCrud_20190713_140607 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablasInscripcionCrud_20190713_140607{}
	m.Created = "20190713_140607"

	migration.Register("CrearTablasInscripcionCrud_20190713_140607", m)
}

// Run the migrations
func (m *CrearTablasInscripcionCrud_20190713_140607) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20190713_140607_crear_tablas_inscripcion_crud.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *CrearTablasInscripcionCrud_20190713_140607) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20190713_140607_crear_tablas_inscripcion_crud.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
