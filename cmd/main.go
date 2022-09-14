package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	piilogger "pii-logger/pkg/piilogger"
	"time"
)

// updated at build time
var Version = "development"

func main() {
	// http://localhost:6060/debug/pprof/heap
	// top
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	version := flag.Bool("version", false, "prints the application version then exits")
	delay := flag.Int("delay", 5, "the time between outputs")
	locale := flag.String("locale", "en-AU", "language locale")
	entitiesFilePath := flag.String("entitiesFilePath", piilogger.DefaultFilePath, "entities file path")
	specificEntities := flag.String("specificEntities", "all", "specific entities to use")
	useSentences := flag.String("useSentences", "yes", "use the entities in natural language sentences if available")
	flag.Parse()

	if *version {
		fmt.Printf("version: %s\n", Version)
		os.Exit(1)
	}

	ticker := time.NewTicker(time.Second * time.Duration(*delay))

	write := piilogger.Initilise(*entitiesFilePath, *locale, *specificEntities, *useSentences)

	value, err := write()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(value)

	for range ticker.C {
		value, err := write()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(value)
	}
}
