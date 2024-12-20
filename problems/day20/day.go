package day20

import (
	"fmt"
	"slices"
	"sync"
	"sync/atomic"

	"alexi.ch/aoc/2024/lib"
)

type Entry struct {
	tile            rune
	coord           lib.Coord
	visited         bool
	distanceToStart int
}

type Day20 struct {
	s1            int
	s2            string
	origMaze      map[lib.Coord]rune
	width, height int
}

func New() Day20 {
	return Day20{s1: 0, s2: "", origMaze: make(map[lib.Coord]rune)}
}

func (d *Day20) Title() string {
	return "Day 20 - Race Condition"
}

func (d *Day20) Setup() {
	// var lines = lib.ReadLines("data/20-test-data.txt")
	var lines = lib.ReadLines("data/20-data.txt")
	d.height = len(lines)
	d.width = len(lines[0])
	for y, line := range lines {
		for x, r := range line {
			d.origMaze[lib.NewCoord2D(x, y)] = r
		}
	}

	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day20) SolveProblem1() {
	d.s1 = 0

	maze, start, target := d.buildMaze()
	// d.printMaze(maze, d.width, d.height)
	d.walkMaze(maze, start, target)

	noShortcutSolution := maze[target].distanceToStart
	// d.s1 = noShortcutSolution
	// return

	// find all stones that can be removed - walls that are only 1 stone thick
	thinWalls := make([]lib.Coord, 0)
	for y := 1; y < d.height-1; y++ {
		for x := 1; x < d.width-1; x++ {
			c := lib.NewCoord2D(x, y)
			if maze[c].tile != '#' {
				continue
			}
			left := maze[lib.NewCoord2D(x-1, y)]
			right := maze[lib.NewCoord2D(x+1, y)]
			up := maze[lib.NewCoord2D(x, y-1)]
			down := maze[lib.NewCoord2D(x, y+1)]
			if left.tile == '.' && right.tile == '.' {
				thinWalls = append(thinWalls, c)
			} else if up.tile == '.' && down.tile == '.' {
				thinWalls = append(thinWalls, c)
			}
		}
	}
	// fmt.Printf("Thin walls: %#d\n", len(thinWalls))

	wg := sync.WaitGroup{}
	wg.Add(len(thinWalls))
	atomicCounter := atomic.Int64{}
	for i, thinWall := range thinWalls {
		go func(wallNr int, wallCoord lib.Coord) {
			maze, start, target := d.buildMaze()
			maze[thinWall].tile = '.'
			// d.printMaze(maze, d.width, d.height)
			d.walkMaze(maze, start, target)

			shortcutSolution := maze[target].distanceToStart
			if noShortcutSolution-shortcutSolution >= 100 {
				atomicCounter.Add(1)
			}
			// fmt.Printf("#%d: No shortcut: %d, shortcut: %d, diff: %d, s1: %d\n", wallNr, noShortcutSolution, shortcutSolution, noShortcutSolution-shortcutSolution, d.s1)
			wg.Done()

		}(i+1, thinWall)
	}
	wg.Wait()
	d.s1 = int(atomicCounter.Load())
}

func (d *Day20) SolveProblem2() {
	// // width, height := 7, 7
	// width, height := 71, 71
	// coords := d.coords
	// // coords := d.coords[:12]
	// // coords := d.coords[:1024]
	// maze := make(map[lib.Coord]*Entry)

	// for y := range height {
	// 	for x := range width {
	// 		maze[lib.NewCoord2D(x, y)] = &Entry{
	// 			tile:            ' ',
	// 			coord:           lib.NewCoord2D(x, y),
	// 			distanceToStart: -1,
	// 			visited:         false,
	// 		}
	// 	}
	// }

	// for i := 1; i < len(coords)+1; i++ {
	// 	coords := d.coords[:i]

	// 	// reset maze
	// 	for y := range height {
	// 		for x := range width {
	// 			maze[lib.NewCoord2D(x, y)].distanceToStart = -1
	// 			maze[lib.NewCoord2D(x, y)].visited = false
	// 			maze[lib.NewCoord2D(x, y)].tile = ' '
	// 		}
	// 	}

	// 	// fill in new coords
	// 	for _, coord := range coords {
	// 		maze[coord].tile = '#'
	// 	}
	// 	start := lib.NewCoord0()
	// 	end := lib.NewCoord2D(width-1, height-1)

	// 	d.walkMaze(maze, start, end)

	// 	if maze[end].distanceToStart == -1 {
	// 		// fmt.Printf("Corrupting entry: %#v\n", coords[len(coords)-1])
	// 		d.s2 = fmt.Sprintf("%d,%d", coords[len(coords)-1].X, coords[len(coords)-1].Y)
	// 		break
	// 	}
	// }

	// fmt.Printf("End node: %#v\n", *(maze[end]))
	// d.s1 = maze[end].distanceToStart
}

func (d *Day20) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day20) Solution2() string {
	return d.s2
}

func (d *Day20) walkMaze(maze map[lib.Coord]*Entry, start lib.Coord, target lib.Coord) {
	maze[start].distanceToStart = 0
	maze[start].visited = false
	unvisited := []*Entry{maze[start]}

	// // all empty tiles are unvisited at the beginning
	// for _, entry := range maze {
	// 	if entry.tile != '#' {
	// 		unvisited = append(unvisited, entry)
	// 	}
	// }
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
			nextCoord := act.coord.AddXY(c[0], c[1])
			nextNode := maze[nextCoord]
			if nextNode == nil {
				continue
			}
			if nextNode.tile == '#' {
				continue
			}
			cost := 1
			// do we need to update the neighbour's number, because it is not yet set or larger than the actual distance?
			if nextNode.distanceToStart < 0 || nextNode.distanceToStart > act.distanceToStart+cost {
				nextNode.distanceToStart = act.distanceToStart + cost
			}
			// add unvisited nodes to the list
			if !nextNode.visited && !slices.Contains(unvisited, nextNode) {
				unvisited = append(unvisited, nextNode)
			}
		}
		if act.coord == target {
			break
		}
		if len(unvisited) == 0 {
			break
		}
	}
	// d.printMaze(d.maze)
	// fmt.Printf("end node: %#v\n", d.maze[d.target])
	// d.s1 = d.maze[d.target].distanceToStart
}

func (d *Day20) printMaze(m map[lib.Coord]*Entry, width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			e := m[lib.NewCoord2D(x, y)]
			if e.tile == '#' {
				fmt.Printf("#")
			} else if e.visited {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Day20) buildMaze() (maze map[lib.Coord]*Entry, start lib.Coord, target lib.Coord) {
	maze = make(map[lib.Coord]*Entry)
	for coord, r := range d.origMaze {
		if r == 'S' {
			start = coord
			r = '.'
		}
		if r == 'E' {
			target = coord
			r = '.'
		}
		maze[coord] = &Entry{
			tile:            r,
			coord:           coord,
			distanceToStart: -1,
			visited:         false,
		}
	}
	return maze, start, target
}
