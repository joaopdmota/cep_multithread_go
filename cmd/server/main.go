package main

import (
	"cep_finder/infra/server/handlers"
	"cep_finder/internal/middlewares"
	"fmt"
	"net/http"

	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(middlewares.Log)
	cepHandler := handlers.NewCepHandler()
	readyHandler := handlers.NewReadyHandler()

	router.HandleFunc("/", readyHandler.ServeHTTP).Methods("GET")
	router.HandleFunc("/cep", cepHandler.FetchCep).Methods("POST")

	fmt.Println("Serving on port 8080...")
	http.ListenAndServe(":8080", router)
}
