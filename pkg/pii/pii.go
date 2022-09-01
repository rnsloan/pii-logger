package pii

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Locale struct {
	EnAU []string `toml:"en-AU"`
}

type Config struct {
	Phone struct {
		Locale
	}
}

func Write() string {
	var config Config
	f := "./pkg/pii/entities.toml"

	_, err := toml.DecodeFile(f, &config)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return config.Phone.EnAU[0]

}
