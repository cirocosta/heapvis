package pkg

import (
	"fmt"
	"os"

	"github.com/google/pprof/profile"
)

type (
	Location  string
	ByteCount uint64
)

func LoadProfile(file string) (p Profile, err error) {
	f, err := os.Open(file)
	if err != nil {
		err = fmt.Errorf("failed to read profile file %s: %w", file, err)
		return
	}

	defer f.Close()

	p, err = profile.Parse(f)
	if err != nil {
		err = fmt.Errorf(err, "failed parsing profile from file %s: %w",
			file, err)
		return
	}

	return
}
