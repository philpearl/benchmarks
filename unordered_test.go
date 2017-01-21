package benchmarks

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

type unordered struct {
	id    string
	value int
}

type unorderedList []unordered

func (ul unorderedList) get(k string) (int, bool) {
	for i := range ul {
		uo := &ul[i]
		if uo.id == k {
			return uo.value, true
		}
	}
	return 0, false
}

func (ul unorderedList) add(id string, value int) unorderedList {
	return append(ul, unordered{id: id, value: value})
}

func BenchmarkUnorderedList(b *testing.B) {

	for _, numToAdd := range []int{1, 10, 25, 100, 500, 1000} {
		b.Run(fmt.Sprintf("add_%d", numToAdd), func(b *testing.B) {

			l := make(unorderedList, 0, numToAdd)

			toadd := make([]string, numToAdd)

			for i := range toadd {
				toadd[i] = strconv.Itoa(int(rand.Int63()))
			}

			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				benchmarkCheckAndAddEntries(b, numToAdd, l, toadd)
			}
		})
	}

}

func benchmarkCheckAndAddEntries(b *testing.B, numToAdd int, l unorderedList, toadd []string) {
	for i, s := range toadd {
		if _, ok := l.get(s); !ok {
			l = l.add(s, i)
		}
	}
}
