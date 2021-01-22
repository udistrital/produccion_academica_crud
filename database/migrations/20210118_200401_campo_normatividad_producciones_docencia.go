package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CampoNormatividadProduccionesDocencia_20210118_200401 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CampoNormatividadProduccionesDocencia_20210118_200401{}
	m.Created = "20210118_200401"

	migration.Register("CampoNormatividadProduccionesDocencia_20210118_200401", m)
}

// Run the migrations
func (m *CampoNormatividadProduccionesDocencia_20210118_200401) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210118_200401_campo_normatividad_producciones_docencia_up.sql")

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
func (m *CampoNormatividadProduccionesDocencia_20210118_200401) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210118_200401_campo_normatividad_producciones_docencia_down.sql")

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
