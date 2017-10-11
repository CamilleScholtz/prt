package ports

import (
	"fmt"
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
		New("/usr/src/prt/6c37/mpv"),
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
	got = p[2].Location
	want = Location{"/usr/src/prt", "6c37", "mpv"}
	if got != want {
		t.Errorf("p[2].Location: Got %s, want %s", got, want)
	}
}

func ExampleLocate() {
	order := []string{"core", "opt", "contrib"}
	all, _ := All("/usr/src/prt")
	p, _ := Locate("firefox", order, all)

	fmt.Println(p[0].Location.Full())
	// Output:
	// /usr/src/prt/opt/firefox
}
