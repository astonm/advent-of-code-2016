package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s INPUT\n", os.Args[0])
		os.Exit(-1)
	}
	part1()
	part2()
}

func part1() {
	var counters []stringCounter
	for _, line := range readFile(os.Args[1], "\n") {
		if counters == nil {
			counters = make([]stringCounter, len(line))
			for i := 0; i < len(line); i++ {
				counters[i] = make(stringCounter)
			}
		}

		for i, c := range line {
			counters[i].count(string(c))
		}
	}

	for i := 0; i < len(counters); i++ {
		mc := counters[i].mostCommon()
		fmt.Print(mc[0])
	}
	fmt.Println()
}

func part2() {
	var counters []stringCounter
	for _, line := range readFile(os.Args[1], "\n") {
		if counters == nil {
			counters = make([]stringCounter, len(line))
			for i := 0; i < len(line); i++ {
				counters[i] = make(stringCounter)
			}
		}

		for i, c := range line {
			counters[i].count(string(c))
		}
	}

	for i := 0; i < len(counters); i++ {
		mc := counters[i].mostCommon()
		fmt.Print(mc[len(mc)-1])
	}
	fmt.Println()
}
