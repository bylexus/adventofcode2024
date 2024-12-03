package problems

import (
	"fmt"

	"alexi.ch/aoc/2024/lib"
)

type Playground struct {
}

func NewPlayground() Playground {
	return Playground{}
}

func (d *Playground) Title() string {
	return "Playground"
}

func (d *Playground) Setup() {
	c1 := lib.NewCoord0()
	c2 := lib.NewCoord2D(-3, 5)
	c3 := lib.NewCoord2D(2, 7)
	fmt.Printf("%s\n", c1)
	fmt.Printf("%s\n", c2)
	fmt.Printf("%s\n", c3)
	fmt.Printf("manhattan dis: %d\n", c2.Manhattan(c3))
}

func (d *Playground) SolveProblem1() {
}

func (d *Playground) SolveProblem2() {
}

func (d *Playground) Solution1() string {
	return fmt.Sprintf("%d", 0)
}

func (d *Playground) Solution2() string {
	return fmt.Sprintf("%d", 0)
}
