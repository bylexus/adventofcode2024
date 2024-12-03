package problems

import (
	"fmt"
	"strconv"
	"strings"

	"alexi.ch/aoc/2024/lib"
	"github.com/bylexus/go-stdlib/eerr"
)

type Day02 struct {
	reports    [][]int
	badReports [][]int
	s1         int
	s2         int
}

func NewDay02() Day02 {
	return Day02{s1: 0, s2: 0, reports: make([][]int, 0), badReports: make([][]int, 0)}
}

func (d *Day02) Title() string {
	return "Day 02 - Red-Nosed Reports"
}

func (d *Day02) Setup() {
	// var lines = lib.ReadLines("data/02-test-data.txt")
	var lines = lib.ReadLines("data/02-data.txt")
	for _, line := range lines {
		nrs := strings.Split(line, " ")
		report := make([]int, 0)
		for _, nr := range nrs {
			n, err := strconv.Atoi(nr)
			eerr.PanicOnErr(err)
			report = append(report, n)
		}
		d.reports = append(d.reports, report)
	}
	// fmt.Printf("%#v\n", d.reports)
}

func (d *Day02) SolveProblem1() {
	saveReports := 0
	for _, report := range d.reports {
		if d.checkReport(report) {
			saveReports++
		} else {
			// save bad report for part 2:
			d.badReports = append(d.badReports, report)
		}
	}
	d.s1 = saveReports
}

func (d *Day02) SolveProblem2() {
	stillSaveReports := d.s1
	// just go through the bad reports from part 1, and remove
	// one number at a time, check if it will be OK:
	for _, report := range d.badReports {
		for i := 0; i < len(report); i++ {
			newReport := lib.Splice(report, i)
			if d.checkReport(newReport) {
				stillSaveReports++
				break
			}
		}
	}
	d.s2 = stillSaveReports
}

func (d *Day02) checkReport(report []int) bool {
	var minDiff int
	var maxDiff int
	if report[0] > report[1] {
		minDiff = 1
		maxDiff = 3
	} else if report[0] < report[1] {
		minDiff = -3
		maxDiff = -1
	} else {
		return false
	}
	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]
		if diff == 0 || diff < minDiff || diff > maxDiff {
			return false
		}
	}
	return true
}

func (d *Day02) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day02) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
