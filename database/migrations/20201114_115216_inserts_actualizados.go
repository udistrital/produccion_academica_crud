package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertsActualizados_20201114_115216 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsActualizados_20201114_115216{}
	m.Created = "20201114_115216"

	migration.Register("InsertsActualizados_20201114_115216", m)
}

// Run the migrations
func (m *InsertsActualizados_20201114_115216) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201114_115216_inserts_actualizados.up.sql")

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
func (m *InsertsActualizados_20201114_115216) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201114_115216_inserts_actualizados.down.sql")

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
