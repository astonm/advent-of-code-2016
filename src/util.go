package main

import (
	"io/ioutil"
	"log"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

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

func (s *Set) len() int {
	return len(*s)
}

func (s *Set) asIntList() []int {
	out := make([]int, 0)
	for v := range *s {
		out = append(out, v.(int))
	}
	return out
}

func (s *Set) asStringList() []string {
	out := make([]string, 0)
	for v := range *s {
		out = append(out, v.(string))
	}
	return out
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

func match(regex, s string) []string {
	m := regexp.MustCompile(regex).FindAllStringSubmatch(s, -1)
	if len(m) == 0 {
		return make([]string, 0)
	}
	return m[0]
}

func replaceAll(regex, src, repl string) string {
	return regexp.MustCompile(regex).ReplaceAllString(src, repl)
}

func intAll(l []int, f func(int) bool) bool {
	for _, v := range l {
		if !f(v) {
			return false
		}
	}
	return true
}

func intAny(l []int, f func(int) bool) bool {
	for _, v := range l {
		if f(v) {
			return true
		}
	}
	return false
}

func stringAll(l []string, f func(string) bool) bool {
	for _, v := range l {
		if !f(v) {
			return false
		}
	}
	return true
}

func stringAny(l []string, f func(string) bool) bool {
	for _, v := range l {
		if f(v) {
			return true
		}
	}
	return false
}

type GridCoords struct {
	Width  int
	Height int
}

func (g GridCoords) Size() int {
	return g.Width * g.Height
}

func (g GridCoords) At(x, y int) int {
	if x < 0 || y < 0 || x >= g.Width || y >= g.Height {
		log.Fatalf("invalid coords (%v, %v)", x, y)
	}
	return y*g.Width + x
}

func (g GridCoords) Coords(p int) (int, int) {
	return p % g.Width, p / g.Width
}

func (g GridCoords) Adj(x, y int) []int {
	out := make([]int, 0)
	if y > 0 {
		out = append(out, g.At(x, y-1))
	}
	if x+1 < g.Width {
		out = append(out, g.At(x+1, y))
	}
	if y+1 < g.Height {
		out = append(out, g.At(x, y+1))
	}
	if x > 0 {
		out = append(out, g.At(x-1, y))
	}
	return out
}

func (g GridCoords) Around(x, y int) []int {
	out := make([]int, 0)
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if x+dx >= 0 && x+dx < g.Width && y+dy >= 0 && y+dy < g.Height {
				out = append(out, g.At(x+dx, y+dy))
			}
		}
	}
	return out
}

type Walker func(x int, y int, newRow bool)

func (g GridCoords) Walk(walker Walker) {
	var newRow bool
	for y := 0; y < g.Height; y++ {
		newRow = true
		for x := 0; x < g.Width; x++ {
			walker(x, y, newRow)
			newRow = false
		}
	}
}

func irange(args ...int) []int {
	var lower, upper int
	var step int = 1

	if len(args) == 1 {
		lower = 0
		upper = args[0]
	} else if len(args) == 2 {
		lower = args[0]
		upper = args[1]
	} else if len(args) == 3 {
		lower = args[0]
		upper = args[1]
		step = args[2]
	} else {
		log.Fatalf("wrong number of args %d", len(args))
	}

	out := make([]int, 0, (upper-lower)/step)
	if lower < upper {
		for i := lower; i < upper; i += step {
			out = append(out, i)
		}
	} else {
		for i := lower; i > upper; i += step {
			out = append(out, i)
		}
	}

	return out
}

func intSliceSelect(src []int, indices []int) []int {
	out := make([]int, 0, len(indices))
	for _, idx := range indices {
		out = append(out, src[idx])
	}
	return out
}

func intSliceRemove(src []int, indices []int) []int {
	skip := newSet()
	for _, i := range indices {
		skip.add(i)
	}

	out := make([]int, 0, len(indices))
	for i, v := range src {
		if !skip.has(i) {
			out = append(out, v)
		}
	}
	return out
}

func stringSliceSelect(src []string, indices []int) []string {
	out := make([]string, 0, len(indices))
	for _, idx := range indices {
		out = append(out, src[idx])
	}
	return out
}

