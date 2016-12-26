package commands

import "github.com/onodera-punpun/prt/config"

// Load config
var c = config.Load()

// Opts
var o []string

// needed by recursive()
// TODO: Probably not the smartest way of hangling this?
var i int

// All of these are []strings, again, needed by recursive()
var allPorts, instPorts, instVers, checkPorts []string
