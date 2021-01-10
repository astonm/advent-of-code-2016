package main

import (
	"bytes"
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
	text := []byte("abcdefgh")
	if os.Args[1][len(os.Args[1])-7:] == "example" {
		text = []byte("abcde")
	}

	lines := readFile(os.Args[1], "\n")
	fmt.Println(string(codecScramble(lines, text, ENCODE)))
}

func part2() {
	text := []byte("fbgdceah")
	if os.Args[1][len(os.Args[1])-7:] == "example" {
		text = []byte("decab")
	}

	lines := readFile(os.Args[1], "\n")
	reversedLines := make([]string, len(lines))
	for i := 0; i < len(lines); i++ {
		reversedLines[i] = lines[len(lines)-1-i]
	}
	fmt.Println(string(codecScramble(reversedLines, text, DECODE)))
}

const ENCODE = true
const DECODE = false

func codecScramble(lines []string, text []byte, encode bool) []byte {
	var m []string
	for _, line := range lines {
		m = match(`swap position (\d+) with position (\d+)`, line)
		if len(m) > 0 {
			text[atoi(m[1])], text[atoi(m[2])] = text[atoi(m[2])], text[atoi(m[1])]
		}

		m = match(`swap letter (\w) with letter (\w)`, line)
		if len(m) > 0 {
			text = bytesMapFunc(text, func(i int, b byte) byte {
				if string(b) == m[1] {
					return m[2][0]
				}
				if string(b) == m[2] {
					return m[1][0]
				}
				return b
			})
		}

		m = match(`rotate (left|right) (\d+) step`, line)
		if len(m) > 0 {
			dist := atoi(m[2])
			if m[1] == "left" {
				dist = -dist
			}
			if !encode {
				dist = -dist
			}
			text = rotateRight(text, dist)
		}

		m = match(`rotate based on position of letter (\w)`, line)
		if len(m) > 0 {
			if encode {
				dist := 1
				idx := bytes.Index(text, []byte(m[1]))
				dist += idx
				if idx >= 4 {
					dist++
				}
				text = rotateRight(text, dist)
			} else {
				// n.b. this is not fool proof! there may be times where reversal is ambiguous
				destIdx := bytes.Index(text, []byte(m[1]))

				for startIdx := 0; startIdx < len(text); startIdx++ {
					dist := 1
					dist += startIdx
					if startIdx >= 4 {
						dist++
					}
					if (startIdx+dist)%len(text) == destIdx {
						text = rotateRight(text, -dist)
						break
					}
				}
			}
		}

		m = match(`reverse positions (\d+) through (\d+)`, line)
		if len(m) > 0 {
			s, e := atoi(m[1]), atoi(m[2])
			text = bytesMapFunc(text, func(i int, b byte) byte {
				if s <= i && i <= e {
					return text[e-(i-s)]
				}
				return b
			})
		}

		m = match(`move position (\d+) to position (\d+)`, line)
		if len(m) > 0 {
			p0, p1 := atoi(m[1]), atoi(m[2])
			if !encode {
				p0, p1 = p1, p0
			}
			toInsert := []byte{text[p0]}
			tmp := append(text[:p0], text[p0+1:]...)
			text = append(tmp[:p1], append(toInsert, tmp[p1:]...)...)
		}
	}
	return text
}

func rotateRight(b []byte, offset int) []byte {
	for offset < 0 {
		offset += len(b)
	}
	offset = offset % len(b)
	left, right := b[:len(b)-offset], b[len(b)-offset:]
	return append(right, left...)
}
