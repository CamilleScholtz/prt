package main

import (
	"fmt"
	"os"

	"github.com/onodera-punpun/prt/cmd"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Missing command, use help for a list of commands!")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help":
		fmt.Println("Usage: prt command [arguments]")
		fmt.Println("")
		fmt.Println("commands:")
		fmt.Println("  depends                 list dependencies recursivly")
		fmt.Println("  diff                    list outdated packages")
		fmt.Println("  info                    print port information")
		fmt.Println("  install                 build and install packages")
		fmt.Println("  list                    list ports and packages")
		fmt.Println("  loc                     print port locations")
		//fmt.Println("  patch                   patch ports")
		fmt.Println("  prov                    search ports for files")
		fmt.Println("  pull                    pull in ports")
		fmt.Println("  sysup                   update outdated packages")
		fmt.Println("  uninstall               uninstall packages")
		fmt.Println("  help                    print help and exit")
		os.Exit(0)
	case "depends":
		cmd.Depends(os.Args[2:])
		os.Exit(0)
	case "diff":
		cmd.Diff(os.Args[2:])
		os.Exit(0)
	case "info":
		cmd.Info(os.Args[2:])
		os.Exit(0)
	case "install":
		cmd.Install(os.Args[2:])
		os.Exit(0)
	case "list":
		cmd.List(os.Args[2:])
		os.Exit(0)
	case "loc":
		cmd.Loc(os.Args[2:])
		os.Exit(0)
		//	case "patch":
		//		cmd.Patch(os.Args[2:])
		//		os.Exit(0)
	case "prov":
		cmd.Prov(os.Args[2:])
		os.Exit(0)
	case "pull":
		cmd.Pull(os.Args[2:])
		os.Exit(0)
	case "sysup":
		cmd.Sysup(os.Args[2:])
		os.Exit(0)
	case "uninstall":
		cmd.Uninstall(os.Args[2:])
		os.Exit(0)
	default:
		fmt.Fprintln(os.Stderr, "Invalid command, use help for a list of commands!")
		os.Exit(1)
	}
}
