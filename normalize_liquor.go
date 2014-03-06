package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/aodin/denver/liquor"
	"log"
	"os"
	"time"
)

var (
	input  = flag.String("input", "", "The input file")
	output = flag.String(
		"output",
		fmt.Sprintf("./licenses_%s.csv", time.Now().Format(`2006-01-02`)),
		"Name of output file",
	)
)

func main() {
	flag.Parse()

	// Normalize the given liquor license file
	if *input == "" {
		log.Fatal("Please select an input file with -input")
	}

	original, err := liquor.ParseLicensesCSV(*input)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Original file contains %d licenses\n", len(original))

	licenses, err := liquor.Normalize(original)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Normalized file contains %d licenses\n", len(licenses))

	// Output to a csv
	// TODO Don't overwrite the input file, ever
	o, err := os.OpenFile(*output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer o.Close()

	// Write the header
	ocsv := csv.NewWriter(o)
	ocsv.Write(liquor.NormalizedHeader())
	for _, license := range licenses {
		ocsv.Write(license.NormalizedCSV())
	}
	ocsv.Flush()
	log.Println("File written:", *output)
}
