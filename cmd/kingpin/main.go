package main

import (
	"errors"

	"github.com/alecthomas/kingpin"
)

var (
	foo bool
	bar bool
	cmd = kingpin.CommandLine
)

func init() {
	cmd.Flag("foo", "foo").BoolVar(&foo)
	cmd.Flag("bar", "bar").BoolVar(&bar)
	cmd.Validate(func(_ *kingpin.Application) error {
		if foo && bar {
			return errors.New("cannot process with both --foo and --bar options")
		}
		return nil
	})
}

func main() {

	kingpin.MustParse(cmd.Parse([]string{"--foo", "--bar"}))
}
