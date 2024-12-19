package day19

import (
	"fmt"
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

	for _, p := range d.patterns {
		towels := make([]string, 0)
		for _, towel := range d.towels {
			if strings.Contains(p, towel) {
				towels = append(towels, towel)
			}
		}
		memory := make(map[string]int)
		results := d.countMatch(p, memory, towels)
		if results > 0 {
			d.s1++
			d.s2 += results
		}
	}
}

func (d *Day19) SolveProblem2() {
	// already solved in part 1!

}

func (d *Day19) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day19) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day19) countMatch(pattern string, memory map[string]int, towels []string) int {
	if val, ok := memory[pattern]; ok {
		return val
	}
	count := 0
	for _, towel := range towels {
		if len(pattern) < len(towel) {
			continue
		}
		if strings.Index(pattern, towel) == 0 {
			if len(pattern) == len(towel) {
				count++
				continue
			}
			nextPattern := pattern[len(towel):]
			if len(nextPattern) > 0 {
				ret := d.countMatch(nextPattern, memory, towels)
				if ret > 0 {
					count += ret
				}
			}
		}
	}
	memory[pattern] = count
	return count
}
