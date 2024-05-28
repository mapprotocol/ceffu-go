package client

import (
	"math/rand"
	"time"
)

type RequestID interface {
	Generate() int64
}

func NewRequestID() RequestID {
	return &defaultRequestID{}
}

type defaultRequestID struct{}

func (d *defaultRequestID) Generate() int64 {
	// For now, just return a random number
	rand.Seed(time.Now().UnixNano())
	return rand.Int63()
}
