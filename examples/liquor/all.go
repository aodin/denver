package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/aodin/denver/licenses"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

func Additions(a, b map[string]*licenses.License) {
	for key, license := range b {
		_, exists := a[key]
		if !exists {
			log.Println("Added:", license)
		}
	}
}

func ById(ls []*licenses.License) (byId map[string]*licenses.License) {
	byId = make(map[string]*licenses.License)
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

var path = flag.String("path", ".", "Path to data files")

var prefix = "liquor_license"

type LicenseFile struct {
	Hash      string
	Name      string
	Path      string
	Timestamp time.Time
}

func (l LicenseFile) String() string {
	return l.Name 
}

func (l LicenseFile) Date() string {
	y, m, d := l.Timestamp.Date()
	return fmt.Sprintf("%d-%02d-%02d", y, m, d)
}

func (l LicenseFile) IsEmpty() bool {
	return l.Path == ""
}

var timestampRegexp = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

func FromFilename(filename string) (LicenseFile, error) {
	rawTimestamp := timestampRegexp.FindString(filename)
	t, err := time.Parse("2006-01-02", rawTimestamp)
	if err != nil {
		return LicenseFile{}, err
	}
	l := LicenseFile{
		Name:      filename,
		Timestamp: t,
	}
	return l, nil
}

type LicenseFiles []LicenseFile

// Implement the sort.Interface for sorting
func (a LicenseFiles) Len() int {
	return len(a)
}

func (a LicenseFiles) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type ByDate struct {
	LicenseFiles
}

func (a ByDate) Less(i, j int) bool {
	x, y := a.LicenseFiles[i], a.LicenseFiles[j]
	// Most recent logs first
	return x.Timestamp.Unix() > y.Timestamp.Unix()
}

func (a LicenseFiles) Sort() {
	sort.Sort(ByDate{a})
}

// Find all license files in the given data directory
func main() {
	flag.Parse()
	log.Println("Path:", *path)

	files, err := ioutil.ReadDir(*path)
	if err != nil {
		panic(err)
	}

	ls := make([]LicenseFile, 0)
	for _, file := range files {
		if !strings.HasPrefix(file.Name(), prefix) {
			continue
		}
		license, err := FromFilename(file.Name())
		if err != nil {
			log.Println("Could not parse filename:", file.Name())
			continue
		}
		license.Path = filepath.Join(*path, file.Name())

		// Generate the md5 hash of the file contents
		f, err := os.Open(license.Path)
		if err != nil {
			log.Println("Cound not open file:", license.Path)
		}
		defer f.Close()

		hasher := md5.New()
		_, err = io.Copy(hasher, f)
		if err != nil {
			log.Println("Cound generate hash for:", license.Name)
		}
		license.Hash = hex.EncodeToString(hasher.Sum(nil))
		ls = append(ls, license)
	}
	LicenseFiles(ls).Sort()

	var prev LicenseFile
	var prevLicenses map[string]*licenses.License

	

	for _, license := range ls {
		// Compare the two files (only if the hashes vary)
		var unique map[string]*licenses.License
		if !prev.IsEmpty() && prev.Hash != license.Hash {
			log.Println("Compare:", prev.Date(), "to", license.Date())

			// Get the licenses from each file
			a, err := licenses.ParseLicensesCSV(license.Path)
			if err != nil {
				panic(err)
			}
			unique = ById(a)
			Additions(prevLicenses, unique)
		} else {
			a, err := licenses.ParseLicensesCSV(license.Path)
			if err != nil {
				panic(err)
			}
			unique = ById(a)
		}
		prev = license
		prevLicenses = unique
	}
}
