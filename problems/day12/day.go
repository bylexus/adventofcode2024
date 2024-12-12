package day12

import (
	"fmt"
	"math/rand"
	"time"

	"alexi.ch/aoc/2024/lib"
	"github.com/bylexus/go-stdlib/eerr"
	"github.com/gdamore/tcell/v2"
)

type FenceMap struct {
	t bool
	r bool
	b bool
	l bool
}

type Plot struct {
	id      rune
	region  *Region
	visited bool
	fences  FenceMap
	coords  lib.Coord
}

type Region struct {
	id    rune
	color tcell.Color
	plots []*Plot
}

func (r *Region) Fences() int {
	fences := 0
	for _, plot := range r.plots {
		if plot.fences.t {
			fences++
		}
		if plot.fences.r {
			fences++
		}
		if plot.fences.b {
			fences++
		}
		if plot.fences.l {
			fences++
		}
	}
	return fences
}

func (r *Region) Area() int {
	return len(r.plots)
}

type Day12 struct {
	s1 int
	s2 int

	width, height int
	garden        map[lib.Coord]*Plot
	regions       []*Region
}

func New() Day12 {
	return Day12{s1: 0, s2: 0, garden: make(map[lib.Coord]*Plot), regions: make([]*Region, 0)}
}

func (d *Day12) Title() string {
	return "Day 12 - Garden Groups"
}

func (d *Day12) Setup() {
	// var lines = lib.ReadLines("data/12-test-data.txt")
	var lines = lib.ReadLines("data/12-data.txt")
	d.height = len(lines)
	d.width = len(lines[0])
	for y, line := range lines {
		for x, r := range line {
			p := Plot{id: r, coords: lib.NewCoord2D(x, y), visited: false, fences: FenceMap{}, region: nil}
			d.garden[lib.NewCoord2D(x, y)] = &p
		}
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day12) SolveProblem1() {
	d.s1 = 0
	// walk through all the garden,
	// flood-fill areas while keeping track of the number of fences and areas.
	// when a new plot is found, we start a new region, and flood-fill it until done.
	// Then we search for the next area on unvisited patches.

	s, err := tcell.NewScreen()
	eerr.PanicOnErr(err)
	s.Init()
	defStyle := tcell.StyleDefault.Background(tcell.NewRGBColor(50, 50, 50)).Foreground(tcell.ColorDarkGray)
	// s.SetStyle(defStyle)

	for _, plot := range d.garden {
		s.SetContent(plot.coords.X, plot.coords.Y, ' ', nil, defStyle)
	}

	for y := 0; y < d.height; y++ {
		for x := 0; x < d.width; x++ {
			actPlot := d.garden[lib.NewCoord2D(x, y)]
			if actPlot.visited {
				continue
			}
			// mark a whole region:
			// as a side-effect, also count the plot fences and on which side they are:
			d.floodFill(actPlot, s)
		}
	}
	for _, r := range d.regions {
		d.s1 += r.Fences() * r.Area()
	}

outer:
	for {
		ev := s.PollEvent()
		switch ev.(type) {
		case *tcell.EventKey:
			break outer
		}
	}

	s.Fini()
}

func (d *Day12) SolveProblem2() {
	d.s2 = 0
	// We re-use the already processed garden patches from Solution 1: We need the fence information from
	// S1 again.
	// check each plot, in all dirs: If top and left is the same, but t/l not, then we found an "outer" facing corner
	// add that to the "inner" facing corners
	for _, r := range d.regions {
		corners := 0
		for _, p := range r.plots {
			corners += d.countCorners(p)
		}
		d.s2 += corners * r.Area()
	}
}

func (d *Day12) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day12) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day12) countFences(coord lib.Coord) FenceMap {
	fenceMap := FenceMap{}
	plot := d.garden[coord]
	for _, vec := range lib.MOVE_VEC_2D_4DIRS {
		nextCoord := coord.Add(lib.NewCoord2D(vec[0], vec[1]))

		if nextPlot, ok := d.garden[nextCoord]; ok {
			if nextPlot.id != plot.id {
				fenceMap = d.vec2FenceMap(vec, fenceMap)
			}
		} else {
			fenceMap = d.vec2FenceMap(vec, fenceMap)
		}
	}
	return fenceMap
}

func (d *Day12) vec2FenceMap(vec []int, f FenceMap) FenceMap {
	// vec is an x/y moving vector, see lib.MOVE_VEC_2D_4DIRS
	if vec[0] == 1 {
		f.r = true
	}
	if vec[0] == -1 {
		f.l = true
	}
	if vec[1] == 1 {
		f.b = true
	}
	if vec[1] == -1 {
		f.t = true
	}
	return f
}

