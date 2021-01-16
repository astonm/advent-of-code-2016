package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s INPUT\n", os.Args[0])
		os.Exit(-1)
	}
	theEnd()
}

func theEnd() {
	lines := readFile(os.Args[1], "\n")
	checkDepth := 42
	for n := 0; ; n++ {
		seq := make([]int, 0, checkDepth)
		runComputer(lines, n, func(out int) bool {
			seq = append(seq, out)
			return len(seq) == checkDepth
		})

		good := true
		for i := 0; i < checkDepth; i++ {
			good = good && seq[i] == i%2
		}
		if good {
			fmt.Println(n, seq)
			return
		}
	}
}

func runComputer(lines []string, initialValue int, output func(int) bool) {
	regs := make(map[string]int)
	regs["a"] = initialValue

	pc := 0

	code := make([][]string, 0)
	for _, line := range lines {
		code = append(code, strings.Split(line, " "))
	}

	for pc < len(code) {
		inst := code[pc]

		if inst[0] == "cpy" {
			if i, err := strconv.Atoi(inst[1]); err == nil {
				regs[inst[2]] = i
			} else {
				regs[inst[2]] = regs[inst[1]]
			}
		}

		if inst[0] == "inc" {
			regs[inst[1]]++
		}

		if inst[0] == "dec" {
			regs[inst[1]]--
		}

		if inst[0] == "jnz" {
			var jnzVal int
			if i, err := strconv.Atoi(inst[1]); err == nil {
				jnzVal = i
			} else {
				jnzVal = regs[inst[1]]
			}

			if jnzVal != 0 {
				var jump int
				if i, err := strconv.Atoi(inst[2]); err == nil {
					jump = i
				} else {
					jump = regs[inst[2]]
				}

				pc += jump
				continue
			}
		}

		if inst[0] == "out" {
			var shouldStop bool
			if i, err := strconv.Atoi(inst[1]); err == nil {
				shouldStop = output(i)
			} else {
				shouldStop = output(regs[inst[1]])
			}

			if shouldStop {
				return
			}
		}

		pc++
	}

	fmt.Println(regs)
}
