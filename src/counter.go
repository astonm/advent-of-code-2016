package main

import (
	"sort"
)

type intCounter map[int]int

func (c intCounter) count(val int) {
	_, exists := c[val]
	if !exists {
		c[val] = 0
	}
	c[val] += 1
}

func (c intCounter) countAll(vals []int) {
	for _, v := range vals {
		c.count(v)
	}
}

func (c intCounter) mostCommon() []int {
	keys := make([]int, 0, len(c))
	for k, _ := range c {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if c[keys[i]] > c[keys[j]] {
			return true
		}
		if c[keys[i]] == c[keys[j]] {
			return keys[i] < keys[j]
		}
		return false
	})
	return keys
}

type stringCounter map[string]int

func (c stringCounter) count(val string) {
	_, exists := c[val]
	if !exists {
		c[val] = 0
	}
	c[val] += 1
}

func (c stringCounter) countAll(vals []string) {
	for _, v := range vals {
		c.count(v)
	}
}

func (c stringCounter) mostCommon() []string {
	keys := make([]string, 0, len(c))
	for k, _ := range c {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if c[keys[i]] > c[keys[j]] {
			return true
		}
		if c[keys[i]] == c[keys[j]] {
			return keys[i] < keys[j]
		}
		return false
	})
	return keys
}
