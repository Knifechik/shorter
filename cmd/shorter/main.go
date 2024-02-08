package main

import (
	"log"
	"net/http"
	"ozontest/cmd/shorter/internal/adapters/repo"
	"ozontest/cmd/shorter/internal/api/httpserv"
	"ozontest/cmd/shorter/internal/app"
)

func main() {
	dbName := "postgres"
	dbInfo := "dbname=postgres password=pass123 user=user123 sslmode=disable host=localhost port=5432"

	db, err := repo.New(dbName, dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	myapp := app.New(db)
	httpAPI := httpserv.New(myapp)

	s := http.Server{
		Addr:    ":8080",
		Handler: httpAPI,
	}

	log.Fatal(s.ListenAndServe())
}
