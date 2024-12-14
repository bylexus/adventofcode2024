package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"alexi.ch/aoc/2024/problems"
	"alexi.ch/aoc/2024/problems/day01"
	"alexi.ch/aoc/2024/problems/day02"
	"alexi.ch/aoc/2024/problems/day03"
	"alexi.ch/aoc/2024/problems/day04"
	"alexi.ch/aoc/2024/problems/day05"
	"alexi.ch/aoc/2024/problems/day06"
	"alexi.ch/aoc/2024/problems/day07"
	"alexi.ch/aoc/2024/problems/day08"
	"alexi.ch/aoc/2024/problems/day09"
	"alexi.ch/aoc/2024/problems/day10"
	"alexi.ch/aoc/2024/problems/day11"
	"alexi.ch/aoc/2024/problems/day12"
	"alexi.ch/aoc/2024/problems/day13"
	"alexi.ch/aoc/2024/problems/day14"
)

func main() {
	tannenbaum()
	var problem_map = map[string](func() problems.Problem){
		"01":         func() problems.Problem { p := day01.NewDay01(); return &p },
		"02":         func() problems.Problem { p := day02.NewDay02(); return &p },
		"03":         func() problems.Problem { p := day03.NewDay03(); return &p },
		"04":         func() problems.Problem { p := day04.NewDay04(); return &p },
		"05":         func() problems.Problem { p := day05.New(); return &p },
		"06":         func() problems.Problem { p := day06.New(); return &p },
		"07":         func() problems.Problem { p := day07.New(); return &p },
		"08":         func() problems.Problem { p := day08.New(); return &p },
		"09":         func() problems.Problem { p := day09.New(); return &p },
		"10":         func() problems.Problem { p := day10.New(); return &p },
		"11":         func() problems.Problem { p := day11.New(); return &p },
		"12":         func() problems.Problem { p := day12.New(); return &p },
		"13":         func() problems.Problem { p := day13.New(); return &p },
		"14":         func() problems.Problem { p := day14.New(); return &p },
		"playground": func() problems.Problem { p := problems.NewPlayground(); return &p },
	}

	var to_solve = os.Args[1:]

	if len(to_solve) == 0 {
		var keys = make([]string, 0)
		for key := range problem_map {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		to_solve = keys
	}

	// Run solving all problems  in parallel:
	var start = time.Now()
	var wg sync.WaitGroup
	wg.Add(len(to_solve))
	for _, p := range to_solve {
		go func(probKey string) {
			defer wg.Done()
			var prob = problem_map[probKey]
			if prob != nil {
				problems.Solve(prob())
			} else {
				panic("Problem not found")
			}
		}(p)
	}
	wg.Wait()
	var duration = time.Since(start)
	fmt.Printf("\n\nFull runtime: %s\n\n", duration)
}

func tannenbaum() {
	var t = strings.Join([]string{
		"\x1B[1;97m",
		"Advent of Code 2024",
		"--------------------",
		"",
		"        \x1B[1;93m*   *",
		"         \\ /",
		"         AoC",
		"         -\x1B[1;91m*\x1B[1;93m-",
		"          \x1B[1;37m|\x1B[0;32m",
		"          *",
		"         /*\\",
		"        /\x1B[1;94m*\x1B[0;32m*\x1B[1;93m*\x1B[0;32m\\",
		"       /\x1B[1;91m*\x1B[0;32m***\x1B[1;94m*\x1B[0;32m\\",
		"      /**\x1B[1;93m*\x1B[0;32m****\\",
		"     /**\x1B[1;94m*\x1B[0;32m***\x1B[1;91m*\x1B[0;32m**\\",
		"    /********\x1B[1;93m*\x1B[0;32m**\\",
		"   /**\x1B[1;91m*\x1B[0;32m*****\x1B[1;94m*\x1B[0;32m****\\",
		"  /**\x1B[1;94m*\x1B[0;32m*\x1B[1;93m*\x1B[0;32m**********\\",
		" /**\x1B[1;94m*\x1B[0;32m*****\x1B[1;93m*\x1B[0;32m**\x1B[1;91m*\x1B[0;32m****\x1B[1;93m*\x1B[0;32m\\",
		"          #",
		"          #",
		"       \x1B[1;97m2-0-2-4",
		"       #######",
		"\x1B[0m",
	}, "\n")
	fmt.Print(t)
}
