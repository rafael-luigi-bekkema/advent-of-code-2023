package main

import (
	"fmt"
	"slices"
)

func day11a() int {
	return day11aTotal(mustReadInput(11))
}

func day11b() int {
	return day11bTotal(mustReadInput(11), 1_000_000)
}

func printMap(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println()
}

func day11aTotal(lines []string) int {
outery:
	for y := len(lines) - 1; y >= 0; y-- {
		line := lines[y]
		for _, chr := range []byte(line) {
			if chr == '#' {
				continue outery
			}
		}
		// row is empty, add an empty row
		lines = append(lines, "")
		copy(lines[y+1:], lines[y:])
		lines[y] = line
	}

outerx:
	for x := len(lines[0]) - 1; x >= 0; x-- {
		for y := range lines {
			if lines[y][x] == '#' {
				continue outerx
			}
		}
		// column is empty, resize all lines
		for sy := range lines {
			lines[sy] = lines[sy][:x] + "." + lines[sy][x:]
		}
	}

	var galaxies []point
	for y, line := range lines {
		for x, chr := range line {
			if chr != '#' {
				continue
			}
			galaxies = append(galaxies, point{y: y, x: x})
		}
	}

	var total int
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			total += abs(g1.y-g2.y) + abs(g1.x-g2.x)
		}
	}

	return total
}

func day11bTotal(lines []string, multiplier int) int {
	ys := make([]bool, len(lines))
	xs := make([]bool, len(lines[0]))

outery:
	for y := len(lines) - 1; y >= 0; y-- {
		line := lines[y]
		for _, chr := range []byte(line) {
			if chr == '#' {
				continue outery
			}
		}
		ys[y] = true
	}

outerx:
	for x := len(lines[0]) - 1; x >= 0; x-- {
		for y := range lines {
			if lines[y][x] == '#' {
				continue outerx
			}
		}
		xs[x] = true
	}

	var galaxies []point
	for y, line := range lines {
		for x, chr := range line {
			if chr != '#' {
				continue
			}
			galaxies = append(galaxies, point{y: y, x: x})
		}
	}

	var total int
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			sys := []int{g1.y, g2.y}
			slices.Sort(sys)
			sxs := []int{g1.x, g2.x}
			slices.Sort(sxs)

			var xtotal, ytotal int
			for x := sxs[0] + 1; x <= sxs[1]; x++ {
				if xs[x] {
					xtotal += multiplier
					continue
				}
				xtotal++
			}
			for y := sys[0] + 1; y <= sys[1]; y++ {
				if ys[y] {
					ytotal += multiplier
					continue
				}
				ytotal++
			}

			total += xtotal + ytotal
		}
	}

	return total
}
