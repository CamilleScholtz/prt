package ports

/*func TestParsePkgfile(t *testing.T) {
	p := New("/usr/src/prt/opt/firefox")

	if err := p.Pkgfile.Parse(); err != nil {
		t.Errorf(err.Error())
	}

	got := p.Pkgfile.Description
	want := "The Mozilla Firefox browser with Alsa support"
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
	want = "60.0.1"
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
		"gtk3", "python", "alsa-lib", "xorg-libxt", "yasm", "mesa3d", "rust"}
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
	want := []string{"https://ftp.mozilla.org/pub/firefox/releases/60.0.1/source/firefox-60.0.1.source.tar.xz",
		"firefox.desktop"}
	if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", want) {
		t.Errorf("p.Pkgfile.Source: Got %s, want %s", got, want)
	}
}

func BenchmarkParsePkgfileSource(b *testing.B) {
	p := New("/usr/src/prt/opt/firefox")

	for i := 0; i < b.N; i++ {
		p.Pkgfile.Parse(true)
	}
}*/
