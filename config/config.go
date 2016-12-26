package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
)

// Config is a stuct with all config values
type Config struct {
	PortDir    string
	Order      []string
	Alias      [][]string
	IndentChar string
	DarkColor  color.Attribute
	LightColor color.Attribute
	Pull       map[string]Pull
}

// Pull is a struct with values related to repos
type Pull struct {
	URL    string
	Branch string
}

// colorFix converts a config color (0..15) to a color compatible color (30..97)
func colorFix(i color.Attribute) (color.Attribute, error) {
	if i > 15 {
		return 0, fmt.Errorf("Could not convert '" + string(i) + "' to color!")
	}

	if i <= 7 {
		i += 30
	} else if i <= 15 {
		i += 82
	}

	return i, nil
}

// Load loads the config
func Load() (*Config, error) {
	// TODO: Use filepath stuff here?
	var config Config
	_, err := toml.DecodeFile("/etc/prt/config.toml", &config)
	if err != nil {
		return nil, fmt.Errorf("Could not decode config!")
	}

	// Convert colors to something usable
	config.DarkColor, err = colorFix(config.DarkColor)
	if err != nil {
		return nil, err
	}
	config.LightColor, err = colorFix(config.LightColor)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
