package problems

import (
	"fmt"

	"alexi.ch/aoc/2024/lib"
)

type Day03 struct {
	s1 uint64
	s2 uint64
}

func NewDay03() Day03 {
	return Day03{s1: 0, s2: 0}
}

func (d *Day03) Title() string {
	return "Day 03 - xxxxx"
}

func (d *Day03) Setup() {
	// var lines = lib.ReadLines("data/03-test-data.txt")
	var lines = lib.ReadLines("data/03-data.txt")
	for _, line := range lines {
		line = line
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day03) SolveProblem1() {
	d.s1 = 0
}

func (d *Day03) SolveProblem2() {
	d.s2 = 0
}

func (d *Day03) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day03) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
