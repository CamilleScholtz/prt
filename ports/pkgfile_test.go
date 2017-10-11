package ports

import (
	"fmt"
	"testing"
)

func TestParsePkgfile(t *testing.T) {
	p := New("/usr/src/prt/opt/firefox")

	if err := p.Pkgfile.Parse(); err != nil {
		t.Errorf(err.Error())
	}

	got := p.Pkgfile.Description
	want := "The Mozilla Firefox browser"
	if got != want {
		t.Errorf("p.Pkgfile.Description: Got %s, want %s", got, want)
	}
	got = p.Pkgfile.URL
	want = "https://www.mozilla.com/firefox/"
	if got != want {
		t.Errorf("p.Pkgfile.URL: Got %s, want %s", got, want)
	}
	got = p.Pkgfile.Maintainer
	want = "Fredrik Rinnestam, fredrik at crux dot nu"
	if got != want {
		t.Errorf("p.Pkgfile.Maintainer: Got %s, want %s", got, want)
	}
	got = p.Pkgfile.Name
	want = "firefox"
	if got != want {
		t.Errorf("p.Pkgfile.Name: Got %s, want %s", got, want)
	}
	got = p.Pkgfile.Version
	want = "52.4.0esr"
	if got != want {
		t.Errorf("p.Pkgfile.Version: Got %s, want %s", got, want)
	}
	got = p.Pkgfile.Release
	want = "1"
	if got != want {
		t.Errorf("p.Pkgfile.Release: Got %s, want %s", got, want)
	}
	gotL := p.Pkgfile.Depends
	wantL := []string{"nss", "autoconf-2.13", "unzip", "zip", "libidl", "gtk",
		"gtk3", "python", "alsa-lib", "xorg-libxt", "yasm", "mesa3d"}
	if fmt.Sprintf("%v", gotL) != fmt.Sprintf("%v", wantL) {
		t.Errorf("p.Pkgfile.Depends: Got %s, want %s", gotL, wantL)
	}
	gotL = p.Pkgfile.Optional
	wantL = []string{}
	if fmt.Sprintf("%v", gotL) != fmt.Sprintf("%v", wantL) {
		t.Errorf("p.Pkgfile.Optional: Got %s, want %s", gotL, wantL)
	}
	gotL = p.Pkgfile.Source
	wantL = []string{}
	if fmt.Sprintf("%v", gotL) != fmt.Sprintf("%v", wantL) {
		t.Errorf("p.Pkgfile.Source: Got %s, want %s", gotL, wantL)
	}
}

func BenchmarkParsePkgfile(b *testing.B) {
	p := New("/usr/src/prt/opt/firefox")

	for i := 0; i < b.N; i++ {
		p.Pkgfile.Parse()
	}
}

func TestParsePkgfileSource(t *testing.T) {
	p := New("/usr/src/prt/opt/firefox")

	if err := p.Pkgfile.Parse(true); err != nil {
		t.Errorf(err.Error())
	}

	got := p.Pkgfile.Source
	want := []string{"https://ftp.mozilla.org/pub/firefox/releases/52.4.0esr/source/firefox-52.4.0esr.source.tar.xz",
		"firefox-install-dir.patch", "firefox.desktop"}
	if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", want) {
		t.Errorf("p.Pkgfile.Source: Got %s, want %s", got, want)
	}
}

func BenchmarkParsePkgfileSource(b *testing.B) {
	p := New("/usr/src/prt/opt/firefox")

	for i := 0; i < b.N; i++ {
		p.Pkgfile.Parse(true)
	}
}

func TestRecursiveDepends(t *testing.T) {
	p := New("/usr/src/prt/opt/firefox")
	o := []string{"punpun", "6c37-dropin", "core", "6c37-git", "6c37-update",
		"6c37", "opt", "xorg", "contrib"}
	a, _ := All("/usr/src/prt")

	p.Pkgfile.Parse()
	got, err := p.Pkgfile.RecursiveDepends([][]Location{}, o, a)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := []string{}
	if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", want) {
		//t.Errorf("p.Pkgfile.RecursiveDepends: Got %s, want %s", got, want)
	}
}

func BenchmarkRecursiveDepends(b *testing.B) {
	p := New("/usr/src/prt/opt/firefox")
	o := []string{"punpun", "6c37-dropin", "core", "6c37-git", "6c37-update",
		"6c37", "opt", "xorg", "contrib"}
	a, _ := All("/usr/src/prt")

	p.Pkgfile.Parse()

	for i := 0; i < b.N; i++ {
		p.Pkgfile.RecursiveDepends([][]Location{}, o, a)
	}
}
