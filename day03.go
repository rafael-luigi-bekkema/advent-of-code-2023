package main

import (
	"strconv"
)

func day3a() int {
	return day3aSumParts(mustReadInput(3))
}

func day3isSymbol(lines []string, dy, dx int) bool {
	if dy < 0 || dx < 0 || dy >= len(lines) || dx >= len(lines[dy]) {
		return false
	}
	s := lines[dy][dx]
	return (s < '0' || s > '9') && s != '.'
}

func day3aSumParts(lines []string) int {
	var sum int

	for y, line := range lines {
		var number []byte
		var hasSymbol bool
		for x := 0; x < len(line); x++ {
			if line[x] < '0' || line[x] > '9' {
				if len(number) > 0 && hasSymbol {
					n, _ := strconv.Atoi(string(number))
					sum += n
				}
				number = number[:0]
				hasSymbol = false
				continue
			}

			number = append(number, line[x])

			for i := 0; i < 9; i++ {
				dx := i%3 - 1
				dy := i/3 - 1
				if dx == 0 && dy == 0 {
					continue
				}

				if day3isSymbol(lines, y+dy, x+dx) {
					hasSymbol = true
					break
				}
			}
		}

		if len(number) > 0 && hasSymbol {
			n, _ := strconv.Atoi(string(number))
			sum += n
		}
	}

	return sum
}

func day3b() int {
	return day3bSumGearRatios(mustReadInput(3))
}

func day3bExtractNumber(line string, x int) (num int, end int) {
	for ; x >= 0; x-- {
		if x == 0 || line[x-1] < '0' || line[x-1] > '9' {
			break
		}
	}

	x2 := x
	for ; x2 < len(line); x2++ {
		if x2 == len(line)-1 || line[x2+1] < '0' || line[x2+1] > '9' {
			break
		}
	}

	n, _ := strconv.Atoi(string(line[x : x2+1]))
	return n, x2
}

func day3bGearRatio(lines []string, y, x int) int {
	var numbers []int
	for i := 0; i < 9; i++ {
		dx := i%3 - 1
		dy := i/3 - 1
		if dx == 0 && dy == 0 {
			continue
		}

		y, x := dy+y, dx+x
		if y < 0 || y >= len(lines) || x < 0 || x >= len(lines[y]) || lines[y][x] < '0' || lines[y][x] > '9' {
			continue
		}

		num, endx := day3bExtractNumber(lines[y], x)

		// Don't count the same number twice
		if dx == 0 && endx >= x+1 || dx == -1 && endx == x+1 {
			i++
		} else if dx == -1 && endx >= x+2 {
			i += 2
		}

		numbers = append(numbers, num)
	}

	if len(numbers) == 2 {
		return numbers[0] * numbers[1]
	}
	return 0
}

func day3bSumGearRatios(lines []string) int {
	var sum int

	for y, line := range lines {
		for x, chr := range []byte(line) {
			if chr != '*' {
				continue
			}

			sum += day3bGearRatio(lines, y, x)
		}
	}

	return sum
}
