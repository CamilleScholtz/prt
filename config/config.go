package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
)

// Config is a stuct with all config values.
type Config struct {
	PortDir    string
	Order      []string
	Alias      [][]string
	IndentChar string
	ErrorChar  string
	DarkColor  color.Attribute
	ErrorColor color.Attribute
	LightColor color.Attribute
	Pull       map[string]Pull
}

// Pull is a struct with values related to repos.
type Pull struct {
	URL    string
	Branch string
}

// colorFix converts a config color (0..15) to a color compatible color (30..97).
func colorFix(i color.Attribute) (color.Attribute, error) {
	if i > 15 {
		return 0, fmt.Errorf("config colorFix: Could not convert '" + string(i) + "' to color!")
	}

	if i <= 7 {
		i += 30
	} else if i <= 15 {
		i += 82
	}

	return i, nil
}

// Decode decodes the config.
func Decode() *Config {
	var c Config
	_, err := toml.DecodeFile("/etc/prt/config.toml", &c)
	// TODO: Once ports is fixed, make this return and err.
	if err != nil {
		fmt.Fprintln(os.Stderr, "config decode /etc/prt/config.toml: "+err.Error())
		os.Exit(1)
	}

	// Convert colors to something usable.
	c.DarkColor, err = colorFix(c.DarkColor)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	c.ErrorColor, err = colorFix(c.ErrorColor)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	c.LightColor, err = colorFix(c.LightColor)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return &c
}
