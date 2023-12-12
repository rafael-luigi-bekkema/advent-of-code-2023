package main

import "fmt"

func day10a() int {
	return day10aLoop(mustReadInput(10))
}

type point struct {
	y, x int
}

func (p point) String() string {
	return fmt.Sprintf("y=%d, x=%d", p.y, p.x)
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

const (
	vertical       byte = '|'
	horizontal     byte = '-'
	north_and_east byte = 'L'
	north_and_west byte = 'J'
	south_and_west byte = '7'
	south_and_east byte = 'F'
	ground         byte = '.'
	start_chr      byte = 'S'
)

func day10connects(lines []string, from, to point) (result bool) {
	if from.y < 0 || from.y >= len(lines) || from.x < 0 || from.x >= len(lines[from.y]) ||
		to.y < 0 || to.y >= len(lines) || to.x < 0 || to.x >= len(lines[to.y]) {
		return false
	}
	fromchr := lines[from.y][from.x]
	tochr := lines[to.y][to.x]

	switch {
	case to.y < from.y: // up
		return (fromchr == vertical || fromchr == north_and_east || fromchr == north_and_west || fromchr == start_chr) &&
			(tochr == vertical || tochr == south_and_west || tochr == south_and_east)
	case to.y > from.y: // down
		return (fromchr == vertical || fromchr == south_and_west || fromchr == south_and_east || fromchr == start_chr) &&
			(tochr == vertical || tochr == north_and_west || tochr == north_and_east)
	case to.x > from.x: // right
		return (fromchr == horizontal || fromchr == north_and_east || fromchr == south_and_east || fromchr == start_chr) &&
			(tochr == horizontal || tochr == north_and_west || tochr == south_and_west)
	case to.x < from.x: // left
		return (fromchr == horizontal || fromchr == north_and_west || fromchr == south_and_west || fromchr == start_chr) &&
			(tochr == horizontal || tochr == north_and_east || tochr == south_and_east)
	default:
		panic("unknown direction")
	}
}

func day10follow(lines []string, start, from, p point, path []point) int {
	for _, dir := range []direction{up, down, left, right} {
		var to point
		switch dir {
		case up:
			to = point{y: p.y - 1, x: p.x}
		case down:
			to = point{y: p.y + 1, x: p.x}
		case left:
			to = point{y: p.y, x: p.x - 1}
		case right:
			to = point{y: p.y, x: p.x + 1}
		}
		if to == from {
			continue
		}
		if to == start {
			// did it
			return len(path)/2 + 1
		}
		if day10connects(lines, p, to) {
			if n := day10follow(lines, start, p, to, append(path, from)); n != 0 {
				return n
			}
		}
	}

	return 0
}

func day10aLoop(lines []string) int {
	var start point

outer:
	for y, line := range lines {
		for x, chr := range []byte(line) {
			if chr == start_chr {
				start = point{y: y, x: x}
				break outer
			}
		}
	}

	result := day10follow(lines, start, point{-1, -1}, start, nil)

	return result
}
