package commands

import "github.com/onodera-punpun/prt/config"

// Load config.
var c = config.Load()

// Opts.
var o []string

// Needed by depends() and build().
// TODO: Probably not the smartest way of handling this?
var i int
var v bool

// Again, needed by depends() and build().
// TODO: Again, probably not the smartest way of handling things
var allPorts, instPorts, instVers, checkPorts, toInst []string
