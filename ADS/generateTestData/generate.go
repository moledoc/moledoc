package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/moledoc/check"
)

// prepForFile prepares int list for a file by converting the list into a string.
func prepForFile(prep []int) string {
	done := make([]string, len(prep))
	for i, elem := range prep {
		done[i] = strconv.Itoa(elem)
	}
	return strings.Join(done, ",")
}

// sortingTestData is a function that generates a line of test data for sorting algorithms.
func sortingTestData(n int, limit int) string {
	if n == 0 {
		return ";;\n"
	}
	p := rand.Intn(100)
	nrs := make([]int, n)
	for i := 0; i < n; i++ {
		sign := 1
		if rand.Intn(100) < p {
			sign = -1
		}
		nrs[i] = sign * rand.Intn(limit)
	}
	input := prepForFile(nrs)
	sort.Ints(nrs[:])
	output := prepForFile(nrs)
	sort.Sort(sort.Reverse(sort.IntSlice(nrs)))
	outputDesc := prepForFile(nrs)
	return fmt.Sprintf("%v;%v;%v\n", input, output, outputDesc)
}

// notIn generates a number, that is not in the given int list.
func notIn(nrs []int, limit int) int {
	for {
		found := false
		attempt := rand.Intn(limit)
		for _, elem := range nrs {
			if elem == attempt {
				found = true
				break
			}
		}
		if !found {
			return attempt
		}
	}
	return -1
}

// searchingTestData is a function that generates a line of test data for searching algorithms.
func searchingTestData(n int, limit int) string {
	if n == 0 {
		return ";;0;0;-1\n"
	}
	exist := make(map[int]bool)
	p := rand.Intn(100)
	nrs := make([]int, n)
	for i := 0; i < n; {
		sign := 1
		if rand.Intn(100) < p {
			sign = -1
		}
		nr := sign * rand.Intn(limit)
		if !exist[nr] {
			nrs[i] = nr
			exist[nr] = true
			i++
		}
	}
	input := prepForFile(nrs)
	var ind int
	var srch int
	var srchSorted int
	if rand.Intn(2) == 1 {
		ind = rand.Intn(n)
		srch = nrs[ind]
		sort.Ints(nrs[:])
		srchSorted = nrs[ind]
	} else {
		ind = -1
		srch = notIn(nrs, limit)
		srchSorted = srch
	}
	sort.Ints(nrs[:])
	inputSorted := prepForFile(nrs)
	return fmt.Sprintf("%v;%v;%v;%v;%v\n", input, inputSorted, srch, srchSorted, ind)
}

func appendToFile(filename string, line string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check.Err(err)
	defer f.Close()
	_, err = f.WriteString(line)
	check.Err(err)
}

func main() {
	nFlag := flag.Int("n", 10, "How many rows of test data is generated")
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	n := *nFlag

	sortingData := "./sorting/test_data.txt"
	searchingData := "./searching/test_data.txt"
	// No error handling, since we do not care, if the file existed or not, we just want to delete it.
	// Basically imitating rm -f
	os.Remove(sortingData)
	os.Remove(searchingData)
	limit := rand.Intn(1 << 20)
	appendToFile(sortingData, sortingTestData(0, limit))
	appendToFile(searchingData, searchingTestData(0, limit))
	for i := 1; i <= n; i *= 10 {
		rand.Seed(time.Now().UTC().UnixNano())
		if i <= 100000 {
			appendToFile(sortingData, sortingTestData(i, limit))
		}
		appendToFile(searchingData, searchingTestData(i, limit))
	}
	fmt.Println("[INFO] Test data generated")
	fmt.Printf("\t%v\n", sortingData)
	fmt.Printf("\t%v\n", searchingData)
}
