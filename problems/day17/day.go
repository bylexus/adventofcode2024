package day17

import (
	"fmt"
	"regexp"
	"strings"

	"alexi.ch/aoc/2024/lib"
	"golang.org/x/exp/slices"
)

type CPU struct {
	a   int
	b   int
	c   int
	ip  int
	mem []int
	out []int
}

func (c *CPU) Step() bool {
	if c.ip < len(c.mem)-1 {
		opcode := c.mem[c.ip]
		operand := c.mem[c.ip+1]
		if c.execute(opcode, operand) {
			c.ip += 2
		}
		return true
	} else {
		return false
	}
}

// executes an opcode and its operand. returns true if the ip should be incremented
func (c *CPU) execute(opcode int, operand int) bool {
	incIp := true
	switch opcode {
	case 0:
		// adv (division)
		// c.a = c.a / int(math.Pow(float64(2), float64(c.getComboOperandValue(operand))))
		c.a = c.a >> c.getComboOperandValue(operand)

	case 1:
		// bxl: bitwise XOR
		c.b = c.b ^ operand

	case 2:
		// bst: bitwise modulo 8 (keep only 3 low bits)
		c.b = c.getComboOperandValue(operand) & 0b111
	case 3:
		// jnz: jump if not zero
		if c.a != 0 {
			incIp = false
			c.ip = operand
		}
	case 4:
		// bxc: bitwise XOR
		c.b = c.b ^ c.c
	case 5:
		// out: output
		c.out = append(c.out, c.getComboOperandValue(operand)&0b111)
	case 6:
		// bdv: (division)
		// c.b = c.a / int(math.Pow(float64(2), float64(c.getComboOperandValue(operand))))
		c.b = c.a >> c.getComboOperandValue(operand)
	case 7:
		// cdv: (division)
		c.c = c.a >> c.getComboOperandValue(operand)
	}
	return incIp
}

func (c *CPU) getComboOperandValue(operand int) int {
	// Combo operands 0 through 3 represent literal values 0 through 3.
	// Combo operand 4 represents the value of register A.
	// Combo operand 5 represents the value of register B.
	// Combo operand 6 represents the value of register C.
	// Combo operand 7 is reserved and will not appear in valid programs.
	if operand >= 0 && operand <= 3 {
		return operand
	}
	if operand == 4 {
		return c.a
	}
	if operand == 5 {
		return c.b
	}
	if operand == 6 {
		return c.c
	}
	panic(fmt.Sprintf("Unknown operand: %d, ip: %d", operand, c.ip))
}

func (c *CPU) OutString() string {
	return strings.Join(lib.Map(&c.out, func(i int) string { return fmt.Sprintf("%d", i) }), ",")
}

type Day17 struct {
	s1  string
	s2  int
	cpu CPU
}

func New() Day17 {
	return Day17{s1: "", s2: 0, cpu: CPU{a: 0, b: 0, c: 0, ip: 0}}
}

func (d *Day17) Title() string {
	return "Day 17 - Chronospatial Computer"
}

func (d *Day17) Setup() {
	// var lines = lib.ReadLines("data/17-test-data.txt")
	// var lines = lib.ReadLines("data/17-test2-data.txt")
	var lines = lib.ReadLines("data/17-data.txt")
	regMatcher := regexp.MustCompile(`Register .: (\d+)`)
	programMatcher := regexp.MustCompile(`Program: (.*)`)
	d.cpu.a = lib.StrToInt(regMatcher.FindStringSubmatch(lines[0])[1])
	d.cpu.b = lib.StrToInt(regMatcher.FindStringSubmatch(lines[1])[1])
	d.cpu.c = lib.StrToInt(regMatcher.FindStringSubmatch(lines[2])[1])
	p := strings.Split(programMatcher.FindStringSubmatch(lines[4])[1], ",")
	for _, s := range p {
		d.cpu.mem = append(d.cpu.mem, lib.StrToInt(s))
	}
	fmt.Printf("%#v\n", d.cpu)
}

func (d *Day17) SolveProblem1() {
	for d.cpu.Step() {
		//
	}
	d.s1 = d.cpu.OutString()
}

func (d *Day17) SolveProblem2() {
	/*
		This is my input's instruction set:
		 0: bst 4; B = A % 8
		 2: bxl 7; B = B XOR 7 (7 = 0x111 -> Flip all bits), here: 7 - B
		 4: cdv 5; C = A / 2^B
		 6: adv 3; A = A / 2^3; --> A = A / 8 --> shift 3 left?
		 8: bxl 7; B = B XOR 7 (Flip all bits); 7 - B
		10: bxc 1; B = B XOR C
		12: out 5; OUT B % 8
		14: jmp 0; If A != 0 goto start
	*/
	// testA := 0
	testA := 35180000000000

	for {
		// d.cpu.a = 117440
		// d.cpu.a = 100000000000000 - 35180000000000
		// d.cpu.a = 35180000000000
		d.cpu.a = testA
		d.cpu.b = 0
		d.cpu.c = 0
		d.cpu.ip = 0
		d.cpu.out = make([]int, 0, 32)
		// fmt.Printf("testA: %d\n", testA)
		for d.cpu.Step() {
			//
		}
		// fmt.Printf("out: %s\n", d.cpu.OutString())
		// fmt.Printf("len(out): %d\n", len(d.cpu.out))
		if slices.Compare(d.cpu.out, d.cpu.mem) == 0 {
			break
		}
		testA++
		// break
	}
	d.s2 = testA
}

func (d *Day17) Solution1() string {
	return d.s1
}

func (d *Day17) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
