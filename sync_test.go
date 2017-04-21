package benchmarks

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkRLock(b *testing.B) {

	lock := sync.RWMutex{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lock.RLock()
		lock.RUnlock()
	}
}

func BenchmarkAtomicLoad(b *testing.B) {
	value := atomic.Value{}

	value.Store("hat")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		v := value.Load().(string)
		if len(v) == 0 {
			b.Fatalf("shouldn't happen")
		}
	}
}
