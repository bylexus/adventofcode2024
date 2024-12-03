package day03

import (
	"fmt"
	"regexp"

	"alexi.ch/aoc/2024/lib"
)

const (
	TYPE_MUL  = "mul"
	TYPE_DONT = "don't()"
	TYPE_DO   = "do()"
)

type DataEntry struct {
	entryType string
	mult1     int
	mult2     int
}

type Day03 struct {
	multiplications []DataEntry
	s1              uint64
	s2              uint64
}

func NewDay03() Day03 {
	return Day03{s1: 0, s2: 0, multiplications: make([]DataEntry, 0)}
}

func (d *Day03) Title() string {
	return "Day 03 - Mull It Over"
}

func (d *Day03) Setup() {
	// var lines = lib.ReadLines("data/03-test-data.txt")
	// var lines = lib.ReadLines("data/03-test2-data.txt")
	var lines = lib.ReadLines("data/03-data.txt")
	// matches: mul(nr,nr), or do(), or don't():
	matcher := regexp.MustCompile(`((mul)\((\d+),(\d+)\))|(do\(\))|(don't\(\))`)
	for _, line := range lines {
		matches := matcher.FindAllStringSubmatch(line, -1)
		// fmt.Printf("%#v\n", matches)
		for _, match := range matches {
			if len(match) >= 5 && match[2] == TYPE_MUL {
				e := DataEntry{entryType: TYPE_MUL, mult1: lib.StrToInt(match[3]), mult2: lib.StrToInt(match[4])}
				d.multiplications = append(d.multiplications, e)
			} else if len(match) >= 1 && match[0] == TYPE_DO {
				e := DataEntry{entryType: TYPE_DO}
				d.multiplications = append(d.multiplications, e)
			} else if len(match) >= 1 && match[0] == TYPE_DONT {
				e := DataEntry{entryType: TYPE_DONT}
				d.multiplications = append(d.multiplications, e)
			}
		}
	}
	// fmt.Printf("%#v\n", d.multiplications)
}

func (d *Day03) SolveProblem1() {
	d.s1 = 0
	for _, m := range d.multiplications {
		if m.entryType == TYPE_MUL {
			d.s1 += uint64(m.mult1 * m.mult2)
		}
	}
}

func (d *Day03) SolveProblem2() {
	d.s2 = 0
	enabled := true
	for _, m := range d.multiplications {
		if m.entryType == TYPE_MUL && enabled {
			d.s2 += uint64(m.mult1 * m.mult2)
		} else if m.entryType == TYPE_DO {
			enabled = true
		} else if m.entryType == TYPE_DONT {
			enabled = false
		}
	}
}

func (d *Day03) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day03) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
