package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"go-web-service/src/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/calculate", handlers.CalculateHandler).Methods("POST")

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}