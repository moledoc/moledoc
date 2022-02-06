package insertion_test

import (
	"fmt"
	"testing"

	sort "github.com/moledoc/moledoc/ADS/Sorting"
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

func TestInsertion(t *testing.T) {
	cases := [][]int{{3, 2, 1}, {-1, -2, -3}, {2, 1, 3}, {5, 6, 2, 3, 1, 10, 4, 7, 8, 9}, {}, {1, 2, 3, 4, 5}, {1}}
	expected := [][]int{{1, 2, 3}, {-3, -2, -1}, {1, 2, 3}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {}, {1, 2, 3, 4, 5}, {1}}
	fmt.Print("Increasing algorithm: ")
	for i := 0; i < len(cases); i++ {
		if !assertList(sort.Insertion(cases[i]), expected[i]) {
			fmt.Println("FAIL")
			t.Fatalf("Expected %v, got %v\n", expected[i], cases[i])
		}
	}
	fmt.Println("OK")
	fmt.Print("Decreasing algorithm: ")
	for i := 0; i < len(cases); i++ {
		if !assertList(sort.InsertionDesc(cases[i]), expected[i]) {
			fmt.Println("FAIL")
			t.Fatalf("Expected %v, got %v\n", expected[i], cases[i])
		}
	}
	fmt.Println("OK")
}
