package pii

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/BurntSushi/toml"
)

type Entities = map[string]map[string][]string

type Locale struct {
	EnAU []string `toml:"en-AU"`
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

func Write(entitiesFilePath string) (string, error) {
	locale := "EnAU"
	entities, err := getEntities(entitiesFilePath)

	if err != nil {
		return "", err
	}

	entityNames := getEntityNames(entities)

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	entityIndex := r.Intn(len(entityNames))
	randomEntity := entityNames[entityIndex]

	entityItems := entities[randomEntity][locale]
	itemIndex := r.Intn(len(entityItems))

	return entityItems[itemIndex], nil
}
