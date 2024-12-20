package day18

import (
	"fmt"
	"regexp"
	"slices"

	"alexi.ch/aoc/2024/lib"
)

type Entry struct {
	tile            rune
	coord           lib.Coord
	visited         bool
	distanceToStart int
}

type Day18 struct {
	s1     int
	s2     string
	coords []lib.Coord
}

func New() Day18 {
	return Day18{s1: 0, s2: ""}
}

func (d *Day18) Title() string {
	return "Day 18 - RAM Run"
}

func (d *Day18) Setup() {
	// var lines = lib.ReadLines("data/18-test-data.txt")
	var lines = lib.ReadLines("data/18-data.txt")
	matcher := regexp.MustCompile(`(\d+),(\d+)`)
	for _, line := range lines {
		matches := matcher.FindStringSubmatch(line)
		if matches != nil {
			x := lib.StrToInt(matches[1])
			y := lib.StrToInt(matches[2])
			d.coords = append(d.coords, lib.NewCoord2D(x, y))
		}
	}

	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day18) SolveProblem1() {
	d.s1 = 0
	// width, height := 7, 7
	width, height := 71, 71
	// coords := d.coords
	// coords := d.coords[:12]
	coords := d.coords[:1024]
	maze := make(map[lib.Coord]*Entry)

	for y := range height {
		for x := range width {
			maze[lib.NewCoord2D(x, y)] = &Entry{
				tile:            ' ',
				coord:           lib.NewCoord2D(x, y),
				distanceToStart: -1,
				visited:         false,
			}
		}
	}

	for _, coord := range coords {
		maze[coord].tile = '#'
	}

	// d.printMaze(maze, width, height)

	start := lib.NewCoord0()
	end := lib.NewCoord2D(width-1, height-1)
	d.walkMaze(maze, start, end)
	// fmt.Printf("End node: %#v\n", *(maze[end]))
	d.s1 = maze[end].distanceToStart
}

func (d *Day18) SolveProblem2() {
	// width, height := 7, 7
	width, height := 71, 71
	coords := d.coords
	// coords := d.coords[:12]
	// coords := d.coords[:1024]
	maze := make(map[lib.Coord]*Entry)

	for y := range height {
		for x := range width {
			maze[lib.NewCoord2D(x, y)] = &Entry{
				tile:            ' ',
				coord:           lib.NewCoord2D(x, y),
				distanceToStart: -1,
				visited:         false,
			}
		}
	}

	for i := 1; i < len(coords)+1; i++ {
		coords := d.coords[:i]

		// reset maze
		for y := range height {
			for x := range width {
				maze[lib.NewCoord2D(x, y)].distanceToStart = -1
				maze[lib.NewCoord2D(x, y)].visited = false
				maze[lib.NewCoord2D(x, y)].tile = ' '
			}
		}

		// fill in new coords
		for _, coord := range coords {
			maze[coord].tile = '#'
		}
		start := lib.NewCoord0()
		end := lib.NewCoord2D(width-1, height-1)

		d.walkMaze(maze, start, end)

		if maze[end].distanceToStart == -1 {
			// fmt.Printf("Corrupting entry: %#v\n", coords[len(coords)-1])
			d.s2 = fmt.Sprintf("%d,%d", coords[len(coords)-1].X, coords[len(coords)-1].Y)
			break
		}
	}

	// fmt.Printf("End node: %#v\n", *(maze[end]))
	// d.s1 = maze[end].distanceToStart
}

func (d *Day18) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day18) Solution2() string {
	return d.s2
}

func (d *Day18) walkMaze(maze map[lib.Coord]*Entry, start lib.Coord, target lib.Coord) {
	maze[start].distanceToStart = 0
	maze[start].visited = false
	unvisited := []*Entry{maze[start]}

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

func (d *Day18) printMaze(m map[lib.Coord]*Entry, width, height int) {
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
