package main

import (
	"accounting/pkg/db"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Can't load env: %v", err)
	}

	err = db.InitDb()
	if err != nil {
		log.Fatalf("Can't connect to DB: %v", err)
	}

	log.Fatal(http.ListenAndServe(":8080", routes()))
}
