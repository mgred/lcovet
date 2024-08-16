package main

import (
	"flag"
	"os"
)

type Config struct {
	Help    bool
	Version bool
  Output  string
  Html    bool
}

func NewConfig() *Config {
	return &Config{}
}

// Returns the file to write the output to.
//
// The file that's returned is either a file given by the
// `-output` option or stdout, which is the default.
func (c *Config) GetOutput() (file *os.File, err error) {
  file = os.Stdout
  if c.Output != "" {
    file, err = os.Create(c.Output)
    if err != nil {
      return
    }
  }
  return
}

func (c *Config) FlagSet(name string) *flag.FlagSet {
	flags := flag.NewFlagSet(name, flag.ExitOnError)
	flags.BoolVar(&c.Help, "help", false, "print usage information")
	flags.BoolVar(&c.Html, "html", false, "print as html")
	flags.BoolVar(&c.Version, "version", false, "print version information")
	flags.StringVar(&c.Output, "output", "", "output file")
	return flags
}
