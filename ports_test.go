package main

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkPorts(b *testing.B) {
	if err := parseConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i := 0; i < b.N; i++ {
		ports()
	}
}
