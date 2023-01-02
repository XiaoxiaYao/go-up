package main

import (
	"backend/pkg/repository"
	"backend/pkg/repository/dbrepo"
	"flag"
	"log"
	"net/http"
)

type application struct {
	DSN string
	DB  repository.DatabaseRepo
}

func main() {
	app := application{}

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Postgre connection")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	mux := app.routes()

	log.Println("😊 Starting server on port 8080. :)")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
