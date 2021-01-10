package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s INPUT\n", os.Args[0])
		os.Exit(-1)
	}
	part1()
	part2()
}

const NUM_IPS int = 4294967296

func inRanges(i int, ranges [][2]int) bool {
	for _, r := range ranges {
		if r[0] <= i && i < r[1] {
			return true
		}
	}
	return false
}

func part1() {
	allRanges := make([][2]int, 0)
	for _, line := range readFile(os.Args[1], "\n") {
		m := strings.Split(line, "-")
		allRanges = append(allRanges, [2]int{atoi(m[0]), atoi(m[1]) + 1})
	}

	for i := range over(NUM_IPS) {
		if !inRanges(i, allRanges) {
			fmt.Println(i)
			return
		}
	}
}

func part2() {
	allRanges := make([][2]int, 0)
	for _, line := range readFile(os.Args[1], "\n") {
		m := strings.Split(line, "-")
		allRanges = append(allRanges, [2]int{atoi(m[0]), atoi(m[1]) + 1})
	}

	sort.Slice(allRanges, func(i, j int) bool {
		return allRanges[i][0] < allRanges[j][0]
	})

	mergedRanges := make([][2]int, 1)
	mergedRanges[0] = allRanges[0]
	for _, r := range allRanges[1:] {
		lr := mergedRanges[len(mergedRanges)-1]
		if r[0] <= lr[1] {
			mergedRanges[len(mergedRanges)-1][1] = intMax(lr[1], r[1])
		} else {
			mergedRanges = append(mergedRanges, r)
		}
	}

	c := 0
	for _, r := range mergedRanges {
		c += r[1] - r[0]
	}
	fmt.Println(NUM_IPS - c)
}
