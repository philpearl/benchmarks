package benchmarks

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

// Speed
// - read file
// - check and add symbol map
// - add to reverse symbol map
// - check and add edge map (*2 for both directions)
// - add to UF
// Symbol is <type>:<cid>:<id>, say about 40 bytes. Id is about 4 bytes
// so we're talking about 44 * 100mm = ~5GB for the forward direction.

// BenchmarkSync. Uncontended lock/unlock
// About 23 ns
func BenchmarkSync(b *testing.B) {
	var s sync.Mutex

	for i := 0; i < b.N; i++ {
		s.Lock()
		s.Unlock()
	}
}

// BenchmarkAlloc. Allocating a small slice
func BenchmarkAlloc(b *testing.B) {
	store := make([][]byte, b.N)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		store[i] = make([]byte, 8)
	}
}

// BenchmarkAddToMap benchmarks adding something to an unsized map
// About 460 ns
func BenchmarkAddToMap(b *testing.B) {
	toadd := make([]string, b.N)

	for i := range toadd {
		toadd[i] = strconv.Itoa(int(rand.Int63()))
	}

	m := make(map[string]int)

	b.ReportAllocs()
	b.ResetTimer()

	for i, s := range toadd {
		m[s] = i
	}
}

// BenchmarkCheckAddToMap benchmarks adding something to an unsized map, checking
// first whether the item is present (it shouldn't be). Minimally slower
// About 460 ns
func BenchmarkCheckAddToMap(b *testing.B) {
	toadd := make([]string, b.N)

	for i := range toadd {
		toadd[i] = strconv.Itoa(int(rand.Int63()))
	}

	m := make(map[string]int)

	b.ReportAllocs()
	b.ResetTimer()

	for i, s := range toadd {
		if _, ok := m[s]; !ok {
			m[s] = i
		}
	}
}

func BenchmarkCheckAddToBigMap(b *testing.B) {
	m := make(map[string]int, 1000000)
	toadd := make([]string, b.N)

	for i := range toadd {
		toadd[i] = strconv.Itoa(int(rand.Int63()))
	}

	for i := 0; i < 1000000; i++ {
		m[strconv.Itoa(int(rand.Int63()))] = i
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i, s := range toadd {
		if _, ok := m[s]; !ok {
			m[s] = i
		}
	}
}

// BenchmarkSizedCheckAddToMap benchmarks adding something to an unsized map, checking
// first whether the item is present (it shouldn't be).
// About 250 ns
func BenchmarkSizedCheckAddToMap(b *testing.B) {
	toadd := make([]string, b.N)

	for i := range toadd {
		toadd[i] = strconv.Itoa(int(rand.Int63()))
	}

	m := make(map[string]int, b.N)

	b.ReportAllocs()
	b.ResetTimer()

	for i, s := range toadd {
		if _, ok := m[s]; !ok {
			m[s] = i
		}
	}
}
