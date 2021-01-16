package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s INPUT\n", os.Args[0])
		os.Exit(-1)
	}
	bothParts()
}

func bothParts() {
	lines := readFile(os.Args[1], "\n")
	width, height := len(lines[0]), len(lines)

	grid := strings.Join(lines, "")
	g := GridCoords{width, height}

	targets := make(map[int]int)
	for i := 0; i < len(grid); i++ {
		c := grid[i]
		if c != '.' && c != '#' {
			targets[int(c-'0')] = i
		}
	}

	distanceCache := make(map[[2]int]int)
	for i := 0; i < len(targets); i++ {
		for j := i; j < len(targets); j++ {
			if j == i {
				distanceCache[[2]int{i, j}] = 0
				continue
			}

			d := findDistance(targets[i], targets[j], grid, g)
			distanceCache[[2]int{i, j}] = d
			distanceCache[[2]int{j, i}] = d
		}
	}

	rest := irange(1, len(targets))
	shortest := MaxInt
	var shortestSeq []int
	for p := range permutations(len(rest), len(rest)) {
		seq := append([]int{0}, intSliceSelect(rest, p)...)
		d := sequenceDistance(seq, distanceCache)
		if d < shortest {
			shortest = d
			shortestSeq = seq
		}
	}
	fmt.Println("part 1:", shortest, shortestSeq)

	shortest = MaxInt
	for p := range permutations(len(rest), len(rest)) {
		seq := append(append([]int{0}, intSliceSelect(rest, p)...), 0)
		d := sequenceDistance(seq, distanceCache)
		if d < shortest {
			shortest = d
			shortestSeq = seq
		}
	}
	fmt.Println("part 2:", shortest, shortestSeq)
}

type State struct {
	pos   int
	steps int
}

func findDistance(start int, dest int, grid string, g GridCoords) int {
	q := make([]State, 1)
	q[0] = State{pos: start}

	var curr State
	seen := newSet()
	for len(q) > 0 {
		curr, q = q[0], q[1:]
		if curr.pos == dest {
			return curr.steps
		}
		if seen.has(curr.pos) {
			continue
		}
		seen.add(curr.pos)

		x, y := g.Coords(curr.pos)
		for _, n := range g.Adj(x, y) {
			if grid[n] != '#' {
				q = append(q, State{n, curr.steps + 1})
			}
		}
	}
	return -1
}

func sequenceDistance(seq []int, distanceCache map[[2]int]int) int {
	s := 0
	for i := 0; i < len(seq)-1; i++ {
		s += distanceCache[[2]int{seq[i], seq[i+1]}]
	}
	return s
}
