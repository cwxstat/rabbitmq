package consumer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/streadway/amqp"

	"github.com/cwxstat/rabbitmq/lib/conn"
	"github.com/cwxstat/rabbitmq/lib/flag"

	"github.com/cwxstat/rabbitmq/lib/consumer"
)

var _ = Describe("Consumer", func() {

	var (
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
		certPath = "../../etc/certs"
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

	Describe("Test Connection", func() {

		Context("Connection", func() {
			It("should send messages", func() {

				f := flag.NewFlags()
				handler := &consumer.HS{}
				c, err := consumer.NewConsumer(f.Exchange,
					f.ExchangeType, f.Queue,
					f.BindingKey, f.ConsumerTag, handler, SetupConn)

				Expect(err).To(BeNil())

				err = c.Shutdown()
				Expect(err).To(BeNil())

			})
		})

	})

})
