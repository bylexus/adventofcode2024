package day21

import (
	"fmt"
	"slices"
	"strings"

	"alexi.ch/aoc/2024/lib"
)

const (
	TYPE_KEYPAD = iota
	TYPE_DIRECTIONPAD
)

type Entry struct {
	tile            rune
	coord           lib.Coord
	visited         bool
	distanceToStart int
	keysToStart     []rune
}

func (e *Entry) String() string {
	str := fmt.Sprintf("%c, Keys: ", e.tile)
	for _, k := range e.keysToStart {
		str += fmt.Sprintf("%c", k)
	}
	return str
}

type Keypad map[lib.Coord]*Entry

func createKeypad() Keypad {
	keypad := make(Keypad)
	keypad[lib.NewCoord2D(0, 0)] = &Entry{tile: '7', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(0, 0)}
	keypad[lib.NewCoord2D(1, 0)] = &Entry{tile: '8', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(1, 0)}
	keypad[lib.NewCoord2D(2, 0)] = &Entry{tile: '9', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(2, 0)}
	keypad[lib.NewCoord2D(0, 1)] = &Entry{tile: '4', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(0, 1)}
	keypad[lib.NewCoord2D(1, 1)] = &Entry{tile: '5', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(1, 1)}
	keypad[lib.NewCoord2D(2, 1)] = &Entry{tile: '6', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(2, 1)}
	keypad[lib.NewCoord2D(0, 2)] = &Entry{tile: '1', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(0, 2)}
	keypad[lib.NewCoord2D(1, 2)] = &Entry{tile: '2', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(1, 2)}
	keypad[lib.NewCoord2D(2, 2)] = &Entry{tile: '3', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(2, 2)}
	keypad[lib.NewCoord2D(1, 3)] = &Entry{tile: '0', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(1, 3)}
	keypad[lib.NewCoord2D(2, 3)] = &Entry{tile: 'A', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(2, 3)}
	return keypad
}

func createDirectionalPad() Keypad {
	keypad := make(Keypad)
	keypad[lib.NewCoord2D(1, 0)] = &Entry{tile: '^', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(1, 0)}
	keypad[lib.NewCoord2D(2, 0)] = &Entry{tile: 'A', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(2, 0)}
	keypad[lib.NewCoord2D(0, 1)] = &Entry{tile: '<', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(0, 1)}
	keypad[lib.NewCoord2D(1, 1)] = &Entry{tile: 'v', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(1, 1)}
	keypad[lib.NewCoord2D(2, 1)] = &Entry{tile: '>', distanceToStart: -1, keysToStart: make([]rune, 0), coord: lib.NewCoord2D(2, 1)}
	return keypad
}

type Day21 struct {
	s1      int
	s2      int
	numbers []string
}

func New() Day21 {
	return Day21{s1: 0, s2: 0}
}

func (d *Day21) Title() string {
	return "Day 21 - Keypad Conundrum"
}

