package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/utils"
)

// Prov lists TODO
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
			fmt.Println("  -i,   --installed       search in installed ports")
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
		if utils.StringInList("i", o) {
			db, err := ioutil.ReadFile("/var/lib/pkg/db")
			if err != nil {
				fmt.Fprintln(os.Stderr, "Could not read '/var/lib/pkg/db'!")
				continue
			}

			// Thanks to CandyGumdrop for the regex!
			r, err := regexp.Compile("(?m).*" + val + ".*")
			if err != nil {
				fmt.Fprintln(os.Stderr, "'"+val+"' is not a valid regex!")
				continue
			}
			match := r.FindAll(db, -1)

			for _, f := range match {
				color.Set(config.Struct.DarkColor)
				fmt.Print(config.Struct.IndentChar)
				color.Unset()
				fmt.Println(string(f))
			}
		} else {

		}
	}
}
