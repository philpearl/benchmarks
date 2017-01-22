package benchmarks

import (
	"testing"

	"runtime"
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

func BenchmarkChannelOneByteMultiReceive(b *testing.B) {
	ch := make(chan byte, 4096)
	wg := sync.WaitGroup{}

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for range ch {

			}
		}()
	}

	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ch <- byte(i)
	}
	close(ch)
	wg.Wait()
}

func BenchmarkChannelOneByteMultiSendRecv(b *testing.B) {
	ch := make(chan byte, 4096)
	wg := sync.WaitGroup{}

	numCPU := runtime.NumCPU()
	wg.Add(numCPU)
	for i := 0; i < numCPU; i++ {
		go func() {
			for range ch {

			}
			wg.Done()
		}()
	}

	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()

	sendWG := sync.WaitGroup{}
	sendWG.Add(numCPU)
	for j := 0; j < numCPU; j++ {
		go func() {
			for i := 0; i < b.N; i += numCPU {
				ch <- byte(i)
			}
			sendWG.Done()
		}()
	}
	sendWG.Wait()
	close(ch)
	wg.Wait()
}
