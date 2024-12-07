package day07

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"alexi.ch/aoc/2024/lib"
	"github.com/bylexus/go-stdlib/efunctional"
)

type Equation struct {
	result   int
	operands []int
}

func (e Equation) Clone() Equation {
	return Equation{result: e.result, operands: slices.Clone(e.operands)}
}

type Day07 struct {
	permutations map[int][][]string
	equations    []Equation
	s1           int
	s2           int
}

func New() Day07 {
	return Day07{s1: 0, s2: 0, permutations: make(map[int][][]string), equations: make([]Equation, 0)}
}

func (d *Day07) Title() string {
	return "Day 07 - Bridge Repair"
}

func (d *Day07) Setup() {
	// var lines = lib.ReadLines("data/07-test-data.txt")
	var lines = lib.ReadLines("data/07-data.txt")
	matcher := regexp.MustCompile(`(\d+):(.*)`)
	for _, line := range lines {
		matches := matcher.FindStringSubmatch(line)
		if matches != nil {
			res := lib.StrToInt(matches[1])
			nrstr := strings.Split(strings.Trim(matches[2], " "), " ")
			nrs := lib.Map(&nrstr, lib.StrToInt)
			d.equations = append(d.equations, Equation{result: res, operands: nrs})
		}

	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day07) SolveProblem1() {
	d.s1 = 0
	// perms := d.getPermutations(3, []string{"+", "*"})
	// fmt.Printf("%#v\n", perms)
outer:
	for _, eq := range d.equations {
		perms := d.getPermutations(len(eq.operands)-1, []string{"+", "*"})
		for _, perm := range perms {
			res := eq.operands[0]
			for i := 1; i < len(eq.operands); i++ {
				op := perm[i-1]
				nr := eq.operands[i]
				if op == "+" {
					res += nr
				} else {
					res *= nr
				}
			}
			if res == eq.result {
				d.s1 += res
				continue outer
			}
		}
	}
}

func (d *Day07) SolveProblem2() {
	d.s2 = 0
	d.permutations = make(map[int][][]string)
outer:
	for _, eq := range d.equations {
		perms := d.getPermutations(len(eq.operands)-1, []string{"+", "*", "||"})
		for _, perm := range perms {
			res := eq.operands[0]
			for i := 1; i < len(eq.operands); i++ {
				op := perm[i-1]
				nr := eq.operands[i]
				if op == "+" {
					res += nr
				} else if op == "||" {
					res = lib.StrToInt(fmt.Sprintf("%d%d", res, nr))
				} else {
					res *= nr
				}
			}
			if res == eq.result {
				d.s2 += res
				continue outer
			}
		}
	}
}

func (d *Day07) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day07) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day07) getPermutations(n int, values []string) [][]string {
	res := make([][]string, 0)
	if n <= 0 {
		return res
	}

	if perms, ok := d.permutations[n]; ok {
		return perms
	}

	singlePerms := make([]string, 0, len(values))
	singlePerms = slices.Concat(singlePerms, values)

	if n == 1 {
		for _, perm := range values {
			res = append(res, []string{perm})
		}
		d.permutations[n] = res
	} else {
		prevPerms := d.getPermutations(n-1, values)
		for _, perm := range singlePerms {
			for _, prevPerm := range prevPerms {
				newPerm := slices.Concat(prevPerm, []string{perm})
				res = append(res, newPerm)
			}
		}
		d.permutations[n] = res
	}

	return res
}

func concatEq(eq Equation, ops []string) (Equation, []string) {
	if idx := slices.Index(ops, "||"); idx >= 0 {
		eq = eq.Clone()
		ops = slices.Clone(ops)
		eq.operands[idx] = lib.StrToInt(fmt.Sprintf("%d%d", eq.operands[idx], eq.operands[idx+1]))
		eq.operands[idx+1] = -1
		ops[idx] = "-1"
		eq.operands = efunctional.Filter(eq.operands, func(i int) bool { return i != -1 })
		ops = efunctional.Filter(ops, func(i string) bool { return i != "-1" })
		return concatEq(eq, ops)
	}

	return eq, ops
}