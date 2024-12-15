package day15

import (
	"fmt"
	"strings"

	"alexi.ch/aoc/2024/lib"
)

var MOVE_MAP = map[rune][]int{
	'^': {0, -1},
	'>': {1, 0},
	'v': {0, 1},
	'<': {-1, 0},
}

type MapEntry struct {
	kind rune
	// coord lib.Coord
}

// type Warehouse struct {
// 	houseMap      [][]*MapEntry
// 	width, height int
// 	robotPos      lib.Coord
// }

type Day15 struct {
	s1        int
	s2        int
	warehouse [][]*MapEntry
	// warehouse2      [][]*MapEntry
	robotPos      lib.Coord
	width, height int
	// width2, height2 int
	movements []rune
}

func New() Day15 {
	return Day15{
		s1:        0,
		s2:        0,
		warehouse: make([][]*MapEntry, 0),
		// warehouse2: make([][]*MapEntry, 0),
		movements: make([]rune, 0),
	}
}

func (d *Day15) Title() string {
	return "Day 15 - Warehouse Woes"
}

func (d *Day15) Setup() {
	// var lines = lib.ReadLines("data/15-test-simple-data.txt")
	// var lines = lib.ReadLines("data/15-2-test-simple-data.txt")
	// var lines = lib.ReadLines("data/15-test-data.txt")
	var lines = lib.ReadLines("data/15-data.txt")
	var line = 0

	for l := strings.TrimSpace(lines[line]); l != ""; l = lines[line] {

		we := make([]*MapEntry, 0)
		// x2 := 0
		for x, r := range l {
			if r == '@' {
				d.robotPos = lib.NewCoord2D(x, line)
				we = append(we, &MapEntry{kind: '.'})
			} else {
				we = append(we, &MapEntry{kind: r})
			}
			// x2++
		}

		d.warehouse = append(d.warehouse, we)
		line++
	}
	d.width = len(d.warehouse[0])
	d.height = len(d.warehouse)
	for ; line < len(lines); line++ {
		for _, r := range lines[line] {
			if r == '^' || r == '>' || r == 'v' || r == '<' {
				d.movements = append(d.movements, r)
			}
		}
	}

	// d.PrintWarehouse()
}

func (d *Day15) SolveProblem1() {
	d.s1 = 0
	d.PrintWarehouse()
	for _, movement := range d.movements {
		d.moveRobot(movement)
	}
	d.PrintWarehouse()

	for y, line := range d.warehouse {
		for x, entry := range line {
			if entry.kind == 'O' {
				d.s1 += x + y*100
			}
		}
	}
}

func (d *Day15) SolveProblem2() {
	d.s2 = 0
}

func (d *Day15) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day15) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day15) PrintWarehouse() {
	for y, line := range d.warehouse {
		for x, e := range line {
			if d.robotPos == lib.NewCoord2D(x, y) {
				fmt.Printf("@")
			} else {
				fmt.Printf("%c", e.kind)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Day15) moveRobot(m rune) {
	// first, check for a free space ahead of the robot, until we hit a wall
	// If we find one, we do:
	// - move all the boxes in front of the robot, starting with the farthest away
	// - move the robot
	// if we do NOT find a free space, the movement is ignored.

	// search free space
	freePos, ok := d.searchFreeSpace(d.robotPos, m)
	if ok {
		// move all boxes
		d.moveBoxes(d.robotPos, freePos, m)

		// move robot
		moveVec := MOVE_MAP[m]
		d.robotPos = d.robotPos.AddXY(moveVec[0], moveVec[1])
	}
}

func (d *Day15) searchFreeSpace(startPos lib.Coord, dir rune) (lib.Coord, bool) {
	dirVec := MOVE_MAP[dir]
	actPos := startPos
	for {
		actPos = actPos.AddXY(dirVec[0], dirVec[1])
		// out of bounds check:
		if actPos.X < 0 || actPos.Y < 0 || actPos.X >= d.width || actPos.Y >= d.height {
			return lib.NewCoord0(), false
		}
		if d.warehouse[actPos.Y][actPos.X].kind == '#' {
			return lib.NewCoord0(), false
		}
		if d.warehouse[actPos.Y][actPos.X].kind == '.' {
			return actPos, true
		}
	}
}

// move boxes one place further in a straight line.
// we assume that startPos is before endPos, in the opposite direction of dir.
func (d *Day15) moveBoxes(startPos, endPos lib.Coord, dir rune) {
	forward := MOVE_MAP[dir]
	backwards := []int{forward[0] * -1, forward[1] * -1}

	for {
		targetPos := endPos
		endPos = endPos.AddXY(backwards[0], backwards[1])
		// swap box with empty space:
		if d.warehouse[targetPos.Y][targetPos.X].kind == '.' && d.warehouse[endPos.Y][endPos.X].kind == 'O' {
			tmp := d.warehouse[targetPos.Y][targetPos.X]
			d.warehouse[targetPos.Y][targetPos.X] = d.warehouse[endPos.Y][endPos.X]
			d.warehouse[endPos.Y][endPos.X] = tmp
		}

		// if we reached the start, we're done:
		if endPos == startPos {
			return
		}
	}
}
