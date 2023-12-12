package main

import "testing"

func TestDay11a(t *testing.T) {
	expect := 10494813
	result := day11a()

	AssertEqual(t, expect, result)
}

func TestDay11b(t *testing.T) {
	expect := 840988812853
	result := day11b()

	AssertEqual(t, expect, result)
}

func TestDay11aExample(t *testing.T) {
	input := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	expect := 374
	result := day11aTotal(input)
	AssertEqual(t, expect, result)
}

func TestDay11bExample(t *testing.T) {
	input := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	expect := 1030
	result := day11bTotal(input, 10)
	AssertEqual(t, expect, result)

	expect = 8410
	result = day11bTotal(input, 100)
	AssertEqual(t, expect, result)
}
