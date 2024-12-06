package day06

import (
	"fmt"
	"maps"
	"sync"
	"sync/atomic"

	"alexi.ch/aoc/2024/lib"
)

const (
	RES_OUT_OF_MAP = 1
	RES_LOOP       = 2
)

type Tile struct {
	tile     rune
	dirsSeen []rune
}

type Map struct {
	theMap     map[lib.Coord]Tile
	width      int
	height     int
	guardStart lib.Coord
	guardPos   lib.Coord
	guardDir   rune
}

var dirVec map[rune]lib.Coord = map[rune]lib.Coord{
	'N': lib.NewCoord2D(0, -1),
	'S': lib.NewCoord2D(0, 1),
	'W': lib.NewCoord2D(-1, 0),
	'E': lib.NewCoord2D(1, 0),
}

func (m Map) String() string {
	str := ""
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if x == m.guardPos.X && y == m.guardPos.Y {
				str += fmt.Sprintf("%c", m.guardDir)
			} else {
				str += fmt.Sprintf("%c", m.theMap[lib.NewCoord2D(x, y)].tile)
			}
		}
		str += "\n"
	}
	str += "\n"
	return str
}

type Day06 struct {
	initialMap Map
	guardMap   Map
	s1         int
	s2         int
}

func New() Day06 {
	return Day06{s1: 0, s2: 0, guardMap: Map{theMap: make(map[lib.Coord]Tile)}}
}

func (d *Day06) Title() string {
	return "Day 06 - Guard Gallivant"
}

func (d *Day06) Setup() {
	// var lines = lib.ReadLines("data/06-test-data.txt")
	var lines = lib.ReadLines("data/06-data.txt")
	d.guardMap.height = len(lines)
	d.guardMap.width = len(lines[0])
	for y, line := range lines {
		for x, r := range line {
			c := lib.NewCoord2D(x, y)
			tile := Tile{tile: r}
			if r == '.' {
				continue
			}
			if r == '^' {
				d.guardMap.guardStart = c
				d.guardMap.guardPos = c
				d.guardMap.guardDir = 'N'
				tile.tile = 'X'
				tile.dirsSeen = []rune{'N'}
			}
			d.guardMap.theMap[c] = tile
		}
	}
	d.initialMap = d.guardMap
	d.initialMap.theMap = maps.Clone(d.guardMap.theMap)
}

func (d *Day06) SolveProblem1() {
	d.s1 = 0
	for d.walkInMap(&d.guardMap) == 0 {
		//
	}
	for _, r := range d.guardMap.theMap {
		if r.tile == 'X' {
			d.s1++
		}
	}
	// fmt.Printf("%v\n", d.guardMap)
}

func (d *Day06) SolveProblem2() {
	d.s2 = 0

	// try all possible obstacle positions in parallel, using
	// goroutines
	wg := sync.WaitGroup{}
	counter := atomic.Int64{}

	// realisation: we only have to place obstacles in the path
	// that was used in the 1st part: he reaches only the
	// tiles that he visited in the 1st part, so no need to
	// place obstacles in the rest of the map
	for obstacleCoord := range d.guardMap.theMap {
		wg.Add(1)
		go func() {
			defer wg.Done()
			theMap := d.initialMap
			theMap.theMap = maps.Clone(d.initialMap.theMap)
			if theMap.guardStart == obstacleCoord {
				return
			}
			if theMap.theMap[obstacleCoord].tile == '#' {
				return
			}
			theMap.theMap[obstacleCoord] = Tile{tile: '#'}
			res := 0

			for {
				res = d.walkInMap(&theMap)
				if res != 0 {
					break
				}
			}
			if res == RES_LOOP {
				counter.Add(1)
			}
		}()

	}

	wg.Wait()
	d.s2 = int(counter.Load())
	// fmt.Printf("%v\n", d.guardMap)

}

func (d *Day06) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day06) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day06) walkInMap(guardMap *Map) int {
	actTile := guardMap.theMap[guardMap.guardPos]
	actDir := guardMap.guardDir
	newPos := guardMap.guardPos.Add(dirVec[actDir])

	// out of the map? stop:
	if newPos.X < 0 || newPos.Y < 0 || newPos.X >= guardMap.width || newPos.Y >= guardMap.height {
		return RES_OUT_OF_MAP
	}

	newPosTile := guardMap.theMap[newPos]
	if newPosTile.tile == '#' {
		guardMap.guardDir = d.turnRightDir(guardMap.guardDir)
		actTile.dirsSeen = append(actTile.dirsSeen, guardMap.guardDir)
		guardMap.theMap[guardMap.guardPos] = actTile
	} else {
		// loop detection
		if lib.Contains(newPosTile.dirsSeen, actDir) {
			return RES_LOOP
		}

		newPosTile.tile = 'X'
		newPosTile.dirsSeen = append(newPosTile.dirsSeen, guardMap.guardDir)
		guardMap.theMap[newPos] = newPosTile
		guardMap.guardPos = newPos

	}
	return 0
}

func (d *Day06) turnRightDir(dir rune) rune {
	switch dir {
	case 'N':
		return 'E'
	case 'E':
		return 'S'
	case 'S':
		return 'W'
	case 'W':
		return 'N'
	default:
		panic("Unknown direction")
	}
}
