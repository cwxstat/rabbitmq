package setup

import (
	"github.com/cwxstat/rabbitmq/lib/conn"
	"github.com/streadway/amqp"
	"os"
)

type Setup struct {
	certPath          string
	caCertificate     string
	clientCertificate string
	clientKey         string
	username          string
	password          string
	port              string
	url               string

	SetupConn func() (*amqp.Connection, error)
}

func NewSetup() *Setup {

	var (
		username string = "pig"
		password string = "P033wor4"
		url      string = "localhost"
		port     string = "5671"
	)

	// If environment set, use those values
	env(&username, &password, &url, &port)

	s := &Setup{certPath: "./etc/certs",
		caCertificate:     "ca_certificate.pem",
		clientCertificate: "client_certificate.pem",
		clientKey:         "key.unencrypted.pem",
		username:          username,
		password:          password,
		port:              port,
		url:               url,
	}

	s.update()

	return s

}

func (s *Setup) update() *Setup {
	s.SetupConn = func() (*amqp.Connection, error) {
		conn := conn.NewCONN()
		conn.CertPath(s.certPath).
			CACertificate(s.caCertificate).
			ClientCertificate(s.clientCertificate).
			ClientKey(s.clientKey).
			Port(s.port).
			Username(s.username).
			Password(s.password).
			URL(s.url)

		return conn.Conn()

	}
	return s
}

func (s *Setup) URL(url string) *Setup {
	s.url = url
	s.update()
	return s
}

func (s *Setup) Username(username string) *Setup {
	s.username = username
	s.update()
	return s
}

func (s *Setup) Port(port string) *Setup {
	s.port = port
	s.update()
	return s
}

func (s *Setup) Password(password string) *Setup {
	s.password = password
	s.update()
	return s
}

func (s *Setup) CertPath(certPath string) *Setup {
	s.certPath = certPath
	s.update()
	return s
}

func env(username, password, url, port *string) {

	if value, ok := os.LookupEnv("RABBITMQ_USERNAME"); ok {
		*username = value
	}
	if value, ok := os.LookupEnv("RABBITMQ_PASSWORD"); ok {
		*password = value
	}
	if value, ok := os.LookupEnv("RABBITMQ_URL"); ok {
		*url = value
	}
	if value, ok := os.LookupEnv("RABBITMQ_PORT"); ok {
		*port = value
	}
}
