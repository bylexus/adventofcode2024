package day10

import (
	"fmt"

	"alexi.ch/aoc/2024/lib"
)

type Day10 struct {
	s1          int
	s2          int
	topologyMap [][]int

	trailHeads map[lib.Coord]int
	// trailEnds  map[lib.Coord]int

}

func New() Day10 {
	return Day10{s1: 0, s2: 0, topologyMap: make([][]int, 0), trailHeads: make(map[lib.Coord]int)}
}

func (d *Day10) Title() string {
	return "Day 10 - Hoof It"
}

func (d *Day10) Setup() {
	// var lines = lib.ReadLines("data/10-test-data.txt")
	var lines = lib.ReadLines("data/10-data.txt")
	for y, line := range lines {
		xLine := make([]int, 0)
		for x, r := range line {
			nr := lib.StrToInt(string(r))
			coord := lib.NewCoord2D(x, y)
			if nr == 0 {
				d.trailHeads[coord] = 0
			}
			xLine = append(xLine, nr)
		}
		d.topologyMap = append(d.topologyMap, xLine)
	}
	// fmt.Printf("%v\n", d.topologyMap)
}

func (d *Day10) SolveProblem1() {
	d.s1 = 0
	for startPoint := range d.trailHeads {
		targetCoords := d.ramble(startPoint)
		// we only need unique end coords per start point, so create a set:
		targetCoordSet := make(map[lib.Coord]int)
		for _, coord := range targetCoords {
			targetCoordSet[coord] += 1
		}
		d.s1 += len(targetCoordSet)
	}
}

func (d *Day10) SolveProblem2() {
	d.s2 = 0
	// Solution 2 in this case is even simpler than solution 1 :-))
	// We do not need unique end coords, but count all reached end coords:
	for startPoint := range d.trailHeads {
		targetCoords := d.ramble(startPoint)
		// fmt.Printf("Start : %#v, reached ends: %#v, score: %d\n\n", startPoint, targetCoords, len(targetCoords))
		d.s2 += len(targetCoords)
	}
}

func (d *Day10) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day10) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day10) ramble(start lib.Coord) []lib.Coord {
	actNr := d.topologyMap[start.Y][start.X]
	allTargetCoords := make([]lib.Coord, 0)

	// recursive break condition: If we sit on a 9, we have reached
	// our destination: we return the coords of that destination
	if actNr == 9 {
		return []lib.Coord{start}
	}

	// if not, we walk in all 4 dirs, if possible, and collect
	// target coords:
	for _, dirVec := range lib.MOVE_VEC_2D_4DIRS {
		nextCoord := start.Add(lib.NewCoord2D(dirVec[0], dirVec[1]))
		if nextCoord.X >= 0 && nextCoord.Y >= 0 && nextCoord.X < len(d.topologyMap[0]) && nextCoord.Y < len(d.topologyMap) {
			nextNr := d.topologyMap[nextCoord.Y][nextCoord.X]
			if actNr+1 == nextNr {
				resultCoords := d.ramble(nextCoord)
				allTargetCoords = append(allTargetCoords, resultCoords...)
			}
		}
	}

	return allTargetCoords
}
