package day11

import (
	"fmt"
	"strconv"
	"strings"

	"alexi.ch/aoc/2024/lib"
)

type Day11 struct {
	s1 int
	s2 int
	// map of stone nr to stone count (so many stones with that nr)
	stones map[int]int
}

func New() Day11 {
	return Day11{s1: 0, s2: 0, stones: make(map[int]int)}
}

func (d *Day11) Title() string {
	return "Day 11 - Plutonian Pebbles"
}

func (d *Day11) Setup() {
	// var lines = lib.ReadLines("data/11-test-data.txt")
	var lines = lib.ReadLines("data/11-data.txt")

	nrStrs := strings.Split(lines[0], " ")
	for _, s := range nrStrs {
		nr := lib.StrToInt(s)
		d.stones[nr]++
	}

	// fmt.Printf("%v\n", d.stones)
}

func (d *Day11) SolveProblem1() {
	d.s1 = 0
	rounds := 25
	stoneCount := 0
	initialStones := d.stones
	for range rounds {
		stoneCount = 0
		newStones := make(map[int]int)
		for nr, count := range d.stones {
			str := strconv.Itoa(nr)
			stoneCount += count
			if nr == 0 {
				newStones[1] = newStones[1] + count
			} else if len(str)%2 == 0 {
				nr1 := lib.StrToInt(str[:len(str)/2])
				nr2 := lib.StrToInt(str[len(str)/2:])
				newStones[nr1] = newStones[nr1] + count
				newStones[nr2] = newStones[nr2] + count
				stoneCount += count
			} else {
				newNr := nr * 2024
				newStones[newNr] = newStones[newNr] + count
			}
		}
		d.stones = newStones
	}
	d.stones = initialStones

	d.s1 = stoneCount
}

func (d *Day11) SolveProblem2() {
	d.s2 = 0
	rounds := 75
	stoneCount := 0
	initialStones := d.stones
	for range rounds {
		stoneCount = 0
		newStones := make(map[int]int)
		for nr, count := range d.stones {
			str := strconv.Itoa(nr)
			stoneCount += count
			if nr == 0 {
				newStones[1] = newStones[1] + count
			} else if len(str)%2 == 0 {
				nr1 := lib.StrToInt(str[:len(str)/2])
				nr2 := lib.StrToInt(str[len(str)/2:])
				newStones[nr1] = newStones[nr1] + count
				newStones[nr2] = newStones[nr2] + count
				stoneCount += count
			} else {
				newNr := nr * 2024
				newStones[newNr] = newStones[newNr] + count
			}
		}
		d.stones = newStones
	}
	d.stones = initialStones
	d.s2 = stoneCount
}

func (d *Day11) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day11) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
