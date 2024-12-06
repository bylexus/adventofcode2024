package day07

import (
	"fmt"

	"alexi.ch/aoc/2024/lib"
)

type Day07 struct {
	s1 int
	s2 int
}

func New() Day07 {
	return Day07{s1: 0, s2: 0}
}

func (d *Day07) Title() string {
	return "Day 07 - xxx"
}

func (d *Day07) Setup() {
	var lines = lib.ReadLines("data/07-test-data.txt")
	// var lines = lib.ReadLines("data/07-data.txt")
	for _, line := range lines {
		line = line
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day07) SolveProblem1() {
	d.s1 = 0
}

func (d *Day07) SolveProblem2() {
	d.s2 = 0
}

func (d *Day07) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day07) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
