package day13

import (
	"cmp"
	"fmt"
	"regexp"
	"slices"

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
		a := 0
		matchingResults := make([][]int, 0)
		maxA := lib.Min(100, lib.Min(machine.X/machine.AX+1, machine.Y/machine.AY+1))
		for a <= maxA {
			b := lib.Max((machine.X-a*machine.AX)/machine.BX, (machine.Y-a*machine.AY)/machine.BY)
			maxB := lib.Min(100, lib.Min(machine.X/machine.BX+1, machine.Y/machine.BY+1))
			for b <= maxB {
				if (a*machine.AX+b*machine.BX == machine.X) && (a*machine.AY+b*machine.BY == machine.Y) {
					matchingResults = append(matchingResults, []int{a, b})
				}
				b += 1
			}
			a += 1
		}

		var minRes []int
		if len(matchingResults) > 0 {
			minRes = slices.MinFunc(matchingResults, func(a, b []int) int {
				resA := 3*a[0] + a[1]
				resB := 3*b[0] + b[1]
				return cmp.Compare(resA, resB)
			})

		}
		res := 0
		if minRes != nil {
			res = 3*minRes[0] + minRes[1]
		}
		d.s1 += res
		// fmt.Printf("Machine: %#v, Min costing presses: %#v, tokens: %d\n", machine, minRes, res)
	}
}

func (d *Day13) SolveProblem2() {
	d.s2 = 0
	// for _, machine := range d.machines {
	// 	machine.X += 10000000000000
	// 	machine.Y += 10000000000000
	// 	a := 0
	// 	matchingResults := make([][]int, 0)
	// 	maxA := lib.Min(machine.X/machine.AX+1, machine.Y/machine.AY+1)
	// 	for a <= maxA {
	// 		b := lib.Max((machine.X-a*machine.AX)/machine.BX, (machine.Y-a*machine.AY)/machine.BY)
	// 		maxB := lib.Min(machine.X/machine.BX+1, machine.Y/machine.BY+1)
	// 		for b <= maxB {
	// 			if (a*machine.AX+b*machine.BX == machine.X) && (a*machine.AY+b*machine.BY == machine.Y) {
	// 				matchingResults = append(matchingResults, []int{a, b})
	// 			}
	// 			b += 1
	// 		}
	// 		a += 1
	// 	}

	// 	var minRes []int
	// 	if len(matchingResults) > 0 {
	// 		minRes = slices.MinFunc(matchingResults, func(a, b []int) int {
	// 			resA := 3*a[0] + a[1]
	// 			resB := 3*b[0] + b[1]
	// 			return cmp.Compare(resA, resB)
	// 		})

	// 	}
	// 	res := 0
	// 	if minRes != nil {
	// 		res = 3*minRes[0] + minRes[1]
	// 	}
	// 	d.s2 += res
	// 	// fmt.Printf("Machine: %#v, Min costing presses: %#v, tokens: %d\n", machine, minRes, res)
	// }

}

func (d *Day13) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day13) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
