package main

import (
	"fmt"
	"math/bits"
	"os"
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

func part1() {
	lines := readFile(os.Args[1], "\n")
	favNum := atoi(lines[0])
	wallAt := wallFunc(favNum)

	start := State{x: 1, y: 1}
	destXY := strings.Split(lines[1], ",")
	dest := State{x: atoi(destXY[0]), y: atoi(destXY[1])}

	fmt.Println(shortestDistance(start, dest, wallAt))
}

func part2() {
	lines := readFile(os.Args[1], "\n")
	favNum := atoi(lines[0])
	wallAt := wallFunc(favNum)

	start := State{x: 1, y: 1}
	fmt.Println(countLocations(start, 50, wallAt))
}

type State struct {
	x     int
	y     int
	steps int
}

func (s State) hash() int {
	return s.x<<32 + s.y
}

type WallAtFunc func(x int, y int) bool

func wallFunc(n int) WallAtFunc {
	return func(x int, y int) bool {
		p := x*x + 3*x + 2*x*y + y + y*y
		p += n
		return bits.OnesCount64(uint64(p))%2 == 1
	}
}

func shortestDistance(start State, dest State, wallAt WallAtFunc) int {
	q := make([]State, 1)
	q[0] = start

	seen := newSet()

	var curr State
	for len(q) > 0 {
		curr, q = q[0], q[1:]

		if curr.x == dest.x && curr.y == dest.y {
			return curr.steps
		}

		nextSteps := []State{
			State{curr.x - 1, curr.y - 0, curr.steps + 1},
			State{curr.x - 0, curr.y - 1, curr.steps + 1},
			State{curr.x + 1, curr.y + 0, curr.steps + 1},
			State{curr.x + 0, curr.y + 1, curr.steps + 1},
		}

		for _, next := range nextSteps {
			if next.x >= 0 && next.y >= 0 && !wallAt(next.x, next.y) && !seen.has(next.hash()) {
				seen.add(next.hash())
				q = append(q, next)
			}
		}
	}
	return -1
}

func countLocations(start State, maxSteps int, wallAt WallAtFunc) int {
	q := make([]State, 1)
	q[0] = start

	seen := newSet()

	var curr State
	for len(q) > 0 {
		curr, q = q[0], q[1:]
		if curr.steps >= maxSteps {
			continue
		}

		nextSteps := []State{
			State{curr.x - 1, curr.y - 0, curr.steps + 1},
			State{curr.x - 0, curr.y - 1, curr.steps + 1},
			State{curr.x + 1, curr.y + 0, curr.steps + 1},
			State{curr.x + 0, curr.y + 1, curr.steps + 1},
		}

		for _, next := range nextSteps {
			if next.x >= 0 && next.y >= 0 && !wallAt(next.x, next.y) && !seen.has(next.hash()) {
				seen.add(next.hash())
				q = append(q, next)
			}
		}
	}
	return seen.len()
}
