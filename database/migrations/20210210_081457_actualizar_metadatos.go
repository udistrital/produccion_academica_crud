package main

import (
	"fmt"
	"github.com/astaxie/beego/migration"
	"io/ioutil"
	"strings"
)

// DO NOT MODIFY
type ActualizarMetadatos_20210210_081457 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ActualizarMetadatos_20210210_081457{}
	m.Created = "20210210_081457"

	migration.Register("ActualizarMetadatos_20210210_081457", m)
}

// Run the migrations
func (m *ActualizarMetadatos_20210210_081457) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210210_081457_actualizar_metadatos_up.sql")

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
func (m *ActualizarMetadatos_20210210_081457) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210210_081457_actualizar_metadatos_down.sql")

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
