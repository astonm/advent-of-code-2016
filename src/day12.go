package main

import (
	"fmt"
	"os"
	"strconv"
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

	code := make([]string, 0)
	for _, line := range readFile(os.Args[1], "\n") {
		code = append(code, line)
	}

	var m []string
	for pc < len(code) {
		inst := code[pc]

		m = match(`cpy ([^\s]+) ([^\s]+)`, inst)
		if len(m) > 0 {
			if i, err := strconv.Atoi(m[1]); err == nil {
				regs[m[2]] = i
			} else {
				regs[m[2]] = regs[m[1]]
			}
		}

		m = match(`inc ([^\s]+)`, inst)
		if len(m) > 0 {
			regs[m[1]]++
		}

		m = match(`dec ([^\s]+)`, inst)
		if len(m) > 0 {
			regs[m[1]]--
		}

		m = match(`jnz ([^\s]+) ([^\s]+)`, inst)
		if len(m) > 0 {
			var jnzVal int
			if i, err := strconv.Atoi(m[1]); err == nil {
				jnzVal = i
			} else {
				jnzVal = regs[m[1]]
			}

			if jnzVal != 0 {
				pc += atoi(m[2])
				continue
			}
		}

		pc++
	}

	fmt.Println(regs)
}
