package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"time"

	cursor "github.com/ahmetalpbalkan/go-cursor"
	"github.com/cavaliercoder/grab"
	humanize "github.com/dustin/go-humanize"
	"github.com/go2c/optparse"
	"github.com/onodera-punpun/prt/ports"
)

// downloadCommand downloads port sources
func downloadCommand(input []string) error {
	// Define valid arguments.
	o := optparse.New()
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	_, err := o.Parse(input)
	if err != nil {
		return fmt.Errorf("invaild argument, use `-h` for a list of arguments")
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt download [arguments]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -h,   --help            print help and exit")

		return nil
	}

	p := ports.New(".")
	if err := p.Pkgfile.Parse(); err != nil {
		return err
	}

	if err := p.Pkgfile.Parse(true); err != nil {
		return err
	}

	var urls = []string{}
	r := regexp.MustCompile("^(http|https|ftp|file)://")
	for _, s := range p.Pkgfile.Source {
		f := path.Join(ports.SrcDir, path.Base(s))

		// Continue if file is not an URL.
		if !r.MatchString(s) {
			continue
		}

		// Continue if file has already been downloaded.
		if _, err := os.Stat(f); err == nil {
			continue
		}

		urls = append(urls, s)
	}

	res, err := grab.GetBatch(config.ConcurrentDownloads, "/tmp", urls...)
	if err != nil {
		return err
	}

	// Monitor downloads.
	dl := make([]*grab.Response, 0, len(urls))
	t := time.NewTicker(200 * time.Millisecond)
	defer t.Stop()

	// Hide cursor.
	fmt.Print(cursor.Hide())

outer:
	for {
		select {
		case r := <-res:
			if r != nil {
				dl = append(dl, r)
			} else {
				if err := downloadPrint(dl); err != nil {
					return err
				}
				break outer
			}
		case <-t.C:
			if err := downloadPrint(dl); err != nil {
				return err
			}
		}
	}

	// Show cursor again.
	fmt.Printf("\033[%dB%s", len(dl)*2, cursor.Show())

	return nil
}

// downloadPrint prints the progress of all downloads to the terminal.
func downloadPrint(downloads []*grab.Response) error {
	for i, d := range downloads {
		f := path.Base(d.Filename)
		c := humanize.Bytes(uint64(d.BytesComplete()))
		m := humanize.Bytes(uint64(d.Size))

		fmt.Printf("Downloading source %d/%d, %s.\n", i+1, len(downloads),
			light(f))
		fmt.Printf("%s%s%s of %s\n", cursor.ClearEntireLine(), dark(
			config.IndentChar), c, m)
	}

	// Move cursor two lines of for each download.
	fmt.Printf("\033[%dA", len(downloads)*2)

	return nil
}
