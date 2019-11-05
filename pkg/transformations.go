package pkg

import (
	"fmt"
)

// Xor performans an "exclusive or"-a-like subtraction, returning `0`-values
// when not found in both sides (or the same), and regular subtraction when
// present in both but not equal.
//
// a: {
//   "foo": {1},
//   "bar": {1},
// }
//
// b: {
//   "bar:  {4},
//   "caz": {4},
// }
//
// => {
//   "foo:  {0},
//   "bar:  {3},
//   "caz": {0},
// }
//
//
func Xor(a, b Profile) (res Profile) {
	res = Profile{}

	for kB, _ := range b {
		res[kB] = Values{}
	}

	for kA, vA := range a {
		vB, found := b[kA]
		if !found {
			res[kA] = Values{}
			continue
		}

		res[kA] = Values{
			vB[0] - vA[0],
			vB[1] - vA[1],
			vB[2] - vA[2],
			vB[3] - vA[3],
		}
	}

	return
}

// Filter takes care of filtering out functions that we don't care about.
//
// A function is considered "in" if specified in `fns`.
//
func Filter(src []Profile, fns ...string) (res []Profile) {
	res = []Profile{}

	for _, profile := range src {

		resProfile := Profile{}

		for _, fn := range fns {
			vals, found := profile[fn]
			if !found {
				continue
			}

			resProfile[fn] = vals
		}

		if len(resProfile) > 0 {
			res = append(res, resProfile)
		}
	}

	return
}

func Normalize(src []Profile) {
	all := map[string]struct{}{}

	for _, profile := range src {
		for k := range profile {
			all[k] = struct{}{}
		}
	}

	for k := range all {
		for _, profile := range src {
			_, found := profile[k]
			if !found {
				profile[k] = Values{}
			}
		}
	}

	return
}

func Delta(src []Profile) (res []Profile, err error) {
	if len(src) < 2 {
		err = fmt.Errorf("a minimum of 2 profiles is required")
		return
	}

	// for each pair -- `xor` it

	return
}
