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
	runIt(0)
	runIt(1)
}

func runIt(initialC int) {
	regs := make(map[string]int)
	regs["c"] = initialC

	pc := 0

	code := make([][]string, 0)
	for _, line := range readFile(os.Args[1], "\n") {
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
				pc += atoi(inst[2])
				continue
			}
		}

		pc++
	}

	fmt.Println(regs)
}
