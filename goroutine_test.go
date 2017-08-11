package benchmarks

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkGoroutine(b *testing.B) {

	for _, numGoRoutines := range []int{1, 2, 4, 8, 16, 32, 100, 1000, 10000, 100000} {
		numGoRoutines := numGoRoutines
		b.Run(strconv.Itoa(numGoRoutines), func(b *testing.B) {
			b.ReportAllocs()

			for k := 0; k < b.N; k++ {
				wg := &sync.WaitGroup{}
				wg.Add(numGoRoutines)
				target := 400000000 / numGoRoutines
				for i := 0; i < numGoRoutines; i++ {
					go func() {
						total := 0
						for j := 0; j < target; j++ {
							total += j
						}
						wg.Done()
						if total == 0 {
							b.Logf("really?")
						}
					}()
				}
				wg.Wait()
			}
		})
	}
}

func BenchmarkGoroutineChannel(b *testing.B) {

	for _, numGoRoutines := range []int{1, 2, 8, 32, 1000, 10000} {
		numGoRoutines := numGoRoutines
		for _, chanBuffer := range []int{0, 1, 2, 8, 32, 1000, 10000} {
			chanBuffer := chanBuffer
			b.Run(fmt.Sprintf("goroutines=%d_chanBuffer=%d", numGoRoutines, chanBuffer), func(b *testing.B) {

				for k := 0; k < b.N; k++ {
					var numZeroWork int64
					// Stop the timer while we're starting our goroutines
					b.StopTimer()
					ch := make(chan int, chanBuffer)
					wg := &sync.WaitGroup{}
					wg.Add(numGoRoutines)
					for i := 0; i < numGoRoutines; i++ {
						go func() {
							total := 0
							for j := range ch {
								total += j
							}
							wg.Done()
							if total == 0 {
								atomic.AddInt64(&numZeroWork, 1)
							}
						}()
					}
					b.StartTimer()

					for i := 0; i < 100000; i++ {
						ch <- i
					}
					close(ch)
					wg.Wait()

					if numZeroWork > 0 && k == 0 {
						b.Logf("%d goroutines did zero work", numZeroWork)
					}
				}
			})
		}
	}
}

func BenchmarkGoroutineChannelWork(b *testing.B) {

	for _, numGoRoutines := range []int{1, 2, 4, 8} {
		numGoRoutines := numGoRoutines
		for _, work := range []int{100, 1000, 10000, 100000} {
			work := work
			b.Run(fmt.Sprintf("goroutines=%d_work=%d", numGoRoutines, work), func(b *testing.B) {

				for k := 0; k < b.N; k++ {
					var numZeroWork int64
					// Stop the timer while we're starting our goroutines
					b.StopTimer()
					ch := make(chan int, 1000)
					wg := &sync.WaitGroup{}
					wg.Add(numGoRoutines)
					for i := 0; i < numGoRoutines; i++ {
						go func() {
							total := 0
							for j := range ch {
								for l := 0; l < work; l++ {
									total += j
								}
							}
							wg.Done()
							if total == 0 {
								atomic.AddInt64(&numZeroWork, 1)
							}
						}()
					}
					b.StartTimer()

					for i := 0; i < 100000; i++ {
						ch <- 1
					}
					close(ch)
					wg.Wait()

					if numZeroWork > 0 && k == 0 {
						b.Logf("%d goroutines did zero work", numZeroWork)
					}
				}
			})
		}
	}
}

func BenchmarkJustWork(b *testing.B) {
	for _, work := range []int{100, 1000, 10000, 100000} {
		work := work
		b.Run(fmt.Sprintf("work=%d", work), func(b *testing.B) {

			for k := 0; k < b.N; k++ {
				total := 0
				for l := 0; l < work; l++ {
					total++
				}
				if total == 0 {
					b.Logf("oops")
				}
			}
		})
	}
}
