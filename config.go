// config.go contains functions that interact with the config file, this is a
// file called `config.toml` found in `/etc/prt/`.

package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/ports"
)

// config is a stuct with all config values. See `runtime/config/config.toml`
// for more information about these values.
var config struct {
	PrtDir string
	PkgDir string
	SrcDir string
	WrkDir string

	Order []string
	Alias [][]ports.Location

	IndentChar string
	ErrorChar  string
	DarkColor  color.Attribute
	ErrorColor color.Attribute
	LightColor color.Attribute

	Repo map[string]repo
}

type location struct {
	repo string
	port string
}

// pull is a struct with values related to repos.
type repo struct {
	URL    string
	Branch string
}

// colorFix converts a config color (0..15) to a color compatible color (so
// 30..97).
func colorFix(i color.Attribute) (color.Attribute, error) {
	if i > 15 {
		return 0, fmt.Errorf("config: Could not convert '" + string(i) +
			"' to color!")
	}

	if i <= 7 {
		i += 30
	} else if i <= 15 {
		i += 82
	}

	return i, nil
}

// pareConfig parses a toml config.
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
