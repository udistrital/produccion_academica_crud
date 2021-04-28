package main

import (
	"fmt"
	"github.com/astaxie/beego/migration"
	"io/ioutil"
	"strings"
)

// DO NOT MODIFY
type ActualizarOrtografia_20210211_161950 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ActualizarOrtografia_20210211_161950{}
	m.Created = "20210211_161950"

	migration.Register("ActualizarOrtografia_20210211_161950", m)
}

// Run the migrations
func (m *ActualizarOrtografia_20210211_161950) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210211_161950_actualizar_ortografia_up.sql")

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
func (m *ActualizarOrtografia_20210211_161950) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
