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
var expectedDesc [][]int = [][]int{{3, 2, 1}, {-1, -2, -3}, {3, 2, 1}, {10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, {}, {5, 4, 3, 2, 1}, {1}}

func TestInsertion(t *testing.T) {
	fmt.Print("Increasing algorithm: ")
	for i := 0; i < len(cases); i++ {
		if sorted := sort.Insertion(cases[i]); !assertList(sorted, expected[i]) {
			fmt.Println("FAIL")
			t.Fatalf("%v: Expected %v, got %v\n", i, expected[i], sorted)
		}
	}
	fmt.Println("OK")

	fmt.Print("Increasing algorithm 2: ")
	for i := 0; i < len(cases); i++ {
		if sorted := sort.Insertion(cases[i]); !assertList(sorted, expected[i]) {
			fmt.Println("FAIL")
			t.Fatalf("%v: Expected %v, got %v\n", i, expected[i], sorted)
		}
	}
	fmt.Println("OK")

	fmt.Print("Decreasing algorithm: ")
	for i := 0; i < len(cases); i++ {
		if sorted := sort.InsertionDesc(cases[i]); !assertList(sorted, expectedDesc[i]) {
			fmt.Println("FAIL")
			t.Fatalf("%v: Expected %v, got %v\n", i, expected[i], sorted)
		}
	}
	fmt.Println("OK")
}
