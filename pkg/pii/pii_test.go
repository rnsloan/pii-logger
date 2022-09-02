package pii

import (
	"regexp"
	"testing"

	"github.com/BurntSushi/toml"
)

func TestEntitiesToml(t *testing.T) {
	var config Config
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

func TestWrite(t *testing.T) {
	write := Initilise("entities.toml", "en-AU")

	_, err := write()

	if err != nil {
		t.Errorf("%v", err)
	}
}
