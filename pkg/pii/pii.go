package pii

import (
	"encoding/json"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

type Entities = map[string]map[string][]string
type EntitiesIndexCache = map[string]map[int]bool

type TomlEntities struct {
	Phone struct {
		Locale
	}
	Name struct {
		Locale
	}
}

// https://stackoverflow.com/a/42849112
func structToMap(entities TomlEntities) (Entities, error) {
	var myMap map[string]map[string][]string
	data, err := json.Marshal(entities)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &myMap)
	return myMap, nil
}

func getEntities(entitiesFilePath string) (Entities, error) {
	var tomlEntities TomlEntities
	_, err := os.Open(entitiesFilePath)

	if err != nil {
		return nil, err
	}

	_, error := toml.DecodeFile(entitiesFilePath, &tomlEntities)

	if error != nil {
		return nil, error
	}

	entities, error := structToMap(tomlEntities)

	if error != nil {
		return nil, error
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

func getRandomItemIndex(cache EntitiesIndexCache, entityName string, entitiesItemLength int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	if _, ok := cache[entityName]; !ok {
		cache[entityName] = make(map[int]bool)
	}

	// clear cache if it contains 75%+ of available entity items
	if float32(len(cache[entityName]))/float32(entitiesItemLength) >= .75 {
		cache[entityName] = make(map[int]bool)
	}

	index := r.Intn(entitiesItemLength)

	if _, ok := cache[entityName][index]; !ok {
		cache[entityName][index] = true
		return index
	} else {
		return getRandomItemIndex(cache, entityName, entitiesItemLength)
	}
}

func Initilise(entitiesFilePath, loc string) func() (string, error) {
	var entitiesCache Entities
	var entitiesNameCache []string
	entitiesIndexCache := make(map[string]map[int]bool)

	return func() (string, error) {
		locale := formatLocale(loc)
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)

		if len(entitiesCache) == 0 || len(entitiesNameCache) == 0 {
			entities, err := getEntities(entitiesFilePath)

			if err != nil {
				return "", err
			}

			entitiesCache = entities
			entitiesNameCache = getEntityNames(entitiesCache)
		}

		entityIndex := r.Intn(len(entitiesNameCache))
		randomEntity := entitiesNameCache[entityIndex]

		entityItems := entitiesCache[randomEntity][locale]
		itemIndex := getRandomItemIndex(entitiesIndexCache, randomEntity, len(entityItems))

		return entityItems[itemIndex], nil
	}
}
