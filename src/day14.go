package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"sort"
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
	search(hash)
}

func part2() {
	search(stretchedHash)
}

func hash(salt string, counter int) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(salt+itoa(counter))))
}

func stretchedHash(salt string, counter int) string {
	out := fmt.Sprintf("%x", md5.Sum([]byte(salt+itoa(counter))))
	for i := 0; i < 2016; i++ {
		out = fmt.Sprintf("%x", md5.Sum([]byte(out)))
	}
	return out
}

func stretch() {}

func findRepeated(s string, repeats int) rune {
	prev := rune(0)
	count := 0
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if prev == c {
			count++
		} else {
			prev = c
			count = 1
		}

		if count >= repeats {
			return c
		}
	}
	return rune(0)
}

func search(hasher func(string, int) string) {
	salt := readFile(os.Args[1], "\n")[0]

	triples := make(map[rune][]int)
	keyIndexes := newSet()

	maxIndex := -1
	for i := 0; i < maxIndex || maxIndex == -1; i++ {
		got := hasher(salt, i)
		triple := findRepeated(got, 3)
		if triple != 0 {
			triples[triple] = append(triples[triple], i)
		}

		quint := findRepeated(got, 5)
		if quint != 0 {
			prev := triples[quint]
			for _, idx := range prev {
				if 0 < i-idx && i-idx <= 1000 {
					keyIndexes.add(idx)
					if keyIndexes.len() == 64 {
						maxIndex = i + 1000 // look for any stragglers
					}
				}
			}
		}
	}

	keyIndexList := keyIndexes.asIntList()
	sort.Ints(keyIndexList)
	fmt.Println(keyIndexList[63])
}
