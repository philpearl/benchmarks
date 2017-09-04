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
	for _, numGoRoutines := range []int{1, 2, 8, 32, 1000, 10000, 100000} {
		for _, buf := range []int{0, 1, 100, 1000, 10000} {

			// Before we start the benchmark we start a pool of goroutines to do the
			// work.
			ch := make(chan struct{}, buf)
			results := make(chan int, 1000)
			wg := sync.WaitGroup{}
			wg.Add(numGoRoutines)
			work := 10000
			for i := 0; i < numGoRoutines; i++ {
				go func() {
					defer wg.Done()
					for range ch {
						total := 0
						for l := 0; l < work; l++ {
							total++
						}
						results <- total
					}
				}()
			}

			b.Run(fmt.Sprintf("goroutines=%d,buf=%d", numGoRoutines, buf), func(b *testing.B) {
				for k := 0; k < b.N; k++ {
					// Send a 1000 requests for work
					for i := 0; i < 1000; i++ {
						ch <- struct{}{}
					}
					total := 0
					for i := 0; i < 1000; i++ {
						total += <-results
					}
					if total != 1000*work {
						b.Logf("total not as expected. %d", total)
					}
				}
			})
			close(ch)
			wg.Wait()
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
		b.Run(fmt.Sprintf("work=%d", work), func(b *testing.B) {
			b.ReportAllocs()

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

/*
We benchmark some work, just adding up. The results are as follows. I used `go test -run ^$ -bench BenchmarkJustWork -count 10` and fed the results into benchstat.

name                    time/op
JustWork/work=100-8     42.3ns ± 2%
JustWork/work=1000-8     316ns ± 2%
JustWork/work=10000-8   3.05µs ± 2%
JustWork/work=100000-8  31.1µs ± 3%

Now, to calculate the overhead of moving work to a goroutine, we repeat the benchmark, but do our adding up on a goroutine. Here's the result when we restrict it to one CPU.

name                         time/op
WorkOnGoroutine/work=100      447ns ± 3%
WorkOnGoroutine/work=1000     800ns ± 2%
WorkOnGoroutine/work=10000   3.49µs ± 1%
WorkOnGoroutine/work=100000  31.2µs ± 2%

To calculate the overhead should be just a case of subtracting one set of numbers from the other

100      447 - 42.3 = 405ns
1000      800 - 316 = 484ns
10000   3.49 - 3.05 = 0.44µs
100000  31.2 - 31.1 = 0.1µs

So this is all a bit rough, but it looks like the overhead of this particular way of moving work to a goroutine is about 400ns. The 100,000 case looks worse than this, but I'm going to put this down to rounding effects.

What happens if we allow the benchmark to run on all available CPUs? Here's the numbers.

name                           time/op
WorkOnGoroutine/work=100-8      559ns ± 2%
WorkOnGoroutine/work=1000-8     985ns ± 1%
WorkOnGoroutine/work=10000-8   4.64µs ± 2%
WorkOnGoroutine/work=100000-8  33.9µs ± 0%

If we work out the overhead again things look much worse.

100      559 - 42.3 = 516.7ns
1000      985 - 316 = 669ns
10000   4.64 - 3.05 = 1.59µs
100000  33.9 - 31.1 = 2.8µs

Not only is the overhead higher, but it gets worse when more work is done on the goroutine. This makes very little sense, as the on the face of it the amount of overhead should be independent of the work.

*/

func BenchmarkWorkOnGoroutine(b *testing.B) {
	for _, work := range []int{100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("work=%d", work), func(b *testing.B) {

			b.ReportAllocs()

			for k := 0; k < b.N; k++ {
				rsp := make(chan int)
				go func(rsp chan<- int, work int) {
					total := 0
					for l := 0; l < work; l++ {
						total++
					}
					rsp <- total
					close(rsp)
				}(rsp, work)
				total := <-rsp

				if total == 0 {
					b.Logf("oops")
				}
			}
		})
	}
}
