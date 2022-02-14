package common

import (
	"fmt"
	"github.com/moledoc/check"
	"os"
	"strconv"
	"strings"
	"time"
)

var Cases [][]int
var Results map[string][]time.Duration = make(map[string][]time.Duration)

// AssertList checks if two integer list are the same.
func AssertList(result []int, expected []int) bool {
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

// StrToIntlst is a function that converts string to an integer list.
// Expected string format: x1,x2,x3,...,xn
func StrToIntlst(str string) []int {
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

// BenchmarkPriner is a function, that prints the benchmarks of algorithm tests.
func BenchmarkPrinter() {
	times := make([][]time.Duration, len(Results))
	keys := make([]string, len(Results))
	var i int
	for key := range Results {
		times[i] = Results[key]
		keys[i] = key
		i++
	}

	os.Remove("benchmarks.txt")
	f, err := os.OpenFile("benchmarks.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check.Err(err)
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%15s", "n"))
	for i := 0; i < len(Results); i++ {
		_, err = f.WriteString(fmt.Sprintf("%15s", keys[i]))
		check.Err(err)
	}
	_, err = f.WriteString("\n")
	check.Err(err)
	for j := 0; j < len(times[0]); j++ {
		_, err = f.WriteString(fmt.Sprintf("%15v", len(Cases[j])))
		for i := 0; i < len(keys); i++ {
			_, err = f.WriteString(fmt.Sprintf("%15s", times[i][j]))
			check.Err(err)
		}
		_, err = f.WriteString("\n")
		check.Err(err)
	}
}
