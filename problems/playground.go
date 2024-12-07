package problems

import (
	"fmt"
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
}

func (d *Playground) SolveProblem1() {
	sl := []int{1, 2, 3}
	fmt.Printf("%#v\n", sl)
	modSlice(sl, 1, 42)
	fmt.Printf("%#v\n", sl)
}

func (d *Playground) SolveProblem2() {
}

func (d *Playground) Solution1() string {
	return fmt.Sprintf("%d", 0)
}

func (d *Playground) Solution2() string {
	return fmt.Sprintf("%d", 0)
}

func modSlice(a []int, idx int, value int) {
	a[idx] = value
}
