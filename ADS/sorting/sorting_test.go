package sorting_test

import (
	"bufio"
	"fmt"
	"github.com/moledoc/check"
	"github.com/moledoc/moledoc/ADS/sorting"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type templ func([]int) []int

var cases [][]int
var expected [][]int
var expectedDesc [][]int
var results map[string][]time.Duration = make(map[string][]time.Duration)

// assertList checks if two integer list are the same.
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

// strToIntlst is a function that converts string to an integer list.
// Expected string format: x1,x2,x3,...,xn
func strToIntlst(str string) []int {
	if str == "" {
		var tmp []int
		return tmp
	}
	lstStr := strings.Split(str, ",")
	lst := make([]int, len(lstStr))
	for i, elem := range lstStr {
		nr, err := strconv.Atoi(elem)
		check.Err(err)
		lst[i] = nr
	}
	return lst
}

// TestData is a function that reads in generated test data
func TestData(t *testing.T) {
	f, err := os.Open("test_data.txt")
	defer f.Close()
	check.Err(err)
	scanner := bufio.NewScanner(f)
	defer check.Scanner(scanner)
	for scanner.Scan() {
		comps := strings.Split(scanner.Text(), ";")
		input := strToIntlst(comps[0])
		output := strToIntlst(comps[1])
		outputDesc := strToIntlst(comps[2])
		cases = append(cases, input)
		expected = append(expected, output)
		expectedDesc = append(expectedDesc, outputDesc)
	}
}

// test is a template for testing different searching algorithms.
func test(t *testing.T, fn templ, algo string, expc [][]int) {
	for i := 0; i < len(cases); i++ {
		start := time.Now()
		sorted := fn(cases[i])
		elapsed := time.Since(start)
		if !assertList(sorted, expc[i]) {
			t.Fatalf("case %v: Expected %v, got %v\n", i, expc[i], sorted)
		}
		results[algo][i] = elapsed
	}
}

func TestInsertion(t *testing.T) {
	algo := "insertion"
	results[algo] = make([]time.Duration, len(cases))
	test(t, sorting.Insertion, algo, expected)
}

func TestInsertion2(t *testing.T) {
	algo := "insertion2"
	results[algo] = make([]time.Duration, len(cases))
	test(t, sorting.Insertion2, algo, expected)
}

func TestInsertionDesc(t *testing.T) {
	algo := "insertionDesc"
	results[algo] = make([]time.Duration, len(cases))
	test(t, sorting.InsertionDesc, algo, expectedDesc)
}

// TestResults is a function, that prints the benchmarks of algorithm tests.
// NB! this functions needs to be last function in the file!
func TestResults(t *testing.T) {
	times := make([][]time.Duration, len(results))
	keys := make([]string, len(results))
	var i int
	for key := range results {
		times[i] = results[key]
		keys[i] = key
		i++
	}

	os.Remove("times_sorting.txt")
	f, err := os.OpenFile("times_sorting.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check.Err(err)
	defer f.Close()

	for i := 0; i < len(results); i++ {
		_, err = f.WriteString(fmt.Sprintf("%15s", keys[i]))
		check.Err(err)
	}
	_, err = f.WriteString("\n")
	check.Err(err)
	for j := 0; j < len(times[0]); j++ {
		for i := 0; i < len(keys); i++ {
			_, err = f.WriteString(fmt.Sprintf("%15s", times[i][j]))
			check.Err(err)
		}
		_, err = f.WriteString("\n")
		check.Err(err)
	}
}
