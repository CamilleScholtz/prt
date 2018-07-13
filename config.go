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

	Order   []string
	Aliases [][]ports.Location

	ConcurrentDownloads int

	IndentChar  string
	WarningChar string

	DarkColor    color.Attribute
	LightColor   color.Attribute
	WarningColor color.Attribute

	Repo map[string]repo
}

// pull is a struct with values related to repos.
type repo struct {
	URL    string
	Branch string
}

var dark, light, warning func(a ...interface{}) string

// pareConfig parses a toml config.
func parseConfig() error {
	_, err := toml.DecodeFile("/etc/prt/config.toml", &config)
	if err != nil {
		return fmt.Errorf("config /etc/prt/config.toml: " + err.Error())
	}

	ports.PrtDir = config.PrtDir
	ports.PkgDir = config.PkgDir
	ports.SrcDir = config.SrcDir
	ports.WrkDir = config.WrkDir

	ports.Order = config.Order
	ports.Aliases = config.Aliases

	dark = color.New(config.DarkColor).SprintFunc()
	light = color.New(config.LightColor).SprintFunc()
	warning = color.New(config.WarningColor).SprintFunc()

	return nil
}