func stringSliceRemove(src []string, indices []int) []string {
	skip := newSet()
	for _, i := range indices {
		skip.add(i)
	}

	out := make([]string, 0, len(indices))
	for i, v := range src {
		if !skip.has(i) {
			out = append(out, v)
		}
	}
	return out
}

func combinations(n, r int) <-chan ([]int) {
	// code borrowed from python's itertools pseudocode:
	// https://docs.python.org/3/library/itertools.html#itertools.combinations
	out := make(chan []int)
	pool := irange(0, n)

	go func() {
		if r > len(pool) {
			close(out)
			return
		}

		indices := irange(0, r)
		out <- intSliceSelect(pool, indices)

		for {
			var i int
			broken := false
			for i = r - 1; i >= 0; i-- {
				if indices[i] != i+n-r {
					broken = true
					break
				}
			}
			if !broken {
				close(out)
				return
			}

			indices[i]++
			for j := i + 1; j < r; j++ {
				indices[j] = indices[j-1] + 1
			}
			out <- intSliceSelect(pool, indices)
		}
	}()
	return out
}

func allCombinations(n int) <-chan ([]int) {
	out := make(chan []int)
	go func() {
		for r := 1; r <= n; r++ {
			for v := range combinations(n, r) {
				out <- v
			}
		}
		close(out)
	}()
	return out
}

func permutations(n, r int) <-chan ([]int) {
	// code borrowed from python's itertools pseudocode:
	// https://docs.python.org/3/library/itertools.html#itertools.permutations
	out := make(chan []int)
	pool := irange(0, n)

	go func() {
		if r > n {
			close(out)
			return
		}

		indices := irange(0, n)
		cycles := irange(n, n-r, -1)
		out <- intSliceSelect(pool, indices[:r])

		for n > 0 {
			broken := false
			for i := r - 1; i >= 0; i-- {
				cycles[i] -= 1
				if cycles[i] == 0 {
					indices = append(indices[:i], append(indices[i+1:], indices[i])...)
					cycles[i] = n - i
				} else {
					j := cycles[i]
					indices[i], indices[len(indices)-j] = indices[len(indices)-j], indices[i]
					out <- intSliceSelect(pool, indices[:r])
					broken = true
					break
				}
			}

			if !broken {
				close(out)
				return
			}
		}
	}()
	return out
}
func intChainIters(iterators ...<-chan ([]int)) <-chan ([]int) {
	out := make(chan []int)
	go func() {
		for _, iterator := range iterators {
			for v := range iterator {
				out <- v
			}
		}
		close(out)
	}()
	return out
}

func crt(rems, mods []int64) int64 {
	// modified from https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
	p := big.NewInt(1)
	for _, mod := range mods {
		p.Mul(p, big.NewInt(mod))
	}

	out := big.NewInt(0)
	for i, mod := range mods {
		n := big.NewInt(mod)
		q := new(big.Int).Div(p, n)

		s := new(big.Int)
		gcd := new(big.Int).GCD(nil, s, n, q)

		if gcd.Int64() != 1 {
			log.Fatalf("%d not coprime", n)
		}

		out.Add(out, s.Mul(big.NewInt(rems[i]), s.Mul(s, q)))
	}
	return out.Mod(out, p).Int64()
}

func bytesMapFunc(s []byte, f func(int, byte) byte) []byte {
	out := make([]byte, len(s))
	for i := 0; i < len(out); i++ {
		out[i] = f(i, s[i])
	}
	return out
}

func intMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func over(args ...int) <-chan int { // e.g. for x := range over(0, 10)
	var lower, upper int
	var step int = 1

	if len(args) == 1 {
		lower = 0
		upper = args[0]
	} else if len(args) == 2 {
		lower = args[0]
		upper = args[1]
	} else if len(args) == 3 {
		lower = args[0]
		upper = args[1]
		step = args[2]
	} else {
		log.Fatalf("wrong number of args %d", len(args))
	}

	out := make(chan int)
	go func() {
		for i := lower; i < upper; i += step {
			out <- i
		}
		close(out)
	}()
	return out
}
