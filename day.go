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
	for _, line := range readFile(os.Args[1], "\n") {
		fmt.Println(line)
	}
}

func part2() {
	for _, line := range readFile(os.Args[1], "\n") {
		fmt.Println(line)
	}
}
