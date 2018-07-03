package main

import (
	"fmt"
	"regexp"

	"github.com/go2c/optparse"
	"github.com/onodera-punpun/prt/ports"
)

// provCommand searches ports for files.
func provCommand(input []string) error {
	// Define valid arguments.
	o := optparse.New()
	argi := o.Bool("installed", 'i', false)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	vals, err := o.Parse(input)
	if err != nil {
		return fmt.Errorf("invaild argument, use `-h` for a list of arguments")
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt prov [arguments] [queries]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -i,   --installed       search in installed ports")
		fmt.Println("  -h,   --help            print help and exit")

		return nil
	}

	// This command needs a value.
	if len(vals) == 0 {
		return fmt.Errorf("please specify a query")
	}

	for _, v := range vals {
		if *argi {
			var db ports.Database
			if err := db.Parse(); err != nil {
				return err
			}

			for _, p := range db.Packages {
				// Search for files.
				var fl []string
				for _, f := range p.Files {
					m, err := regexp.MatchString(v, f)
					if err != nil {
						return err
					}
					if m {
						fl = append(fl, f)
					}
				}

				// Print port location.
				if len(fl) > 0 {
					fmt.Println(p.Name)
				}

				// Print matched files.
				for _, f := range fl {
					fmt.Printf("%s%s\n", dark(config.IndentChar), f)
				}
			}
		} else {
			all, err := ports.All()
			if err != nil {
				return err
			}

			for _, p := range all {
				if err := p.Footprint.Parse(); err != nil {
					continue
				}

				// Search for files.
				var fl []string
				for _, f := range p.Footprint.Files {
					m, err := regexp.MatchString(v, f.Path)
					if err != nil {
						return err
					}
					if m {
						fl = append(fl, f.Path)
					}
				}

				// Print port location.
				if len(fl) > 0 {
					fmt.Println(p.Location.Base())
				}

				// Print matched files.
				for _, f := range fl {
					fmt.Printf("%s%s\n", dark(config.IndentChar), f)
				}
			}
		}
	}

	return nil
}
