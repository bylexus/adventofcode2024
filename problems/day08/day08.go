package day08

import (
	"fmt"

	"alexi.ch/aoc/2024/lib"
)

type Day08 struct {
	width  int
	height int
	// keep all antennas as map of rune to coordinates:
	// this allows easy access to all coords for a given antenna type
	antennas map[rune][]lib.Coord
	// keep track of the antinodes, mark each coordinate with the number of times it was seen
	antinotes map[lib.Coord]int
	s1        int
	s2        int
}

func New() Day08 {
	return Day08{s1: 0, s2: 0, antennas: make(map[rune][]lib.Coord), antinotes: make(map[lib.Coord]int)}
}

func (d *Day08) Title() string {
	return "Day 08 - Resonant Collinearity"
}

func (d *Day08) Setup() {
	// var lines = lib.ReadLines("data/08-test-data.txt")
	var lines = lib.ReadLines("data/08-data.txt")
	d.height = len(lines)
	d.width = len(lines[0])
	for y, line := range lines {
		for x, r := range line {
			if r != '.' {
				d.antennas[r] = append(d.antennas[r], lib.NewCoord2D(x, y))
			}
		}
	}
	// fmt.Printf("%#v\n", d.antennas)
}

func (d *Day08) SolveProblem1() {
	d.s1 = 0
	// just loop over all coords of the same antenna(s),
	// and calculate the antinodes location: Go through all coords,
	// and calc the double distance from A to B, and mark it in the
	// antinodes map:
	for _, coords := range d.antennas {
		for i := 0; i < len(coords); i++ {
			for j := 0; j < len(coords); j++ {
				coordA := coords[i]
				coordB := coords[j]
				if coordA == coordB {
					continue
				}
				distX := coordB.X - coordA.X
				distY := coordB.Y - coordA.Y
				coordB.X += distX
				coordB.Y += distY
				if coordB.X < 0 || coordB.Y < 0 || coordB.X >= d.width || coordB.Y >= d.height {
					continue
				}
				d.antinotes[coordB]++
			}
		}
	}
	d.s1 = len(d.antinotes)
}

func (d *Day08) SolveProblem2() {
	d.s2 = 0
	d.antinotes = make(map[lib.Coord]int)

	// again just loop over all coords of the same antenna(s),
	// and calculate the antinodes location: Go through all coords,
	// and calc ALL the distances from A to B, until it is not on the map anymore, and mark it in the:
	//
	// seems way too easy for a day 08 problem?
	for _, coords := range d.antennas {
		for i := 0; i < len(coords); i++ {
			for j := 0; j < len(coords); j++ {
				coordA := coords[i]
				coordB := coords[j]
				if coordA == coordB {
					continue
				}
				d.antinotes[coordA]++
				d.antinotes[coordB]++

				distX := coordB.X - coordA.X
				distY := coordB.Y - coordA.Y
				for {
					coordB.X += distX
					coordB.Y += distY
					if coordB.X < 0 || coordB.Y < 0 || coordB.X >= d.width || coordB.Y >= d.height {
						break
					}
					d.antinotes[coordB]++
				}
			}
		}
	}
	d.s2 = len(d.antinotes)
}

func (d *Day08) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day08) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
