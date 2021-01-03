package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s INPUT\n", os.Args[0])
		os.Exit(-1)
	}
	bothParts()
}

type State map[string][]int

func put(s State, val int, dest string) {
	_, exists := s[dest]
	if !exists {
		s[dest] = make([]int, 0)
	}
	s[dest] = append(s[dest], val)
}

func bothParts() {
	state := make(State)
	insts := make([]string, 0)
	for _, line := range readFile(os.Args[1], "\n") {
		insts = append(insts, line)
	}

	var part1 string

	for len(insts) > 0 {
		inst := insts[0]
		insts = insts[1:]

		var m []string

		m = match(`value (\d+) goes to (bot \d+)`, inst)
		if len(m) > 0 {
			put(state, atoi(m[1]), m[2])
			continue
		}

		m = match(`(bot \d+) gives (low|high) to (bot \d+|output \d+) and (low|high) to (bot \d+)`, inst)
		if len(m) > 0 {
			srcVals := state[m[1]]
			if len(srcVals) < 2 {
				insts = append(insts, inst)
				continue
			}

			sort.Ints(srcVals)
			firstGive := 0
			if m[2] == "high" {
				firstGive = 1
			}

			if m[2] == m[4] {
				log.Fatalf("got %v and %v from instruction %v, expected them to be opposites", m[2], m[4], inst)
			}

			put(state, srcVals[firstGive], m[3])
			put(state, srcVals[1-firstGive], m[5])
			state[m[1]] = make([]int, 0)
		}

		for bot, chips := range state {
			if len(chips) >= 2 {
				if chips[0] == 61 && chips[1] == 17 || chips[0] == 61 && chips[1] == 17 {
					part1 = bot
				}
			}
		}
	}

	fmt.Println(part1)
	fmt.Println(state["output 0"][0] * state["output 1"][0] * state["output 2"][0])
}
