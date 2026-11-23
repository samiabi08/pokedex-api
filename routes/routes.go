package routes

import (
	"net/http"
	"pokedex/controller"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()
	file := http.FileServer(http.Dir("assets"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", file))
	mux.Handle("GET /{$}", http.RedirectHandler("/pokemon", http.StatusMovedPermanently))
	mux.HandleFunc("GET /pokemon", controller.ListPokemon)
	mux.HandleFunc("GET /pokemon/{id}", controller.ShowPokemon)
	mux.HandleFunc("/", controller.NotFound)
	return mux
}
