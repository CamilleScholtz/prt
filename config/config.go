package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

// Initialize global variables
var Struct Config

// Initialize config options
type Config struct {
	PrtDir string
}

func InitConfig() {
	// Read out config
	config, err := ioutil.ReadFile("./config.toml")
	if err != nil {
		fmt.Println("Could not read config.")
		os.Exit(1)
	}

	// Decode config
	if _, err := toml.Decode(string(config), &Struct); err != nil {
		fmt.Println("Could not decode config.")
		os.Exit(1)
	}
}
