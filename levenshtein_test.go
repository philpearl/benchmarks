package benchmarks

import (
	"testing"

	agnivade "github.com/agnivade/levenshtein"
	arbovm "github.com/arbovm/levenshtein"
	dgryski "github.com/dgryski/trifles/leven"
	honzab "github.com/honzab/levenshtein"
	kse "github.com/kse/levenshtein"
	philpearl "github.com/philpearl/levenshtein"
	texttheater "github.com/texttheater/golang-levenshtein/levenshtein"
)

func BenchmarkLevenshtein_philpearl(b *testing.B) {
	s1 := "frederick"
	s2 := "fredelstick"
	total := 0

	b.ReportAllocs()
	b.ResetTimer()

	c := &philpearl.Context{}

	for i := 0; i < b.N; i++ {
		total += c.Distance(s1, s2)
	}

	if total == 0 {
		b.Logf("total is %d", total)
	}
}

func BenchmarkLevenshtein_long(b *testing.B) {
	s1 := "frederick"
	s2 := "fredelstick"
	total := 0

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		total += levDist(s1, s2)
	}

	if total == 0 {
		b.Logf("total is %d", total)
	}
}

func BenchmarkLevenshtein_dgryski(b *testing.B) {
	s1 := "frederick"
	s2 := "fredelstick"
	total := 0

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		total += dgryski.Levenshtein([]rune(s1), []rune(s2))
	}

	if total == 0 {
		b.Logf("total is %d", total)
	}
}

func BenchmarkLevenshtein_texttheater(b *testing.B) {
	s1 := []rune("frederick")
	s2 := []rune("fredelstick")
	total := 0

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		total += texttheater.DistanceForStrings(s1, s2, texttheater.DefaultOptions)
	}

	if total == 0 {
		b.Logf("total is %d", total)
	}
}

func BenchmarkLevenshtein_kse(b *testing.B) {
	s1 := []byte("frederick")
	s2 := []byte("fredelstick")
	total := 0

	lev := kse.New(1, 1, 1)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		total += lev.Dist(s1, s2)
	}

	if total == 0 {
		b.Logf("total is %d", total)
	}
}

func BenchmarkLevenshtein_honzab(b *testing.B) {
	s1 := "frederick"
	s2 := "fredelstick"
	total := 0

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		total += honzab.Distance(s1, s2)
	}

	if total == 0 {
		b.Logf("total is %d", total)
	}
}

func BenchmarkLevenshtein_arbovm(b *testing.B) {
	s1 := "frederick"
	s2 := "fredelstick"
	total := 0

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		total += arbovm.Distance(s1, s2)
	}

	if total == 0 {
		b.Logf("total is %d", total)
	}
}

func BenchmarkLevenshtein_agnivade(b *testing.B) {
	s1 := "frederick"
	s2 := "fredelstick"
	total := 0

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		total += agnivade.ComputeDistance(s1, s2)
	}

	if total == 0 {
		b.Logf("total is %d", total)
	}
}
