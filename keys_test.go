package benchmarks

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bmizerany/assert"
)

func BenchmarkKeySprintf(b *testing.B) {
	key1 := "key1"
	key2 := "key2"
	key3 := "key3"

	results := make([]string, 0, b.N)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		results = append(results, fmt.Sprintf("%s§%s§%s", key1, key2, key3))
		key1, key2, key3 = key3, key1, key2
	}
}

func BenchmarkKeyAdd(b *testing.B) {
	key1 := "key1"
	key2 := "key2"
	key3 := "key3"

	results := make([]string, 0, b.N)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		results = append(results, key1+"§"+key2+"§"+key3)
		key1, key2, key3 = key3, key1, key2
	}
}

func BenchmarkKeyExtractSplit(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		key := strings.Split("blah§sitting§pretty", "§")[1]
		if key != "sitting" {
			b.Fatal("key does not match")
		}
	}
}

func BenchmarkKeyExtract(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		key := keyExtract("blah§sitting§pretty", '§', 1)
		if key != "sitting" {
			b.Fatal("key does not match")
		}
	}
}

func TestKeyExtract(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		field int
		exp   string
	}{
		{
			name:  "first",
			key:   "ab§ac§ad",
			field: 0,
			exp:   "ab",
		},
		{
			name:  "second",
			key:   "ab§ac§ad",
			field: 1,
			exp:   "ac",
		},
		{
			name:  "third",
			key:   "ab§ac§ad",
			field: 2,
			exp:   "ad",
		},
		{
			name:  "empty1",
			key:   "§zz§ad",
			field: 0,
			exp:   "",
		},
		{
			name:  "empty2",
			key:   "ab§§ad",
			field: 1,
			exp:   "",
		},
		{
			name:  "empty3",
			key:   "ab§aa§",
			field: 2,
			exp:   "",
		},
		{
			name:  "pastempty1",
			key:   "§zz§ad",
			field: 1,
			exp:   "zz",
		},
		{
			name:  "pastempty2",
			key:   "ab§§ad",
			field: 2,
			exp:   "ad",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := keyExtract(test.key, '§', test.field)
			assert.Equal(t, test.exp, actual)
		})
	}
}
