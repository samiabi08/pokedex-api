package controller

import (
	"errors"
	"log"
	"net/http"
	"pokedex/helpers"
	"pokedex/models"
	"pokedex/services"
	"strings"
)

type ListPage struct {
	Search  string
	Pokemon []models.Pokemon
}

func ListPokemon(w http.ResponseWriter, r *http.Request) {
	recherche := strings.TrimSpace(r.URL.Query().Get("nom"))
	pokemon, err := services.RecherchePokemonParNom(recherche)
	if err != nil {
		RenderAPIError(w, err)
		return
	}
	helpers.Rendu(w, http.StatusOK, "list", ListPage{
		Search:  recherche,
		Pokemon: pokemon,
	})
}

func ShowPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := services.RecherchePokemonParId(r.PathValue("id"))
	if err != nil {
		RenderAPIError(w, err)
		return
	}
	helpers.Rendu(w, http.StatusOK, "pokemon", pokemon)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	helpers.RenduErreur(w, http.StatusNotFound, "Cette page n'existe pas.")
}

func RenderAPIError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrBadRequest):
		helpers.RenduErreur(w, http.StatusBadRequest, "Numéro de Pokémon invalide.")
	case errors.Is(err, models.ErrNotFound):
		helpers.RenduErreur(w, http.StatusNotFound, "Ce Pokémon n'existe pas.")
	case errors.Is(err, models.ErrUnavailable):
		log.Println("Tyradex :", err)
		helpers.RenduErreur(w, http.StatusServiceUnavailable, "Service momentanément indisponible.")
	default:
		log.Println("Erreur inattendue :", err)
		helpers.RenduErreur(w, http.StatusInternalServerError, "Une erreur est survenue.")
	}
}
