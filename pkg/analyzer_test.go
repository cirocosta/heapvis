package pkg_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cirocosta/heapvis/pkg"
)

var _ = Describe("Analyzer", func() {

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
		})
	})

})
