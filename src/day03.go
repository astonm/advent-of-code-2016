package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	combos := [][]int{
		[]int{0, 1, 2},
		[]int{0, 2, 1},
		[]int{1, 2, 0},
	}

	c := 0
	for _, line := range readFile(os.Args[1], "\n") {
		line = strings.TrimSpace(line)
		n := findAll(`\d+`, line)

		good := true
		for _, combo := range combos {
			good = good && (atoi(n[combo[0]])+atoi(n[combo[1]]) > atoi(n[combo[2]]))
		}
		if good {
			c += 1
		}
	}
	fmt.Println(c)
}

func part2() {
	combos := [][]int{
		[]int{0, 1, 2},
		[]int{0, 2, 1},
		[]int{1, 2, 0},
	}

	c := 0
	var all_nums []int

	for _, line := range readFile(os.Args[1], "no break") {
		for _, s := range findAll(`\d+`, line) {
			all_nums = append(all_nums, atoi(s))
		}

		for i := 0; i < len(all_nums)/9; i++ {
			for j := 0; j < 3; j++ {
				b := i*9 + j

				base := all_nums[b:]
				good := true
				for _, combo := range combos {
					good = good && base[combo[0]*3]+base[combo[1]*3] > base[combo[2]*3]
				}
				if good {
					c += 1
				}
			}
		}

	}
	fmt.Println(c)
}