func (d *Day12) floodFill(plot *Plot, s tcell.Screen) {
	if plot.visited {
		return
	}
	plot.visited = true
	plot.fences = d.countFences(plot.coords)
	region := plot.region
	// no region yet? So we found a new one, create it:
	if region == nil {
		region = &Region{id: plot.id, plots: make([]*Plot, 0), color: createRandomColor()}
		d.regions = append(d.regions, region)
	}
	region.plots = append(region.plots, plot)
	plot.region = region

	s.SetContent(plot.coords.X, plot.coords.Y, '█', nil, tcell.StyleDefault.Foreground(region.color))
	s.Show()
	time.Sleep(1 * time.Millisecond)

	// look in all 4 dirs, and check if we still have to flood-fill:
	for _, vec := range lib.MOVE_VEC_2D_4DIRS {
		nextCoord := plot.coords.Add(lib.NewCoord2D(vec[0], vec[1]))

		if nextPlot, ok := d.garden[nextCoord]; ok {
			if !nextPlot.visited {
				if plot.id == nextPlot.id {
					// we set the next plot to the same region,
					// as it belongs to us:
					nextPlot.region = region
					d.floodFill(nextPlot, s)
				} else {
					// pass: we do not walk into another region. This is the main loop's
					// task

				}
			}
		}
	}
}

// Count all corners the actual spot is in. This is somewhat tricky, as you can see below:
func (d *Day12) countCorners(p *Plot) int {
	corners := 0
	/*
			We fetch all 8 patches surrounding the actual plot:
		0,0 +-+-+-+
			|C|C|C|
			+-+-+-+
			|C|X|C| <-- X marks the actual spot
			+-+-+-+
			|C|C|C|
			+-+-+-+ 2,2
	*/
	var surroundingPlots [][]*Plot = make([][]*Plot, 3)
	for y := 0; y < 3; y++ {
		surroundingPlots[y] = make([]*Plot, 3)
		for x := 0; x < 3; x++ {
			surroundingPlots[y][x] = d.garden[p.coords.Add(lib.NewCoord2D(x-1, y-1))]
		}
	}
	left := surroundingPlots[1][0]
	right := surroundingPlots[1][2]
	top := surroundingPlots[0][1]
	bottom := surroundingPlots[2][1]
	topleft := surroundingPlots[0][0]
	topright := surroundingPlots[0][2]
	bottomleft := surroundingPlots[2][0]
	bottomright := surroundingPlots[2][2]

	/*
			  Differentiate between "inner" and "outer" corners:

		        +-+  /-- this is an outer corner: the "C" marked sits at the outer-facing corner: has a C top and left, but not diagonally
		        |C| /
		        + +-+
		    --> |C C|
		        +-+-+
				^
				|
				+---- this is an inner corner: the C is walled-in on one corner, with 2 walls on the sides
	*/

	// "inner" corners: if a plot has 2 perpendicular fences, then that is a inner corner (walls e.g. top/left)
	if p.fences.t && p.fences.r {
		corners++
	}
	if p.fences.t && p.fences.l {
		corners++
	}
	if p.fences.b && p.fences.r {
		corners++
	}
	if p.fences.b && p.fences.l {
		corners++
	}

	// outer corners:

	// check diagonaltop left:
	/*
	    +-+
	   A|C|
	  +-+ +
	  |C C| <-- this is our actual spot we check
	  +-+-+
	*/
	if left != nil && left.id == p.id && top != nil && top.id == p.id && topleft != nil && topleft.id != p.id {
		corners++
	}

	// check diagonaltop right
	/*
			  +-+-+
			  |C|A|<-- this is our actual spot we check
			  + +-+
		   -->|C C|
			  +-+-+
	*/
	if right != nil && right.id == p.id && top != nil && top.id == p.id && topright != nil && topright.id != p.id {
		corners++
	}

	// check diagonal bottom right
	/*
			  +-+-+
		   -->|C C|
			  + +-+
		      |C|A|<-- this is our actual spot we check
			  +-+-+
	*/
	if right != nil && right.id == p.id && bottom != nil && bottom.id == p.id && bottomright != nil && bottomright.id != p.id {
		corners++
	}

	// check diagonal bottom left
	/*
			  +-+-+
		      |C C| <--
			  +-+ +
		   +->|A|C|
		   |  +-+-+
		   |
		   +----------- this is our actual spot we check
	*/
	if left != nil && left.id == p.id && bottom != nil && bottom.id == p.id && bottomleft != nil && bottomleft.id != p.id {
		corners++
	}

	return corners
}

func createRandomColor() tcell.Color {
	return tcell.NewRGBColor(int32(rand.Intn(256)), int32(rand.Intn(256)), int32(rand.Intn(256)))
}
