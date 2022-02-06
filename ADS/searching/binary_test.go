package binary_test

import (
	search "github.com/moledoc/moledoc/ADS/searching"
	"testing"
)

var cases [][]int = [][]int{{1, 2, 3}, {-3, -2, -1}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {}, {1, 2, 3, 4, 5}, {1}}
var srchFor []int = []int{2, -1, 3, 5, 4, 1, 11}
var expected []int = []int{1, 2, 2, -1, 3, 0, -1}

func TestBinary(t *testing.T) {
	for i := 0; i < len(cases); i++ {
		if at := search.Binary(cases[i], srchFor[i]); at != expected[i] {
			t.Fatalf("%v: Expected %v, got %v\n", i, expected[i], at)
		}
	}
}
