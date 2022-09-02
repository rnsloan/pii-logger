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
	flag.Parse()

	entitiesFilePath := "./pkg/pii/entities.toml"

	ticker := time.NewTicker(time.Second * time.Duration(*delay))

	write := pii.Initilise(entitiesFilePath, *locale)

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
