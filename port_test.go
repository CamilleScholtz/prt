package main

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func BenchmarkDepends(b *testing.B) {
	if err := parseConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get all ports.
	all, err := ports()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var p port
	p.Location = path.Join(config.PrtDir, "opt/firefox")
	if err := p.parsePkgfile(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i := 0; i < b.N; i++ {
		p.depends(false, all)
	}
}

func BenchmarkParseFootprint(b *testing.B) {
	if err := parseConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var p port
	p.Location = path.Join(config.PrtDir, "opt/firefox")

	for i := 0; i < b.N; i++ {
		p.parseFootprint()
	}
}

func BenchmarkParseMd5sum(b *testing.B) {
	if err := parseConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var p port
	p.Location = path.Join(config.PrtDir, "opt/firefox")

	for i := 0; i < b.N; i++ {
		p.parseMd5sum()
	}
}

func BenchmarkParsePkgfile(b *testing.B) {
	if err := parseConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var p port
	p.Location = path.Join(config.PrtDir, "opt/firefox")

	for i := 0; i < b.N; i++ {
		p.parsePkgfile()
	}
}

/*func BenchmarkParsePkgfileSh(b *testing.B) {
	if err := parseConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var p port
	p.Location = path.Join(config.PrtDir, "opt/firefox")

	for i := 0; i < b.N; i++ {
		p.parsePkgfileSh()
	}
}*/

func BenchmarkLocation(b *testing.B) {
	if err := parseConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get all ports.
	all, err := ports()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i := 0; i < b.N; i++ {
		location("opt/firefox", all)
	}
}
