package handle_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHandle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handle Suite")
}
