package benchmarks

import "testing"
import "sync"

func BenchmarkChannelDirectSend(b *testing.B) {
	benchmark := func(b *testing.B, buffer int) {
		ch := make(chan int, buffer)
		wg := sync.WaitGroup{}
		var total int

		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				total += v
			}
		}()

		for i := 0; i < b.N; i++ {
			ch <- 1
		}

		close(ch)
		wg.Wait()
	}

	b.ReportAllocs()

	b.Run("direct", func(b *testing.B) {
		benchmark(b, 0)
	})

	b.Run("buffered", func(b *testing.B) {
		benchmark(b, 1)
	})

	b.Run("buffered-10", func(b *testing.B) {
		benchmark(b, 10)
	})
}
