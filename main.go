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
		//fmt.Println("  depends                 list dependencies recursively")
		//fmt.Println("  diff                    list outdated packages")
		//fmt.Println("  graph                   generate dependency graph")
		//fmt.Println("  info                    print port information")
		//fmt.Println("  install                 build and install ports and their dependencies")
		//fmt.Println("  list                    list ports and packages")
		fmt.Println("  loc                     print port locations")
		//fmt.Println("  patch                   patch ports")
		//fmt.Println("  prov                    search ports for files")
		fmt.Println("  pull                    pull in ports")
		//fmt.Println("  sysup                   update outdated packages")
		//fmt.Println("  uninstall               uninstall packages")
		fmt.Println("  help                    print help and exit")
	//case "depends":
	//	err = depends(os.Args[2:])
	//case "diff":
	//	err = diff(os.Args[2:])
	//case "graph":
	//	err = graph(os.Args[2:])
	//case "info":
	//	err = info(os.Args[2:])
	//case "install":
	//	err =install(os.Args[2:])
	//case "list":
	//	err = list(os.Args[2:])
	case "loc":
		err = loc(os.Args[2:])
	//case "patch":
	//	err = patch(os.Args[2:])
	//case "prov":
	//	err = prov(os.Args[2:])
	case "pull":
		err = pull(os.Args[2:])
	//case "sysup":
	//	err = sysup(os.Args[2:])
	//case "uninstall":
	//	err = uninstall(os.Args[2:])
	default:
		err = fmt.Errorf("invalid command, use `help` for a list of commands")
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s!\n", capitalize(err.Error()))
		os.Exit(1)
	}
	os.Exit(0)
}
