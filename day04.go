package main

import (
	"math"
	"strconv"
	"strings"
)

func toNums(line string) []int {
	fields := strings.Fields(line)
	nums := make([]int, len(fields))
	for i, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	return nums
}

func day4aCardPoints(line string) int {
	_, rest, _ := strings.Cut(line, ": ")

	swinning, shave, _ := strings.Cut(rest, " | ")

	winning := intersect(toNums(swinning), toNums(shave))

	return int(math.Pow(2, float64(len(winning)-1)))
}

func day4a() int {
	return sumtotal(mustReadInput(4), day4aCardPoints)
}
