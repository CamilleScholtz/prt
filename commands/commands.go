package commands

import "github.com/onodera-punpun/prt/config"

// Initialize some variables
var conf, _ = config.Load()
var i int
var o []string
var allPorts, instPorts, instVers, checkPorts []string
