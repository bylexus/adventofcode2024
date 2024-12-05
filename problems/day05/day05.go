package day05

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"alexi.ch/aoc/2024/lib"
)

type Rules map[int][]int

type Day05 struct {
	followers  Rules
	preceeders Rules
	updates    [][]int
	errorRules [][]int
	s1         int
	s2         int
}

func New() Day05 {
	return Day05{s1: 0, s2: 0, followers: make(Rules), preceeders: make(Rules)}
}

func (d *Day05) Title() string {
	return "Day 05 - Print Queue"
}

func (d *Day05) Setup() {
	// var lines = lib.ReadLines("data/05-test-data.txt")
	var lines = lib.ReadLines("data/05-data.txt")
	rulesMatcher := regexp.MustCompile(`(\d+)\|(\d+)`)
	updateMatcher := regexp.MustCompile(`,`)
	for _, line := range lines {
		rules := rulesMatcher.FindStringSubmatch(line)
		if len(rules) == 3 {
			first := lib.StrToInt(rules[1])
			second := lib.StrToInt(rules[2])

			firstEntry, ok := d.followers[first]
			if !ok {
				firstEntry = make([]int, 0)
			}
			firstEntry = append(firstEntry, second)
			d.followers[first] = firstEntry

			secondEntry, ok := d.preceeders[second]
			if !ok {
				secondEntry = make([]int, 0)
			}
			secondEntry = append(secondEntry, first)
			d.preceeders[second] = secondEntry
		}

		// part 2 of input: updates
		if updateMatcher.MatchString(line) {
			updateStrings := strings.Split(line, ",")
			updates := make([]int, 0)
			for _, update := range updateStrings {
				updates = append(updates, lib.StrToInt(update))
			}
			d.updates = append(d.updates, updates)
		}
	}
	// fmt.Printf("%#v\n", d.preceeders)
	// fmt.Printf("%#v\n", d.followers)
	// fmt.Printf("%#v\n", d.updates)
}

func (d *Day05) SolveProblem1() {
	d.s1 = 0
	for _, update := range d.updates {
		correct := true
		for i, nr := range update {
			preceedingElements := update[:i]
			if !d.elementsPreceeding(preceedingElements, nr) {
				correct = false
				break
			}
			followingElements := update[i+1:]
			if !d.elementsFollowing(followingElements, nr) {
				correct = false
				break
			}
		}
		if correct {
			// fmt.Printf("Update correct: %#v\n", update)
			d.s1 += update[len(update)/2]
		} else {
			d.errorRules = append(d.errorRules, update)
			// fmt.Printf("Update INcorrect: %#v\n", update)
		}
	}
}

func (d *Day05) SolveProblem2() {
	// seems to be a (bubble) sort problem: just sort them, by using a element compare function:
	// el1 < el2 if el1 is not contained in el2's follower's list
	// The algorithm below is an optimized bubble sort.
	d.s2 = 0
	for _, update := range d.errorRules {
		// fmt.Printf("wrong update: %#v\n", update)

		// The Go! Way: Use pre-defined sort function with element compare function
		slices.SortFunc(update, func(a, b int) int {
			if d.aBeforeB(a, b) {
				return -1
			}
			return 1
		})

		// Manual solution: Bubble Sort!
		// bubble sort the list by the following element compare function:
		// the follower (el[j+1]) must not have the preceeder (el[j]) in its follower's list:

		// for i := 0; i < len(update); i++ {
		// 	swapped := false
		// 	for j := 0; j < len(update)-i-1; j++ {
		// 		preceeder := update[j]
		// 		follower := update[j+1]
		// 		if !d.aBeforeB(preceeder, follower) {
		// 			// swap elements: j+1 comes before j
		// 			update[j], update[j+1] = update[j+1], update[j]
		// 			swapped = true
		// 		}
		// 	}
		// 	if !swapped {
		// 		break
		// 	}
		// }
		// fmt.Printf("corrected update: %#v\n", update)
		d.s2 += update[len(update)/2]
	}
}

func (d *Day05) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day05) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

// Preceeding elements all must come before the tested element el.
// We do this by making sure that the preceeding elements are NOT present
// in the follower's list of el.
func (d *Day05) elementsPreceeding(preceedingElements []int, el int) bool {
	for _, preceedingElement := range preceedingElements {
		if !d.aBeforeB(preceedingElement, el) {
			return false
		}
	}
	return true
}

// Following elements all must come after the tested element el.
// We do this by making sure that the following elements are NOT present
// in the preceeder's list of el.
func (d *Day05) elementsFollowing(followingElements []int, el int) bool {
	for _, followingElement := range followingElements {
		if !d.aBeforeB(el, followingElement) {
			return false
		}
	}
	return true
}

// Compares 2 docs and returns true if doc1 comes before doc2.
func (d *Day05) aBeforeB(doc1, doc2 int) bool {
	// doc1 must not be part of doc2's followers list:
	return !lib.Contains(d.followers[doc2], doc1)
}
