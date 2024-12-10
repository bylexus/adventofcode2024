package day10

import (
	"fmt"
	"time"

	"alexi.ch/aoc/2024/lib"
	"github.com/bylexus/go-stdlib/eerr"
	"github.com/gdamore/tcell/v2"
)

type Day10 struct {
	s1          int
	s2          int
	topologyMap [][]int

	trailHeads map[lib.Coord]int
	// trailEnds  map[lib.Coord]int

}

func New() Day10 {
	return Day10{s1: 0, s2: 0, topologyMap: make([][]int, 0), trailHeads: make(map[lib.Coord]int)}
}

func (d *Day10) Title() string {
	return "Day 10 - Hoof It"
}

func (d *Day10) Setup() {
	// var lines = lib.ReadLines("data/10-test-data.txt")
	var lines = lib.ReadLines("data/10-data.txt")
	for y, line := range lines {
		xLine := make([]int, 0)
		for x, r := range line {
			nr := lib.StrToInt(string(r))
			coord := lib.NewCoord2D(x, y)
			if nr == 0 {
				d.trailHeads[coord] = 0
			}
			xLine = append(xLine, nr)
		}
		d.topologyMap = append(d.topologyMap, xLine)
	}
	// fmt.Printf("%v\n", d.topologyMap)
}

func (d *Day10) SolveProblem1() {
	s, err := tcell.NewScreen()
	eerr.PanicOnErr(err)
	err = s.Init()
	eerr.PanicOnErr(err)

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault)
	highlightStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorRed).Bold(true)
	s.SetStyle(defStyle)
	s.Clear()

	d.drawMap(s, defStyle)
	s.Show()

	go func() {
		d.s1 = 0
		for startPoint := range d.trailHeads {
			targetCoords := d.ramble(startPoint, s, highlightStyle)
			// we only need unique end coords per start point, so create a set:
			targetCoordSet := make(map[lib.Coord]int)
			for _, coord := range targetCoords {
				targetCoordSet[coord] += 1
			}
			d.s1 += len(targetCoordSet)
		}

	}()

	for {
		ev := s.PollEvent()
		switch ev.(type) {
		case *tcell.EventKey:
			s.Fini()
			return
		}
	}

}

func (d *Day10) SolveProblem2() {
	d.s2 = 0
	// Solution 2 in this case is even simpler than solution 1 :-))
	// We do not need unique end coords, but count all reached end coords:
	// for startPoint := range d.trailHeads {
	// 	targetCoords := d.ramble(startPoint)
	// 	d.s2 += len(targetCoords)
	// }
}

func (d *Day10) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day10) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day10) ramble(start lib.Coord, s tcell.Screen, style tcell.Style) []lib.Coord {
	actNr := d.topologyMap[start.Y][start.X]
	allTargetCoords := make([]lib.Coord, 0)

	style = style.Foreground(tcell.NewRGBColor(int32((actNr+10)*12), int32(0), int32(0))).Background(tcell.NewRGBColor(int32(actNr*25), int32(actNr*25), int32(actNr*25)))
	s.SetContent(start.X, start.Y, '*', nil, style)
	s.Show()
	time.Sleep(20 * time.Millisecond)

	// recursive break condition: If we sit on a 9, we have reached
	// our destination: we return the coords of that destination
	if actNr == 9 {
		return []lib.Coord{start}
	}

	// if not, we walk in all 4 dirs, if possible, and collect
	// target coords:
	for _, dirVec := range lib.MOVE_VEC_2D_4DIRS {
		nextCoord := start.Add(lib.NewCoord2D(dirVec[0], dirVec[1]))
		if nextCoord.X >= 0 && nextCoord.Y >= 0 && nextCoord.X < len(d.topologyMap[0]) && nextCoord.Y < len(d.topologyMap) {
			nextNr := d.topologyMap[nextCoord.Y][nextCoord.X]
			if actNr+1 == nextNr {
				resultCoords := d.ramble(nextCoord, s, style)
				allTargetCoords = append(allTargetCoords, resultCoords...)
			}
		}
	}

	return allTargetCoords
}

func (d *Day10) drawMap(s tcell.Screen, style tcell.Style) {
	s.SetContent(0, 0, 'A', nil, style)

	for y, line := range d.topologyMap {
		for x, nr := range line {
			style = style.Foreground(tcell.NewRGBColor(int32(nr*25), int32(nr*25), int32(nr*25)))
			s.SetContent(x, y, '█', nil, style)
		}
	}
}
