package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type locationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	/* \/only really interested in this part\/ */
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		/* ^only really interested in this part^ */
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetAreaEncounters(url string) ([]string, error) {
	var areaInfo locationArea
	var areaEncounters []string
	entry, ok := cache.Get(url)
	if ok { // if cached
		json.Unmarshal(entry, &areaInfo)
	} else {
		res, err := http.Get(url)
		if err != nil {
			return []string{}, err
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		cache.Add(url, body)
		if err != nil {
			return []string{}, err
		}
		json.Unmarshal(body, &areaInfo)
	}

	for _, encounter := range areaInfo.PokemonEncounters {
		areaEncounters = append(areaEncounters, encounter.Pokemon.Name)
	}
	return areaEncounters, nil
}
