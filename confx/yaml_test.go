package confx_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/senpan/xtools/confx"
)

type YamlStruct struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var _ = Describe("Yaml", func() {
	BeforeEach(func() {
		confx.Reset()
		confx.InitConfig(confx.WithConfig("testdata/conf.yaml"), confx.WithConfigPathPrefix(pwd))
	})

	Context("GetConf", func() {
		It("empty", func() {
			conf := confx.GetConf("Yaml", "empty")
			Expect(conf).Should(Equal(""))
		})

		It("GetConf", func() {
			conf := confx.GetConf("Yaml", "name")
			Expect(conf).Should(Equal("yaml"))
		})
	})

	Context("GetConfByDefault", func() {
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

	Context("GetConfToSlice", func() {
		It("empty", func() {
			conf := confx.GetConfToSlice("Yaml", "empty")
			Expect(len(conf)).Should(Equal(0))
		})

		It("correct", func() {
			conf := confx.GetConfToSlice("Yaml", "hosts")
			Expect(len(conf)).Should(Equal(3))
			Expect(strings.Join(conf, ",")).Should(Equal("127.0.0.1,127.0.0.2,127.0.0.3"))
		})
	})

	Context("GetConfToMap", func() {
		It("empty", func() {
			data := confx.GetConfToMap("empty")
			Expect(len(data)).Should(Equal(0))
		})

		It("correct", func() {
			data := confx.GetConfToMap("YamlStringMap")
			Expect(data["name"]).Should(Equal("yaml"))
			Expect(data["host"]).Should(Equal("127.0.0.1"))
			Expect(data["empty"]).Should(Equal(""))
		})
	})

	Context("GetConfToArrayMap", func() {
		It("empty", func() {
			data := confx.GetConfToArrayMap("empty")
			Expect(len(data)).Should(Equal(0))
		})

		It("correct", func() {
			data := confx.GetConfToArrayMap("YamlArrayMap")
			Expect(len(data)).Should(Equal(1))
			Expect(data["name"][0]).Should(Equal("yaml1"))
			Expect(data["name"][1]).Should(Equal("yaml2"))
			Expect(data["empty"]).Should(BeNil())
		})
	})

	Context("ParseConfToStruct", func() {
		It("invalid", func() {
			var data []YamlStruct
			err := confx.ParseConfToStruct("YamlObject1", &data)
			Expect(err).Should(BeNil())
			Expect(len(data)).Should(Equal(0))
		})

		It("correct", func() {
			var data []YamlStruct
			err := confx.ParseConfToStruct("YamlObject", &data)
			Expect(err).Should(BeNil())
			Expect(len(data)).Should(Equal(3))
			Expect(data[0].Host).Should(Equal("public1"))
			Expect(data[0].Port).Should(Equal("9092"))
			Expect(data[1].Host).Should(Equal("public2"))
			Expect(data[1].Port).Should(Equal("9093"))
			Expect(data[2].Host).Should(Equal("public3"))
			Expect(data[2].Port).Should(Equal("9094"))
		})
	})
})
