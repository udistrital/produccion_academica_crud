package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CorreccionMetadatos_20201214_133126 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CorreccionMetadatos_20201214_133126{}
	m.Created = "20201214_133126"

	migration.Register("CorreccionMetadatos_20201214_133126", m)
}

// Run the migrations
func (m *CorreccionMetadatos_20201214_133126) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201214_133126_correccion_metadatos.up.sql")

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
func (m *CorreccionMetadatos_20201214_133126) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201214_133126_correccion_metadatos.down.sql")

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
