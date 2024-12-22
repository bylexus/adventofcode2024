package day22

import (
	"fmt"

	"alexi.ch/aoc/2024/lib"
)

type DataPoint struct {
	hash    int
	digit   int
	diff    int
	diffSeq []int
}

type Day22 struct {
	s1     int
	s2     int
	inputs []int
}

func New() Day22 {
	return Day22{s1: 0, s2: 0}
}

func (d *Day22) Title() string {
	return "Day 22 - Monkey Market"
}

func (d *Day22) Setup() {
	// var lines = lib.ReadLines("data/22-test-data.txt")
	// var lines = lib.ReadLines("data/22-test2-data.txt")
	var lines = lib.ReadLines("data/22-data.txt")
	for _, line := range lines {
		if len(line) > 0 {
			d.inputs = append(d.inputs, lib.StrToInt(line))
		}
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day22) SolveProblem1() {
	d.s1 = 0
	// straight forward: just calc 2000 hashes, then sum them up
	for _, i := range d.inputs {
		h := i
		for range 2000 {
			h = hash(h)
		}
		d.s1 += h
		// fmt.Printf("%d: %d\n", i, h)
	}
}

func (d *Day22) SolveProblem2() {
	d.s2 = 0
	// We need a bit of book-keeping:
	// for each input, we calc some values:
	// - the hash
	// - the digit (last byte of the hash)
	// - the difference to the previous digit
	// - the last 4 differences, the diffSeq
	//
	// then we also sum up the digits for each unique diffSequences: This is the amount of
	// bananas we get for each unique sequence.
	//
	// We also need to make sure to just take the FIRST occurence of each sequence per input.
	// In the end we have a hash map of sequences => points, and we just need to find the maximum points.

	// keep a sum of all digit points per unique sequence:
	sequencePoints := make(map[string]int)

	for _, input := range d.inputs {
		h := input
		// local book-keeping: seen sequences for this input. Needed to make sure we only
		// process the first occurence:
		seenSeq := make(map[string]bool)

		// initial data point per input:
		res := make([]DataPoint, 2001)
		res[0] = DataPoint{hash: h, digit: h % 10, diff: 0, diffSeq: []int{}}
		for i := range 2000 {
			h = hash(h)
			dp := DataPoint{
				hash:  h,
				digit: h % 10,
				diff:  h%10 - res[i].digit,
			}
			// create last diff sequence:
			diffs := make([]int, 0)
			if i >= 2 {
				diffs = append(diffs, res[i-2].diff)
				diffs = append(diffs, res[i-1].diff)
				diffs = append(diffs, res[i].diff)
				diffs = append(diffs, dp.diff)
			}
			dp.diffSeq = diffs
			res[i+1] = dp

			// sum up points for each unique sequence:
			k := fmt.Sprintf("%v", dp.diffSeq)
			if len(dp.diffSeq) == 4 {
				// ... but only if it is the first occurence of this sequence in that input:
				if _, ok := seenSeq[k]; !ok {
					seenSeq[k] = true
					sequencePoints[k] += dp.digit
				}
			}
		}
	}
	maxP := 0
	for _, v := range sequencePoints {
		if v > maxP {
			maxP = v
		}
	}
	d.s2 = maxP
}

func (d *Day22) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day22) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func hash(start int) int {
	mod := 16777216 // 2^24
	act := start
	res := ((act * 64) ^ act) % mod
	res = ((res / 32) ^ res) % mod
	res = ((res * 2048) ^ res) % mod

	return res
}
