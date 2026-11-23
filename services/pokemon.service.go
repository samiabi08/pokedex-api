package services

import (
	"pokedex/api"
	"pokedex/models"
	"strconv"
	"strings"
)

func RecherchePokemonParId(idParam string) (*models.Pokemon, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		return nil, models.ErrBadRequest
	}
	pokemon, err := api.GetPokemon(id)
	if err != nil {
		return nil, err
	}
	if pokemon.Name.Fr == "" {
		return nil, models.ErrNotFound
	}
	return &pokemon, nil
}

func RecherchePokemonParNom(nom string) ([]models.Pokemon, error) {
	liste, err := api.GetAllPokemon()
	if err != nil {
		return nil, err
	}
	nom = strings.ToLower(strings.TrimSpace(nom))
	resultat := []models.Pokemon{}
	for _, value := range liste {
		if value.Id < 1 {
			continue
		}
		if nom != "" && !strings.Contains(strings.ToLower(value.Name.Fr), nom) {
			continue
		}
		resultat = append(resultat, value)
	}
	return resultat, nil
}
