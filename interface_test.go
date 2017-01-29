package benchmarks

import (
	"testing"
)

type benchType interface {
	get() int
}

type typeA struct {
	a int
}

func (a *typeA) get() int {
	return a.a
}

type typeB struct {
	b int
}

func (b *typeB) get() int {
	return b.b
}

func BenchmarkInterfaceTypeAssertion(b *testing.B) {
	var aCount, bCount int

	var x interface{}

	x = typeA{}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		switch x.(type) {
		case *typeA:
			aCount++
		case typeB:
			bCount++
		}
	}
}

func BenchmarkInterfaceCall(b *testing.B) {
	var x benchType
	x = &typeA{}
	b.ReportAllocs()
	b.ResetTimer()

	total := 0
	for i := 0; i < b.N; i++ {
		total += x.get()
	}

	if total > 0 {
		b.Logf("total is %d", total)
	}
}

func BenchmarkInterfaceCallComparison(b *testing.B) {
	x := &typeA{}
	b.ReportAllocs()
	b.ResetTimer()

	total := 0
	for i := 0; i < b.N; i++ {
		total += x.get()
	}

	if total > 0 {
		b.Logf("total is %d", total)
	}
}

func BenchmarkInterfaceCallTypeAssertion(b *testing.B) {
	var x interface{}
	x = typeA{}
	b.ReportAllocs()
	b.ResetTimer()

	total := 0
	for i := 0; i < b.N; i++ {
		switch x := x.(type) {
		case *typeA:
			total += x.get()
		case typeB:
			total += x.get()
		}
	}

	if total > 0 {
		b.Logf("total is %d", total)
	}
}

func BenchmarkInterfaceStore(b *testing.B) {
	store := make([]interface{}, 1)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		store[0] = typeA{}
	}
}

func BenchmarkInterfaceStorePointer(b *testing.B) {
	store := make([]interface{}, 1)
	val := &typeA{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		store[0] = val
	}
}
