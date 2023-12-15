package main

import (
	"bytes"
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

func day12b() int {
	return day12bTotal(mustReadInput(12), true)
}

type day12job struct {
	osprings   string
	nums       []int
	numSums    []int
	lastDmgIdx int
	len        int
	lennums    int
}

func day12newjob(springs string, nums []int) *day12job {
	numSums := make([]int, len(nums))
	var sum int
	for i := len(numSums) - 1; i >= 0; i-- {
		sum += nums[i]
		if i < len(numSums)-1 {
			sum++
		}
		numSums[i] = sum
	}

	return &day12job{springs, nums, numSums, strings.LastIndex(springs, "#"), len(springs), len(nums)}
}

func (job *day12job) day12aRecur(numidx, pos int, posibility []byte) int {
	var total int

	if (job.len - pos) < job.numSums[numidx] {
		return 0
	}

	num := job.nums[numidx]
	maxnum := job.len - num
outer:
	for trypos := pos; trypos <= maxnum; trypos++ {
		upto := trypos + num

		// if we have damaged spring before this point
		// there is no point in continuing
		for i := pos; i < trypos; i++ {
			if job.osprings[i] == damaged {
				break outer
			}
		}

		// see if none of the selected spots
		// are fine in the original
		for i := upto - 1; i >= trypos; i-- {
			if job.osprings[i] == fine {
				trypos = i
				continue outer
			}
		}

		// the spot after the selection should be 'fine'
		if upto < job.len && job.osprings[upto] == damaged {
			continue outer
		}

		ppossibility := make([]byte, len(posibility))
		copy(ppossibility, posibility)
		for i := trypos; i < upto-1; i++ {
			ppossibility[i] = '#'
		}

		// if we are at the last number
		if numidx == job.lennums-1 {
			// check if there should be damaged springs after our selection
			// if so skip to that point
			if job.lastDmgIdx >= upto {
				trypos = job.lastDmgIdx - num
				continue
			}

			fmt.Println(string(ppossibility))

			total++
			continue
		}

		total += job.day12aRecur(numidx+1, upto+1, ppossibility)
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
	return day12newjob(springs, nums).day12aRecur(0, 0, bytes.Repeat([]byte{fine}, len(springs)))
}

func day12bRow(line string) int {
	springs, info, ok := strings.Cut(line, " ")
	if !ok {
		panic("invalid line: " + line)
	}
	springs = strings.Join([]string{springs, springs, springs, springs, springs}, "?")
	info = strings.Join([]string{info, info, info, info, info}, ",")

	nums := strsToNums(strings.Split(info, ","))

	return day12newjob(springs, nums).day12aRecur(0, 0, bytes.Repeat([]byte{fine}, len(springs)))
}

func day12bTotal(lines []string, withfile bool) int {
	jobs := make(chan string)
	type jobresult struct {
		job    string
		result int
	}
	result := make(chan jobresult)
	for i := 0; i < 10; i++ {
		go func() {
			for job := range jobs {
				result <- jobresult{job, day12bRow(job)}
			}
		}()
	}

	go func() {
		for _, line := range lines {
			jobs <- line
		}
		close(jobs)
	}()

	var total int
	for i := 0; i < len(lines); i++ {
		res := <-result
		total += res.result
		fmt.Printf("Done %.01f %%\n", (float64(i)/float64(len(lines)))*100)
	}
	return total
}