func (d *Day21) Setup() {
	var lines = lib.ReadLines("data/21-test-data.txt")
	// var lines = lib.ReadLines("data/21-data.txt")
	for _, line := range lines {
		if len(line) > 0 {
			d.numbers = append(d.numbers, line)
		}
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day21) SolveProblem1() {
	d.s1 = 0
	allNumberKeys := make(map[rune]Keypad)
	masterKeypad := createKeypad()
	for _, pad := range masterKeypad {
		key := pad.tile
		start := pad.coord
		keypad := createKeypad()
		walkKeypad(keypad, start)
		allNumberKeys[key] = keypad
	}

	allDirectionalKeys := make(map[rune]Keypad)
	masterKeypad = createDirectionalPad()
	for _, pad := range masterKeypad {
		key := pad.tile
		start := pad.coord
		keypad := createDirectionalPad()
		walkKeypad(keypad, start)
		allDirectionalKeys[key] = keypad
	}

	for _, number := range d.numbers {
		// number := "029A"
		seq := getSequenceForKeys(allNumberKeys, allDirectionalKeys, number)
		length := len(seq)
		nrs := lib.StrToInt(strings.ReplaceAll(number, "A", ""))
		fmt.Printf("Number: %s, Length: %d, Nrs: %d, total: %d, seq: %s\n", number, length, nrs, nrs*length, string(seq))
		d.s1 += nrs * length
	}
}

func (d *Day21) SolveProblem2() {
	d.s2 = 0
}

func (d *Day21) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day21) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func walkKeypad(pad map[lib.Coord]*Entry, start lib.Coord) {
	pad[start].distanceToStart = 0
	pad[start].visited = false
	unvisited := []*Entry{pad[start]}

	for {
		// find the smallest unvisited node
		act := slices.MinFunc(unvisited, func(a, b *Entry) int {
			if a.distanceToStart == -1 {
				return 1
			}
			if b.distanceToStart == -1 {
				return -1
			}
			return a.distanceToStart - b.distanceToStart
		})
		if act.distanceToStart == -1 {
			break
		}
		act.visited = true

		unvisited = lib.Splice(unvisited, slices.Index(unvisited, act))

		// update neighbour node's distance:
		for _, c := range lib.MOVE_VEC_2D_4DIRS {
			nextDir := vec2Dir(c)
			nextCoord := act.coord.AddXY(c[0], c[1])
			nextNode := pad[nextCoord]
			if nextNode == nil {
				continue
			}
			cost := 1
			// do we need to update the neighbour's number, because it is not yet set or larger than the actual distance?
			if nextNode.distanceToStart < 0 || nextNode.distanceToStart > act.distanceToStart+cost {
				nextNode.distanceToStart = act.distanceToStart + cost
				nextNode.keysToStart = append([]rune{}, act.keysToStart...)
				nextNode.keysToStart = append(nextNode.keysToStart, nextDir)
			}
			// add unvisited nodes to the list
			if !nextNode.visited && !slices.Contains(unvisited, nextNode) {
				unvisited = append(unvisited, nextNode)
			}
		}
		if len(unvisited) == 0 {
			break
		}
	}

	dirs := pad[start].keysToStart
	dirPrefs := []rune{'<', '^', '>', 'v', 'A'}
	slices.SortFunc(dirs, func(a, b rune) int {
		return slices.Index(dirPrefs, a) - slices.Index(dirPrefs, b)
	})
	pad[start].keysToStart = dirs
	// d.printMaze(d.maze)
	// fmt.Printf("end node: %#v\n", d.maze[d.target])
	// d.s1 = d.maze[d.target].distanceToStart
}

func vec2Dir(dir []int) rune {
	if dir[0] == 0 && dir[1] == -1 {
		return '^'
	}
	if dir[0] == 0 && dir[1] == 1 {
		return 'v'
	}
	if dir[0] == -1 && dir[1] == 0 {
		return '<'
	}
	if dir[0] == 1 && dir[1] == 0 {
		return '>'
	}
	panic("Invalid direction")
}

func getKeyEntry(keypad Keypad, key rune) *Entry {
	for _, e := range keypad {
		if e.tile == key {
			return e
		}
	}
	panic("Key not found")
}

func getSequenceForKeys(numKeys map[rune]Keypad, dirKeys map[rune]Keypad, numbers string) []rune {
	// get dirs for 1st robot, operating the real keypad, for the given number sequence:
	seq := make([]rune, 0)
	startKey := 'A'
	for _, r := range numbers {
		keyEntry := getKeyEntry(numKeys[startKey], r)
		dirs := keyEntry.keysToStart
		dirPrefs := []rune{'^', '>', 'v', '<', 'A'}
		slices.SortFunc(dirs, func(a, b rune) int {
			return slices.Index(dirPrefs, a) - slices.Index(dirPrefs, b)
		})
		seq = append(seq, dirs...)
		seq = append(seq, 'A')
		startKey = r
	}
	// fmt.Printf("Sequence for %s: %v\n", numbers, string(seq))

	// get dirs for 2st robot for key sequence:
	seq2 := make([]rune, 0)
	startKey = 'A'
	for _, r := range seq {
		keyEntry := getKeyEntry(dirKeys[startKey], r)
		dirs := keyEntry.keysToStart
		dirPrefs := []rune{'^', '>', '<', 'v', 'A'}
		slices.SortFunc(dirs, func(a, b rune) int {
			return slices.Index(dirPrefs, a) - slices.Index(dirPrefs, b)
		})
		seq2 = append(seq2, dirs...)
		seq2 = append(seq2, 'A')
		startKey = r
	}
	// fmt.Printf("Sequence for %s: %v\n", string(seq), string(seq2))

	// get dirs for myself: I can press the 1st btn directly
	seq3 := make([]rune, 0)
	startKey = 'A'
	for _, r := range seq2 {
		keyEntry := getKeyEntry(dirKeys[startKey], r)
		dirs := keyEntry.keysToStart
		dirPrefs := []rune{'^', '>', '<', 'v', 'A'}
		slices.SortFunc(dirs, func(a, b rune) int {
			return slices.Index(dirPrefs, a) - slices.Index(dirPrefs, b)
		})
		seq3 = append(seq3, dirs...)
		seq3 = append(seq3, 'A')
		startKey = r
	}
	// fmt.Printf("Sequence for %s: %v\n", string(seq2), string(seq3))
	// fmt.Printf("Seq len: %d\n", len(seq3))
	return seq3
}
