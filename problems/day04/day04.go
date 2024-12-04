package day04

import (
	"fmt"

	"alexi.ch/aoc/2024/lib"
)

var searchVectors = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

type Day04 struct {
	field [][]rune
	s1    int
	s2    int
}

func NewDay04() Day04 {
	return Day04{s1: 0, s2: 0, field: make([][]rune, 0)}
}

func (d *Day04) Title() string {
	return "Day 04 - Ceres Search"
}

func (d *Day04) Setup() {
	// var lines = lib.ReadLines("data/04-test-data.txt")
	var lines = lib.ReadLines("data/04-data.txt")
	for _, line := range lines {
		fieldLine := make([]rune, 0)
		for _, r := range line {
			fieldLine = append(fieldLine, r)
		}
		d.field = append(d.field, fieldLine)
	}
	// fmt.Printf("%#v\n", d.field)
}

func (d *Day04) SolveProblem1() {
	d.s1 = 0
	for y := 0; y < len(d.field); y++ {
		for x := 0; x < len(d.field[y]); x++ {
			if d.field[y][x] == 'X' {
				d.s1 += d.searchXmasFrom(x, y)
			}
		}
	}
}

func (d *Day04) SolveProblem2() {
	d.s2 = 0
	for y := 0; y < len(d.field); y++ {
		for x := 0; x < len(d.field[y]); x++ {
			if d.field[y][x] == 'A' {
				d.s2 += d.searchMasAsXFrom(x, y)
			}
		}
	}
}

func (d *Day04) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day04) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day04) searchXmasFrom(startX, startY int) int {
	findings := 0
	search := "XMAS"

	for _, sv := range searchVectors {
		x := startX
		y := startY
		found := true
		for i := 0; i < 4; i++ {
			// in the field:
			if y >= 0 && x >= 0 && y < len(d.field) && x < len(d.field[y]) {
				if d.field[y][x] != rune(search[i]) {
					found = false
					break
				}
			} else {
				// out of the field, not found
				found = false
				break
			}

			x += sv[0]
			y += sv[1]
		}
		if found {
			findings++
		}
	}
	return findings
}

func (d *Day04) searchMasAsXFrom(startX, startY int) int {
	/*
			 We're looking for a figure like this:

			 M.S
		     .A.
		     M.S

			 The A is always the center, and startX, startY should point
			 to an A already.
			 Then we just need to check the 2 top corners, and if the 2 bottom corners have  the opposite char.
			 for all corners.
	*/

	if d.field[startY][startX] != 'A' {
		return 0
	}
	if startX >= 1 && startY >= 1 && startX < len(d.field[0])-1 && startY < len(d.field)-1 {
		// M or S on the top corners
		if d.field[startY-1][startX-1] != 'M' && d.field[startY-1][startX-1] != 'S' {
			return 0
		}
		if d.field[startY-1][startX+1] != 'M' && d.field[startY-1][startX+1] != 'S' {
			return 0
		}

		// The opposite letter on the bottom corners
		if d.field[startY-1][startX-1] == 'M' && d.field[startY+1][startX+1] != 'S' {
			return 0
		}
		if d.field[startY-1][startX+1] == 'M' && d.field[startY+1][startX-1] != 'S' {
			return 0
		}
		if d.field[startY-1][startX-1] == 'S' && d.field[startY+1][startX+1] != 'M' {
			return 0
		}
		if d.field[startY-1][startX+1] == 'S' && d.field[startY+1][startX-1] != 'M' {
			return 0
		}
		// yep, all good, this must be an X-MAS!
		return 1
	}
	return 0
}
