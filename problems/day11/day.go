package day11

import (
	"fmt"

	"alexi.ch/aoc/2024/lib"
)

type Day11 struct {
	s1 int
	s2 int
}

func New() Day11 {
	return Day11{s1: 0, s2: 0}
}

func (d *Day11) Title() string {
	return "Day 11 - xxxxxx"
}

func (d *Day11) Setup() {
	// var lines = lib.ReadLines("data/11-test-data.txt")
	var lines = lib.ReadLines("data/11-data.txt")
	for _, line := range lines {
		line = line
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day11) SolveProblem1() {
	d.s1 = 0
}

func (d *Day11) SolveProblem2() {
	d.s2 = 0
}

func (d *Day11) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day11) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
