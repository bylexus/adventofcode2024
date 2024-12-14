package day15

import (
	"fmt"

	"alexi.ch/aoc/2024/lib"
)

type Day15 struct {
	s1 int
	s2 int
}

func New() Day15 {
	return Day15{s1: 0, s2: 0}
}

func (d *Day15) Title() string {
	return "Day 15 - xxxxxx"
}

func (d *Day15) Setup() {
	var lines = lib.ReadLines("data/15-test-data.txt")
	// var lines = lib.ReadLines("data/15-data.txt")
	for _, line := range lines {
		line = line
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day15) SolveProblem1() {
	d.s1 = 0
}

func (d *Day15) SolveProblem2() {
	d.s2 = 0
}

func (d *Day15) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day15) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
