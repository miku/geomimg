package geomimg

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRandomString(t *testing.T) {
	rand.Seed(0)
	var cases = []struct {
		length int
		result string
	}{
		{0, ""},
		{1, "c"},
		{2, "ub"},
		{16, "yhizzkakbleeanso"},
	}
	for _, c := range cases {
		result := RandomString(c.length)
		if c.result != result {
			t.Fatalf("want %v, got %v", c.result, result)
		}
	}
}

func BenchmarkRandomString(b *testing.B) {
	var bms = []int{1, 2, 8, 16, 32, 128}
	for _, bm := range bms {
		name := fmt.Sprintf("len-%d", bm)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RandomString(bm)
			}
		})
	}
}
