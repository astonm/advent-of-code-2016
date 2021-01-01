package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func readFile(path, delim string) []string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	contents := strings.TrimSpace(string(b))

	return strings.Split(contents, delim)
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func atof(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func ri(c complex128) (float64, float64) {
	return real(c), imag(c)
}

type Set map[interface{}]struct{}

func newSet() *Set {
	s := make(Set)
	return &s
}

func (s *Set) add(elm interface{}) {
	(*s)[elm] = struct{}{}
}

func (s *Set) has(elm interface{}) bool {
	_, exists := (*s)[elm]
	return exists
}

func (s *Set) remove(elm interface{}) {
	delete(*s, elm)
}

func findAll(regex, s string) []string {
	return regexp.MustCompile(regex).FindAllString(s, -1)
}

func find(regex, s string) string {
	all := regexp.MustCompile(regex).FindAllString(s, -1)
	if len(all) > 0 {
		return all[0]
	}
	return ""
}
