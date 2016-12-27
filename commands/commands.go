package commands

import "github.com/onodera-punpun/prt/config"

// TODO: PThis whole global var thing isn't really clean

// Load config.
var c = config.Load()

// Opts.
var o []string

// Needed by depends() and build(), I always assign i to ints anyway/
var i int

// List with all ports.
var all []string

// List with checked ports.
var cp []string

// List with installed ports.
var inst []string

// List with installed port versions.
var instv []string

// List with ports to install.
var toInst []string
