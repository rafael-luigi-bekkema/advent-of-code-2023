package main

import (
	"fmt"
	"strings"
)

const (
	damaged = '#'
	fine    = '.'
	unknown = '?'
)

func day12a() int {
	return sumtotal(mustReadInput(12), day12aRow)
}

func day12aRecur(osprings, nsprings string, nums []int, numidx, pos int) int {
	var total int
outer:
	for trypos := pos; trypos < len(nsprings); trypos++ {
		upto := trypos + nums[numidx]
		nnsprings := make([]byte, len(nsprings))
		copy(nnsprings, nsprings)
		for i := trypos; i < upto; i++ {
			if i >= len(osprings) || osprings[i] == fine {
				continue outer
			}
			nnsprings[i] = damaged
		}

		for i := pos; i < trypos; i++ {
			if osprings[i] == damaged {
				continue outer
			}
		}

		if upto < len(osprings) && osprings[upto] == damaged {
			continue outer
		}

		if numidx == len(nums)-1 {
			// Check if in the original there is a damanged spring after this point
			// If so it is not a possiblity
			for i := upto; i < len(osprings); i++ {
				chr := osprings[i]
				jchr := nnsprings[i]
				if (jchr == damaged && chr == fine) || (chr == damaged && jchr == fine) {
					nnsprings[i] = 'W'
					fmt.Println("wrong", osprings, string(nnsprings))
					continue outer
				}
			}
			total++
			// fmt.Println("got one", osprings, nums, string(nnsprings))
			continue
		}
		total += day12aRecur(osprings, string(nnsprings), nums, numidx+1, upto+1)
	}

	return total
}

func day12aRow(line string) int {
	springs, info, ok := strings.Cut(line, " ")
	if !ok {
		panic("invalid line: " + line)
	}
	nums := strsToNums(strings.Split(info, ","))

	// fmt.Println()
	nsprings := strings.Repeat(".", len(springs))
	return day12aRecur(springs, nsprings, nums, 0, 0)
}
