package sorting_test

import (
	"bufio"
	"github.com/moledoc/check"
	"github.com/moledoc/moledoc/ADS/common"
	"github.com/moledoc/moledoc/ADS/sorting"
	"os"
	"strings"
	"testing"
	"time"
)

type templ func([]int) []int

// var common.Cases [][]int
var expected [][]int
var expectedDesc [][]int

// var common.Results map[string][]time.Duration = make(map[string][]time.Duration)

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
		common.Cases = append(common.Cases, input)
		expected = append(expected, output)
		expectedDesc = append(expectedDesc, outputDesc)
	}
}

// test is a template for testing different searching algorithms.
func test(t *testing.T, fn templ, algo string, expc [][]int) {
	for i := 0; i < len(common.Cases); i++ {
		start := time.Now()
		sorted := fn(common.Cases[i])
		elapsed := time.Since(start)
		if !common.AssertList(sorted, expc[i]) {
			t.Fatalf("case %v: Expected %v, got %v\n", i, expc[i], sorted)
		}
		common.Results[algo][i] = elapsed
	}
}

func TestInsertion(t *testing.T) {
	algo := "insertion"
	common.Results[algo] = make([]time.Duration, len(common.Cases))
	test(t, sorting.Insertion, algo, expected)
}

func TestInsertion2(t *testing.T) {
	algo := "insertion2"
	common.Results[algo] = make([]time.Duration, len(common.Cases))
	test(t, sorting.Insertion2, algo, expected)
}

func TestInsertionDesc(t *testing.T) {
	algo := "insertionDesc"
	common.Results[algo] = make([]time.Duration, len(common.Cases))
	test(t, sorting.InsertionDesc, algo, expectedDesc)
}

// Testcommon.Results is a function, that prints the benchmarks of algorithm tests.
// NB! this functions needs to be last function in the file!
func TestResults(t *testing.T) {
	common.BenchmarkPrinter()
}
