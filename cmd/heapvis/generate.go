package main

import (
	"github.com/cirocosta/heapvis/pkg"
)

type generateCommand struct {
	Profile string `long:"profile" required:"true" description:"pprof profile to read"`
}

func (c *generateCommand) Execute(args []string) (err error) {
	_, err = pkg.LoadProfile(c.Profile)
	return
}
