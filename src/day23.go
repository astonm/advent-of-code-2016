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
	part1()
	part2()
}

func part1() {
	regs := make(map[string]int)
	regs["a"] = 7

	runComputer(readFile(os.Args[1], "\n"), regs)
}

func part2() {
	regs := make(map[string]int)
	regs["a"] = 12

	runComputer(readFile(os.Args[1], "\n"), regs)
}

func runComputer(lines []string, regs map[string]int) {
	pc := 0

	code := make([][]string, 0)
	for _, line := range lines {
		code = append(code, strings.Split(line, " "))
	}

	for pc < len(code) {
		inst := code[pc]

		if dest, r1, r2, skip := detectMulAdd(code, pc); dest != "" {
			regs[dest] += regs[r1] * regs[r2]
			regs[r1] = 0
			regs[r2] = 0
			pc += skip
			continue
		}

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

		if inst[0] == "mul" {
			regs[inst[1]] = regs[inst[1]] * regs[inst[2]]
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

		if inst[0] == "tgl" {
			togglePc := pc + regs[inst[1]]
			if togglePc >= 0 && togglePc < len(code) {
				toToggle := code[togglePc]

				if len(toToggle) == 2 {
					if toToggle[0] == "inc" {
						toToggle[0] = "dec"
					} else {
						toToggle[0] = "inc"
					}
				}

				if len(toToggle) == 3 {
					if toToggle[0] == "jnz" {
						toToggle[0] = "cpy"
					} else {
						toToggle[0] = "jnz"
					}
				}
			}
		}
		pc++
	}

	fmt.Println(regs)
}

func detectMulAdd(code [][]string, pc int) (dest string, r1 string, r2 string, skip int) {
	instrs := code[pc:]
	if len(instrs) >= 5 {
		var dest, r1, r2 string
		if instrs[0][0] == "inc" {
			dest = instrs[0][1]
		}
		if instrs[1][0] == "dec" && instrs[2][0] == "jnz" && instrs[2][2] == "-2" {
			if instrs[1][1] == instrs[2][1] {
				r1 = instrs[1][1]
			}
		}
		if instrs[3][0] == "dec" && instrs[4][0] == "jnz" && instrs[4][2] == "-5" {
			if instrs[3][1] == instrs[4][1] {
				r2 = instrs[3][1]
			}
		}

		if dest != "" && r1 != "" && r2 != "" {
			return dest, r1, r2, 5
		}
	}
	return "", "", "", 0
}
