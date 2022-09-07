package pii

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/BurntSushi/toml"
)

var fileName = "test_file.toml"

func contains[K comparable](s []K, e K) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

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

[IPAddress]
"en-AU" = ["47.124.160.94", "135.34.220.4"]
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
			t.Errorf("could not find entity: '%s' in entities file", v)
		}
	}

	if len(config.Phone.ENAU) < 1 {
		t.Errorf("config.Phone.EnAU expected slice got: %s", config.Phone.ENAU)
	}
}

func TestGetEntities(t *testing.T) {
	file := createTomlFile()

	entities, err := getEntities(file)

	if err != nil {
		t.Errorf("%v", err)
	}

	if entities["Name"]["ENAU"] == nil {
		t.Errorf("could not find expected entity in toml file [\"Name\"][\"ENAU\"]")
	}

	removeTomlFile()
}

func TestGetEntityNames(t *testing.T) {
	entities := map[string]map[string][]string{"Name": {"ENAU": []string{"john"}}, "Phone": {"ENAU": []string{"0487439000"}}}
	entityNames := getEntityNames(entities, All)
	expected := []string{"Name", "Phone"}

	for _, c := range []struct {
		in string
	}{
		{"Name"},
		{"Phone"},
	} {
		got := contains(entityNames, c.in)
		if !got {
			t.Errorf("could not find (%q) in %q", c.in, entityNames)
		}
	}

	if !contains(entityNames, expected[0]) || !contains(entityNames, expected[1]) {
		t.Errorf("expected: '%s' , got: %s", entityNames, expected)
	}

	entityNamesWithExclusion := getEntityNames(entities, "name")
	if !reflect.DeepEqual(entityNamesWithExclusion, []string{"Name"}) {
		t.Errorf("expected entity 'phone' to be excluded. Found: %s", entityNamesWithExclusion)
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

func TestGenerateEntityItem(t *testing.T) {
	reg := "04[0-9]{8}"
	item, err := generateEntityItem(fmt.Sprintf("/%s/", reg))

	if err != nil {
		t.Errorf("%v", err)
	}

	match, _ := regexp.MatchString(reg, item)

	if !match {
		t.Errorf("expected %s to match the regular expression: %s", item, reg)
	}
}

func TestWrite(t *testing.T) {
	write := Initilise("entities.toml", "en-AU", All)

	_, err := write()

	if err != nil {
		t.Errorf("%v", err)
	}
}
