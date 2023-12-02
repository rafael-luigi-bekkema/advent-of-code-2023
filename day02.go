package main

import (
	"fmt"
	"strings"
)

func day2aPossible(line string) bool {
	const red, green, blue = 12, 13, 14

	_, rest, _ := strings.Cut(line, ": ")
	for _, set := range strings.Split(rest, "; ") {
		for _, item := range strings.Split(set, ", ") {
			var count int
			var color string
			fmt.Sscanf(item, "%d %s", &count, &color)

			if color == "red" && count > red || color == "green" && count > green || color == "blue" && count > blue {
				return false
			}
		}
	}

	return true
}

func day2aWork(lines []string) int {
	var total int
	for i, line := range lines {
		if day2aPossible(line) {
			total += i + 1
		}
	}
	return total
}

func day2a() int {
	return day2aWork(mustReadInput(2))
}

func day2bPower(line string) int {
	var minRed, minGreen, minBlue int

	_, rest, _ := strings.Cut(line, ": ")
	for _, set := range strings.Split(rest, "; ") {
		for _, item := range strings.Split(set, ", ") {
			var count int
			var color string
			fmt.Sscanf(item, "%d %s", &count, &color)

			switch {
			case color == "red" && count > minRed:
				minRed = count
			case color == "green" && count > minGreen:
				minGreen = count
			case color == "blue" && count > minBlue:
				minBlue = count
			}
		}
	}

	return minRed * minGreen * minBlue
}

func day2bWork(lines []string) int {
	var total int
	for _, line := range lines {
		total += day2bPower(line)
	}
	return total
}

func day2b() int {
	return day2bWork(mustReadInput(2))
}
