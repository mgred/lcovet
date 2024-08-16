package main

import (
	"fmt"
	"os"

	"github.com/mgred/lcovet/internal/lcovet"
)

const APPNAME = "lcovet"

func main() {
	c := NewConfig()
	flags := c.FlagSet(APPNAME)
	flags.Usage = usage
	flags.Parse(os.Args[1:])

	if c.Help {
		usage()
		os.Exit(0)
	}

	if c.Version {
		fmt.Printf("%s: %s\n", APPNAME, GetVersion())
		os.Exit(0)
	}

  output, err := c.GetOutput()
  handleError(err)
  defer func() {
    handleError(output.Close())
  }()

	parser := lcovet.NewParser(os.Stdin)
	result := parser.Parse()
  formatter := lcovet.NewFormatter(result)

  if c.Html {
    formatter.Html(output)
  } else {
    formatter.Simple(output)
  }
}

func handleError(e error) {
  if e != nil {
    fmt.Printf("ERROR %v", e)
    os.Exit(3)
  }
}

func usage() {
	fmt.Fprintf(os.Stdout, `%[1]s [options]

SYNOPSIS
    %[1]s [-html] [-json] [-output=<file>] [<path>...]

OPTIONS
    -json             Format to JSON
    -html             Format to HTML
    -output           File to write output to
    -help             Print help message
    -version          Print version

`, APPNAME)
}
