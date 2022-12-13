package confx_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var pwd string

func TestConfx(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Confx Suite")
}

var _ = BeforeSuite(func() {
	pwd, _ = os.Getwd()
})
