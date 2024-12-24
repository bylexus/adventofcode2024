package day24

import (
	"fmt"
	"regexp"
	"slices"

	"alexi.ch/aoc/2024/lib"
	"github.com/bylexus/go-stdlib/emaps"
)

var wires map[string]*Wire

type Gate interface {
	GetValue() int
}

type AndGate struct {
	valueCalculated bool
	value           int
	wireA           string
	wireB           string
}

func (g *AndGate) GetValue() int {
	if !g.valueCalculated {
		g.value = wires[g.wireA].GetValue() & wires[g.wireB].GetValue()
		g.valueCalculated = true
		return g.value
	} else {
		return g.value
	}
}

type OrGate struct {
	valueCalculated bool
	value           int
	wireA           string
	wireB           string
}

func (g *OrGate) GetValue() int {
	if !g.valueCalculated {
		g.value = (wires[g.wireA].GetValue() | wires[g.wireB].GetValue()) & 1
		g.valueCalculated = true
		return g.value
	} else {
		return g.value
	}
}

type XOrGate struct {
	valueCalculated bool
	value           int
	wireA           string
	wireB           string
}

func (g *XOrGate) GetValue() int {
	if !g.valueCalculated {
		g.value = (wires[g.wireA].GetValue() ^ wires[g.wireB].GetValue()) & 1
		g.valueCalculated = true
		return g.value
	} else {
		return g.value
	}
}

type Wire struct {
	name            string
	valueCalculated bool
	value           int
	inputGate       Gate
}

func (w *Wire) GetValue() int {
	if !w.valueCalculated {
		w.value = w.inputGate.GetValue()
		w.valueCalculated = true
		return w.value
	} else {
		return w.value
	}
}

func (w *Wire) String() string {
	return fmt.Sprintf("%s: v: %d", w.name, w.value)
}

type Day24 struct {
	s1 int
	s2 int
}

func New() Day24 {
	return Day24{s1: 0, s2: 0}
}

func (d *Day24) Title() string {
	return "Day 24 - Crossed Wires"
}

func (d *Day24) Setup() {
	wires = make(map[string]*Wire)
	// var lines = lib.ReadLines("data/24-test-data.txt")
	var lines = lib.ReadLines("data/24-data.txt")
	valueMatcher := regexp.MustCompile(`(.*): (\d)`)
	connectionMatcher := regexp.MustCompile(`(.*) (.*) (.*) -> (.*)`)
	for _, line := range lines {
		valueMatch := valueMatcher.FindStringSubmatch(line)
		if valueMatch != nil {
			wire := valueMatch[1]
			value := lib.StrToInt(valueMatch[2])
			wires[wire] = &Wire{
				name:            wire,
				valueCalculated: true,
				value:           value,
			}
		}

		connMatch := connectionMatcher.FindStringSubmatch(line)
		if connMatch != nil {
			wireA := connMatch[1]
			gateStr := connMatch[2]
			wireB := connMatch[3]
			targetWire := connMatch[4]
			var gate Gate
			if gateStr == "AND" {
				gate = &AndGate{
					valueCalculated: false,
					value:           -1,
					wireA:           wireA,
					wireB:           wireB,
				}
			} else if gateStr == "OR" {
				gate = &OrGate{
					valueCalculated: false,
					value:           -1,
					wireA:           wireA,
					wireB:           wireB,
				}
			} else if gateStr == "XOR" {
				gate = &XOrGate{
					valueCalculated: false,
					value:           -1,
					wireA:           wireA,
					wireB:           wireB,
				}
			}
			w := Wire{
				name:            targetWire,
				valueCalculated: false,
				value:           -1,
				inputGate:       gate,
			}
			wires[targetWire] = &w
		}
	}
	// fmt.Printf("%v\n", wires)
}

func (d *Day24) SolveProblem1() {
	d.s1 = 0
	wireNames := emaps.GetMapKeys(&wires)
	slices.Sort(wireNames)
	res := 0
	for _, wName := range slices.Backward(wireNames) {
		wire := wires[wName]
		if wName[0] == 'z' {
			value := wire.GetValue()
			res = (res<<1 | value)

			// fmt.Printf("Wire %s: %d\n", wName, value)
		}
	}
	d.s1 = res
}

func (d *Day24) SolveProblem2() {
	d.s2 = 0
}

func (d *Day24) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day24) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
