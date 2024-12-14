package day13

import (
	"fmt"
	"math"
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
		a, b := solve2(machine)
		d.s1 += 3*a + b
	}
}

func (d *Day13) SolveProblem2() {
	d.s2 = 0
	for _, machine := range d.machines {
		machine.X += 10000000000000
		machine.Y += 10000000000000

		a, b := solve2(machine)
		d.s2 += 3*a + b
	}
}

func (d *Day13) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day13) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

// Solves a linear equation with 2 unknown variables:
// we have a linear equation system:
// each button has 2 values that increase the X/Y axis.
// The Target X and Y are both reached by pressing Button A x times + Button B y times.

// We reach X by pressing x*A + y*B = X
// We reach Y by pressing x*A + y*B = Y
// so our movements to reach X and Y can be expressed as:
//
// x*Ax + y*Bx = Prize X
// x*Ay + y*By = Prize Y

// so we can solve this with a linear system solving approach to find.
// Let's rename the variables a bit to make it simpler:
//
// ax + by = c  (a = Ax, b = Bx, c = Prize X)
// dx + ey = f  (d = Ay, e = By, f = Prize Y)

// I'm using a subsitution approach: separate y, then replace y in the other formula:
// y = -(a*x - c)/b
// x = (b*e*(c/b-f/e))/(a*e-b*d)
// So we can first calculate x without any dependency to y, then insert x to the y formula.
// Example:
//
// Input:
// Button A: X+94, Y+34
// Button B: X+22, Y+67
// Prize: X=8400, Y=5400
//
// a = 94, b = 22, c = 8400
// d = 34, e = 67, f = 5400
// --> x = (22*67*(8400/22-5400/67))/(94*67-22*34) = 80
//
//	y = -(94*80 - 8400) / 22 = 40
//
// --> when both x and y are an integer number, it is a solution.
func solve2(m Machine) (int, int) {
	a := float64(m.AX)
	b := float64(m.BX)
	c := float64(m.X)
	d := float64(m.AY)
	e := float64(m.BY)
	f := float64(m.Y)
	x := math.Round((b * e * (c/b - f/e)) / (a*e - b*d))
	y := math.Round(-(a*x - c) / b)
	if int(x)*m.AX+int(y)*m.BX == m.X && int(x)*m.AY+int(y)*m.BY == m.Y {
		return int(x), int(y)
	}
	return 0, 0
}
