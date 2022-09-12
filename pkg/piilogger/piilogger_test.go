package piilogger

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
)

var rawToml = `
[phone]
en-AU.values = ["0487439000", "+61487439000", "0487 439 000"]
en-AU.sentences = ["my phone number is %s."]

[name]
en-AU.values = ["John Richards", "Danielle Wong"]

[IPAddress]
en-AU.values = ["47.124.160.94", "135.34.220.4"]
`

func contains[K comparable](s []K, e K) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TestWrite(t *testing.T) {
	write := Initilise("entities.toml", "en-AU", All, "no")

	_, err := write()

	if err != nil {
		t.Errorf("%v", err)
	}
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

	if len(config.Phone.ENAU.Values) < 1 {
		t.Errorf("config.Phone.EnAU expected slice got: %s", config.Phone.ENAU)
	}
}

func TestGetEntities(t *testing.T) {
	entities, err := getEntities(rawToml)

	if err != nil {
		t.Errorf("%v", err)
	}

	if entities["Name"]["ENAU"]["Values"] == nil {
		t.Errorf("could not find expected entity in toml file [\"Name\"][\"ENAU\"][\"Values\"]")
	}
}

func TestGetEntityNames(t *testing.T) {
	entities := map[string]map[string]map[string][]string{"Name": {"ENAU": {"Values": []string{"john"}}}, "Phone": {"ENAU": {"Values": []string{"0487439000"}}}}
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

func TestGetRandomValueIndex(t *testing.T) {
	entitiesIndexCache := make(map[string]map[int]bool)
	s := make(map[int]bool)

	index1 := getRandomValueIndex(entitiesIndexCache, "Name", 6)
	index2 := getRandomValueIndex(entitiesIndexCache, "Name", 6)
	index3 := getRandomValueIndex(entitiesIndexCache, "Name", 6)
	index4 := getRandomValueIndex(entitiesIndexCache, "Name", 6)

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

func TestGenerateEntityValue(t *testing.T) {
	reg := "04[0-9]{8}"
	value, err := generateEntityValueFromRegEx(fmt.Sprintf("/%s/", reg))

	if err != nil {
		t.Errorf("%v", err)
	}

	match, _ := regexp.MatchString(reg, value)

	if !match {
		t.Errorf("expected %s to match the regular expression: %s", value, reg)
	}
}

func TestSentences(t *testing.T) {
	write := Initilise("entities.toml", "en-AU", All, "no")
	out, _ := write()

	if strings.Contains(out, "my") {
		t.Errorf("expected output to not be a sentence. Got: %s", out)
	}

	write2 := Initilise("entities.toml", "en-AU", "Phone", "always")
	out2, _ := write2()

	if !strings.Contains(out2, "number") {
		t.Errorf("expected output to be a sentence. Got: %s", out2)
	}
}
