package handle

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Handle", func() {

	var (
		h       HS
		destDir = "/tmp/handleTest"
	)

	BeforeEach(func() {
		os.RemoveAll(destDir)
		h = HS{DestDir: destDir}

	})

	AfterEach(func() {
		os.RemoveAll(destDir)

	})

	Describe("Check private functions", func() {

		Context("Adding fist and second", func() {

			It("should create successfully", func() {
				Expect(h.createDirIfNotExist()).To(BeNil())

				_, err := os.Stat(h.DestDir)
				Expect(err).To(BeNil())
			})

		})

	})
})
