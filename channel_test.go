package benchmarks

import (
	"testing"

	"sync"
)

func BenchmarkChannelOneByte(b *testing.B) {
	ch := make(chan byte, 4096)
	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		for range ch {

		}
	}()

	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ch <- byte(i)
	}
	close(ch)
	wg.Wait()
}
