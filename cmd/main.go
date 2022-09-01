package main

import (
	"flag"
	"fmt"
	"pii-logger/pkg/pii"
	"time"
)

func main() {
	delay := flag.Int("delay", 5, "the time between outputs")
	flag.Parse()

	ticker := time.NewTicker(time.Second * time.Duration(*delay))

	fmt.Println(pii.Write())

	for range ticker.C {
		fmt.Println(pii.Write())
	}
}
