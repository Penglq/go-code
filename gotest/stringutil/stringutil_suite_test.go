package stringutil_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStringutil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stringutil Suite")
}
