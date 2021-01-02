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
	part1()
	part2()
}

func part1() {
	for _, line := range readFile(os.Args[1], "\n") {
		out := ""
		for len(line) > 0 {
			if line[0] != '(' {
				out += string(line[0])
				line = line[1:]
				continue
			}

			right_paren := strings.Index(line, ")")
			spec := strings.Split(line[1:right_paren], "x")
			line = line[right_paren+1:]

			count := atoi(spec[0])
			times := atoi(spec[1])

			repeatable := line[:count]
			out += strings.Repeat(repeatable, times)
			line = line[len(repeatable):]
		}
		fmt.Println(out, len(out))
	}
}

func decompressedLength(s string) int {
	out := 0
	for len(s) > 0 {
		if s[0] != '(' {
			out++
			s = s[1:]
			continue
		}

		right_paren := strings.Index(s, ")")
		spec := strings.Split(s[1:right_paren], "x")
		s = s[right_paren+1:]

		count := atoi(spec[0])
		times := atoi(spec[1])

		repeatable := s[:count]
		out += times * decompressedLength(repeatable)
		s = s[len(repeatable):]
	}
	return out
}

func part2() {
	for _, line := range readFile(os.Args[1], "\n") {
		fmt.Println(decompressedLength(line))
	}
}
