package pkg_test

import (
	"github.com/cirocosta/heapvis/pkg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("transformations", func() {

	DescribeTable("Xor",
		func(a, b, res pkg.Profile) {
			Expect(pkg.Xor(a, b)).To(Equal(res))
		},
		Entry("empty",
			pkg.Profile{},
			pkg.Profile{},
			pkg.Profile{},
		),
		Entry("only in A",
			pkg.Profile{"fn": pkg.Values{1, 2}},
			pkg.Profile{},
			pkg.Profile{"fn": pkg.Values{0, 0}},
		),
		Entry("only in B",
			pkg.Profile{},
			pkg.Profile{"fn": pkg.Values{1, 2}},
			pkg.Profile{"fn": pkg.Values{0, 0}},
		),
		Entry("in both",
			pkg.Profile{"fn": pkg.Values{1, 1}},
			pkg.Profile{"fn": pkg.Values{3, 3}},
			pkg.Profile{"fn": pkg.Values{2, 2}},
		),
	)

	FDescribe("filter", func() {

		type filterTest struct {
			src, res []pkg.Profile
			fns      []string
		}

		DescribeTable("when",
			func(tc filterTest) {
				Expect(
					pkg.Filter(tc.src, tc.fns...),
				).To(Equal(tc.res))
			},
			Entry("no fns, all out", filterTest{
				src: []pkg.Profile{
					{"fn1": pkg.Values{1, 2}},
					{"fn2": pkg.Values{1, 2}},
				},
				res: []pkg.Profile{},
				fns: []string{},
			}),
			Entry("inexistent fn, all out", filterTest{
				src: []pkg.Profile{
					{"fn1": pkg.Values{1, 2}},
					{"fn2": pkg.Values{1, 2}},
				},
				res: []pkg.Profile{},
				fns: []string{"dsaiuhsi"},
			}),
			Entry("w/ fn, rest out", filterTest{
				src: []pkg.Profile{
					{"fn1": pkg.Values{1}},
					{"fn2": pkg.Values{2}},
					{"fn3": pkg.Values{3}},
				},
				res: []pkg.Profile{
					{"fn2": pkg.Values{2}},
				},
				fns: []string{"fn2"},
			}),
			Entry("w/ fn in multiple profiles, rest out", filterTest{
				src: []pkg.Profile{
					{
						"fn1": pkg.Values{1},
						"fn2": pkg.Values{2},
					},
					{"fn2": pkg.Values{22}},
					{"fn3": pkg.Values{3}},
				},
				res: []pkg.Profile{
					{"fn2": pkg.Values{2}},
					{"fn2": pkg.Values{22}},
				},
				fns: []string{"fn2"},
			}),
			Entry("w/ multiple fn, rest out", filterTest{
				src: []pkg.Profile{
					{"fn1": pkg.Values{1}},
					{"fn2": pkg.Values{2}},
					{"fn3": pkg.Values{3}},
				},
				res: []pkg.Profile{
					{"fn2": pkg.Values{2}},
					{"fn3": pkg.Values{3}},
				},
				fns: []string{"fn2", "fn3"},
			}),
			Entry("w/ multiple fn in multiple profiles, rest out", filterTest{
				src: []pkg.Profile{
					{"fn1": pkg.Values{1}},
					{
						"fn2":   pkg.Values{2},
						"fn3":   pkg.Values{3},
						"fnFOO": pkg.Values{99},
					},
					{
						"fn2":   pkg.Values{22},
						"fn3":   pkg.Values{33},
						"fnFOO": pkg.Values{99},
					},
				},
				res: []pkg.Profile{
					{
						"fn2": pkg.Values{2},
						"fn3": pkg.Values{3},
					},
					{
						"fn2": pkg.Values{22},
						"fn3": pkg.Values{33},
					},
				},
				fns: []string{"fn2", "fn3"},
			}),
		)
	})

	Describe("Delta", func() {

		var (
			src, profiles []pkg.Profile
			err           error
		)

		JustBeforeEach(func() {
			profiles, err = pkg.Delta(src)
		})

		Context("no profiles", func() {
			It("fails", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("1 profile", func() {
			It("fails", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("2 or more profiles", func() {
			Context("having all fns matching", func() {
				BeforeEach(func() {
					src = []pkg.Profile{
						{"fn1": pkg.Values{1, 1, 1, 1}},
						{"fn1": pkg.Values{4, 4, 4, 4}},
					}
				})

				It("succeeds", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("computes the deltas for each fn on each value", func() {
					Expect(profiles).To(ConsistOf([]pkg.Profile{
						{
							"fn1": pkg.Values{3, 3, 3, 3},
						},
					}))
				})

				It("reduces profile count by 1", func() {
					Expect(profiles).To(HaveLen(len(src) - 1))
				})
			})
		})
	})

})
