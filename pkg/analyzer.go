package pkg

import (
	"fmt"
	"io"
	"os"

	pprof "github.com/google/pprof/profile"
)

type SampleType uint8

const (
	AllocObjects SampleType = iota // alloc_objects
	AllocSpace                     // alloc_space
	InUseObjects                   // inuse_objects
	InUseSpace                     // inuse_space
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
	Profile map[string][4]int64
)

const (
	CSVHeader = "fn,alloc_objects,alloc_space,inuse_objects,inuse_space"
)

func ToCSV(w io.Writer, profiles []Profile) (err error) {
	_, err = io.WriteString(w, CSVHeader+"\n")
	if err != nil {
		return
	}

	for _, profile := range profiles {
		for fn, vals := range profile {
			_, err = fmt.Fprintf(w, "%s,%d,%d,%d,%d\n", fn,
				vals[0], vals[1], vals[2], vals[3],
			)

			if err != nil {
				return
			}
		}
	}

	return
}

// FromPprof converts a pprof memory profile to `Profile`.
//
func FromPprof(src *pprof.Profile) (profile Profile, err error) {
	profile = Profile{}

	for _, sample := range src.Sample {

		var (
			fn       = sample.Location[0].Line[0].Function.Name
			existing = profile[fn]
		)

		profile[fn] = [4]int64{
			existing[0] + sample.Value[0],
			existing[1] + sample.Value[1],
			existing[2] + sample.Value[2],
			existing[3] + sample.Value[3],
		}
	}

	return
}

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
