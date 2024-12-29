package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"testrepo/internal/stuff"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("environment variable $PORT must be set")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", stuff.IndexHandler)
	http.ListenAndServe(":"+port, mux)
}
