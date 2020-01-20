package main

import (
	"errors"

	"github.com/alecthomas/kingpin"
)

func main() {
	var (
		foo bool
		bar bool
	)
	app := kingpin.New("kingpin", "tryout kingpin package")
	app.Flag("foo", "foo").BoolVar(&foo)
	app.Flag("bar", "bar").BoolVar(&bar)
	app.Validate(func(_ *kingpin.Application) error {
		if foo && bar {
			return errors.New("cannot process with both --foo and --bar options")
		}
		return nil
	})

	kingpin.MustParse(app.Parse([]string{"--foo", "--bar"}))
}
