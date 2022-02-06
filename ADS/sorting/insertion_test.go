package insertion_test

import (
	"fmt"
	"testing"

	sort "github.com/moledoc/moledoc/ADS/sorting"
)

func assertList(result []int, expected []int) bool {
	if len(result) != len(expected) {
		return false
	}
	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			return false
		}
	}
	return true
}

var cases [][]int = [][]int{{3, 2, 1}, {-1, -2, -3}, {2, 1, 3}, {5, 6, 2, 3, 1, 10, 4, 7, 8, 9}, {}, {1, 2, 3, 4, 5}, {1}}
var expected [][]int = [][]int{{1, 2, 3}, {-3, -2, -1}, {1, 2, 3}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {}, {1, 2, 3, 4, 5}, {1}}

func TestInsertion(t *testing.T) {
	fmt.Print("Increasing algorithm: ")
	for i := 0; i < len(cases); i++ {
		if !assertList(sort.Insertion(cases[i]), expected[i]) {
			fmt.Println("FAIL")
			t.Fatalf("%v: Expected %v, got %v\n", i, expected[i], cases[i])
		}
	}
	fmt.Println("OK")
	fmt.Print("Decreasing algorithm: ")
	for i := 0; i < len(cases); i++ {
		if !assertList(sort.InsertionDesc(cases[i]), expected[i]) {
			fmt.Println("FAIL")
			t.Fatalf("%v: Expected %v, got %v\n", i, expected[i], cases[i])
		}
	}
	fmt.Println("OK")
}
