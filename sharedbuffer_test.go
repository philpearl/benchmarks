package benchmarks

import (
	"runtime"
	"sync"
	"testing"
)

func TestSharedBuffer(t *testing.T) {
	sb := newSharedBuffer(10)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		var val byte
		var closed bool
		for i := 0; !closed; i++ {
			val, closed = sb.get()

			if !closed && val != byte(i) {
				t.Errorf("Expected %d, got %d", i, val)
			}
		}
	}()

	for i := 0; i < 100; i++ {
		sb.put(byte(i))
	}
	sb.close()

	wg.Wait()
}

// BenchmarkSharedBuffer tests a home-grown channel equivalent to see how it
// performs relative to a real channel. See BenchmarkChannelOneByte()
func BenchmarkSharedBuffer(b *testing.B) {
	sb := newSharedBuffer(4096)
	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()
		var closed bool

		for !closed {
			_, closed = sb.get()
		}
	}()

	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sb.put(byte(i))
	}
	sb.close()
	wg.Wait()
}

func BenchmarkSharedBufferMulti(b *testing.B) {
	sb := newSharedBuffer(4096)
	wg := sync.WaitGroup{}

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			var closed bool

			for !closed {
				_, closed = sb.get()
			}
		}()
	}

	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sb.put(byte(i))
	}
	sb.close()
	wg.Wait()
}
