package main

import (
	"fmt"
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
		// Continue if file is not an URL.
		if !r.MatchString(s) {
			continue
		}

		urls = append(urls, s)
	}
	if len(urls) == 0 {
		return nil
	}

	b, err := grab.GetBatch(config.ConcurrentDownloads, ports.SrcDir, urls...)
	if err != nil {
		return err
	}

	// Monitor downloads.
	rl := make([]*grab.Response, 0, len(urls))
	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()

	// Hide cursor.
	fmt.Print(cursor.Hide())
	defer fmt.Print(cursor.Show())

	var complete int
	for complete != len(urls) {
		select {
		case r := <-b:
			if r != nil {
				rl = append(rl, r)
			}
		case <-t.C:
			complete = 0
			for i, r := range rl {
				if r.IsComplete() {
					// TODO: A more descriptive error message.
					if err := r.Err(); err != nil {
						return err
					}

					complete++
				}

				f := path.Base(r.Filename)
				c := humanize.Bytes(uint64(r.BytesComplete()))
				m := humanize.Bytes(uint64(r.Size))

				fmt.Printf("Downloading source %d/%d, %s.\n", i+1, len(rl),
					light(f))
				fmt.Printf("%s%s%s of %s\n", cursor.ClearEntireLine(), dark(
					config.IndentChar), c, m)
			}

			// Move cursor two lines of for each download.
			fmt.Print(cursor.MoveUp(len(rl) * 2))
		}
	}

	// Move the cursor to the bottom.
	// TODO: When prt gets interrupted this doesn't get called.
	fmt.Print(cursor.MoveDown(len(rl) * 2))

	return nil
}
