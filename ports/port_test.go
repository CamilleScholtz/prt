package ports

import (
	"testing"
)

func TestNew(t *testing.T) {
	p := New("/usr/src/prt/opt/firefox")

	got := p.Location
	want := Location{"/usr/src/prt", "opt", "firefox"}
	if got != want {
		t.Errorf("p.Location: Got %s, want %s", got, want)
	}
}

func BenchmarkParseDepends(b *testing.B) {
	p := New("/usr/src/prt/opt/firefox")
	o := []string{"punpun", "6c37-dropin", "core", "6c37-git", "6c37-update",
		"6c37", "opt", "xorg", "contrib"}
	a, _ := All("/usr/src/prt")

	p.Pkgfile.Parse()

	for i := 0; i < b.N; i++ {
		p.ParseDepends([][]Location{}, o, a)
	}
}
