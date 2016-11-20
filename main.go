package main

import (
	"fmt"
	"os"

	"github.com/onodera-punpun/prt/commands"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Missing command, use help for a list of commands.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help":
		fmt.Println("Usage: prt command [arguments]")
		fmt.Println("")
		fmt.Println("commands:")
		fmt.Println("  depends                 list dependencies recursivly")
		fmt.Println("  diff                    list outdated packages")
		fmt.Println("  build                   build and install ports")
		fmt.Println("  info                    print port information")
		fmt.Println("  list                    list ports")
		fmt.Println("  location                prints port locations")
		fmt.Println("  patch                   patch ports")
		fmt.Println("  provide                 search ports for files")
		fmt.Println("  remove                  remove installed ports")
		fmt.Println("  pull                    pull in ports")
		fmt.Println("  sysup                   update outdated packages")
		fmt.Println("  help                    print help and exit")
		os.Exit(0)
	case "depends":
		commands.Depends(os.Args[1:])
		os.Exit(0)
		//	case "build":
		//		commands.Build(os.Args[1:])
		//		os.Exit(0)
	case "info":
		commands.Info(os.Args[1:])
		os.Exit(0)
		//	case "list":
		//		commands.List(os.Args[1:])
		//		os.Exit(0)
		//	case "location":
		//		commands.Location(os.Args[1:])
		//		os.Exit(0)
		//	case "patch":
		//		commands.Patch(os.Args[1:])
		//		os.Exit(0)
		//	case "provide":
		//		commands.Provide(os.Args[1:])
		//		os.Exit(0)
		//	case "remove":
		//		commands.Remove(os.Args[1:])
		//		os.Exit(0)
		//	case "pull":
		//		commands.Pull(os.Args[1:])
		//		os.Exit(0)
		//	case "sysup":
		//		commands.Sysup(os.Args[1:])
		//		os.Exit(0)
	default:
		fmt.Println("Invalid command, use help for a list of commands.")
		os.Exit(1)
	}
}
