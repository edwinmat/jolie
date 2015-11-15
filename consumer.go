package main

import (
	"github.com/edwinmat/server-shared"
)

type Consumer interface {
	Configure() error
	Consume() (*shared.ConsumerQueues, error)
}
