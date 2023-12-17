package main

import "testing"

func TestDay13a(t *testing.T) {
	expect := 33520
	result := day13a()

	AssertEqual(t, expect, result)
}

func TestDay13aPuzzle(t *testing.T) {
	inputs := []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
		"",
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}

	expect := 405
	result := day13aTotal(inputs)

	AssertEqual(t, expect, result)
}
