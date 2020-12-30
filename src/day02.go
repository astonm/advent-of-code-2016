package main

import (
	"fmt"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {
	num := 5
	for _, line := range readFile(os.Args[1], "\n") {
		for _, c := range line {
			if c == 'U' && num-3 > 0 {
				num -= 3
			}
			if c == 'D' && num+3 < 10 {
				num += 3
			}
			if c == 'L' && num%3 != 1 {
				num -= 1
			}
			if c == 'R' && num%3 != 0 {
				num += 1
			}
		}
		fmt.Print(num)
	}
	fmt.Println()
}

func part2() {
	inds := map[rune]int{
		'U': 0,
		'D': 1,
		'L': 2,
		'R': 3,
	}

	dirs := map[int][]int{ // num => UDLR
		1:  []int{-1, 3, -1, -1},
		2:  []int{-1, 6, -1, 3},
		3:  []int{1, 7, 2, 4},
		4:  []int{-1, 8, 3, -1},
		5:  []int{-1, -1, -1, 6},
		6:  []int{2, 10, 5, 7},
		7:  []int{3, 11, 6, 8},
		8:  []int{4, 12, 7, 9},
		9:  []int{-1, -1, 8, -1},
		10: []int{6, -1, -1, 11},
		11: []int{7, 13, 10, 12},
		12: []int{8, -1, 11, -1},
		13: []int{11, -1, -1, -1},
	}

	num := 5
	for _, line := range readFile(os.Args[1], "\n") {
		for _, c := range line {
			ind := inds[c]
			got := dirs[num][ind]
			if got != -1 {
				num = got
			}
		}
		fmt.Printf("%X", num)
	}
	fmt.Println()
}
