package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type locationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationPage(url string) (locationAreaResponse, error) {
	entry, ok := cache.Get(url)
	var locations locationAreaResponse
	if ok { // if in cache
		json.Unmarshal(entry, &locations)
		return locations, nil
	}
	// not in cache
	res, err := http.Get(url)
	if err != nil {
		return locationAreaResponse{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	cache.Add(url, body)
	if err != nil {
		return locationAreaResponse{}, err
	}
	json.Unmarshal(body, &locations)
	return locations, nil
}
