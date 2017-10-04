package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go2c/optparse"
	colorful "github.com/lucasb-eyer/go-colorful"
)

// graph generates a dependency grap.
func graph(input []string) error {
	// Define valid arguments.
	o := optparse.New()
	argd := o.Bool("duplicate", 'd', false)
	argn := o.Bool("no-alias", 'n', false)
	argt := o.String("type", 't', "svg")
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	_, err := o.Parse(input)
	if err != nil {
		return fmt.Errorf("invaild argument, use -h for a list of arguments")
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt graph [arguments]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -d,   --duplicate       graph duplicates as well")
		fmt.Println("  -n,   --no-alias        disable aliasing")
		fmt.Println("  -t,   --type            filetype to use")
		fmt.Println("  -h,   --help            print help and exit")

		return nil
	}

	// Get all ports.
	all, err := ports()
	if err != nil {
		return err
	}

	var p port
	p.Location = "."
	if err := p.parsePkgfile(); err != nil {
		return err
	}
	p.depends(!*argn, all)

	// Set file to write to.
	f, err := os.OpenFile(p.Pkgfile.Name+".dot",
		os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	if err != nil {
		return err
	}

	// Prettify graph.
	fmt.Fprintf(f, "digraph G {\n")
	fmt.Fprintf(f, "\tgraph [\n")
	fmt.Fprintf(f, "\t\t%s=\"%s\"\n", "tcenter", "true")
	fmt.Fprintf(f, "\t\t%s=\"%f\"\n", "pad", 2.0)
	fmt.Fprintf(f, "\t]\n\n")

	// Prettify nodes.
	fmt.Fprint(f, "\tnode [\n")
	fmt.Fprintf(f, "\t\t%s=\"%s\"\n", "constraint", "false")
	fmt.Fprintf(f, "\t\t%s=\"%s\"\n", "fontcolor", "#111e38")
	fmt.Fprintf(f, "\t\t%s=\"%d\"\n", "penwidth", 3)
	fmt.Fprintf(f, "\t\t%s=\"%s\"\n", "shape", "box")
	fmt.Fprintf(f, "\t]\n\n")

	// Prettify edges.
	fmt.Fprintf(f, "\tedge [\n")
	fmt.Fprintf(f, "\t\t%s=\"%s\"\n", "arrowhead", "dot")
	fmt.Fprintf(f, "\t\t%s=\"%s\"\n", "color", "#cee0e3")
	fmt.Fprintf(f, "\t\t%s=\"%s\"\n", "headport", "n")
	fmt.Fprintf(f, "\t\t%s=\"%d\"\n", "penwidth", 2)
	fmt.Fprintf(f, "\t]\n\n")

	var c []string
	op := p.getBaseDir()
	pl := p.Depends
	pal, _ := colorful.SoftPalette(128)
	var i int
	var recursive func()
	recursive = func() {
		for _, p := range pl {
			if !stringInList(p.Pkgfile.Name, c) {
				fmt.Fprintf(f, "\tnode [color=\"%s\"]\n", pal[i].Hex())
				fmt.Fprintf(f, "\t\"%s\"->\"%s\"\n", op, p.getBaseDir())

				// Append to checked ports.
				if !*argd {
					c = append(c, p.Pkgfile.Name)
				}
			}

			if len(p.Depends) > 0 {
				pl = p.Depends
				recursive()
			}

			i++
			if i > 128 {
				i = 0
			}
			op = p.getBaseDir()
		}
	}
	recursive()
	fmt.Fprintln(f, "}")

	f.Close()
	if *argt == "dot" {
		return nil
	}

	// Convert to graph.
	cmd := exec.Command("dot", p.Pkgfile.Name+".dot", "-T", *argt, "-o",
		p.Pkgfile.Name+"."+*argt)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("graphviz dot %s: Something went wrong",
			p.Pkgfile.Name+"."+*argt)
	}

	return nil
}
