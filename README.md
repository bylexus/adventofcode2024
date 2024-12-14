# Advent of Code 2024

Welcome to AoC 2024! Another year, another try! but remember, Program Or Die! Muahahahahaaa!

Eagerly, we're all awaiting [Advent of Code, Edition 2024!](https://adventofcode.com/2024/). Finally, it's here.
What a magical time of the year!

My AoC 2024 solutions in GO, again.

## Run problems

All problems can be run by its day index, e.g:

```
$ go run 01
```

or all together:

```
$ go run
```

## Define Problems

1) Create a struct in the `problems` package that implements the `Problem` interface, e.g.:

```go
package problems

import (
	"fmt"
	"alexi.ch/aoc/2024/lib"
)

type DayXX struct {
	s1 uint64
	s2 uint64
}

func NewDayXX() DayXX {
	return DayXX{s1: 0, s2: 0}
}

func (d *DayXX) Title() string {
	return "Day XX - Title comes here"
}

func (d *DayXX) Setup() {
	// var lines = lib.ReadLines("data/01-test.txt")
	var lines = lib.ReadLines("data/01-data.txt")
	for _, line := range lines {
		line = line
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *DayXX) SolveProblem1() {
	d.s1 = 0
}

func (d *DayXX) SolveProblem2() {
	d.s2 = 0
}

func (d *DayXX) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *DayXX) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
```

2) import and instantiate the struct in the main program `aoc.go`:

```go
// aoc.go
var dayXX = problems.NewDayXX()
problem_map["XX"] = &dayXX
```


## Some Notes to certain days

### Day 12: Flood-Fill vizualization

Day 12 was a flood-fill problem. I visualized the areas to fill:

![day12](./day12-flood-the-garden.png)

### Day 14: Find the tree!

The idea was to find a specific arrangement of pixels after an unknown round of
modifications, here, we were looking for a Christmas Tree!

My first idea was to just output the image after each round, and check them visually, maybe there will be a
pattern in the pixel movement?

After viewing some hundred ascii art images manually, I gave up: The pixels did just jump around randomly, and
did not form anything useful.

My second idea was that the tree must be in the center: The first problem indicated that the 
center lines were important - so I looked for a long center arrangement of the pixels - and could
not find any image whithin millions of rounds. So that was not it.

So I did it the hard way: I generated real images (pngs) for the first 10'000 rounds, and viewed them
im the MacOS finder in a grid - and voilà, there it was! deeply buried in some 1000's of images....
And of course it was NOT a big tree in the center....

So after I know how the tree must look, I could find it programmatically.

Here it is:

![Christmas Tree](./day14_tree_round_6668.png)

That was a fun one!
