package pokeapi

import (
	"github.com/quanchobi/pokedexcli/internal/pokecache"
	"time"
)

var cache *pokecache.Cache = pokecache.NewCache(5 * time.Second)
