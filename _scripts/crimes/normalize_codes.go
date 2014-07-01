package main

import (
	"flag"
	"fmt"
	"github.com/aodin/csv2"
	"github.com/aodin/denver/crime"
	"log"
	"os"
	"time"
)

var (
	input  = flag.String("i", "", "The input file")
	output = flag.String(
		"o",
		fmt.Sprintf("./codes_%s.csv", time.Now().Format(`2006-01-02`)),
		"Name of output file",
	)
)

func main() {
	flag.Parse()

	// Normalize the given liquor license file
	if *input == "" {
		log.Fatal("Please select an input file with -i")
	}

	raw, err := crime.ParseOffenseCodesCSV(*input)
	if err != nil {
		log.Fatal(err)
	}

	codes := crime.ConvertRawCodes(raw)
	log.Printf("Converted %d codes\n", len(codes))

	// Output to a csv
	o, err := os.OpenFile(*output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer o.Close()

	// Write the header
	ocsv := csv.NewWriter(o)
	ocsv.WriteHeader(&codes)
	ocsv.Marshal(&codes)
	log.Println("File written:", *output)
}
