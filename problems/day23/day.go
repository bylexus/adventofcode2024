package day23

import (
	"fmt"
	"regexp"
	"slices"

	"alexi.ch/aoc/2024/lib"
)

// defines a map of node names to a list of connected nodes
// so for each node, we know its neighbours
type EdgeMap map[string][]string

type Day23 struct {
	s1      int
	s2      int
	edgeMap EdgeMap
}

func New() Day23 {
	return Day23{s1: 0, s2: 0, edgeMap: make(EdgeMap)}
}

func (d *Day23) Title() string {
	return "Day 23 - LAN Party"
}

func (d *Day23) Setup() {
	// var lines = lib.ReadLines("data/23-test-data.txt")
	var lines = lib.ReadLines("data/23-data.txt")
	matcher := regexp.MustCompile(`^(.*)-(.*)$`)
	for _, line := range lines {
		match := matcher.FindStringSubmatch(line)
		if match != nil {
			left := match[1]
			right := match[2]
			d.edgeMap[left] = append(d.edgeMap[left], right)
			d.edgeMap[right] = append(d.edgeMap[right], left)
		}
	}
	// fmt.Printf("%v\n", d.edgeMap)
}

func (d *Day23) SolveProblem1() {
	nodeTrios := make(map[string][]string)
	d.s1 = 0
	for node, myEdges := range d.edgeMap {
		if len(myEdges) >= 2 {
			for _, otherNode := range myEdges {
				otherNodesEdges := d.edgeMap[otherNode]
				for _, otherNodesEdge := range otherNodesEdges {
					if slices.Contains(myEdges, otherNodesEdge) {
						trio := []string{node, otherNode, otherNodesEdge}
						slices.Sort(trio)
						nodeTrios[fmt.Sprintf("%v", trio)] = trio
					}
				}
			}
		}
	}

	fmt.Printf("trios: %d: %v\n", len(nodeTrios), nodeTrios)
	solution := make([][]string, 0)
	for _, trio := range nodeTrios {
		for _, node := range trio {
			if node[0] == 't' {
				solution = append(solution, trio)
				break
			}
		}

	}
	fmt.Printf("with chief: %d: %v\n", len(solution), solution)
	d.s1 = len(solution)

}

func (d *Day23) SolveProblem2() {
	d.s2 = 0
}

func (d *Day23) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day23) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
