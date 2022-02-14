package main

import (
	"fmt"
	"github.com/moledoc/check"
	"github.com/moledoc/walks"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var root string

var timings []string

func dummy(path string) {
}

func getTimings(path string) {
	if strings.Contains(path, "benchmarks.txt") {
		timings = append(timings, path)
	}
}

func printOptions() {
	for i, elem := range timings {
		fmt.Printf("%v: %v\n", i, elem)
	}
	fmt.Println("-1: quit")
}

func readTimings(filename string) {
	contents, err := ioutil.ReadFile(filename)
	check.Err(err)
	fmt.Println(string(contents))
}

func main() {
	rootLoc, err := os.Getwd()
	check.Err(err)
	root = strings.Replace(rootLoc, "\\", "/", -1)
	// Walk current working directory recursively.
	walks.WalkLinear(rootLoc, getTimings, dummy, -1, 0)
	fmt.Println("Show benchmarks")
	var option string
	n := len(timings)
	for {
		printOptions()
		fmt.Scanln(&option)
		opt, err := strconv.Atoi(option)
		if err != nil || opt >= n || opt < -1 {
			fmt.Println("Incorrect option")
			continue
		}
		if opt == -1 {
			break
		}
		readTimings(timings[opt])
	}
}
