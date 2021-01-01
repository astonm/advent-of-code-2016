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
	sum := 0

	for _, line := range readFile(os.Args[1], "\n") {
		data := line[:len(line)-7]
		checksum := line[len(line)-6 : len(line)-1]

		c := make(stringCounter)
		for _, s := range findAll(`[a-z]`, data) {
			c.count(s)
		}

		top := c.mostCommon()[:5]
		if strings.Join(top, "") == checksum {
			sum += atoi(find(`\d+`, data))
		}
	}
	fmt.Println(sum)
}

func rotate(c string, offset int) string {
	offset = offset % 26
	r0 := c[0]
	r1 := ((r0 - 'a') + byte(offset)) % 26
	return string(r1 + 'a')
}

func part2() {
	for _, line := range readFile(os.Args[1], "\n") {
		data := line[:len(line)-7]

		sectorId := atoi(find(`\d+`, data))
		text := ""
		for _, s := range findAll(`[^\d]`, data) {
			if s == "-" {
				text += " "
			} else {
				text += rotate(s, sectorId)
			}
		}

		if strings.Index(text, "north") > -1 {
			fmt.Println(text, sectorId)
		}
	}
}
