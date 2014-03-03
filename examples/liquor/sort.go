package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/aodin/denver/liquor"
	"os"
	"strings"
)

func AddFilenameSuffix(filename, suffix string) string {
	tokens := strings.SplitN(filename, ".", 2)
	// Careful, tokens[1] may not exist
	if len(tokens) == 2 {
		return fmt.Sprintf("%s_%s.%s", tokens[0], suffix, tokens[1])
	}
	return fmt.Sprintf("%s_%s", tokens[0], suffix)
}

func main() {
	flag.Parse()

	// Get the non-keyword arguments
	args := flag.Args()

	// There should be two files to compare
	if len(args) != 1 {
		fmt.Println("Please enter a file to sort")
		os.Exit(1)
	}

	licenses, err := liquor.ParseLicensesCSV(args[0])
	if err != nil {
		panic(err)
	}

	liquor.Licenses(licenses).Sort()

	// Write the sorted licenses to a CSV
	filename := AddFilenameSuffix(args[0], "sorted")
	output, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	outputCSV := csv.NewWriter(output)

	// No header is written
	for _, row := range licenses {
		if err = outputCSV.Write(row.CSV()); err != nil {
			panic(err)
		}
	}
}
