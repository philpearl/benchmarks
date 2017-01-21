package benchmarks

import (
	"testing"
)

func BenchmarkCopy(b *testing.B) {
	from := make([]byte, b.N)
	to := make([]byte, b.N)

	b.ReportAllocs()
	b.ResetTimer()
	b.SetBytes(1)

	copy(to, from)
}
