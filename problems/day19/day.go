package day19

import (
	"fmt"
	"slices"
	"strings"

	"alexi.ch/aoc/2024/lib"
)

type Day19 struct {
	s1       int
	s2       int
	towels   []string
	patterns []string
}

func New() Day19 {
	return Day19{s1: 0, s2: 0}
}

func (d *Day19) Title() string {
	return "Day 19 - Linen Layout"
}

func (d *Day19) Setup() {
	// var lines = lib.ReadLines("data/19-test-data.txt")
	var lines = lib.ReadLines("data/19-data.txt")
	towels := strings.Split(lines[0], ",")
	towels = lib.Map(&towels, strings.TrimSpace)
	d.towels = towels
	for _, line := range lines[2:] {
		if len(line) > 0 {
			d.patterns = append(d.patterns, line)
		}
	}
	// fmt.Printf("%v\n", d.towels)
	// fmt.Printf("%v\n", d.patterns)
}

func (d *Day19) SolveProblem1() {
	d.s1 = 0

	// results := make([][]string, 0)
	// p := d.patterns[5]
	// p := d.patterns[0]
	for _, p := range d.patterns {
		towels := make([]string, 0)
		for _, towel := range d.towels {
			if strings.Contains(p, towel) {
				towels = append(towels, towel)
			}
		}
		slices.SortFunc(towels, func(i, j string) int {
			return len(i) - len(j)
		})
		// fmt.Printf("pattern: %s, %#v\n", p, towels)
		// memory := make(map[string][][]string)
		// results := d.findTowels(p, memory, towels)
		// if len(results) > 0 {
		// 	d.s1++
		// }
		memory := make(map[string]bool)
		results := d.hasMatch(p, memory, towels)
		if results {
			// fmt.Printf("Good pattern: %s\n", p)
			d.s1++
		} else {
			// fmt.Printf("Bad pattern: %s\n", p)
		}
		// fmt.Printf("pattern: %s, %#v\n", p, results)
		// fmt.Printf("len: %d\n", len(results))
	}
}

func (d *Day19) SolveProblem2() {
	d.s2 = 0
}

func (d *Day19) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day19) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day19) findTowels(pattern string, memory map[string][][]string, towels []string) [][]string {
	// fmt.Printf("pattern: %s\n", pattern)
	// if val, ok := memory[pattern]; ok {
	// 	return val
	// }
	if len(pattern) == 0 {
		return nil
	}
	newRes := make([][]string, 0)
	found := false
	for _, towel := range towels {
		if len(pattern) < len(towel) {
			continue
		}
		if strings.Index(pattern, towel) == 0 {
			nextPattern := pattern[len(towel):]
			var ret [][]string = nil
			if len(nextPattern) > 0 {
				ret = d.findTowels(nextPattern, memory, towels)
				if ret == nil {
					continue
				}
			}
			found = true
			for i, res := range ret {
				if len(res) > 0 {
					res = append([]string{towel}, res...)
				} else {
					res = []string{towel}
				}
				ret[i] = res
			}
			if len(ret) == 0 {
				ret = [][]string{{towel}}
			}
			newRes = append(newRes, ret...)
		}
	}
	// memory[pattern] = newRes
	if found {
		return newRes
	} else {
		return nil
	}
}

func (d *Day19) hasMatch(pattern string, memory map[string]bool, towels []string) bool {
	// fmt.Printf("pattern: %s\n", pattern)
	if val, ok := memory[pattern]; ok {
		return val
	}
	if len(pattern) == 0 {
		return false
	}
	// found := false
	for _, towel := range towels {
		if len(pattern) < len(towel) {
			continue
		}
		if strings.Index(pattern, towel) == 0 {
			if len(pattern) == len(towel) {
				memory[pattern] = true
				return true
			}
			nextPattern := pattern[len(towel):]
			// var ret bool = false
			if len(nextPattern) > 0 {
				ret := d.hasMatch(nextPattern, memory, towels)
				if ret {
					memory[pattern] = true
					return true
				}
			}
			// return false
		}
	}
	memory[pattern] = false
	return false
}
