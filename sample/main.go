// generates a sample profile that shows high mem allocation in `fn`.
//

package main

import (
	"flag"
	"io"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var (
	profile = flag.String("profile", "profile.pprof", "file to write the profile to")
)

func init() {
	// From https://golang.org/pkg/runtime/#pkg-variables:
	//
	//   > "The profiler aims to sample an average of one allocation per
	//   >  MemProfileRate bytes allocated."
	//
	// ps.: could be set using `GODEBUG` too with `memprofilerate`.
	//
	runtime.MemProfileRate = 1
}

func fn(w io.Writer) {
	dumbArray := []string{}
	for i := 0; i < 1000; i++ {
		dumbArray = append(dumbArray, "adhiaudhsaiu")
	}

	runtime.GC()

	err := pprof.WriteHeapProfile(w)
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()

	if *profile == "" {
		log.Fatalf("must specify `-profile`")
		return
	}

	f, err := os.Create(*profile)
	if err != nil {
		panic(err)
	}

	fn(f)

	defer f.Close()

}
