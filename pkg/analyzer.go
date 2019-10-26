package pkg

import (
	"fmt"
	"os"

	pprof "github.com/google/pprof/profile"
)

type SampleType uint8

const (
	AllocObjects SampleType = iota + 1 // alloc_objjects
	AllocSpace                         // alloc_space
	InUseObjects                       // inuse_objects
	InUseSpace                         // inuse_space
)

// Represents the visualization of a profile, s
//
//   Showing top 10 nodes out of 27
//         flat  flat%   sum%        cum   cum%
//        5376B 76.10% 76.10%      5376B 76.10%  runtime.malg
//        1024B 14.50% 90.60%      1792B 25.37%  runtime.allocm
//         256B  3.62% 94.22%       256B  3.62%  runtime.allgadd
//         192B  2.72% 96.94%       288B  4.08%  runtime.gcBgMarkWorker
//         104B  1.47% 98.41%       104B  1.47%  os.newFile
//          96B  1.36% 99.77%        96B  1.36%  runtime.acquireSudog
//            0     0% 99.77%       120B  1.70%  main.main
//            0     0% 99.77%       120B  1.70%  os.Create
//            0     0% 99.77%       120B  1.70%  os.OpenFile
//            0     0% 99.77%       120B  1.70%  os.openFileNolog
//

type (
	// Profile represents a summarized representation of a pprof capture.
	//
	// It maps a given source line location to a statistic.
	//
	Profile struct {
		Type SampleType
		Data map[string]int64
	}
)

// LoadProfiles ...
//
func LoadProfiles(files []string) (profiles []Profile, err error) {
	var (
		pProfile *pprof.Profile
		profile  Profile
	)

	for _, file := range files {
		pProfile, err = loadPprofProfile(file)
		if err != nil {
			err = fmt.Errorf("failed to load pprof profile %s: %w", file, err)
			return
		}

		profile, err = FromPprof(pProfile)
		if err != nil {
			err = fmt.Errorf("failed to convert from pprof to internal format: %w",
				err)
			return
		}

		profiles = append(profiles, profile)
	}

	return
}

// FromPprof converts a pprof memory profile to `Profile`.
//k
//
func FromPprof(src *pprof.Profile) (profile Profile, err error) {
	profile.Data = map[string]int64{}

	var fn string

	for _, sample := range src.Sample {
		fn = sample.Location[0].Line[0].Function.Name
		profile.Data[fn] = profile.Data[fn] + sample.Value[0]
	}

	return
}

// loadPprofProfile ...
//
func loadPprofProfile(file string) (profile *pprof.Profile, err error) {
	f, err := os.Open(file)
	if err != nil {
		err = fmt.Errorf("failed to read profile file %s: %w", file, err)
		return
	}

	defer f.Close()

	profile, err = pprof.Parse(f)
	if err != nil {
		err = fmt.Errorf("failed parsing profile from file %s: %w",
			file, err)
		return
	}

	return
}
