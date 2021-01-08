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
	bothParts()
}

func bothParts() {
	in := readFile(os.Args[1], "->")

	text := []byte(in[0])
	size := atoi(in[1])

	for len(text) < size {
		text = doubleFlip(text)
	}
	text = text[:size]

	fmt.Println(checksum(text))
}

func doubleFlip(s []byte) []byte {
	flipped := bytesMapFunc(s, func(i int, r byte) byte {
		invert := map[byte]byte{'0': '1', '1': '0'}
		return invert[s[len(s)-1-i]]
	})
	return append(append(s, '0'), flipped...)
}

func checksum(s []byte) string {
	var cs []byte
	for {
		cs = make([]byte, 0, len(s)/2)
		for i := 0; i < len(s); i += 2 {
			if s[i] == s[i+1] {
				cs = append(cs, '1')
			} else {
				cs = append(cs, '0')
			}
		}
		if len(cs)%2 == 1 {
			break
		}
		s = cs
	}
	return string(cs)
}
