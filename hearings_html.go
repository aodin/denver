package main

import (
	"encoding/json"
	"fmt"
	"github.com/aodin/denver/geocode"
	"github.com/aodin/denver/liquor"
	"io/ioutil"
	"flag"
	"os"
)

// Convert a saved html file of hearings to JSON

func main() {
	path := flag.String("i", "", "input file")
	flag.Parse()

	if *path == "" {
		panic("Enter a file to parse")
	}

	f, err := os.Open(*path)
	if err != nil {
		panic(err)
	}

	contents, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	// Parse the contents
	raws, err := liquor.ParseHearingsHTML(contents)
	if err != nil {
		panic(err)
	}

	hearings, err := liquor.CleanHearings(raws, geocode.Google)
	if err != nil {
		panic(err)
	}

	// Pretty print the output!
	b, err := json.MarshalIndent(hearings, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", b)
}
