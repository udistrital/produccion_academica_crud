package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ActualizarEsquemaProduccionAcademica_20201112_164236 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ActualizarEsquemaProduccionAcademica_20201112_164236{}
	m.Created = "20201112_164236"

	migration.Register("ActualizarEsquemaProduccionAcademica_20201112_164236", m)
}

// Run the migrations
func (m *ActualizarEsquemaProduccionAcademica_20201112_164236) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201112_164236_actualizar_esquema_produccion_academica.up.sql")

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
func (m *ActualizarEsquemaProduccionAcademica_20201112_164236) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201112_164236_actualizar_esquema_produccion_academica.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
