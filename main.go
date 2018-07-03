package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr,
			"Missing command, use help for a list of commands!")
		os.Exit(1)
	}

	if err := parseConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var err error
	switch os.Args[1] {
	case "help":
		fmt.Println("Usage: prt command [arguments]")
		fmt.Println("")
		fmt.Println("commands:")
		fmt.Println("  depends                 list dependencies recursively")
		fmt.Println("  diff                    list outdated packages")
		fmt.Println("  graph                   generate dependency graph")
		fmt.Println("  info                    print port information")
		//fmt.Println("  install                 build and install ports and their dependencies")
		fmt.Println("  list                    list ports and packages")
		fmt.Println("  loc                     print port locations")
		//fmt.Println("  make                    make package from port")
		//fmt.Println("  patch                   patch ports")
		fmt.Println("  prov                    search ports for files")
		fmt.Println("  pull                    pull in ports")
		//fmt.Println("  sysup                   update outdated packages")
		//fmt.Println("  uninstall               uninstall packages")
		fmt.Println("  help                    print help and exit")
	case "depends":
		err = dependsCommand(os.Args[2:])
	case "diff":
		err = diffCommand(os.Args[2:])
	case "graph":
		err = graphCommand(os.Args[2:])
	case "info":
		err = infoCommand(os.Args[2:])
	//case "install":
	//	err = installCommand(os.Args[2:])
	case "list":
		err = listCommand(os.Args[2:])
	case "loc":
		err = locCommand(os.Args[2:])
	//case "make":
	//	err = makeCommand(os.Args[2:])
	//case "patch":
	//	err = patchCommand(os.Args[2:])
	case "prov":
		err = provCommand(os.Args[2:])
	case "pull":
		err = pullCommand(os.Args[2:])
	//case "sysup":
	//	err = sysupCommand(os.Args[2:])
	//case "uninstall":
	//	err = uninstallCommand(os.Args[2:])
	default:
		err = fmt.Errorf("invalid command, use `help` for a list of commands")
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s!\n", capitalize(err.Error()))
		os.Exit(1)
	}
	os.Exit(0)
}
