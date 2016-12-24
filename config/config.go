package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
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
	DarkColor  color.Attribute
	LightColor color.Attribute
	Pull       map[string]Pull
}

// Pull is a struct with info about repos
type Pull struct {
	URL    string
	Branch string
}

// colorFix converts a config color (0..15) to a color compatible color (30..97)
func colorFix(i color.Attribute) (color.Attribute, error) {
	if i > 15 {
		return 0, fmt.Errorf("Could not convert '" + string(i) + "' to color!")
	}

	var c color.Attribute
	switch i {
	case 0:
		c = 30
	case 1:
		c = 31
	case 2:
		c = 32
	case 3:
		c = 33
	case 4:
		c = 34
	case 5:
		c = 35
	case 6:
		c = 36
	case 7:
		c = 37
	case 8:
		c = 90
	case 9:
		c = 91
	case 10:
		c = 92
	case 11:
		c = 93
	case 12:
		c = 94
	case 13:
		c = 95
	case 14:
		c = 96
	case 15:
		c = 97
	}

	return c, nil
}

func init() {
	_, err := toml.DecodeFile("/etc/prt/config.toml", &Struct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not decode config!")
		os.Exit(1)
	}

	// Convert colors to something usable
	Struct.DarkColor, err = colorFix(Struct.DarkColor)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	Struct.LightColor, err = colorFix(Struct.LightColor)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
