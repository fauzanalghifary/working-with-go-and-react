// This is the entry point for the application
package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {

	// set application config
	var app application

	// read from command line

	// connect to database

	app.Domain = "example.com"

	log.Println("Starting server on port", port)

	// start server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}

}
