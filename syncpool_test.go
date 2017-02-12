package benchmarks

import (
	"sync"
	"testing"
)

func BenchmarkSyncPool(b *testing.B) {

	type dummy struct {
		thing int
	}

	pool := &sync.Pool{
		New: func() interface{} { return &dummy{} },
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		a := pool.Get().(*dummy)
		a.thing = i
		pool.Put(a)
	}
}
