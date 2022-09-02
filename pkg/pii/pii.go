package pii

import (
	"encoding/json"
	"math/rand"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

type Entities = map[string]map[string][]string

type Locale struct {
	ENAU []string `toml:"en-AU"`
}

type Config struct {
	Phone struct {
		Locale
	}
	Name struct {
		Locale
	}
}

// https://stackoverflow.com/a/42849112
func structToMap(config Config) (Entities, error) {
	var myMap map[string]map[string][]string
	data, err := json.Marshal(config)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &myMap)
	return myMap, nil
}

func getEntities(entitiesFilePath string) (Entities, error) {
	var config Config
	_, err := toml.DecodeFile(entitiesFilePath, &config)

	if err != nil {
		return nil, err
	}

	entities, err := structToMap(config)

	if err != nil {
		return nil, err
	}

	return entities, nil
}

func getEntityNames(entities Entities) []string {
	entityNames := make([]string, len(entities))
	i := 0
	for k := range entities {
		entityNames[i] = k
		i++
	}
	return entityNames
}

// 'en-AU' -> 'ENAU'
func formatLocale(locale string) string {
	converted := strings.ReplaceAll(locale, "-", "")
	converted = strings.ToUpper(converted)
	return converted
}

func Initilise(entitiesFilePath string, loc string) func() (string, error) {
	var entitiesCache Entities
	var entitiesNameCache []string

	return func() (string, error) {
		locale := formatLocale(loc)

		if len(entitiesCache) == 0 || len(entitiesNameCache) == 0 {
			entities, err := getEntities(entitiesFilePath)
			entitiesCache = entities

			if err != nil {
				return "", err
			}

			entitiesNameCache = getEntityNames(entitiesCache)
		}

		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)

		entityIndex := r.Intn(len(entitiesNameCache))
		randomEntity := entitiesNameCache[entityIndex]

		entityItems := entitiesCache[randomEntity][locale]
		itemIndex := r.Intn(len(entityItems))

		return entityItems[itemIndex], nil
	}
}
