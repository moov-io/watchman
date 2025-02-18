package coder

import (
	"sync"

	"github.com/vmihailenco/msgpack/v5"
)

var (
	encoderPool = sync.Pool{
		New: func() any {
			return NewEncoder()
		},
	}
)

func NewEncoder() *msgpack.Encoder {
	return msgpack.GetEncoder()
}

func GetEncoder() *msgpack.Encoder {
	sb, ok := encoderPool.Get().(*msgpack.Encoder)
	if !ok {
		return NewEncoder()
	}
	return sb
}

func SaveEncoder(sb *msgpack.Encoder) {
	encoderPool.Put(sb)
}
