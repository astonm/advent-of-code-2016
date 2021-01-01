package main

import (
	"crypto/md5"
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

var hex = "0123456789abcdef"

func nextCharacter(prefix string, start int) (int, string) {
	for {
		s := md5.Sum([]byte(prefix + itoa(start)))
		if s[0] == 0 && s[1] == 0 && s[2]&0xf0 == 0 {
			return start + 1, string(hex[s[2]&0xf])
		}
		start++
	}
}

func part1() {
	for _, line := range readFile(os.Args[1], "\n") {
		out := ""
		start := 0
		var c string
		for i := 0; i < 8; i++ {
			start, c = nextCharacter(line, start)
			out += c
		}
		fmt.Println(out)
	}
}

func nextCharacterAndPosition(prefix string, start int) (int, int, string) {
	for {
		s := md5.Sum([]byte(prefix + itoa(start)))
		if s[0] == 0 && s[1] == 0 && s[2]&0xf0 == 0 {
			return start + 1, int(s[2] & 0xf), string(hex[s[3]&0xf0>>4])
		}
		start++
	}
}

func part2() {
	for _, line := range readFile(os.Args[1], "\n") {
		out := []string{"_", "_", "_", "_", "_", "_", "_", "_"}
		start := 0
		var ind int
		var c string
		for i := 0; i < 8; {
			start, ind, c = nextCharacterAndPosition(line, start)
			if ind < len(out) && out[ind] == "_" {
				out[ind] = c
				i++
			}
		}
		fmt.Println(strings.Join(out, ""))
	}
}
