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
	flag.Parse()

	entitiesFilePath := "./pkg/pii/entities.toml"

	ticker := time.NewTicker(time.Second * time.Duration(*delay))

	write := pii.Initilise(entitiesFilePath)

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
