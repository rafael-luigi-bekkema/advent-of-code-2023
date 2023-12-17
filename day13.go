package main

func day13a() int {
	return day13aTotal(mustReadInput(13))
}

func day13aCheckVertical(puzzle []string, x int) int {
	for px := x + 1; px < len(puzzle[0]); px++ {
		nx := x - (px - x - 1)
		if nx < 0 {
			break
		}

		for y := 0; y < len(puzzle); y++ {
			if puzzle[y][px] != puzzle[y][nx] {
				return 0
			}
		}
	}
	return x + 1
}

func day13aCheckHorizontal(puzzle []string, y int) int {
	for py := y + 1; py < len(puzzle); py++ {
		ny := y - (py - y - 1)
		if ny < 0 {
			break
		}

		for x := 0; x < len(puzzle[0]); x++ {
			if puzzle[py][x] != puzzle[ny][x] {
				return 0
			}
		}
	}
	return 100 * (y + 1)
}

func day13aValue(num int, puzzle []string) int {
	var total int
outery:
	for y := 0; y < len(puzzle)-1; y++ {
		for x := 0; x < len(puzzle[0]); x++ {
			if puzzle[y][x] != puzzle[y+1][x] {
				continue outery
			}
		}

		total += day13aCheckHorizontal(puzzle, y)
	}

outerx:
	for x := 0; x < len(puzzle[0])-1; x++ {
		for y := 0; y < len(puzzle); y++ {
			if puzzle[y][x] != puzzle[y][x+1] {
				continue outerx
			}
		}

		total += day13aCheckVertical(puzzle, x)
	}

	return total
}

func day13aTotal(lines []string) int {
	var puzzle []string
	var total int

	num := 1
	for _, line := range lines {
		if line == "" {
			total += day13aValue(num, puzzle)
			puzzle = nil
			num++
			continue
		}

		puzzle = append(puzzle, line)
	}

	total += day13aValue(num, puzzle)

	return total
}
