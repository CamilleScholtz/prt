package ports

import (
	"testing"
)

func BenchmarkParse(b *testing.B) {
	p := New("/usr/src/prt/core/bash")

	for i := 0; i < b.N; i++ {
		p.Pkgfile.Parse()
	}
}

func BenchmarkParseStrict(b *testing.B) {
	p := New("/usr/src/prt/core/bash")

	for i := 0; i < b.N; i++ {
		p.Pkgfile.Parse(true)
	}
}
