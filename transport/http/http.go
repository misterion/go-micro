package http

// This is a hack

import (
	"github.com/micro/go-micro/transport"
)

func NewTransport(addrs []string, opt ...transport.Option) transport.Transport {
	return transport.NewTransport(addrs, opt...)
}