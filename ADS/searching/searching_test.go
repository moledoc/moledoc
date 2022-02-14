package searching_test

import (
	"bufio"
	"github.com/moledoc/check"
	"github.com/moledoc/moledoc/ADS/common"
	"github.com/moledoc/moledoc/ADS/searching"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type templ func([]int, int) int

// var common.Cases [][]int
var CasesSorted [][]int
var srchFor []int
var srchForSorted []int
var expected []int

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
		inputSorted := common.StrToIntlst(comps[1])
		srch, err := strconv.Atoi(comps[2])
		check.Err(err)
		srchSorted, err := strconv.Atoi(comps[3])
		check.Err(err)
		ind, err := strconv.Atoi(comps[4])
		check.Err(err)
		common.Cases = append(common.Cases, input)
		CasesSorted = append(CasesSorted, inputSorted)
		srchFor = append(srchFor, srch)
		srchForSorted = append(srchForSorted, srchSorted)
		expected = append(expected, ind)
	}
}

// test is a template for testing different searching algorithms.
func test(t *testing.T, fn templ, algo string, sorted bool) {
	for i := 0; i < len(common.Cases); i++ {
		start := time.Now()
		var at int
		if sorted {
			at = fn(CasesSorted[i], srchForSorted[i])
		} else {
			at = fn(common.Cases[i], srchFor[i])
		}
		elapsed := time.Since(start)
		if at != expected[i] {
			t.Fatalf("case %v: Expected %v, got %v\n", i, expected[i], at)
		}
		common.Results[algo][i] = elapsed
	}
}

func TestLinear(t *testing.T) {
	algo := "linear"
	common.Results[algo] = make([]time.Duration, len(common.Cases))
	test(t, searching.Linear, algo, false)
}

func TestLinearSorted(t *testing.T) {
	algo := "linearSorted"
	common.Results[algo] = make([]time.Duration, len(common.Cases))
	test(t, searching.Linear, algo, true)
}

func TestBinary(t *testing.T) {
	algo := "binary"
	common.Results[algo] = make([]time.Duration, len(common.Cases))
	test(t, searching.Binary, algo, true)
}

// Testcommon.Results is a function, that prints the benchmarks of algorithm tests.
// NB! this functions needs to be last function in the file!
func TestResults(t *testing.T) {
	common.BenchmarkPrinter()
}
