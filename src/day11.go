package main

import (
	"fmt"
	"os"
	"sort"
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
	var final *State
	if len(find(".example", os.Args[1])) > 0 {
		final = solve(EXAMPLE)

		final.printPath()
		fmt.Println("------------")
	} else {
		final = solve(INPUT)
	}
	fmt.Println(final.numSteps)
}

func part2() {
	INPUT.floors[0] = append(INPUT.floors[0], []string{"EG", "EM", "DG", "DM"}...)
	fmt.Println(solve(INPUT).numSteps)
}

func solve(initial *State) (final *State) {
	queue := make([]*State, 1)
	queue[0] = initial

	seen := newSet()

	var curr *State

	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]

		if seen.has(curr.hash()) {
			continue
		}
		seen.add(curr.hash())

		done := true
		for i := 0; i < len(curr.floors)-1; i++ {
			done = done && len(curr.floors[i]) == 0
		}
		if done {
			break
		}

		for _, elevatorOffset := range []int{-1, 1} {
			nextElevatorFloor := curr.elevatorFloor + elevatorOffset
			if nextElevatorFloor < 0 || nextElevatorFloor >= len(curr.floors) {
				continue
			}

			movable := curr.floors[curr.elevatorFloor]
			n := len(movable)

			for indices := range intChainIters(combinations(n, 1), combinations(n, 2)) {
				attempt := stringSliceSelect(movable, indices)
				if validCombo(attempt) {
					next := curr.copy()

					next.numSteps = curr.numSteps + 1
					next.elevatorFloor = nextElevatorFloor
					next.floors[next.elevatorFloor] = append(next.floors[next.elevatorFloor], attempt...)
					next.floors[curr.elevatorFloor] = stringSliceRemove(next.floors[curr.elevatorFloor], indices)

					if !validCombo(next.floors[next.elevatorFloor]) || !validCombo(next.floors[curr.elevatorFloor]) {
						continue
					}

					next.prev = curr
					sort.Strings(next.floors[next.elevatorFloor])
					sort.Strings(next.floors[curr.elevatorFloor])
					queue = append(queue, next)
				}
			}

		}
	}

	return curr
}

func validCombo(combo []string) bool {
	s := newSet()
	for _, c := range combo {
		s.add(c)
	}

	numUnpairedMicrochips := 0
	numGenerators := 0
	for _, x := range combo {
		if x[1] == 'M' && !s.has(string(x[0])+"G") {
			numUnpairedMicrochips++
		}
		if x[1] == 'G' {
			numGenerators++
		}
	}
	return numUnpairedMicrochips == 0 || numGenerators == 0
}

type State struct {
	elevatorFloor int
	floors        [][]string
	numSteps      int
	prev          *State
}

func (s *State) copy() *State {
	out := State{}
	out.elevatorFloor = s.elevatorFloor
	out.floors = make([][]string, len(s.floors))
	for i := 0; i < len(s.floors); i++ {
		out.floors[i] = make([]string, len(s.floors[i]))
		copy(out.floors[i], s.floors[i])
	}
	out.numSteps = s.numSteps
	out.prev = nil
	return &out
}

func (s *State) hash() string {
	return fmt.Sprintf("%d %v", s.elevatorFloor, s.floors)
}

func (s *State) print() {
	fmt.Println()
	fmt.Printf("Step %v: E @ %v\n", s.numSteps, s.elevatorFloor)
	for i := len(s.floors) - 1; i >= 0; i-- {
		fmt.Printf("%v: %v\n", i, s.floors[i])
	}
}

func (s *State) printPath() {
	p := s
	for p != nil {
		p.print()
		if p == p.prev {
			panic("wtf")
		}
		p = p.prev
	}
}

var EXAMPLE *State = &State{
	floors: [][]string{
		[]string{"HM", "LM"},
		[]string{"HG"},
		[]string{"LG"},
		[]string{},
	},
}

var INPUT *State = &State{
	floors: [][]string{
		[]string{"SG", "SM", "PG", "PM"},
		[]string{"TG", "RG", "RM", "CG", "CM"},
		[]string{"TM"},
		[]string{},
	},
}
