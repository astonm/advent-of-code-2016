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
	// part1()
	part2()
}

func part1() {
	for _, line := range readFile(os.Args[1], "\n") {
		fmt.Println(line, elfGame1(atoi(line))+1)
	}
}

func part2() {
	for _, line := range readFile(os.Args[1], "\n") {
		fmt.Println(line, elfGame2(atoi(line))+1)
	}
}

func elfGame1(n int) int { // zero-indexed output
	if n == 1 || n == 2 {
		return 0
	}

	var o int
	if n%2 == 0 {
		o = elfGame1(n / 2)
	} else {
		sn := (n + 1) / 2
		o = elfGame1(sn)
		o = (sn - 1 + o) % sn
	}
	return 2 * o
}

func elfGame2(n int) int { // zero-indexed output
	if n == 1 || n == 2 {
		return 0
	}

	if n == 3 {
		return 2
	}

	last := 2
	for i := 4; i <= n; i++ {

		o := last
		o += 1

		var cutoff int
		if i%2 == 1 {
			cutoff = (i - 1) / 2
		} else {
			cutoff = i / 2
		}

		if o >= cutoff {
			o = o + 1
		}

		last = o % i
	}
	return last
}
