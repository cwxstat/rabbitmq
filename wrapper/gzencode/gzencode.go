package gzencode

import (
	"github.com/cwxstat/rabbitmq/lib/compress"
	"github.com/cwxstat/rabbitmq/lib/encode"

	"github.com/cwxstat/rabbitmq/lib/flag"

	"github.com/cwxstat/rabbitmq/lib/handle"
	"github.com/cwxstat/rabbitmq/lib/consumer"
	"github.com/cwxstat/rabbitmq/lib/producer"
	"github.com/cwxstat/rabbitmq/lib/setup"
)

type GZ struct {
	dirIN       string
	dirINGZ     string
	dirINIgnore string
	dirOUT      string
	routingKey  string
	queue       string
	consumerQ   string
	handler     *handle.HS

	setup *setup.Setup
	f     *flag.Flags
	c     *consumer.Consumer
	err   error
}

func NewGZ() *GZ {
	g := &GZ{}
	return g
}

func (g *GZ) CertPath(certsdir string) *GZ {
	g.setup = setup.NewSetup().CertPath(certsdir)
	return g
}

func (g *GZ) CertURL(url string) *GZ {
	g.setup.URL(url)
	return g
}

func (g *GZ) DirIn(dir string) *GZ {
	g.dirIN = dir
	if g.dirINGZ == "" {
		g.dirINGZ = g.dirIN + ".tar.gz"

	}
	if g.dirINIgnore == "" {
		g.dirINIgnore = dir
	}
	if g.routingKey == "" {
		g.routingKey = "routeKeyTargz"
	}
	if g.queue == "" {
		g.queue = "targz"
		if g.consumerQ == "" {
			g.consumerQ = "targz"
		}
	}
	if g.f == nil {
		g.f = flag.NewFlags()
	}

	return g
}

func (g *GZ) ConsumerQ(q string) *GZ {
	g.consumerQ = q
	return g
}

func (g *GZ) DestDir(dir string) *GZ {
	g.dirOUT = dir
	return g
}

func (g *GZ) HandleFile(gzfile, destdir string) *GZ {
	g.handler = &handle.HS{}
	g.handler.DestDir = destdir
	g.handler.File = gzfile

	return g
}

func (g *GZ) Produce() error {
	err := compress.Compress(g.dirIN, g.dirINIgnore)
	if err != nil {
		return err
	}
	result, err := encode.ReadEncode(g.dirINGZ)
	if err != nil {
		return err
	}

	if g.f == nil {
		g.f = flag.NewFlags()
	}
	err = producer.NewPublish(g.f.Exchange, g.f.ExchangeType,
		g.routingKey, result, true, producer.ConfirmOne,
		g.setup.SetupConn)

	if err != nil {
		return err
	}

	return nil
}

func (g *GZ) Consume() error {

	if g.f == nil {
		g.f = flag.NewFlags()
	}

	if g.c, g.err = consumer.NewConsumer(g.f.Exchange,
		g.f.ExchangeType, g.queue,
		g.routingKey, g.f.ConsumerTag, g.handler, g.setup.SetupConn); g.err != nil {
		return g.err
	}

	return nil
}

func (g *GZ) Shutdown() error {
	return g.c.Shutdown()
}
