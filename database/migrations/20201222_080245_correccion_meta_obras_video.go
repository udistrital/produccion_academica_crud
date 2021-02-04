package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CorreccionMetaObrasVideo_20201222_080245 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CorreccionMetaObrasVideo_20201222_080245{}
	m.Created = "20201222_080245"

	migration.Register("CorreccionMetaObrasVideo_20201222_080245", m)
}

// Run the migrations
func (m *CorreccionMetaObrasVideo_20201222_080245) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201222_080245_correccion_meta_obras_video.up.sql")

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
func (m *CorreccionMetaObrasVideo_20201222_080245) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201222_080245_correccion_meta_obras_video.down.sql")

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
