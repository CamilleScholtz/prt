package ports

import (
	"testing"
)

/*func TestAlias(t *testing.T) {
	a := [][]Location{
		[]Location{
			Location{"/usr/src/prt", "core", "openssl"},
			Location{"/usr/src/prt", "6c37", "libressl"},
		},
	}
}*/

func TestLocate(t *testing.T) {
	o := []string{"opt", "6c37-git", "contrib", "6c37"}
	a := []Port{
		New("/usr/src/prt/contrib/mpv"),
		New("/usr/src/prt/opt/firefox"),
		New("/usr/src/prt/6c37-git/mpv"),
	}

	p, err := Locate("firefox", o, a)
	if err != nil {
		t.Errorf(err.Error())
	}

	got := p[0].Location
	want := Location{"/usr/src/prt", "opt", "firefox"}
	if got != want {
		t.Errorf("p[0].Location: Got %s, want %s", got, want)
	}

	p, err = Locate("mpv", o, a)
	if err != nil {
		t.Errorf(err.Error())
	}

	got = p[0].Location
	want = Location{"/usr/src/prt", "6c37-git", "mpv"}
	if got != want {
		t.Errorf("p[0].Location: Got %s, want %s", got, want)
	}
	got = p[1].Location
	want = Location{"/usr/src/prt", "contrib", "mpv"}
	if got != want {
		t.Errorf("p[1].Location: Got %s, want %s", got, want)
	}
}

func BenchmarkLocate(b *testing.B) {
	o := []string{"opt", "6c37-git", "contrib", "6c37"}
	a := []Port{
		New("/usr/src/prt/contrib/mpv"),
		New("/usr/src/prt/opt/firefox"),
		New("/usr/src/prt/6c37-git/mpv"),
	}

	for i := 0; i < b.N; i++ {
		Locate("firefox", o, a)
	}
}

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		All("/usr/srt/prt")
	}
}
