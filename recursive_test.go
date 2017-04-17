package main

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkRecursive(b *testing.B) {
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
	l, err := fullLocation("opt/firefox")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p, err := parsePort(l, "Pkgfile")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < b.N; i++ {
		recursive(p, make(map[key][]string), false, all, inst)
	}
}
