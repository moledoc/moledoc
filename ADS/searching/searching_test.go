package searching_test

import (
	"bufio"
	"fmt"
	"github.com/moledoc/check"
	search "github.com/moledoc/moledoc/ADS/searching"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type templ func([]int, int) int

var cases [][]int
var casesSorted [][]int
var srchFor []int
var srchForSorted []int
var expected []int
var results map[string][]time.Duration = make(map[string][]time.Duration)

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
		inputSorted := strToIntlst(comps[1])
		srch, err := strconv.Atoi(comps[2])
		check.Err(err)
		srchSorted, err := strconv.Atoi(comps[3])
		check.Err(err)
		ind, err := strconv.Atoi(comps[4])
		check.Err(err)
		cases = append(cases, input)
		casesSorted = append(casesSorted, inputSorted)
		srchFor = append(srchFor, srch)
		srchForSorted = append(srchForSorted, srchSorted)
		expected = append(expected, ind)
	}
}

// test is a template for testing different searching algorithms.
func test(t *testing.T, fn templ, algo string, sorted bool) {
	for i := 0; i < len(cases); i++ {
		start := time.Now()
		var at int
		if sorted {
			at = fn(casesSorted[i], srchForSorted[i])
		} else {
			at = fn(cases[i], srchFor[i])
		}
		elapsed := time.Since(start)
		if at != expected[i] {
			t.Fatalf("case %v: Expected %v, got %v\n", i, expected[i], at)
		}
		results[algo][i] = elapsed
	}
}

func TestLinear(t *testing.T) {
	algo := "linear"
	results[algo] = make([]time.Duration, len(cases))
	test(t, search.Linear, algo, false)
}

func TestLinearSorted(t *testing.T) {
	algo := "linearSorted"
	results[algo] = make([]time.Duration, len(cases))
	test(t, search.Linear, algo, true)
}

func TestBinary(t *testing.T) {
	algo := "binary"
	results[algo] = make([]time.Duration, len(cases))
	test(t, search.Binary, algo, true)
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

	os.Remove("times_searching.txt")
	f, err := os.OpenFile("times_searching.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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
