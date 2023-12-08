package main

import (
	"regexp"
)

func day8a() int {
	return day8aSteps(mustReadInput(8))
}

func day8aSteps(lines []string) int {
	directions := lines[0]

	type inS struct{ left, right string }
	ins := map[string]inS{}

	expr := regexp.MustCompile(`[A-Z]{3}`)

	for _, line := range lines[2:] {
		nums := expr.FindAllString(line, 3)

		ins[nums[0]] = inS{left: nums[1], right: nums[2]}
	}

	var current = "AAA"
	var steps int

outer:
	for {
		for _, dir := range directions {
			in, ok := ins[current]
			if !ok {
				panic("unknown location: " + current)
			}

			steps += 1

			switch dir {
			case 'L':
				current = in.left
			case 'R':
				current = in.right
			default:
				panic("unknown direction: " + string(dir))
			}

			if current == "ZZZ" {
				break outer
			}
		}
	}

	return steps
}
