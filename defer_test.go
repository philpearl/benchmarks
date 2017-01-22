package benchmarks

import (
	"testing"
)

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			a := 1
			defer func() {
				a--
			}()
		}()
	}
}
