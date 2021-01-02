package main

import (
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
	c := 0
	for _, line := range readFile(os.Args[1], "\n") {
		if supportsTLS(line) {
			c++
		}
	}
	fmt.Println(c)
}

func part2() {
	c := 0
	for _, line := range readFile(os.Args[1], "\n") {
		if supportsSSL(line) {
			c++
		}
	}
	fmt.Println(c)
}

func hasABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			return true
		}
	}
	return false
}

func supportsTLS(addr string) bool {
	return hasABBA(addr) && !stringAny(findAll(`\[\w+\]`, addr), hasABBA)
}

func findABAs(s string) []string {
	abas := make([]string, 0)
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] && s[i+1] != '[' {
			abas = append(abas, s[i:i+3])
		}
	}
	return abas
}

func supportsSSL(addr string) bool {
	bracedRegex := `\[([a-z]+)\]`
	for _, aba := range findABAs(replaceAll(bracedRegex, addr, "[]")) {
		bab := string(aba[1]) + string(aba[0]) + string(aba[1])

		foundBAB := stringAny(findAll(bracedRegex, addr), func(s string) bool {
			return find(bab, s) != ""
		})
		if foundBAB {
			return true
		}
	}
	return false
}
