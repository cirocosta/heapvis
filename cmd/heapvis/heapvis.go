package main

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

var heapvis struct {
	Generate generateCommand `command:"generate"`
}

func writer(value string) (w io.Writer, err error) {
	var file *os.File

	if value == "-" {
		w = os.Stdout
		return
	}

	file, err = os.Create(value)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to create file %s", value)
		return
	}

	w = file
	return
}

func reader(value string) (r io.Reader, err error) {
	var file *os.File

	if value == "-" {
		r = os.Stdin
		return
	}

	file, err = os.Open(value)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to open dpkg status file at %s", value)
		return
	}

	r = file
	return
}
