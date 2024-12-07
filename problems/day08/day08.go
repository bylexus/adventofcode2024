package day08

import (
	"fmt"

	"alexi.ch/aoc/2024/lib"
)

type Day08 struct {
	s1 int
	s2 int
}

func New() Day08 {
	return Day08{s1: 0, s2: 0}
}

func (d *Day08) Title() string {
	return "Day 08 - xxxxxx"
}

func (d *Day08) Setup() {
	var lines = lib.ReadLines("data/08-test-data.txt")
	// var lines = lib.ReadLines("data/08-data.txt")
	for _, line := range lines {
		line = line
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day08) SolveProblem1() {
	d.s1 = 0
}

func (d *Day08) SolveProblem2() {
	d.s2 = 0
}

func (d *Day08) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day08) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
