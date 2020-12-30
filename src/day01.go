package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}

func rotate(p complex128) complex128 {
	return complex(-imag(p), real(p))
}

func part1() {
	for _, line := range readFile(os.Args[1], "\n") {
		directions := strings.Split(line, ", ")

		p := 0 + 0i
		v := 0 + 1i

		for _, dir := range directions {
			if dir[0] == 'L' {
				v = rotate(v)
			} else {
				v = rotate(rotate(rotate(v)))
			}

			dist := dir[1:]
			p += complex(atof(dist), 0) * v
		}

		r, i := ri(p)
		fmt.Println(math.Abs(r) + math.Abs(i))
	}
}

func part2() {
	for _, line := range readFile(os.Args[1], "\n") {
		directions := strings.Split(line, ", ")

		p := 0 + 0i
		v := 0 + 1i

		seen := newSet()
		seen.add(p)

	OUTER:
		for _, dir := range directions {
			if dir[0] == 'L' {
				v = rotate(v)
			} else {
				v = rotate(rotate(rotate(v)))
			}

			dist := dir[1:]

			for i := 0; i < atoi(dist); i++ {
				p += v

				if seen.has(p) {
					r, i := ri(p)
					fmt.Printf("repeat at (%v, %v), dist %v\n", r, i, math.Abs(r)+math.Abs(i))
					break OUTER
				}
				seen.add(p)
			}
		}
	}
}
