package main

import (
	"crypto/md5"
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
		start := State{x: 0, y: 0}
		dest := State{x: 3, y: 3}
		fmt.Println(bestPath(start, dest, line, intMin))
	}
}

func part2() {
	for _, line := range readFile(os.Args[1], "\n") {
		start := State{x: 0, y: 0}
		dest := State{x: 3, y: 3}
		fmt.Println(len(bestPath(start, dest, line, intMax)))
	}
}

type State struct {
	x    int
	y    int
	path string
}

func (s State) hash() string {
	return fmt.Sprintf("%v,%v,%v", s.x, s.y, s.path)
}

func bestPath(start State, dest State, passcode string, choosePath func(int, int) int) string {
	q := make([]State, 1)
	q[0] = start

	seen := newSet()
	best := ""

	var curr State
	for len(q) > 0 {
		curr, q = q[0], q[1:]

		if curr.x == dest.x && curr.y == dest.y {
			if best == "" || choosePath(len(curr.path), len(best)) != len(best) {
				best = curr.path
			}
			continue
		}

		nextSteps := []State{
			State{curr.x - 0, curr.y - 1, curr.path + "U"},
			State{curr.x + 0, curr.y + 1, curr.path + "D"},
			State{curr.x - 1, curr.y - 0, curr.path + "L"},
			State{curr.x + 1, curr.y + 0, curr.path + "R"},
		}

		s := fmt.Sprintf("%x", md5.Sum([]byte(passcode+curr.path)))
		canGo := make([]bool, 4)
		for i, c := range s[:4] {
			canGo[i] = c >= 'b'
		}

		for dir, next := range nextSteps {
			if next.x >= 0 && next.y >= 0 && next.x < 4 && next.y < 4 {
				if canGo[dir] && !seen.has(next.hash()) {
					seen.add(next.hash())
					q = append(q, next)
				}
			}
		}
	}
	return best
}
