package sorting_test

import (
	"bufio"
	"fmt"
	"github.com/moledoc/check"
	"github.com/moledoc/moledoc/ADS/common"
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

// TestData is a function that reads in generated test data
func TestData(t *testing.T) {
	f, err := os.Open("test_data.txt")
	defer f.Close()
	check.Err(err)
	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	defer check.Scanner(scanner)
	for scanner.Scan() {
		comps := strings.Split(scanner.Text(), ";")
		input := common.StrToIntlst(comps[0])
		output := common.StrToIntlst(comps[1])
		outputDesc := common.StrToIntlst(comps[2])
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
		if !common.AssertList(sorted, expc[i]) {
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
	common.BenchmarkPrinter()
}
