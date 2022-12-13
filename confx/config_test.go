package confx_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/senpan/xtools/confx"
)

// For multiple config files, use the Ordered option
var _ = Describe("Config", Ordered, func() {
	BeforeEach(func() {
		confx.Reset()
		confx.InitConfig(confx.WithConfig("testdata/conf.ini"), confx.WithConfigPathPrefix(pwd))
	})

	Context("GetConf", func() {
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
			confx.InitConfig(confx.WithConfig("testdata/conf.yaml"), confx.WithConfigPathPrefix(pwd))
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

	Context("Ini Env", func() {
		BeforeEach(func() {
			confx.Reset()
			_ = os.Setenv("XTOOLS_ENV", "INI")
			confx.InitConfig(confx.WithEnv(), confx.WithConfig("testdata/conf.ini"), confx.WithConfigPathPrefix(pwd))
		})

		It("correct", func() {
			confx.Reset()
			_ = os.Setenv("XTOOLS_ENV", "INI")
			confx.InitConfig(confx.WithEnv(), confx.WithConfig("testdata/conf.ini"), confx.WithConfigPathPrefix(pwd))
			conf := confx.GetConfByDefault("iniEnv", "env", "")
			Expect(conf).Should(Equal("INI"))
		})

		It("correct2", func() {
			confx.Reset()
			_ = os.Unsetenv("XTOOLS_ENV")
			confx.InitConfig(confx.WithEnv(), confx.WithConfig("testdata/conf.ini"), confx.WithConfigPathPrefix(pwd))
			conf := confx.GetConfByDefault("iniEnv", "env", "")
			Expect(conf).Should(Equal(""))
		})

		It("correct3", func() {
			confx.Reset()
			confx.InitConfig(confx.WithEnv(), confx.WithConfig("testdata/conf.ini"), confx.WithConfigPathPrefix(pwd))
			conf := confx.GetConfByDefault("iniEnv", "env2", "")
			Expect(conf).Should(Equal(""))
		})

		It("correct4", func() {
			confx.Reset()
			_ = os.Setenv("XTOOLS_ENV", "INI")
			confx.InitConfig(confx.WithConfig("testdata/conf.ini"), confx.WithConfigPathPrefix(pwd))
			conf := confx.GetConfByDefault("iniEnv", "env", "")
			Expect(conf).Should(Equal("${XTOOLS_ENV}"))
		})
	})

	Context("Yaml Env", func() {
		It("correct", func() {
			confx.Reset()
			_ = os.Setenv("XTOOLS_ENV", "YAML")
			confx.InitConfig(confx.WithEnv(), confx.WithConfig("testdata/conf.yaml"), confx.WithConfigPathPrefix(pwd))
			conf := confx.GetConfByDefault("YamlEnv", "env", "")
			Expect(conf).Should(Equal("YAML"))
		})

		It("correct2", func() {
			confx.Reset()
			_ = os.Unsetenv("XTOOLS_ENV")
			confx.InitConfig(confx.WithEnv(), confx.WithConfig("testdata/conf.yaml"), confx.WithConfigPathPrefix(pwd))
			conf := confx.GetConfByDefault("YamlEnv", "env", "")
			Expect(conf).Should(Equal(""))
		})

		It("correct3", func() {
			confx.Reset()
			_ = os.Setenv("XTOOLS_ENV", "INI")
			confx.InitConfig(confx.WithEnv(), confx.WithConfig("testdata/conf.yaml"), confx.WithConfigPathPrefix(pwd))
			conf := confx.GetConfByDefault("YamlEnv", "env2", "")
			Expect(conf).Should(Equal(""))
		})

		It("correct4", func() {
			confx.Reset()
			confx.InitConfig(confx.WithConfig("testdata/conf.yaml"), confx.WithConfigPathPrefix(pwd))
			conf := confx.GetConfByDefault("YamlEnv", "env", "")
			Expect(conf).Should(Equal("${XTOOLS_ENV}"))
		})
	})

	AfterAll(func() {
		_ = os.Unsetenv("XTOOLS_ENV")
	})
})
