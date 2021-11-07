package conn_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cwxstat/rabbitmq/lib/conn"

	"github.com/streadway/amqp"
)

var _ = Describe("Conn", func() {

	var (
		certPath          string
		caCertificate     string
		clientCertificate string
		clientKey         string
		username          string
		password          string
		port              string
		url               string

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
		url = "localhost"

		SetupConn = func() (*amqp.Connection, error) {
			conn := conn.NewCONN()
			conn.CertPath(certPath).
				CACertificate(caCertificate).
				ClientCertificate(clientCertificate).
				ClientKey(clientKey).
				Port(port).
				Username(username).
				Password(password).
				URL(url)

			return conn.Conn()

		}

	})

	Describe("Test Connection", func() {

		Context("Configure settings", func() {
			It("should connect", func() {

				_, err := SetupConn()
				Expect(err).To(BeNil())
			})
		})

	})

})
