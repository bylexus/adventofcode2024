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
}

type Warehouse struct {
	houseMap [][]*MapEntry
	robotPos lib.Coord
}

func (w Warehouse) Width() int {
	return len(w.houseMap[0])
}

func (w Warehouse) Height() int {
	return len(w.houseMap)
}

type Day15 struct {
	s1         int
	s2         int
	warehouse  Warehouse
	warehouse2 Warehouse
	movements  []rune
}

func New() Day15 {
	return Day15{
		s1:         0,
		s2:         0,
		warehouse:  Warehouse{},
		warehouse2: Warehouse{},
		movements:  make([]rune, 0),
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

	// parse house map
	// create 2 warehouses: the standard one for part 1,
	// and the double-width for part 2
	for l := strings.TrimSpace(lines[line]); l != ""; l = lines[line] {

		we := make([]*MapEntry, 0)
		we2 := make([]*MapEntry, 0)
		x2 := 0
		for x, r := range l {
			if r == '@' {
				d.warehouse.robotPos = lib.NewCoord2D(x, line)
				d.warehouse2.robotPos = lib.NewCoord2D(x2, line)
				we = append(we, &MapEntry{kind: '.'})
				// double width for 2nd warehouse
				we2 = append(we2, &MapEntry{kind: '.'})
				we2 = append(we2, &MapEntry{kind: '.'})
			} else if r == '.' || r == '#' {
				we = append(we, &MapEntry{kind: r})
				we2 = append(we2, &MapEntry{kind: r})
				we2 = append(we2, &MapEntry{kind: r})
			} else if r == 'O' {
				we = append(we, &MapEntry{kind: r})
				we2 = append(we2, &MapEntry{kind: '['})
				we2 = append(we2, &MapEntry{kind: ']'})
			}
			x2 += 2
		}

		d.warehouse.houseMap = append(d.warehouse.houseMap, we)
		d.warehouse2.houseMap = append(d.warehouse2.houseMap, we2)
		line++
	}

	for ; line < len(lines); line++ {
		for _, r := range lines[line] {
			if r == '^' || r == '>' || r == 'v' || r == '<' {
				d.movements = append(d.movements, r)
			}
		}
	}

	// d.PrintWarehouse(d.warehouse)
	// d.PrintWarehouse(d.warehouse2)
}

func (d *Day15) SolveProblem1() {
	d.s1 = 0
	// d.PrintWarehouse(d.warehouse)
	for _, movement := range d.movements {
		d.warehouse.moveRobot(movement)
	}
	// d.PrintWarehouse(d.warehouse)

	for y, line := range d.warehouse.houseMap {
		for x, entry := range line {
			if entry.kind == 'O' {
				d.s1 += x + y*100
			}
		}
	}
}

func (d *Day15) SolveProblem2() {
	d.s2 = 0
	// d.PrintWarehouse(d.warehouse2)
	for _, movement := range d.movements {
		d.warehouse2.moveRobotWideWarehouse(movement)
		// d.PrintWarehouse(d.warehouse2)
	}
	for y, line := range d.warehouse2.houseMap {
		for x, entry := range line {
			if entry.kind == '[' {
				d.s2 += x + y*100
			}
		}
	}
}

func (d *Day15) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day15) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day15) PrintWarehouse(w Warehouse) {
	for y, line := range w.houseMap {
		for x, e := range line {
			if w.robotPos == lib.NewCoord2D(x, y) {
				fmt.Printf("@")
			} else {
				fmt.Printf("%c", e.kind)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (w *Warehouse) moveRobot(m rune) {
	// first, check for a free space ahead of the robot, until we hit a wall
	// If we find one, we do:
	// - move all the boxes in front of the robot, starting with the farthest away
	// - move the robot
	// if we do NOT find a free space, the movement is ignored.

	// search free space
	freePos, ok := w.searchFreeSpace(w.robotPos, m)
	if ok {
		// move all boxes
		w.moveBoxes(w.robotPos, freePos, m)

		// move robot
		moveVec := MOVE_MAP[m]
		w.robotPos = w.robotPos.AddXY(moveVec[0], moveVec[1])
	}
}

// This is the movement method of the wide warehouse
func (w *Warehouse) moveRobotWideWarehouse(m rune) {
	// first, check for a free space ahead of the robot, until we hit a wall
	// If we find one, we do:
	// - move all the boxes in front of the robot, starting with the farthest away
	// - move the robot
	// if we do NOT find a free space, the movement is ignored.

	// search free space
	moveVec := MOVE_MAP[m]
	nextPos := w.robotPos.AddXY(moveVec[0], moveVec[1])
	ok := w.movingPossibleWideWarehouse(nextPos, m, true)
	if ok {
		// move all boxes in front:
		w.moveBoxesWideWarehouse(nextPos, m, true)

		// move robot
		w.robotPos = nextPos
	}
}

func (w *Warehouse) searchFreeSpace(startPos lib.Coord, dir rune) (lib.Coord, bool) {
	dirVec := MOVE_MAP[dir]
	actPos := startPos
	for {
		actPos = actPos.AddXY(dirVec[0], dirVec[1])
		// out of bounds check:
		if actPos.X < 0 || actPos.Y < 0 || actPos.X >= w.Width() || actPos.Y >= w.Height() {
			return lib.NewCoord0(), false
		}
		if w.houseMap[actPos.Y][actPos.X].kind == '#' {
			return lib.NewCoord0(), false
		}
		if w.houseMap[actPos.Y][actPos.X].kind == '.' {
			return actPos, true
		}
	}
}

func (w *Warehouse) movingPossibleWideWarehouse(startPos lib.Coord, dir rune, checkNeighbourCell bool) bool {
	dirVec := MOVE_MAP[dir]
	actPos := startPos

	// left/right is a bit different than up/down:
	// in any case, we just look at the next boxes, and ask them recursively if can
	// be moved.
	// - in case of left/right: just ask the box part left/right of the act pos if  movement is possible
	// - in case of up/down: we need to ask both box parts in front of the act pos if movement is possible.
	//   Both parts must be movable, no partial movement possible.

	// in any case, if we actually are on an empty spot or a wall, return true/false (recursive 1-option):
	// out of bounds:
	if startPos.X < 0 || startPos.Y < 0 || startPos.X >= w.Width() || startPos.Y >= w.Height() {
		return false
	}
	// empty space:
	if w.houseMap[startPos.Y][startPos.X].kind == '.' {
		return true
	}
	// wall:
	if w.houseMap[startPos.Y][startPos.X].kind == '#' {
		return false
	}

	if dir == '<' || dir == '>' {
		nextPos := actPos.AddXY(dirVec[0], dirVec[1])
		return w.movingPossibleWideWarehouse(nextPos, dir, false)
	} else {
		nextPos := actPos.AddXY(dirVec[0], dirVec[1])
		ok := w.movingPossibleWideWarehouse(nextPos, dir, true)
		if checkNeighbourCell {
			if w.houseMap[actPos.Y][actPos.X].kind == '[' {
				nextPos = actPos.AddXY(1, 0)
			} else if w.houseMap[actPos.Y][actPos.X].kind == ']' {
				nextPos = actPos.AddXY(-1, 0)
			} else {
				panic("This should really not happen....")
			}
			ok2 := w.movingPossibleWideWarehouse(nextPos, dir, false)
			return ok && ok2
		}
		return ok
	}
}

// this function recursively move boxes in the wide warehouse.
// Please be advised that this function do NOT check if movement is possible!
// This has to be done first using movingPossibleWideWarehause function!
func (w *Warehouse) moveBoxesWideWarehouse(startPos lib.Coord, dir rune, checkNeighbourCell bool) {
	dirVec := MOVE_MAP[dir]
	actPos := startPos

	// left/right is a bit different than up/down:
	// in any case, we first move the next boxes recursively, before moving myself.
	// - in case of left/right: just ask the box part left/right of the act pos to move, then move
	// - in case of up/down: we need to ask both box parts in front of the act pos to move.

	// in any case, if we actually are on an empty spot or a wall, return true/false (recursive 1-option):
	// out of bounds:
	if startPos.X < 0 || startPos.Y < 0 || startPos.X >= w.Width() || startPos.Y >= w.Height() {
		panic("Oops! I should not be here! I am off the warehouse!")
	}
	// wall:
	if w.houseMap[startPos.Y][startPos.X].kind == '#' {
		panic("Oops! I should not be here: I hit a wall!")
	}
	// empty space:
	if w.houseMap[startPos.Y][startPos.X].kind == '.' {
		return
	}

	if dir == '<' || dir == '>' {
		nextPos := actPos.AddXY(dirVec[0], dirVec[1])
		// move boxes before me, recursively:
		w.moveBoxesWideWarehouse(nextPos, dir, false)
		// now, move myself:
		tmp := w.houseMap[nextPos.Y][nextPos.X]
		w.houseMap[nextPos.Y][nextPos.X] = w.houseMap[startPos.Y][startPos.X]
		w.houseMap[startPos.Y][startPos.X] = tmp

	} else {
		nextPos := actPos.AddXY(dirVec[0], dirVec[1])
		// move boxes before me, recursively:
		w.moveBoxesWideWarehouse(nextPos, dir, true)
		// now, move myself:
		actPart := w.houseMap[startPos.Y][startPos.X].kind
		tmp := w.houseMap[nextPos.Y][nextPos.X]
		w.houseMap[nextPos.Y][nextPos.X] = w.houseMap[startPos.Y][startPos.X]
		w.houseMap[startPos.Y][startPos.X] = tmp

		if checkNeighbourCell {
			if actPart == '[' {
				nextPos = actPos.AddXY(1, 0)
			} else if actPart == ']' {
				nextPos = actPos.AddXY(-1, 0)
			} else {
				panic("This should really not happen....")
			}
			// now, move the other part:
			w.moveBoxesWideWarehouse(nextPos, dir, false)
		}
	}
}

// move boxes one place further in a straight line.
// we assume that startPos is before endPos, in the opposite direction of dir.
func (w *Warehouse) moveBoxes(startPos, endPos lib.Coord, dir rune) {
	forward := MOVE_MAP[dir]
	backwards := []int{forward[0] * -1, forward[1] * -1}

	for {
		targetPos := endPos
		endPos = endPos.AddXY(backwards[0], backwards[1])
		// swap box with empty space:
		if w.houseMap[targetPos.Y][targetPos.X].kind == '.' && w.houseMap[endPos.Y][endPos.X].kind == 'O' {
			tmp := w.houseMap[targetPos.Y][targetPos.X]
			w.houseMap[targetPos.Y][targetPos.X] = w.houseMap[endPos.Y][endPos.X]
			w.houseMap[endPos.Y][endPos.X] = tmp
		}

		// if we reached the start, we're done:
		if endPos == startPos {
			return
		}
	}
}
