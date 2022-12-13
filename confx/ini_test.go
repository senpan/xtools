package confx_test

import (
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/senpan/xtools/confx"
)

type IniStruct struct {
	Max     int           `ini:"max"`
	Port    string        `ini:"port"`
	Rate    float32       `ini:"rate"`
	Hosts   []string      `ini:"hosts" delim:" "`
	Timeout time.Duration `ini:"timeout"`
}

var _ = Describe("Ini", func() {
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
		It("empty", func() {
			conf := confx.GetConfByDefault("iniConfig", "empty", "empty")
			Expect(conf).Should(Equal("empty"))
		})

		It("empty2", func() {
			conf := confx.GetConfByDefault("iniConfig", "empty", "")
			Expect(conf).Should(Equal(""))
		})

		It("correct", func() {
			conf := confx.GetConfByDefault("iniConfig", "name", "")
			Expect(conf).Should(Equal("iniConfig"))
		})
	})

	Context("GetConfToSlice", func() {
		It("empty", func() {
			conf := confx.GetConfToSlice("iniConfig", "empty")
			Expect(len(conf)).Should(Equal(0))
		})

		It("correct", func() {
			conf := confx.GetConfToSlice("iniConfig", "hosts")
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
			data := confx.GetConfToMap("iniConfigStringMap")
			Expect(data["name"]).Should(Equal("iniConfig"))
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
			data := confx.GetConfToArrayMap("iniConfigArrayMap")
			Expect(len(data)).Should(Equal(1))
			Expect(data["name"][0]).Should(Equal("iniConfig1"))
			Expect(data["name"][1]).Should(Equal("iniConfig2"))
			Expect(data["empty"]).Should(BeNil())
		})
	})

	Context("ParseConfToStruct", func() {
		It("invalid", func() {
			var data IniStruct
			err := confx.ParseConfToStruct("iniConfigObject1", &data)
			Expect(err).Should(BeNil())
			Expect(data.Max).Should(Equal(0))
		})

		It("correct", func() {
			var data IniStruct
			err := confx.ParseConfToStruct("iniConfigObject", &data)
			Expect(err).Should(BeNil())
			Expect(data.Max).Should(Equal(101))
			Expect(data.Port).Should(Equal("9900"))
			Expect(data.Rate).Should(Equal(float32(1.01)))
			Expect(data.Hosts[0]).Should(Equal("127.0.0.1"))
			Expect(data.Hosts[1]).Should(Equal("127.0.0.2"))
			Expect(data.Timeout).Should(Equal(time.Second * 5))
		})
	})
})
