package main

import (
	"flag"
	"fmt"
	"os"
	"pii-logger/pkg/pii"
	"time"
)

func main() {
	delay := flag.Int("delay", 5, "the time between outputs")
	locale := flag.String("locale", "en-AU", "language locale")
	entitiesFilePath := flag.String("entitiesFilePath", "./pkg/pii/entities.toml", "entities file path")
	specificEntities := flag.String("specificEntities", "all", "specific entities to use")
	flag.Parse()

	ticker := time.NewTicker(time.Second * time.Duration(*delay))

	write := pii.Initilise(*entitiesFilePath, *locale, *specificEntities)

	item, err := write()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(item)

	for range ticker.C {
		item, err := write()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(item)
	}
}
