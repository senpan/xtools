package confx_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/senpan/xtools/confx"
)

func TestConfx(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Confx Suite")
}

var _ = BeforeSuite(func() {
	confx.SetConfPathPrefixByPwd()
})
