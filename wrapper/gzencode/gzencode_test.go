package gzencode_test

import (
	"fmt"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cwxstat/rabbitmq/wrapper/gzencode"
)

func ReadFile(file string) (string, error) {
	dat, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

var _ = Describe("Gzencode", func() {

	var (
		dir         string
		desDir      string
		dataWritten string
	)

	BeforeEach(func() {
		dir = "/tmp/compress"
		desDir = "/tmp/junk"
		dataWritten = fmt.Sprintf("data written: %v", time.Now())
		os.MkdirAll(dir+"/plus/stuff", os.FileMode(0755))
		os.MkdirAll(dir+"/plus/stuff2/two2", os.FileMode(0755))
		os.MkdirAll(desDir, os.FileMode(0755))
		os.WriteFile(dir+"/plus/stuff/file", []byte(dataWritten), 0644)
		os.WriteFile(dir+"/plus/stuff2/two2/file2", []byte(dataWritten), 0644)

	})

	AfterEach(func() {
		os.RemoveAll(dir)
		os.RemoveAll(desDir)
		os.RemoveAll("/tmp/compress.tar.gz")
		os.RemoveAll("/tmp/junk.tar.gz")
		os.RemoveAll("/tmp/junk")
	})

	Describe("Check genz", func() {

		Context("Should take 3 things: dir, desDir, handleFile", func() {
			It("should ...", func() {

				g := gzencode.NewGZ()
				g.CertPath("../../etc/certs")
				g.DirIn(dir)

				err := g.Produce()
				Expect(err).To(BeNil())

				// In case it picks up files....
				os.RemoveAll("/tmp/compress.tar.gz")
				os.RemoveAll("/tmp/junk.tar.gz")

				// Consume part
				g.DestDir(desDir)
				g.ConsumerQ("cq")
				g.HandleFile("/tmp/junk3.tar.gz", "/tmp/junk")
				err = g.Consume()

				time.Sleep(2 * time.Second)

				Expect(err).To(BeNil())

				result, err := ReadFile("/tmp/junk/plus/stuff2/two2/file2")
				Expect(err).To(BeNil())
				Expect(result).To(Equal(dataWritten))

			})
		})

	})

})
