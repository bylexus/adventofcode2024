package day13

import (
	"fmt"
	"regexp"

	"alexi.ch/aoc/2024/lib"
)

type Machine struct {
	AX int
	AY int
	BX int
	BY int
	X  int
	Y  int
}

type Day13 struct {
	s1 int
	s2 int

	machines []Machine
}

func New() Day13 {
	return Day13{s1: 0, s2: 0}
}

func (d *Day13) Title() string {
	return "Day 13 - Claw Contraption"
}

func (d *Day13) Setup() {
	// var lines = lib.ReadLines("data/13-test-data.txt")
	var lines = lib.ReadLines("data/13-data.txt")

	btnMatcher := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
	priceMatcher := regexp.MustCompile(`X=(\d+), Y=(\d+)`)
	for l := 0; l < len(lines); l += 4 {
		btnA := btnMatcher.FindStringSubmatch(lines[l])
		btnB := btnMatcher.FindStringSubmatch(lines[l+1])
		price := priceMatcher.FindStringSubmatch(lines[l+2])

		machine := Machine{
			AX: lib.StrToInt(btnA[1]),
			AY: lib.StrToInt(btnA[2]),
			BX: lib.StrToInt(btnB[1]),
			BY: lib.StrToInt(btnB[2]),
			X:  lib.StrToInt(price[1]),
			Y:  lib.StrToInt(price[2]),
		}
		d.machines = append(d.machines, machine)
	}

	// fmt.Printf("%#v\n", d.machines)
}

func (d *Day13) SolveProblem1() {
	d.s1 = 0
	for _, machine := range d.machines {
		a, b := solve(machine)
		d.s1 += 3*a + b
	}
}

func (d *Day13) SolveProblem2() {
	d.s2 = 0
	for _, machine := range d.machines {
		machine.X += 10000000000000
		machine.Y += 10000000000000

		a, b := solve(machine)
		d.s2 += 3*a + b
	}
}

func (d *Day13) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day13) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

// Solves the system of linear equations using Cramer's Rule
// See https://en.wikipedia.org/wiki/Cramer's_rule#Explicit_formulas_for_small_systems
func solve(m Machine) (int, int) {
	A, B := 0, 0
	if (m.AY*m.BX - m.AX*m.BY) != 0 {
		B = ((m.AY * m.X) - (m.AX * m.Y)) / (m.AY*m.BX - m.AX*m.BY)
	}
	if m.AX != 0 {
		A = (m.X - (B * m.BX)) / m.AX
	}

	if ((m.AX*A)+(m.BX*B)) == m.X && ((m.AY*A)+(m.BY*B)) == m.Y {
		return A, B
	}

	return 0, 0
}
