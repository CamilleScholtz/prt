package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
)

// config is a stuct with all config values. See `runtime/config/config.toml` for
// more information about these values.
var config struct {
	PrtDir string
	PkgDir string
	SrcDir string
	WrkDir string

	Order []string
	Alias [][]string

	IndentChar string
	ErrorChar  string
	DarkColor  color.Attribute
	ErrorColor color.Attribute
	LightColor color.Attribute

	Pull map[string]Pull
}

// Pull is a struct with values related to repos.
// TODO: Why does this need to be a global?
type Pull struct {
	URL    string
	Branch string
}

// colorFix converts a config color (0..15) to a color compatible color (30..97).
func colorFix(i color.Attribute) (color.Attribute, error) {
	if i > 15 {
		return 0, fmt.Errorf("config: Could not convert '" + string(i) + "' to color!")
	}

	if i <= 7 {
		i += 30
	} else if i <= 15 {
		i += 82
	}

	return i, nil
}

// decodeConfig parses a toml config.
func parseConfig() error {
	_, err := toml.DecodeFile("/etc/prt/config.toml", &config)
	if err != nil {
		return fmt.Errorf("config /etc/prt/config.toml: " + err.Error())
	}

	config.DarkColor, err = colorFix(config.DarkColor)
	if err != nil {
		return err
	}
	config.ErrorColor, err = colorFix(config.ErrorColor)
	if err != nil {
		return err
	}
	config.LightColor, err = colorFix(config.LightColor)
	if err != nil {
		return err
	}

	return nil
}
