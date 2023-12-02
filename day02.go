package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2aPossible(line string) int {
	const red, green, blue = 12, 13, 14

	game, rest, _ := strings.Cut(line, ": ")
	for _, set := range strings.Split(rest, "; ") {
		for _, item := range strings.Split(set, ", ") {
			var count int
			var color string
			fmt.Sscanf(item, "%d %s", &count, &color)

			if color == "red" && count > red || color == "green" && count > green || color == "blue" && count > blue {
				return 0
			}
		}
	}

	_, num, _ := strings.Cut(game, " ")
	inum, _ := strconv.Atoi(num)

	return inum
}

func day2a() int {
	return sumtotal(mustReadInput(2), day2aPossible)
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

func day2b() int {
	return sumtotal(mustReadInput(2), day2bPower)
}
