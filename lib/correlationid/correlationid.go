package correlationid

import (
	"fmt"
	//log "github.com/sirupsen/logrus"

	"github.com/streadway/amqp"
)

// TODO: finish this
//   ref: https://www.rabbitmq.com/tutorials/tutorial-six-go.html
type CorrelationIdStruct struct {
	CorrelationIdFunc    func() string
	CorrelationIdConsume func(*amqp.Channel) (<-chan amqp.Delivery, error)
}

func NewCorrelationId(qname ...string) *CorrelationIdStruct {
	c := &CorrelationIdStruct{}
	var count int64
	var _qname string = "qreturn"
	if len(qname) > 0 {
		_qname = qname[0]
	}

	c.CorrelationIdFunc = func() string {
		count += 1
		return fmt.Sprintf("cid(%d)", count)
	}
	c.CorrelationIdConsume = func(ch *amqp.Channel) (<-chan amqp.Delivery, error) {
		msgs, err := ch.Consume(
			_qname, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		return msgs, err
	}

	return c
}
