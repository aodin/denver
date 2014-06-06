package main

import (
	"flag"
	"fmt"
	"github.com/aodin/denver/addresses"
	"os"
)

// Parse a csv of addresses
func main() {
	path := flag.String("f", "", "input file")
	flag.Parse()

	if *path == "" {
		panic("Please provide a csv of addresses")
	}

	f, err := os.Open(*path)
	if err != nil {
		panic(err)
	}

	as, err := addresses.Parse(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Parsed %d addresses\n", len(as))
	// Do whatever you want with them
}
