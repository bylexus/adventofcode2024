package day10

import (
	"fmt"

	"alexi.ch/aoc/2024/lib"
)

type Day10 struct {
	s1 int
	s2 int
}

func New() Day10 {
	return Day10{s1: 0, s2: 0}
}

func (d *Day10) Title() string {
	return "Day 10 - xxxxxx"
}

func (d *Day10) Setup() {
	var lines = lib.ReadLines("data/10-test-data.txt")
	// var lines = lib.ReadLines("data/10-data.txt")
	for _, line := range lines {
		line = line
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day10) SolveProblem1() {
	d.s1 = 0
}

func (d *Day10) SolveProblem2() {
	d.s2 = 0
}

func (d *Day10) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day10) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
