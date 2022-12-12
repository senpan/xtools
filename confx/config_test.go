package confx_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/senpan/xtools/confx"
	"github.com/senpan/xtools/flagx"
)

// For multiple profiles, use the Ordered option
var _ = Describe("Config", Ordered, func() {
	Context("GetConf", func() {
		BeforeEach(func() {
			confx.Reset()
			flagx.SetConfig("testdata/conf.ini")
		})

		It("empty", func() {
			conf := confx.GetConf("iniConfig", "empty")
			Expect(conf).Should(Equal(""))
		})

		It("GetConf", func() {
			conf := confx.GetConf("iniConfig", "name")
			Expect(conf).Should(Equal("iniConfig"))
		})
	})

	Context("GetConfByDefault", func() {
		BeforeEach(func() {
			confx.Reset()
			flagx.SetConfig("testdata/conf.yaml")
		})

		It("empty", func() {
			conf := confx.GetConfByDefault("Yaml", "empty", "empty")
			Expect(conf).Should(Equal("empty"))
		})

		It("empty2", func() {
			conf := confx.GetConfByDefault("Yaml", "empty", "")
			Expect(conf).Should(Equal(""))
		})

		It("correct", func() {
			conf := confx.GetConfByDefault("Yaml", "name", "")
			Expect(conf).Should(Equal("yaml"))
		})
	})
})
