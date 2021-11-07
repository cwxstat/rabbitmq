package conn

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/streadway/amqp"
)

type CONN struct {
	certPath          string
	caCertificate     string
	clientCertificate string
	clientKey         string
	username          string
	password          string
	port              string
	url               string
	conn              *amqp.Connection
}

func NewCONN() *CONN {
	return &CONN{}
}

func (conn *CONN) Conn() (*amqp.Connection, error) {
	err := conn.load()
	if err != nil {
		return nil, err
	}
	return conn.conn, nil
}

func (conn *CONN) load() error {
	cfg := new(tls.Config)

	// see at the top
	cfg.RootCAs = x509.NewCertPool()

	ca, err := ioutil.ReadFile(conn.certPath + "/" + conn.caCertificate)
	if err != nil {
		return err
	}
	cfg.RootCAs.AppendCertsFromPEM(ca)

	cert, err := tls.LoadX509KeyPair(conn.certPath+"/"+conn.clientCertificate,
		conn.certPath+"/"+conn.clientKey)
	if err != nil {
		return err
	}

	cfg.Certificates = append(cfg.Certificates, cert)
	conn.conn, err = amqp.DialTLS("amqps://"+conn.username+":"+conn.password+"@"+conn.url+":"+conn.port, cfg)
	if err != nil {
		return err
	}

	return nil
}

func (conn *CONN) CertPath(certPath string) *CONN {
	conn.certPath = certPath
	return conn
}

func (conn *CONN) Username(username string) *CONN {
	conn.username = username
	return conn
}

func (conn *CONN) Password(password string) *CONN {
	conn.password = password
	return conn
}

func (conn *CONN) Port(port string) *CONN {
	conn.port = port
	return conn
}

func (conn *CONN) CACertificate(caCertificate string) *CONN {
	conn.caCertificate = caCertificate
	return conn
}

func (conn *CONN) ClientCertificate(clientCertificate string) *CONN {
	conn.clientCertificate = clientCertificate
	return conn
}

func (conn *CONN) ClientKey(clientKey string) *CONN {
	conn.clientKey = clientKey
	return conn
}

func (conn *CONN) URL(url string) *CONN {
	conn.url = url
	return conn
}
