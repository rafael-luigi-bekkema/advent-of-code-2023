package main

import (
	"fmt"
	"regexp"
)

func day8a() int {
	return day8aSteps(mustReadInput(8))
}

func day8b() int {
	return day8bSteps(mustReadInput(8))
}

type day8instruction struct{ left, right string }

var expr = regexp.MustCompile(`[A-Z0-9]{3}`)

func parseDay8(lines []string) map[string]day8instruction {
	ins := map[string]day8instruction{}

	for _, line := range lines[2:] {
		nums := expr.FindAllString(line, 3)

		ins[nums[0]] = day8instruction{left: nums[1], right: nums[2]}
	}

	return ins
}

func day8aSteps(lines []string) int {
	directions := lines[0]
	ins := parseDay8(lines)

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

func day8bSteps(lines []string) int {
	directions := lines[0]
	ins := parseDay8(lines)

	var positions []string
	for from := range ins {
		if from[2] == 'A' {
			positions = append(positions, from)
		}
	}

	starts := make([]string, len(positions))
	copy(starts, positions)

	var steps int
	firstzs := map[string]int{}
	secondzs := map[string]int{}

outer:
	for {
		for _, dir := range directions {
			steps += 1

			for i, current := range positions {
				in, ok := ins[current]
				if !ok {
					panic("unknown location: " + current)
				}

				switch dir {
				case 'L':
					positions[i] = in.left
				case 'R':
					positions[i] = in.right
				default:
					panic("unknown direction: " + string(dir))
				}

				if positions[i][2] != 'Z' {
					continue
				}

				firstz, ok := firstzs[starts[i]]
				_, ok2 := secondzs[starts[i]]
				if ok && !ok2 {
					secondzs[starts[i]] = steps - firstz
				}
				if !ok {
					firstzs[starts[i]] = steps
				}
			}

			if len(secondzs) == len(starts) {
				break outer
			}
		}
	}

	var sstep int
	var found bool
	for j := 0; j < 1000_000_000; j++ {
		sstep = firstzs[starts[0]] + j*secondzs[starts[0]]

		allmatches := true
		for _, start := range starts[1:] {
			if (sstep-firstzs[start])%secondzs[start] != 0 {
				allmatches = false
				break
			}
		}
		if allmatches {
			fmt.Println(sstep, j)
			found = true
			break
		}
	}

	if !found {
		panic("not found")
	}

	return sstep
}
