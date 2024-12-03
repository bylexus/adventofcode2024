package day04

import (
	"fmt"

	"alexi.ch/aoc/2024/lib"
)

type Day04 struct {
	s1 uint64
	s2 uint64
}

func NewDay04() Day04 {
	return Day04{s1: 0, s2: 0}
}

func (d *Day04) Title() string {
	return "Day 04 - xxxxxx"
}

func (d *Day04) Setup() {
	// var lines = lib.ReadLines("data/04-test-data.txt")
	var lines = lib.ReadLines("data/04-data.txt")
	for _, line := range lines {
		line = line
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day04) SolveProblem1() {
	d.s1 = 0
}

func (d *Day04) SolveProblem2() {
	d.s2 = 0
}

func (d *Day04) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day04) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
