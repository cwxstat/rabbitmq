package consumer

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Handler interface {
	Handle(<-chan amqp.Delivery, chan error)
}

type Consumer struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	tag       string
	done      chan error
	setupConn func() (*amqp.Connection, error)
	handler   Handler
}

func (c *Consumer) setup(exchange, exchangeType,
	queueName, key, ctag string) (<-chan amqp.Delivery, error) {
	var err error

	if c.conn, err = c.setupConn(); err != nil {
		return nil, fmt.Errorf("Dial: %s", err)
	}

	go func() {
		fmt.Printf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	log.Printf("got Connection, getting Channel")
	if c.channel, err = c.conn.Channel(); err != nil {
		return nil, fmt.Errorf("Channel: %s", err)
	}

	log.Printf("got Channel, declaring Exchange (%q)", exchange)
	log.Printf("ExchangeType (%q)", exchangeType)
	if err = c.channel.ExchangeDeclare(
		exchange,     // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return nil, fmt.Errorf("Exchange Declare: %s", err)
	}

	log.Printf("declared Exchange, declaring Queue %q", queueName)
	queue, err := c.channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Declare: %s", err)
	}

	log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		queue.Name, queue.Messages, queue.Consumers, key)

	if err = c.channel.QueueBind(
		queue.Name, // name of the queue
		key,        // bindingKey
		exchange,   // sourceExchange
		false,      // noWait
		nil,        // arguments
	); err != nil {
		return nil, fmt.Errorf("Queue Bind: %s", err)
	}

	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", c.tag)
	deliveries, err := c.channel.Consume(
		queue.Name, // name
		c.tag,      // consumerTag,
		false,      // noAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Consume: %s", err)
	}
	return deliveries, nil
}

func NewConsumer(exchange, exchangeType,
	queueName, key, ctag string,
	handler Handler,
	setupConn func() (*amqp.Connection, error)) (*Consumer, error) {

	c := &Consumer{
		conn:      nil,
		channel:   nil,
		tag:       ctag,
		done:      make(chan error),
		setupConn: setupConn,
		handler:   handler,
	}

	deliveries, err := c.setup(exchange, exchangeType,
		queueName, key, ctag)
	if err != nil {
		return nil, err
	}

	go c.handler.Handle(deliveries, c.done)

	return c, nil
}

func (c *Consumer) Shutdown() error {
	// will close() the deliveries channel
	if err := c.channel.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("Consumer cancel failed: %s", err)
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer log.Printf("AMQP shutdown OK")

	// wait for handle() to exit
	return <-c.done
}

type HS struct {
	count int64
}

func (h *HS) Handle(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		h.count += 1
		log.Printf(
			"got count(%d) %dB delivery: [%v] %q: %q",
			h.count,
			len(d.Body),
			d.DeliveryTag,
			d.Body,
			d.AppId,
		)
		d.Ack(false)

	}
	log.Printf("handle: deliveries channel closed")
	done <- nil
}
