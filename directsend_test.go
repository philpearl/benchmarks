package benchmarks

import (
	"sync"
	"sync/atomic"
	"testing"
)

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

func BenchmarkAtomicAdd(b *testing.B) {
	var total int32

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			atomic.AddInt32(&total, 1)
		}
	})
}

func BenchmarkChannelDirectSend2(b *testing.B) {
	benchmark := func(b *testing.B, buffer int) {
		out := make(chan int, buffer)
		back := make(chan int, buffer)

		go func() {
			for v := range out {
				back <- v
			}
		}()

		for i := 0; i < b.N; i++ {
			out <- 1
			<-back
		}

		close(out)
		close(back)
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

func BenchmarkChannelDirectSend3(b *testing.B) {
	ch := make(chan int, b.N)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ch <- i
	}
	close(ch)

	k := 0
	for j := range ch {
		k += j
	}
	if k != b.N*(b.N-1)/2 {
		b.Errorf("K not as expected k=%d, b.N=%d", k, b.N)
	}
}
