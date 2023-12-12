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

func day12b() int {
	return day12bTotal(mustReadInput(12))
}

type day12job struct {
	osprings   string
	nums       []int
	lastDmgIdx int
	len        int
	lennums    int
}

func day12newjob(springs string, nums []int) *day12job {
	return &day12job{springs, nums, strings.LastIndex(springs, "#"), len(springs), len(nums)}
}

func (job *day12job) day12aRecur(numidx, pos int) int {
	var total int

	if (job.len - pos) < sum(job.nums[numidx:])+len(job.nums[numidx:])-1 {
		return 0
	}

	num := job.nums[numidx]
	maxnum := job.len - num
outer:
	for trypos := pos; trypos <= maxnum; trypos++ {
		upto := trypos + num

		for i := pos; i < trypos; i++ {
			if job.osprings[i] == damaged {
				break outer
			}
		}

		for i := upto - 1; i >= trypos; i-- {
			if job.osprings[i] == fine {
				trypos = i
				continue outer
			}
		}

		if upto < job.len && job.osprings[upto] == damaged {
			continue outer
		}

		if numidx == job.lennums-1 {
			if job.lastDmgIdx >= upto {
				trypos = job.lastDmgIdx - num
				continue
			}

			total++
			continue
		}

		total += job.day12aRecur(numidx+1, upto+1)
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
	return day12newjob(springs, nums).day12aRecur(0, 0)
}

func day12bRow(line string) int {
	springs, info, ok := strings.Cut(line, " ")
	if !ok {
		panic("invalid line: " + line)
	}
	springs = strings.Join([]string{springs, springs, springs, springs, springs}, "?")
	info = strings.Join([]string{info, info, info, info, info}, ",")

	nums := strsToNums(strings.Split(info, ","))

	return day12newjob(springs, nums).day12aRecur(0, 0)
}

func day12bTotal(lines []string) int {
	jobs := make(chan string)
	result := make(chan int)
	for i := 0; i < 200; i++ {
		go func() {
			for job := range jobs {
				result <- day12bRow(job)
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
		total += <-result
		fmt.Printf("Done %.01f %%\n", (float64(i)/float64(len(lines)))*100)
	}
	return total
}
