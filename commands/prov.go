package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// Prov searches ports for files
func Prov(args []string) {
	// Define opts
	shortopts := "hi"
	longopts := []string{
		"--help",
		"--installed",
	}

	// Read out opts
	opts, vals, err := getopt.Getopt(args, shortopts, longopts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	for _, opt := range opts {
		switch opt[0] {
		case "-h", "--help":
			fmt.Println("Usage: prt print [arguments] [queries]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -i,   --installed       search in installed ports only")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-i", "--installed":
			o = append(o, "i")
		}
	}

	if len(vals) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify a query!")
		os.Exit(1)
	}

	for _, val := range vals {
		r, err := regexp.Compile(val)
		if err != nil {
			fmt.Fprintln(os.Stderr, "'"+val+"' is not a valid regex!")
			continue
		}

		// TODO: Use Alias and Loc here to always display repo info?
		if utils.StringInList("i", o) {
			// TODO: Should I use filepath stuff here?
			db, err := os.Open("/var/lib/pkg/db")
			if err != nil {
				fmt.Fprintln(os.Stderr, "Could not read '/var/lib/pkg/db'!")
				continue
			}
			s := bufio.NewScanner(db)

			var blank bool
			var name string
			var files [][]string
			for s.Scan() {
				if blank {
					name = s.Text()
					blank = false
				} else if s.Text() == "" {
					blank = true
				} else if r.MatchString(s.Text()) {
					files = append(files, []string{name, s.Text()})
				}
			}

			var oldName string
			for _, file := range files {
				// Print port name
				if oldName != file[0] {
					fmt.Println(file[0])
				}

				// Print files
				color.Set(c.DarkColor)
				fmt.Print(c.IndentChar)
				color.Unset()
				fmt.Println(file[1])

				oldName = file[0]
			}

			db.Close()
		} else {
			allPorts, err := ports.All()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			for _, name := range allPorts {
				f, err := os.Open(filepath.Join(c.PortDir, name, ".footprint"))
				if err != nil {
					fmt.Fprintln(os.Stderr, "Could not read '"+filepath.Join(c.PortDir, name, ".footprint")+"'!")
					continue
				}
				s := bufio.NewScanner(f)

				var files []string
				for s.Scan() {
					if r.MatchString(s.Text()) {
						files = append(files, s.Text())
					}
				}

				// Print port name
				if len(files) > 0 {
					fmt.Println(name)
				}

				// Print files
				for _, file := range files {
					color.Set(c.DarkColor)
					fmt.Print(c.IndentChar)
					color.Unset()
					fmt.Println(strings.Split(file, "\t")[2])
				}

				f.Close()
			}
		}
	}
}
