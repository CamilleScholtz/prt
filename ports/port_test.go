package ports

import "testing"

func TestNew(t *testing.T) {
	p := New("/usr/src/prt/opt/firefox")

	got := p.Location
	want := Location{"/usr/src/prt", "opt", "firefox"}
	if got != want {
		t.Errorf("p.Location: Got %s, want %s", got, want)
	}
}
