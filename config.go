package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

// Initialize global variables
var Config configStruct

// Initialize config options
type configStruct struct {
	PortDir   string
	RepoOrder []string
}

func InitConfig() {
	// Read out config
	configFile, err := ioutil.ReadFile("./config/config.toml")
	if err != nil {
		fmt.Println("Could not read config.")
		os.Exit(1)
	}

	// Decode config
	if _, err := toml.Decode(string(configFile), &Config); err != nil {
		fmt.Println("Could not decode config.")
		os.Exit(1)
	}
}
