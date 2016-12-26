package commands

import "github.com/onodera-punpun/prt/config"

// Load config
var c = config.Load()

// Opts
var o []string

// Needed by recursive()
// TODO: Probably not the smartest way of handling this?
var i int

// Again, needed by recursive()
var allPorts, instPorts, instVers, checkPorts []string
