package main

import (
	"github.com/edwinmat/server-shared"
)

type PacketHandler interface {
	Configure() error
	Handle(*shared.ConsumerQueues)
}
