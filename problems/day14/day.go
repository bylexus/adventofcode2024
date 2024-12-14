package day14

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"

	"alexi.ch/aoc/2024/lib"
)

type Robot struct {
	startPos lib.Coord
	actPos   lib.Coord
	velocity lib.Coord
}

func (r *Robot) CalcNewPos(width, height, rounds int) {
	endX := (r.startPos.X + r.velocity.X*rounds) % width
	endY := (r.startPos.Y + r.velocity.Y*rounds) % height
	if endX < 0 {
		endX = width + endX
	}
	if endY < 0 {
		endY = height + endY
	}
	r.actPos = lib.NewCoord2D(endX, endY)
}

type Day14 struct {
	s1 int
	s2 int

	robots []Robot
}

func New() Day14 {
	return Day14{s1: 0, s2: 0, robots: make([]Robot, 0)}
}

func (d *Day14) Title() string {
	return "Day 14 - Restroom Redoubt"
}

func (d *Day14) Setup() {
	// var lines = lib.ReadLines("data/14-test-data.txt")
	var lines = lib.ReadLines("data/14-data.txt")
	matcher := regexp.MustCompile(`p=([-\d]+),([-\d]+) v=([-\d]+),([-\d]+)`)
	for _, line := range lines {
		infos := matcher.FindStringSubmatch(line)
		if infos != nil {
			d.robots = append(d.robots, Robot{
				startPos: lib.NewCoord2D(lib.StrToInt(infos[1]), lib.StrToInt(infos[2])),
				actPos:   lib.NewCoord2D(lib.StrToInt(infos[1]), lib.StrToInt(infos[2])),
				velocity: lib.NewCoord2D(lib.StrToInt(infos[3]), lib.StrToInt(infos[4])),
			})
		}

	}
	// fmt.Printf("%#v\n", d.robots)
}

func (d *Day14) SolveProblem1() {
	d.s1 = 0
	// width := 11
	// height := 7
	width := 101
	height := 103
	rounds := 100
	midX := width / 2
	midY := height / 2

	q1, q2, q3, q4 := 0, 0, 0, 0

	for i, robot := range d.robots {
		robot.CalcNewPos(width, height, rounds)
		d.robots[i] = robot
		if robot.actPos.X < midX && robot.actPos.Y < midY {
			q1++
		} else if robot.actPos.X > midX && robot.actPos.Y < midY {
			q2++
		} else if robot.actPos.X < midX && robot.actPos.Y > midY {
			q3++
		} else if robot.actPos.X > midX && robot.actPos.Y > midY {
			q4++
		}
	}
	d.s1 = q1 * q2 * q3 * q4
}

func (d *Day14) SolveProblem2() {
	d.s2 = 0
	width := 101
	height := 103
	rounds := 0

	for {
		rounds++
		coords := make(map[lib.Coord]bool)
		for i, robot := range d.robots {
			robot.CalcNewPos(width, height, rounds)
			d.robots[i] = robot
			// create a set of end coords, for later anaylzing
			coords[robot.actPos] = true
		}

		// Draw an Image of the actual state of robots:
		// img := image.NewRGBA(image.Rect(0, 0, 101, 103))
		// for _, robot := range d.robots {
		// 	img.Set(robot.actPos.X, robot.actPos.Y, color.Black)
		// 	// fmt.Printf("Robot %d moved to (%d, %d)\n", i, endX, endY)
		// }
		// // Save the image to a PNG file
		// f, err := os.Create(fmt.Sprintf("day14_output/%8d.png", rounds))
		// if err != nil {
		// 	panic(err)
		// }
		// png.Encode(f, img)
		// f.Close()

		// find a line of 20 pxs in a row: There must be the border of the tree:
		for y := 0; y < height; y++ {
			for x := 0; x < width-20; x++ {
				found := true
				for i := 0; i < 20; i++ {
					if _, ok := coords[lib.NewCoord2D(x+i, y)]; !ok {
						found = false
						break
					}
				}
				if found {
					// fmt.Println("Found the christmas tree in Round %d!\n", rounds)
					d.s2 = rounds
					// Draw an Image of the actual state of robots:
					img := image.NewRGBA(image.Rect(0, 0, 101, 103))
					for y := 0; y < height; y++ {
						for x := 0; x < width; x++ {
							img.Set(x, y, color.White)
						}
					}

					for _, robot := range d.robots {
						img.Set(robot.actPos.X, robot.actPos.Y, color.RGBA{0, 180, 0, 255})
					}
					// Save the image to a PNG file
					f, err := os.Create(fmt.Sprintf("day14_tree_round_%d.png", rounds))
					if err != nil {
						panic(err)
					}
					png.Encode(f, img)
					f.Close()
					return
				}
			}
		}

		// fmt.Println(d.drawField(width, height, d.robots))
		// fmt.Printf("Round %d\n", rounds)
		// time.Sleep(50 * time.Millisecond)
		// check the vertical middle: is it filled with robots?
		// yes := true
		// for y := 40; y < height-40; y++ {
		// 	midHasRobot := false
		// 	for _, r := range d.robots {
		// 		if r.actPos.Y == y && r.actPos.X == midX {
		// 			midHasRobot = true
		// 			break
		// 		}
		// 	}
		// 	if !midHasRobot {
		// 		yes = false
		// 		break
		// 	}
		// }
		// if yes {
		// 	fmt.Println(d.drawField(width, height, d.robots))
		// 	fmt.Printf("Round %d\n", rounds)
		// 	break
		// }

		// time.Sleep(50 * time.Millisecond)
		// if rounds >= 2024 {
		// 	break

		// }
	}
}

func (d *Day14) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day14) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day14) drawField(width, height int, robots []Robot) string {
	out := ""

	field := make([][]rune, height)
	for y := 0; y < height; y++ {
		line := make([]rune, width)
		for x := 0; x < width; x++ {
			line[x] = '.'
		}
		field[y] = line
	}
	for _, robot := range robots {
		field[robot.actPos.Y][robot.actPos.X] = '*'
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			out += string(field[y][x])
		}
		out += "\n"
	}
	return out
}
