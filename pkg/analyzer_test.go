package pkg_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	"github.com/cirocosta/heapvis/pkg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Analyzer", func() {

	Describe("ToCSV", func() {

		var (
			buf     *bytes.Buffer
			profile pkg.Profile
			lines   []string
		)

		BeforeEach(func() {
			buf = new(bytes.Buffer)
		})

		JustBeforeEach(func() {
			err := profile.ToCSV(buf)
			Expect(err).ToNot(HaveOccurred())

			lines = readLines(buf)
		})

		Context("empty profile", func() {
			It("writes just headers", func() {
				Expect(lines).To(HaveLen(1))

				header := lines[0]

				Expect(header).To(Equal(pkg.CSVHeader))
			})
		})

		Context("having samples", func() {

			BeforeEach(func() {
				profile = pkg.Profile{
					"fmt.Printf": [4]int64{
						1, 2, 3, 4,
					},
				}
			})

			It("writes headers + lines", func() {
				Expect(lines).To(HaveLen(2))

				firstRow := lines[1]
				Expect(firstRow).To(Equal(("fmt.Printf,1,2,3,4")))
			})
		})
	})

	Describe("LoadProfiles", func() {

		var (
			files    []string
			err      error
			profiles []pkg.Profile
		)

		JustBeforeEach(func() {
			profiles, err = pkg.LoadProfiles(files)
		})

		Context("with inexistent file", func() {
			BeforeEach(func() {
				files = []string{"dhsaiuehadsiuhj"}
			})

			It("fails", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("existing invalid profile", func() {
			BeforeEach(func() {
				files = []string{"testdata/invalid.txt"}
			})

			It("fails", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("with valid profile", func() {
			BeforeEach(func() {
				files = []string{"testdata/heap"}
			})

			It("succeeds", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(profiles).To(HaveLen(1))
			})

			It("captures profiling info", func() {
				profile := profiles[0]
				fmt.Printf("profile=%+v\n", profile)
			})
		})
	})

})

func readLines(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	res := []string{}

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	err := scanner.Err()
	if err != nil {
		panic(err)
	}

	return res
}
