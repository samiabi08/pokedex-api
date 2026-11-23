package main

import (
	"log"
	"net/http"

	"pokedex/helpers"
	"pokedex/routes"
)

func main() {
	helpers.LoadTemplates()
	router := routes.SetupRoutes()
	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
