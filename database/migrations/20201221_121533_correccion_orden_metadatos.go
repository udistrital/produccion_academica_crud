package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CorreccionOrdenMetadatos_20201221_121533 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CorreccionOrdenMetadatos_20201221_121533{}
	m.Created = "20201221_121533"

	migration.Register("CorreccionOrdenMetadatos_20201221_121533", m)
}

// Run the migrations
func (m *CorreccionOrdenMetadatos_20201221_121533) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201221_121533_correccion_orden_metadatos.up.sql")

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
func (m *CorreccionOrdenMetadatos_20201221_121533) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201221_121533_correccion_orden_metadatos.down.sql")

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
