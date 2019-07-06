package randata

import (
	"math/rand"
	"sync"
)

// Boolgen is an interface that implements the generation of a random bool
// value (bool type)
type Boolgen interface {
	RandomBool() bool
}

type boolgen struct {
	sync.Mutex
	src       rand.Source
	cache     int64
	remaining int
}

func (b *boolgen) RandomBool() bool {
	b.Lock()
	defer b.Unlock()
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}
