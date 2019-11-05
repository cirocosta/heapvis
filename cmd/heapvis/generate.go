package main

import (
	"fmt"
	"path/filepath"

	"github.com/cirocosta/heapvis/pkg"
)

type generateCommand struct {
	Functions []string `long:"function"                description:"fns to filter by"`
	Output    string   `long:"output"                  default:"-"`
	Paths     []string `long:"profile" required:"true" description:"pprof profile to read"`

	ShowFuncs bool `long:"show-funcs"`
}

func (c *generateCommand) Execute(args []string) (err error) {

	var (
		paths   = []string{}
		matches []string
	)

	for _, path := range c.Paths {
		matches, err = filepath.Glob(path)
		if err != nil {
			return
		}

		paths = append(paths, matches...)
	}

	profiles, err := pkg.LoadProfiles(paths)
	if err != nil {
		return
	}

	pkg.Normalize(profiles)

	if len(c.Functions) > 0 {
		profiles = pkg.Filter(profiles, c.Functions...)
	}

	w, err := writer(c.Output)
	if err != nil {
		return
	}

	if c.ShowFuncs {
		funcs := map[string]struct{}{}

		for _, profile := range profiles {
			for k := range profile {
				funcs[k] = struct{}{}
			}
		}

		for k := range funcs {
			fmt.Fprintf(w, "%s\n", k)
		}

		return
	}

	err = pkg.ToCSV(w, profiles)
	return

}
