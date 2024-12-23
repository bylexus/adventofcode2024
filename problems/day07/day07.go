package day07

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"sync"
	"sync/atomic"

	"alexi.ch/aoc/2024/lib"
)

type Equation struct {
	result   uint64
	operands []uint64
}

func (e Equation) Clone() Equation {
	return Equation{result: e.result, operands: slices.Clone(e.operands)}
}

type Day07 struct {
	equations []Equation
	s1        uint64
	goodS1Eqs []uint64
	s2        uint64
}

func New() Day07 {
	return Day07{s1: 0, s2: 0, equations: make([]Equation, 0)}
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
			res := lib.StrToUint64(matches[1])
			nrstr := strings.Split(strings.Trim(matches[2], " "), " ")
			nrs := lib.Map(&nrstr, lib.StrToUint64)
			d.equations = append(d.equations, Equation{result: res, operands: nrs})
		}

	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day07) SolveProblem1() {
	d.s1 = 0
	permFunc := lib.Memoize(func(n int) [][]string {
		return lib.Permutations([]string{"*", "+"}, n)
	})
outer:
	for eqI, eq := range d.equations {
		perms := permFunc(len(eq.operands) - 1)
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
				d.goodS1Eqs = append(d.goodS1Eqs, uint64(eqI))
				continue outer
			}
		}
	}
}

func (d *Day07) SolveProblem2() {
	d.s2 = d.s1
	// permFunc := lib.PermutationsBuilder([]string{"||", "*", "+"})
	permFunc := lib.Memoize(func(n int) [][]string {
		return lib.Permutations([]string{"||", "*", "+"}, n)
	})

	wg := sync.WaitGroup{}
	sum := atomic.Uint64{}
	for eqI, eq := range d.equations {
		if slices.Contains(d.goodS1Eqs, uint64(eqI)) {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()

			perms := permFunc(len(eq.operands) - 1)
			for _, perm := range perms {
				res := eq.operands[0]
				for i := 1; i < len(eq.operands); i++ {
					op := perm[i-1]
					nr := eq.operands[i]
					if op == "+" {
						res += nr
					} else if op == "||" {
						res = lib.StrToUint64(fmt.Sprintf("%d%d", res, nr))
					} else {
						res *= nr
					}
				}
				if res == eq.result {
					sum.Add(res)
					// d.s2 += res
					// continue outer
					return
				}
			}
		}()
	}
	wg.Wait()
	d.s2 = d.s1 + sum.Load()
}

func (d *Day07) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day07) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
