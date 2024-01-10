package services

import (
	"echo-test/models"
	"encoding/json"
	"io"
	"net/http"
)

func GetAllKanto() (*models.PokemonResponse, error) {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject models.PokemonResponse
	json.Unmarshal(responseData, &responseObject)

	return &responseObject, nil
}
