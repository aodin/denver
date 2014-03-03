package main

import (
	"flag"
	"fmt"
	"github.com/aodin/denver/liquor"
	"log"
	"os"
)

// Track:
// * Additions
// * Deletions
// * Updates

// Determine the difference between the maps
// The a map should be the older map
func Additions(a, b map[string]*liquor.License) {
	for key, license := range b {
		if _, exists := a[key]; !exists {
			log.Println("Added:", license)
		}
	}
}

func Deletions(a, b map[string]*liquor.License) {
	for key, license := range a {
		if _, exists := b[key]; !exists {
			log.Println("Deleted:", license)
		}
	}
}

func Changes(a, b map[string]*liquor.License) {
	for key, license := range a {
		if _, exists := b[key]; exists {
			if !license.Equals(b[key]) {
				log.Println("Changed:", key)
			}
		}
	}
}

func ById(ls []*liquor.License) (byId map[string]*liquor.License) {
	byId = make(map[string]*liquor.License)
	for _, license := range ls {
		_, exists := byId[license.UniqueId]
		if exists {
			// log.Printf("Unique Id %s already exists on line %d\n", license.UniqueId, i + 2)
			// TODO Are the licenses the same?
		} else {
			byId[license.UniqueId] = license
		}
	}
	return
}


func main() {
	flag.Parse()

	// Get the non-keyword arguments
	args := flag.Args()

	// There should be two files to compare
	if len(args) != 2 {
		fmt.Println("Please enter two files to compare")
		os.Exit(1)
	}

	// Get the licenses from each file
	a, err := liquor.ParseLicensesCSV(args[0])
	if err != nil {
		panic(err)
	}
	b, err := liquor.ParseLicensesCSV(args[1])
	if err != nil {
		panic(err)
	}

	// How many unique ids are there?
	uniqueA := ById(a)
	uniqueB := ById(b)

	Additions(uniqueA, uniqueB)
	Deletions(uniqueA, uniqueB)
	Changes(uniqueA, uniqueB)
}
