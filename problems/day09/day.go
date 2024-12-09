package day09

import (
	"fmt"

	"alexi.ch/aoc/2024/lib"
)

const (
	FTYPE_EMPTY = iota
	FTYPE_FILE
)

type FInfo struct {
	id    int
	size  int
	ftype int
}

type Day09 struct {
	s1          int
	s2          int
	file        []*FInfo
	initialFile []*FInfo
}

func New() Day09 {
	return Day09{s1: 0, s2: 0}
}

func (d *Day09) Title() string {
	return "Day 09 - Disk Fragmenter"
}

func (d *Day09) Setup() {
	// var lines = lib.ReadLines("data/09-test-data.txt")
	var lines = lib.ReadLines("data/09-data.txt")
	input := lines[0]
	id := 0
	for i, nrS := range input {
		nr := lib.StrToInt(string(nrS))
		var fInfo FInfo
		if i%2 == 0 {
			// file
			fInfo = FInfo{id: id, size: nr, ftype: FTYPE_FILE}
			id++
		} else {
			// empty space
			fInfo = FInfo{id: -1, size: nr, ftype: FTYPE_EMPTY}

		}
		for j := 0; j < nr; j++ {
			d.file = append(d.file, &fInfo)
			d.initialFile = append(d.initialFile, &fInfo)
		}
	}
	// fmt.Printf("%v\n", d.numbers)
	// fmt.Println(d.FileToStr(d.file))
}

func (d *Day09) SolveProblem1() {
	d.s1 = 0
	// freeBlockIdx := 0
	for blockIdx := len(d.file) - 1; blockIdx >= 0; blockIdx-- {
		blk := d.file[blockIdx]
		if blk.ftype == FTYPE_EMPTY {
			continue
		}
		// find free 1-block from the beginning of the file:
		freeBlockIdx := d.findNextFreeBlock(1)

		// no free space? we're done.
		if blockIdx <= freeBlockIdx {
			break
		}

		// swap file block with empty block:
		d.file[freeBlockIdx], d.file[blockIdx] = d.file[blockIdx], d.file[freeBlockIdx]
	}

	// calc solution
	for i, fInfo := range d.file {
		if fInfo.ftype == FTYPE_FILE {
			d.s1 += fInfo.id * i
		}
	}
}

func (d *Day09) SolveProblem2() {
	d.s2 = 0
	d.file = d.initialFile

	for blockIdx := len(d.file) - 1; blockIdx >= 0; blockIdx-- {
		blk := d.file[blockIdx]
		if blk.ftype == FTYPE_EMPTY {
			continue
		}
		// now, blockIdx stands at the END of a file block

		// find free block of the length of the actual block size:
		freeBlockIdx := d.findNextFreeBlock(blk.size)

		// if the next free block is behind our current block, we're skipping the whole block. No free space before the block.
		if blockIdx <= freeBlockIdx {
			// +1 because our loop decreases blockIdx again
			blockIdx = blockIdx - blk.size + 1
			continue
		}

		// swap all file blocks to empty block:
		for i := 0; i < blk.size; i++ {
			freeIdx := freeBlockIdx + i
			fileIdx := blockIdx - i
			d.file[freeIdx], d.file[fileIdx] = d.file[fileIdx], d.file[freeIdx]
		}
		// +1 because our loop decreases blockIdx again
		blockIdx = blockIdx - blk.size + 1

	}

	// calc solution:
	for i, fInfo := range d.file {
		if fInfo.ftype == FTYPE_FILE {
			d.s2 += fInfo.id * i
		}
	}
}

func (d *Day09) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day09) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day09) FileToStr(f []*FInfo) string {
	out := ""
	for _, fInfo := range f {
		if fInfo.ftype == FTYPE_FILE {
			out += fmt.Sprintf("%d", fInfo.id)
		} else {
			out += "."
		}
	}
	return out
}

func (d *Day09) findNextFreeBlock(length int) int {
	// searches from the beginning of the file for a free block of the given length,
	// returns the start index of the free block
outer:
	for i := 0; i < len(d.file); i++ {
		if d.file[i].ftype == FTYPE_EMPTY {
			// we found the first empty block. Now see if it is big enough:
			emptyStart := i
			l := length
			for ; l > 0; l-- {
				if i < len(d.file) && d.file[i].ftype == FTYPE_FILE {
					// nope, too short
					continue outer
				}
				i++
			}
			// good, block is big enough. Return start index.
			return emptyStart
		}
	}
	// no free block found, return the file length to indicate we have found nothing:
	return len(d.file)
}
