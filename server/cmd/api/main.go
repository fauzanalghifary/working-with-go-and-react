// This is the entry point for the application
package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN    string
	Domain string
	DB     repository.DatabaseRepo
}

func main() {

	// set application config
	var app application

	// read from command line
	// WILL DESTROY LATERS
	flag.StringVar(&app.DSN, "dsn", "postgresql://postgres:j60hgNg2fJJRONghdxI0@containers-us-west-141.railway.app:8057/railway", "Postgres Connection")
	flag.Parse()

	// connect to database
	conn, err := app.connectToDB()

	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	defer app.DB.Connection().Close()

	app.Domain = "example.com"

	log.Println("Starting server on port", port)

	// start server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}

}
