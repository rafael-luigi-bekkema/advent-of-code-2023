package main

import "strings"

func day15a() int {
	line := mustReadInput(15)[0]
	return day15aTotal(line)
}

func day15aHash(input string) int {
	var curr int

	for _, chr := range []byte(input) {
		curr += int(chr)
		curr *= 17
		curr %= 256
	}

	return curr
}

func day15aTotal(input string) int {
	var total int

	for _, line := range strings.Split(input, ",") {
		total += day15aHash(line)
	}

	return total
}
