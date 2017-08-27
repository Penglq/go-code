package gotest_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGotest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gotest Suite")
}
