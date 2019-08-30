package pkg

import (
	"os"

	"github.com/google/pprof/profile"
	"github.com/pkg/errors"
)

type (
	Location  string
	ByteCount uint64
)

type Profile struct {
	Points map[Location]Count
}

func LoadProfiles(files []string) (profiles []*profile.Profile, err error) {
	profiles = make([]*profile.Profile, len(files))

	// TODO perform this concurrently?
	//
	for idx, file := range files {
		profiles[idx], err = loadProfile(file)
		if err != nil {
			return
		}
	}

	return
}

func loadProfile(file string) (p Profile, err error) {
	f, err := os.Open(file)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to read profile file %s", file)
		return
	}

	defer f.Close()

	p, err = profile.Parse(f)
	if err != nil {
		err = errors.Wrapf(err,
			"failed parsing profile from file %s", file)
		return
	}

	return
}

func NewProfile(p *profile.Profile) Profile {
	return
}
