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

func p(x, y int) string {
	return itoa(x) + "," + itoa(y)
}

func part1() {
	trapSet := newSet()
	trapSet.add("^^.")
	trapSet.add(".^^")
	trapSet.add("^..")
	trapSet.add("..^")

	for _, firstRow := range readFile(os.Args[1], "\n") {
		size := len(firstRow)
		grid := make(map[string]string)
		for y := 0; y < size; y++ {
			for x := 0; x < size; x++ {
				if y == 0 {
					grid[p(x, y)] = string(firstRow[x])
				} else {
					l := (grid[p(x-1, y-1)] + ".")[:1]
					c := (grid[p(x+0, y-1)] + ".")[:1]
					r := (grid[p(x+1, y-1)] + ".")[:1]

					if trapSet.has(l + c + r) {
						grid[p(x, y)] = "^"
					} else {
						grid[p(x, y)] = "."
					}
				}
				fmt.Print(grid[p(x, y)])
			}
			fmt.Println()
		}

		safeCount := 0
		for y := 0; y < 40; y++ {
			for x := 0; x < size; x++ {
				if grid[p(x, y)] == "." {
					safeCount++
				}
			}
		}
		fmt.Println(safeCount)
	}
}

func part2() {
	trapSet := newSet()
	trapSet.add("^^.")
	trapSet.add(".^^")
	trapSet.add("^..")
	trapSet.add("..^")

	for _, firstRow := range readFile(os.Args[1], "\n") {
		size := len(firstRow)
		lastLine := make(map[int]string)

		safeCount := 0
		for y := 0; y < 400000; y++ {
			nextLine := make(map[int]string)
			for x := 0; x < size; x++ {
				if y == 0 {
					nextLine[x] = string(firstRow[x])
				} else {
					l := (lastLine[x-1] + ".")[:1]
					c := (lastLine[x+0] + ".")[:1]
					r := (lastLine[x+1] + ".")[:1]

					if trapSet.has(l + c + r) {
						nextLine[x] = "^"
					} else {
						nextLine[x] = "."
					}
				}

				if nextLine[x] == "." {
					safeCount++
				}
			}
			lastLine = nextLine
		}
		fmt.Println(safeCount)
	}
}
