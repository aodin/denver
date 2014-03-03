package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func ConvertLicensesCSV(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)

	// Skip the header
	var header []string
	header, err = r.Read()
	if err != nil {
		return err
	}

	// Update the header
	header[18] = "LATITUDE"
	header[19] = "LONGITUDE"

	entries, err := r.ReadAll()
	log.Printf("Converting %d entries\n", len(entries))

	// Build a single string of the coordinates to convert
	// TODO There must be a better way
	xys := make([]string, len(entries))

	// x and y are index 18 and 19
	for i, entry := range entries {
		xys[i] = fmt.Sprintf("%s %s", entry[18], entry[19])
	}

	cmd := exec.Command("gdaltransform", "-s_srs", "EPSG:2232", "-t_srs", "EPSG:4326")

	// Create a single argument string
	cmd.Stdin = strings.NewReader(strings.Join(xys, "\n"))
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return err
	}

	// Split the result into the lat longs
	coords := strings.Split(out.String(), "\n")
	log.Println("results:", len(coords))

	// Replace the entries x and y with lat and long
	var latlong []string
	for i, entry := range entries {
		// We get x, y, z
		latlong = strings.SplitN(coords[i], " ", 3)
		entry[18] = latlong[1]
		entry[19] = latlong[0]
	}

	o, err := os.OpenFile("./output.csv", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer o.Close()

	ocsv := csv.NewWriter(o)
	ocsv.Write(header)
	ocsv.WriteAll(entries)
	return nil
}

var file = flag.String("file", "", "path to file")

func main() {
	flag.Parse()

	if err := ConvertLicensesCSV(*file); err != nil {
		panic(err)
	}
}
