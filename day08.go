package main

import (
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

	current := "AAA"
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
	insructions := parseDay8(lines)

	type ghost struct {
		location  string
		loopStart int
		loopSize  int
	}

	var ghosts []ghost
	for from := range insructions {
		if from[2] == 'A' {
			ghosts = append(ghosts, ghost{
				location:  from,
				loopStart: -1,
				loopSize:  -1,
			})
		}
	}

	var step int
	var loopsFound int

outer:
	for {
		for _, dir := range directions {
			step += 1

			for i, ghost := range ghosts {
				in, ok := insructions[ghost.location]
				if !ok {
					panic("unknown location: " + ghost.location)
				}

				switch dir {
				case 'L':
					ghosts[i].location = in.left
				case 'R':
					ghosts[i].location = in.right
				default:
					panic("unknown direction: " + string(dir))
				}

				if ghosts[i].location[2] != 'Z' {
					continue
				}

				if ghosts[i].loopStart != -1 && ghosts[i].loopSize == -1 {
					ghosts[i].loopSize = step - ghosts[i].loopStart
					loopsFound++
				}
				if ghosts[i].loopStart == -1 {
					ghosts[i].loopStart = step
				}
			}

			if loopsFound == len(ghosts) {
				break outer
			}
		}
	}

	var found bool
	for j := 0; true; j++ {
		step = ghosts[0].loopStart + j*ghosts[0].loopSize

		allmatches := true
		for _, ghost := range ghosts[1:] {
			if (step-ghost.loopStart)%ghost.loopSize != 0 {
				allmatches = false
				break
			}
		}
		if allmatches {
			found = true
			break
		}
	}

	if !found {
		panic("solution not found")
	}

	return step
}
