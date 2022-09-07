package pii

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	regen "github.com/zach-klippenstein/goregen"
)

const All = "all"

type Entities = map[string]map[string][]string
type EntitiesIndexCache = map[string]map[int]bool

type TomlEntities struct {
	Phone struct {
		Locale
	}
	Name struct {
		Locale
	}
	IPAddress struct {
		Locale
	}
}

// https://stackoverflow.com/a/56129336
func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
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

func getEntityNames(entities Entities, specificEntities string) []string {
	entityAllowList := []string{}

	contains := func(sl []string, item string) bool {
		for i := range sl {
			if strings.EqualFold(sl[i], item) {
				return true
			}
		}
		return false
	}

	if specificEntities != All {
		if strings.Contains(specificEntities, ",") {
			entityAllowList = strings.Split(specificEntities, ",")
		} else {
			entityAllowList = []string{specificEntities}
		}
	}

	entityNames := []string{}

	i := 0
	for k := range entities {
		if len(entityAllowList) > 0 {
			if contains(entityAllowList, k) {
				entityNames = append(entityNames, k)
			}
		} else {
			entityNames = append(entityNames, k)
		}
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

func getRandomValueIndex(cache EntitiesIndexCache, entityName string, entitiesValuesLength int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	if _, ok := cache[entityName]; !ok {
		cache[entityName] = make(map[int]bool)
	}

	// clear cache if it contains 75%+ of available entity values
	if float32(len(cache[entityName]))/float32(entitiesValuesLength) >= .75 {
		cache[entityName] = make(map[int]bool)
	}

	index := r.Intn(entitiesValuesLength)

	// check if the index has been used already
	if _, ok := cache[entityName][index]; !ok {
		cache[entityName][index] = true
		return index
	} else {
		return getRandomValueIndex(cache, entityName, entitiesValuesLength)
	}
}

func generateEntityValue(value string) (string, error) {
	regex := substr(value, 1, len(value)-2)
	regexValue, err := regen.Generate(regex)

	if err != nil {
		return "", err
	}

	return regexValue, nil
}

func Initilise(entitiesFilePath, loc string, specificEntities string) func() (string, error) {
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
			entitiesNameCache = getEntityNames(entitiesCache, specificEntities)
		}

		var randomEntity string
		var entityValues []string
		entityHasValues := false
		circuitBreaker := 0

		// using the supplied locale, find an entity with items
		for !entityHasValues {
			index := r.Intn(len(entitiesNameCache))
			randomEntity = entitiesNameCache[index]
			entityValues = entitiesCache[randomEntity][locale]

			if len(entityValues) > 0 {
				entityHasValues = true
			}

			circuitBreaker++
			if circuitBreaker == 10000 {
				return "", fmt.Errorf("using locale: %s , could not find an Entity with values. Check the toml file", locale)
			}
		}

		valueIndex := getRandomValueIndex(entitiesIndexCache, randomEntity, len(entityValues))
		value := entityValues[valueIndex]

		// 47 = '/'
		if value[0] == 47 && value[len(value)-1] == 47 {
			return generateEntityValue(value)
		}

		return value, nil
	}
}
