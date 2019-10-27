package main

import (
	"path/filepath"

	"github.com/cirocosta/heapvis/pkg"
)

type generateCommand struct {
	Paths  []string `long:"profile" required:"true" description:"pprof profile to read"`
	Output string   `long:"output" default:"-"`
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

	w, err := writer(c.Output)
	if err != nil {
		return
	}

	err = pkg.ToCSV(w, profiles)
	return

}
