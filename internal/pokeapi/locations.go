package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type locationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationPage(url string) (locationAreaResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return locationAreaResponse{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreaResponse{}, err
	}

	var locations locationAreaResponse
	// we area expecting a single page of data, so unmarshal is fine.
	json.Unmarshal(body, &locations)

	return locations, nil
}
