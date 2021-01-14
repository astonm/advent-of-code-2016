package main

import (
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
	nodes := getNodes()

	c := 0
	for a := 0; a < len(nodes); a++ {
		for b := 0; b < len(nodes); b++ {
			if a != b && nodes[a].used > 0 && nodes[b].free >= nodes[a].used {
				c++
			}
		}
	}
	fmt.Println(c)
}

func part2() {
	nodes := getNodes()
	max := nodes[len(nodes)-1]
	g := GridCoords{max.x + 1, max.y + 1}

	q := make([]State, 1)
	q[0] = State{
		nodes:  nodes,
		dataAt: g.At(max.x, 0),
	}

	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			p := g.At(x, y)
			if q[0].nodes[p].used == 0 {
				q[0].emptyAt = p
			}
		}
	}

	seen := newSet()
	var curr State
	for len(q) > 0 {
		curr, q = q[0], q[1:]
		if seen.has(curr.hash()) {
			continue
		}
		seen.add(curr.hash())

		if curr.dataAt == 0 {
			fmt.Println(curr.steps)
			return
		}

		for _, node := range []int{curr.emptyAt, curr.dataAt} {
			for _, neighbor := range g.Adj(curr.nodes[node].x, curr.nodes[node].y) {
				nextNodes := make([]Node, len(curr.nodes))
				copy(nextNodes, curr.nodes)

				var to, from *Node
				if curr.nodes[neighbor].used > 0 && curr.nodes[node].free >= curr.nodes[neighbor].used {
					from = &nextNodes[neighbor]
					to = &nextNodes[node]
				} else if curr.nodes[node].used > 0 && curr.nodes[neighbor].free >= curr.nodes[node].used {
					from = &nextNodes[node]
					to = &nextNodes[neighbor]
				}

				if to == nil || from == nil {
					continue
				}

				to.used += from.used
				to.free -= from.used
				from.free += from.used
				from.used = 0

				nextState := State{
					nodes: nextNodes,
					steps: curr.steps + 1,
				}

				if g.At(from.x, from.y) == curr.dataAt {
					nextState.dataAt = g.At(to.x, to.y)
				} else {
					nextState.dataAt = curr.dataAt
				}

				if g.At(to.x, to.y) == curr.emptyAt {
					nextState.emptyAt = g.At(from.x, from.y)
				} else {
					nextState.emptyAt = curr.emptyAt
				}

				q = append(q, nextState)
			}
		}
	}
}

type Node struct {
	y, x int
	size int
	used int
	free int
}

func getNodes() []Node {
	nodes := make([]Node, 0)
	for _, line := range readFile(os.Args[1], "\n")[2:] {
		node := Node{}

		d := findAll(`\d+`, line)
		node.x = atoi(d[0])
		node.y = atoi(d[1])

		node.size = atoi(d[2])
		node.used = atoi(d[3])
		node.free = atoi(d[4])

		nodes = append(nodes, node)
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].y < nodes[j].y || nodes[i].y == nodes[j].y && nodes[i].x < nodes[j].x
	})

	return nodes
}

type State struct {
	nodes   []Node
	dataAt  int
	emptyAt int
	steps   int
}

func (s State) hash() string {
	return fmt.Sprintf("%v %v", s.dataAt, s.emptyAt)
}
