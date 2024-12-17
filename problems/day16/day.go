package day16

import (
	"fmt"
	"slices"

	"alexi.ch/aoc/2024/lib"
)

type Entry struct {
	tile            rune
	coord           lib.Coord
	facing          rune
	visited         bool
	distanceToStart int
}

type Day16 struct {
	s1            int
	s2            int
	maze          map[lib.Coord]*Entry
	start         lib.Coord
	target        lib.Coord
	width, height int
}

func New() Day16 {
	return Day16{s1: 0, s2: 0, maze: make(map[lib.Coord]*Entry, 0)}
}

func (d *Day16) Title() string {
	return "Day 16 - Reindeer Maze"
}

func (d *Day16) Setup() {
	// var lines = lib.ReadLines("data/16-test-data.txt")
	var lines = lib.ReadLines("data/16-data.txt")
	d.width = len(lines[0])
	d.height = len(lines)
	for y, line := range lines {
		for x, r := range line {
			tile := r
			if tile == 'S' {
				tile = '.'
				d.start = lib.NewCoord2D(x, y)
			}
			if tile == 'E' {
				tile = '.'
				d.target = lib.NewCoord2D(x, y)
			}
			d.maze[lib.NewCoord2D(x, y)] = &Entry{
				tile:            tile,
				coord:           lib.NewCoord2D(x, y),
				distanceToStart: -1,
				visited:         false,
			}
		}
	}
	// fmt.Printf("%v\n", d.numbers)
	// d.printMaze(d.maze)
}

func (d *Day16) SolveProblem1() {
	d.s1 = 0
	// distance to start for start node is, obviously, 0.
	d.maze[d.start].distanceToStart = 0
	d.maze[d.start].facing = 'E'
	unvisited := make([]*Entry, 0)
	for _, entry := range d.maze {
		if entry.tile == '.' {
			unvisited = append(unvisited, entry)
		}
	}
	// sort: by distance to start: lowest entry in front
	for len(unvisited) > 0 {
		slices.SortFunc(unvisited, func(a, b *Entry) int {
			if a.distanceToStart == -1 {
				return 1
			}
			if b.distanceToStart == -1 {
				return -1
			}
			return a.distanceToStart - b.distanceToStart
		})

		act := unvisited[0]
		act.visited = true
		if act.distanceToStart == -1 {
			break
		}
		unvisited = append([]*Entry{}, unvisited[1:]...)

		// update neighbour node's distance:
		for _, c := range lib.MOVE_VEC_2D_4DIRS {
			nextCoord := act.coord.AddXY(c[0], c[1])
			if d.maze[nextCoord].tile == '#' {
				continue
			}
			// calc cost to move to next tile, based on act facing
			cost := 1
			if act.facing == 'E' && c[1] == -1 {
				cost = 1001
			}
			if act.facing == 'E' && c[1] == 1 {
				cost = 1001
			}
			if act.facing == 'E' && c[0] == -1 {
				cost = 2001
			}
			if act.facing == 'N' && c[0] == -1 {
				cost = 1001
			}
			if act.facing == 'N' && c[0] == 1 {
				cost = 1001
			}
			if act.facing == 'N' && c[1] == 1 {
				cost = 2001
			}
			if act.facing == 'W' && c[1] == -1 {
				cost = 1001
			}
			if act.facing == 'W' && c[1] == 1 {
				cost = 1001
			}
			if act.facing == 'W' && c[0] == 1 {
				cost = 2001
			}
			if act.facing == 'S' && c[0] == -1 {
				cost = 1001
			}
			if act.facing == 'S' && c[0] == 1 {
				cost = 1001
			}
			if act.facing == 'S' && c[1] == -1 {
				cost = 2001
			}
			// do we need to update the neighbour's number, because it is not yet set or larger than the actual distance?
			if d.maze[nextCoord].distanceToStart < 0 || d.maze[nextCoord].distanceToStart > d.maze[act.coord].distanceToStart+cost {
				d.maze[nextCoord].distanceToStart = d.maze[act.coord].distanceToStart + cost
				if c[0] == -1 {
					d.maze[nextCoord].facing = 'W'
				}
				if c[0] == 1 {
					d.maze[nextCoord].facing = 'E'
				}
				if c[1] == 1 {
					d.maze[nextCoord].facing = 'S'
				}
				if c[1] == -1 {
					d.maze[nextCoord].facing = 'N'
				}
			}
		}
		// if act.coord == d.target {
		// 	break
		// }
	}
	// d.printMaze(d.maze)
	fmt.Printf("end node: %#v\n", d.maze[d.target])
	d.s1 = d.maze[d.target].distanceToStart
}

func (d *Day16) SolveProblem2() {
	d.s2 = 0

	// now we walk backwards through the maze: we already found the shortest path for every tile
	// in solution 1. Now we do the following backwards:
	// 1. start at the target
	// 2. go back in the direction we came from, considering all the dirs with the (same) smallest number, put them in a queue
	// 3. repeat for each queued tile, until the queue is empty
	tilesOnShortestPath := make([]*Entry, 0)
	tilesOnShortestPath = append(tilesOnShortestPath, d.maze[d.target])
	queue := make([]*Entry, 0)
	queue = append(queue, d.maze[d.target])

	for len(queue) > 0 {
		act := queue[0]
		queue = append([]*Entry{}, queue[1:]...)
		// select all surrounding tiles that have a smaller distance than the actual one
		for _, c := range lib.MOVE_VEC_2D_4DIRS {
			nextCoord := act.coord.AddXY(c[0], c[1])
			if d.maze[nextCoord].tile == '#' {
				continue
			}
			if d.maze[nextCoord].distanceToStart < act.distanceToStart || d.maze[nextCoord].distanceToStart-999 == act.distanceToStart {
				if !slices.Contains(tilesOnShortestPath, d.maze[nextCoord]) {
					tilesOnShortestPath = append(tilesOnShortestPath, d.maze[nextCoord])
					queue = append(queue, d.maze[nextCoord])
				}
			}
		}
	}
	d.s2 = len(tilesOnShortestPath)
}

func (d *Day16) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day16) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day16) printMaze(m map[lib.Coord]*Entry) {
	for y := 0; y < d.height; y++ {
		for x := 0; x < d.width; x++ {
			e := m[lib.NewCoord2D(x, y)]
			if e.tile == '#' {
				fmt.Printf("#####")
			} else {
				fmt.Printf("%5d", e.distanceToStart)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
