package pii

import (
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/BurntSushi/toml"
)

var fileName = "test_file.toml"

func createTomlFile() string {
	path, err1 := os.Getwd()
	if err1 != nil {
		panic(err1)
	}

	file := path + "/" + fileName

	f, err2 := os.Create(path + "/" + fileName)
	if err2 != nil {
		panic(err2)
	}

	toml := `
[phone]
"en-AU" = ["0487439000", "+61487439000", "0487 439 000"]

[name]
"en-AU" = ["John Richards", "Danielle Wong"]
`

	d := []byte(toml)
	_, err3 := f.Write(d)
	if err3 != nil {
		panic(err3)
	}

	return file
}

func removeTomlFile() {
	path, err1 := os.Getwd()
	if err1 != nil {
		panic(err1)
	}

	os.Remove(path + "/" + fileName)
}

// test the default supplied 'entities.toml' file
func TestEntitiesToml(t *testing.T) {
	var config TomlEntities
	f := "entities.toml"

	entities, err := toml.DecodeFile(f, &config)

	if err != nil {
		t.Fatal(err)
	}

	entityNames := []string{"phone", "name"}

	for _, v := range entityNames {
		if !entities.IsDefined(v) {
			t.Errorf("Could not find entity: '%s' in entities file", v)
		}
	}

	match, _ := regexp.MatchString("^0[0-9]+", config.Phone.ENAU[0])
	if !match {
		t.Errorf("config.Phone.EnAU[0] expected phone number got: %s", config.Phone.ENAU[0])
	}
}

func TestGetEntities(t *testing.T) {
	file := createTomlFile()

	entities, err := getEntities(file)

	if err != nil {
		t.Errorf("%v", err)
	}

	if entities["Name"]["ENAU"] == nil {
		t.Errorf("Could not find expected entity in toml file [\"Name\"][\"ENAU\"]")
	}

	removeTomlFile()
}

func TestGetEntityNames(t *testing.T) {
	entities := map[string]map[string][]string{"Name": {"ENAU": []string{"john"}}, "Phone": {"ENAU": []string{"0487439000"}}}
	entityNames := getEntityNames(entities)
	expected := []string{"Name", "Phone"}

	if !reflect.DeepEqual(entityNames, expected) {
		t.Errorf("Expected: '%s' , got: %s", entityNames, expected)
	}
}

func TestGetRandomItemIndex(t *testing.T) {
	entitiesIndexCache := make(map[string]map[int]bool)
	s := make(map[int]bool)

	index1 := getRandomItemIndex(entitiesIndexCache, "Name", 6)
	index2 := getRandomItemIndex(entitiesIndexCache, "Name", 6)
	index3 := getRandomItemIndex(entitiesIndexCache, "Name", 6)
	index4 := getRandomItemIndex(entitiesIndexCache, "Name", 6)

	s[index1] = true
	s[index2] = true
	s[index3] = true
	s[index4] = true

	if len(s) != 4 {
		t.Errorf("expected map with four unique keys. Got: %#v", s)
	}
}

func TestFormatLocale(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"en-AU", "ENAU"},
		{"ha-Latn-NG", "HALATNNG"},
	} {
		got := formatLocale(c.in)
		if got != c.want {
			t.Errorf("formatLocale(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestWrite(t *testing.T) {
	write := Initilise("entities.toml", "en-AU")

	_, err := write()

	if err != nil {
		t.Errorf("%v", err)
	}
}
