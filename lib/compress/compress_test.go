package compress_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"

	"github.com/cwxstat/rabbitmq/lib/compress"
)

var _ = Describe("Compress", func() {

	var (
		dir         string
		desDir      string
		dataWritten string
	)

	BeforeEach(func() {
		dir = "/tmp/compress"
		desDir = "/tmp/junk"
		dataWritten = "\ndata written\n"
		os.MkdirAll(dir+"/plus/stuff", os.FileMode(0755))
		os.MkdirAll(dir+"/plus/stuff2/two2", os.FileMode(0755))
		os.MkdirAll("/tmp/junk", os.FileMode(0755))
		os.WriteFile(dir+"/plus/stuff/file", []byte(dataWritten), 0644)
		os.WriteFile(dir+"/plus/stuff2/two2/file2", []byte(dataWritten), 0644)

	})

	AfterEach(func() {
		os.RemoveAll(dir)
		os.RemoveAll("/tmp/junk")
		os.RemoveAll("/tmp/compress.tar.gz")
	})

	Describe("Check compress", func() {

		Context("Should compress", func() {
			It("should ...", func() {

				// Note remove prefix is "/tmp/compress"
				err := compress.Compress(dir, "/tmp/compress")
				Expect(err).To(BeNil())
				err = compress.UnCompress("/tmp/compress.tar.gz", desDir)
				Expect(err).To(BeNil())
				result, err := os.ReadFile("/tmp/junk/plus/stuff2/two2/file2")
				Expect(err).To(BeNil())
				Expect(string(result)).To(Equal(dataWritten))

			})
		})

	})

})
