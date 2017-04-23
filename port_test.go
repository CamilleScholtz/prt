package main

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkGetDepends(b *testing.B) {
	if err := parseConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	all, err := ports()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inst, err := instPorts()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p, err := parsePort(fullLocation("opt/firefox"), "Pkgfile")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < b.N; i++ {
		recursive(p, make(map[string][]string), false, all, inst)
	}
}

func BenchmarkParseFootprint(b *testing.B) {
	if err := parseConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var p port
	p.Location = fullLocation("opt/firefox")

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
	p.Location = fullLocation("opt/firefox")

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
	p.Location = fullLocation("opt/firefox")

	for i := 0; i < b.N; i++ {
		p.parsePkgfile()
	}
}

func BenchmarkLocation(b *testing.B) {
	if err := parseConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	all, err := ports()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < b.N; i++ {
		location("opt/firefox", all)
	}
}
