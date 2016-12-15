package config

import (
	"fmt"
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
	_, err := toml.DecodeFile("./runtime/config.toml", &Struct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not decode config!")
		os.Exit(1)
	}
}
