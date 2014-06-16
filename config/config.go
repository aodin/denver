package config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Config struct {
	Port     int            `json:"port"`
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// Return a string of credentials approriate for Go's sql.Open() func
func (db DatabaseConfig) Credentials() string {
	// TODO Different credentials for different drivers
	return fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s",
		db.Host,
		db.Port,
		db.Name,
		db.User,
		db.Password,
	)
}

// By default, the parser will look for a file called settings.json in
// current directory.
func Parse() (Config, error) {
	return ParseFile("./settings.json")
}

func ParseFile(filename string) (Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	return parse(f)
}

func parse(f io.Reader) (Config, error) {
	var c Config
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		return c, err
	}
	if err = json.Unmarshal(contents, &c); err != nil {
		return c, err
	}
	// TODO Allow flag values to override configuration?
	return c, nil
}
