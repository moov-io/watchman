package coder

import (
	"sync"

	"github.com/vmihailenco/msgpack/v5"
)

var (
	decoderPool = sync.Pool{
		New: func() any {
			return NewDecoder()
		},
	}
)

func NewDecoder() *msgpack.Decoder {
	return msgpack.GetDecoder()
}

func GetDecoder() *msgpack.Decoder {
	sb, ok := decoderPool.Get().(*msgpack.Decoder)
	if !ok {
		return NewDecoder()
	}
	return sb
}

func SaveDecoder(sb *msgpack.Decoder) {
	decoderPool.Put(sb)
}
