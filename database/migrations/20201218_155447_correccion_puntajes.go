package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CorreccionPuntajes_20201218_155447 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CorreccionPuntajes_20201218_155447{}
	m.Created = "20201218_155447"

	migration.Register("CorreccionPuntajes_20201218_155447", m)
}

// Run the migrations
func (m *CorreccionPuntajes_20201218_155447) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201218_155447_correccion_puntajes.up.sql")

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
func (m *CorreccionPuntajes_20201218_155447) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201218_155447_correccion_puntajes.down.sql")

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
