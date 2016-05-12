package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/lmroz/goldi/goldigen"
)

var _ = Describe("Config", func() {
	Describe("NewConfig", func() {
		It("should set the default type registration function name", func() {
			config := main.NewConfig("package_name", "", "", "")
			Expect(config.FunctionName).To(Equal(main.DefaultFunctionName))
		})

		It("should panic if the package name is empty", func() {
			Expect(func() { main.NewConfig("", "", "", "") }).To(Panic())
		})
	})

	Describe("PackageName", func() {
		It("should only return the package name", func() {
			config := main.NewConfig("github.com/fgrosse/servo", "", "", "")
			Expect(config.Package).To(Equal("github.com/fgrosse/servo"))
			Expect(config.PackageName()).To(Equal("servo"))
		})
	})

	Describe("OutputName", func() {
		It("should return the output file base bane", func() {
			config := main.NewConfig("github.com/fgrosse/servo", "", "/home/fgrosse/goldi/config/types.yml", "/home/fgrosse/goldi/types.go")
			Expect(config.OutputName()).To(Equal("types.go"))
		})
	})

	Describe("InputName", func() {
		It("should return the input file name relative to the output file", func() {
			config := main.NewConfig("github.com/fgrosse/servo", "", "/home/fgrosse/goldi/config/types.yml", "/home/fgrosse/goldi/types.go")
			Expect(config.InputName()).To(Equal("config/types.yml"))
		})
	})
})
