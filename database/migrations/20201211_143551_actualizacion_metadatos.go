package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ActualizacionMetadatos_20201211_143551 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ActualizacionMetadatos_20201211_143551{}
	m.Created = "20201211_143551"

	migration.Register("ActualizacionMetadatos_20201211_143551", m)
}

// Run the migrations
func (m *ActualizacionMetadatos_20201211_143551) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201211_143551_actualizacion_metadatos.up.sql")

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
func (m *ActualizacionMetadatos_20201211_143551) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201211_143551_actualizacion_metadatos.down.sql")

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
