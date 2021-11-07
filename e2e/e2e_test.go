package e2e_test

import (
	"os"

	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cwxstat/rabbitmq/lib/compress"
	"github.com/cwxstat/rabbitmq/lib/encode"

	"github.com/streadway/amqp"

	"github.com/cwxstat/rabbitmq/lib/conn"
	"github.com/cwxstat/rabbitmq/lib/flag"

	"github.com/cwxstat/rabbitmq/lib/consumer"
	"github.com/cwxstat/rabbitmq/lib/producer"
	//"github.com/cwxstat/rabbitmq/e2e"
)

// FIXME: (mmc) This is really bad... works; but, needs to be cleaned up.
var _ = Describe("E2e", func() {

	var (
		dir         string
		desDir      string
		dataWritten string

		certPath          string
		caCertificate     string
		clientCertificate string
		clientKey         string
		username          string
		password          string
		port              string

		SetupConn func() (*amqp.Connection, error)
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

		certPath = "../etc/certs"
		caCertificate = "ca_certificate.pem"
		clientCertificate = "client_certificate.pem"
		clientKey = "key.unencrypted.pem"
		username = "pig"
		password = "P033wor4"
		port = "5671"

		SetupConn = func() (*amqp.Connection, error) {
			conn := conn.NewCONN()
			conn.CertPath(certPath).
				CACertificate(caCertificate).
				ClientCertificate(clientCertificate).
				ClientKey(clientKey).
				Port(port).
				Username(username).
				Password(password)

			return conn.Conn()

		}

	})
	AfterEach(func() {
		os.RemoveAll(dir)
		os.RemoveAll(desDir)
		os.RemoveAll("/tmp/compress.tar.gz")
		os.RemoveAll("/tmp/junk.tar.gz")
	})

	Describe("Check e2e", func() {

		Context("Should ...", func() {
			It("should ...", func() {

				// Note remove prefix is "/tmp/compress"
				err := compress.Compress(dir, "/tmp/compress")
				Expect(err).To(BeNil())

				result, err := encode.ReadEncode("/tmp/compress.tar.gz")
				Expect(err).To(BeNil())

				f := flag.NewFlags()

				err = producer.NewPublish(f.Exchange, f.ExchangeType,
					"test-gzip", result, true, producer.ConfirmOne,
					SetupConn)

				Expect(err).To(BeNil())
				os.RemoveAll("/tmp/compress.tar.gz")

				handler := &encode.HS{}
				handler.File = "/tmp/junk.tar.gz"
				handler.DestDir = desDir

				c, err := consumer.NewConsumer(f.Exchange,
					f.ExchangeType, "testq",
					"test-gzip", f.ConsumerTag, handler, SetupConn)

				Expect(err).To(BeNil())

				time.Sleep(2 * time.Second)
				err = c.Shutdown()
				Expect(err).To(BeNil())

			})
		})

	})

})
