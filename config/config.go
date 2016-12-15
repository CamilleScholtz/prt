package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

// Struct contains the config values
// TODO: Is a struct really needed?
var Struct config

// Initialize config options
type config struct {
	PortDir    string
	Order      []string
	Alias      [][]string
	IndentChar string
	Pull       map[string]pull
}

type pull struct {
	URL    string
	Branch string
}

func init() {
	// Read out config
	f, err := ioutil.ReadFile("./runtime/config.toml")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read config!")
		os.Exit(1)
	}

	// Decode config
	if _, err := toml.Decode(string(f), &Struct); err != nil {
		fmt.Fprintln(os.Stderr, "Could not decode config!")
		os.Exit(1)
	}
}
