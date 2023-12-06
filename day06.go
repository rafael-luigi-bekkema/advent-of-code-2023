package main

import (
	"strconv"
	"strings"
)

func day6a() int {
	return day6aMultiplyWins(mustReadInput(6))
}

func day6aMultiplyWins(line []string) int {
	times := strsToNums(strings.Fields(line[0])[1:])
	distances := strsToNums(strings.Fields(line[1])[1:])

	var mul int

	for i, time := range times {
		distance := distances[i]
		wins := day6aWins(time, distance)
		if i == 0 {
			mul = wins
		} else {
			mul *= wins
		}
	}

	return mul
}

func day6aWins(time, minDistance int) int {
	wins := 0

	for i := 0; i <= time; i++ {
		remain := time - i
		speed := i

		distance := remain * speed
		if distance > minDistance {
			wins++
		}
	}

	return wins
}

func day6b() int {
	return day6bTotalWins(mustReadInput(6))
}

func day6bTotalWins(line []string) int {
	time, _ := strconv.Atoi(strings.Join(strings.Fields(line[0])[1:], ""))
	distance, _ := strconv.Atoi(strings.Join(strings.Fields(line[1])[1:], ""))

	return day6aWins(time, distance)
}
