package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ActualizarNormatividadProduccionesDocencia_20210118_214609 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ActualizarNormatividadProduccionesDocencia_20210118_214609{}
	m.Created = "20210118_214609"

	migration.Register("ActualizarNormatividadProduccionesDocencia_20210118_214609", m)
}

// Run the migrations
func (m *ActualizarNormatividadProduccionesDocencia_20210118_214609) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210118_214609_actualizar_normatividad_producciones_docencia_up.sql")

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
func (m *ActualizarNormatividadProduccionesDocencia_20210118_214609) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
