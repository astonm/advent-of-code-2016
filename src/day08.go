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
	g, grid := part1()
	part2(g, grid)
}

func part1() (GridCoords, []int) {
	g := GridCoords{Width: 50, Height: 6}
	grid := make([]int, g.Size())
	temp := make([]int, g.Size())
	for _, line := range readFile(os.Args[1], "\n") {
		var m []string
		copy(temp, grid)

		m = match(`rect (\d+)x(\d+)`, line)
		if len(m) > 0 {
			w := atoi(m[1])
			h := atoi(m[2])

			for x := 0; x < w; x++ {
				for y := 0; y < h; y++ {
					temp[g.At(x, y)] = 1
				}
			}
		}

		m = match(`rotate row y=(\d+) by (\d+)`, line)
		if len(m) > 0 {
			row := atoi(m[1])
			rot := atoi(m[2])

			for x := 0; x < g.Width; x++ {
				temp[g.At(x, row)] = grid[g.At((x+g.Width-rot)%g.Width, row)]
			}
		}

		m = match(`rotate column x=(\d+) by (\d+)`, line)
		if len(m) > 0 {
			col := atoi(m[1])
			rot := atoi(m[2])

			for y := 0; y < g.Height; y++ {
				temp[g.At(col, y)] = grid[g.At(col, (y+g.Height-rot)%g.Height)]
			}
		}
		grid, temp = temp, grid
	}
	c := 0
	for i := 0; i < g.Size(); i++ {
		c += grid[i]
	}
	fmt.Println(c)
	return g, grid
}

func part2(g GridCoords, grid []int) {
	g.Walk(func(x int, y int, newRow bool) {
		if newRow {
			fmt.Println()
		}

		if grid[g.At(x, y)] == 1 {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
	})
	fmt.Println("\n")
}
