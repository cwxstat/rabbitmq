package flag

import (
	"time"
)

type Flags struct {
	Exchange     string
	ExchangeType string
	Queue        string
	BindingKey   string
	ConsumerTag  string
	Lifetime     time.Duration
}

func NewFlags() *Flags {
	return &Flags{Exchange: "test-exchange",
		ExchangeType: "fanout",
		Queue:        "test-queue",
		BindingKey:   "test-key",
		ConsumerTag:  "simple-consumer",
		Lifetime:     300 * time.Second}
}
