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
	bothParts()
}

func bothParts() {
	discs := make([]int64, 0)
	offsets := make([]int64, 0)

	count := int64(0)
	for _, line := range readFile(os.Args[1], "\n") {
		m := match(`Disc #(\d+) has (\d+) positions; at time=0, it is at position (\d+).`, line)
		n := atoi(m[1])
		numPositions := atoi(m[2])
		start := atoi(m[3])

		discs = append(discs, int64(numPositions))
		offsets = append(offsets, -int64(n+start))
		count++
	}

	fmt.Println(crt(offsets, discs))

	discs = append(discs, 11)
	offsets = append(offsets, -count-1)
	fmt.Println(crt(offsets, discs))
}
