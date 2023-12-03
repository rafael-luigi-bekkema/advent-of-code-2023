package main

import "testing"

func TestDay3a(t *testing.T) {
	expect := 525911
	result := day3a()

	AssertEqual(t, expect, result)
}

func TestDay3aSumParts(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	expect := 4361
	result := day3aSumParts(input)

	AssertEqual(t, expect, result)
}

func TestDay3b(t *testing.T) {
	expect := 75805607
	result := day3b()

	AssertEqual(t, expect, result)
}

func TestDay3bSumGearRatios(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	expect := 467835

	result := day3bSumGearRatios(input)

	AssertEqual(t, expect, result)
}
