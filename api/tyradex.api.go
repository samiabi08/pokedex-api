package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pokedex/models"

	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

const baseURL = "https://tyradex.app/api/v1"

func get(endpoint string, dest any) error {
	request, err := http.NewRequest(http.MethodGet, baseURL+endpoint, nil)
	if err != nil {
		return err
	}
	request.Header.Set("User-Agent", "DataDex/1.0")
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("%w : %v", models.ErrUnavailable, err)
	}
	defer response.Body.Close()
	switch {
	case response.StatusCode == http.StatusOK:
	case response.StatusCode == http.StatusNotFound:
		return models.ErrNotFound
	case response.StatusCode >= 500:
		return fmt.Errorf("%w : %s", models.ErrUnavailable, response.Status)
	default:
		return fmt.Errorf("réponse inattendue : %s", response.Status)
	}

	return json.NewDecoder(response.Body).Decode(dest)
}

func GetPokemon(id int) (models.Pokemon, error) {
	var pokemon models.Pokemon
	err := get(fmt.Sprintf("/pokemon/%d", id), &pokemon)
	return pokemon, err

}

func GetAllPokemon() ([]models.Pokemon, error) {
	var pokemonList []models.Pokemon
	err := get("/pokemon", &pokemonList)
	return pokemonList, err
}
