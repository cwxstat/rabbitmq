package gzencode_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGzencode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gzencode Suite")
}
