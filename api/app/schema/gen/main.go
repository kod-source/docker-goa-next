package main

import (
	"log"

	"github.com/kod-source/docker-goa-next/app/schema"
	"github.com/shogo82148/myddlmaker"
)

func main() {
	// create a new DDL maker.
	m, err := myddlmaker.New(&myddlmaker.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m.AddStructs(&schema.User{})
	m.AddStructs(&schema.Post{})
	m.AddStructs(&schema.Comment{})
	m.AddStructs(&schema.Like{})
	m.AddStructs(&schema.Room{})
	m.AddStructs(&schema.UserRoom{})
	m.AddStructs(&schema.Thread{})
	m.AddStructs(&schema.Message{})

	// generate an SQL file.
	if err := m.GenerateFile(); err != nil {
		log.Fatal(err)
	}

	// generate Go source code for basic SQL operations
	// such as INSERT, SELECT, and UPDATE.
	if err := m.GenerateGoFile(); err != nil {
		log.Fatal(err)
	}
}
