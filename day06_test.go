package main

import "testing"

func TestDay6a(t *testing.T) {
	expect := 741000
	result := day6a()

	AssertEqual(t, expect, result)
}

func TestDay6aMultiplyWins(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	expect := 288
	result := day6aMultiplyWins(input)

	AssertEqual(t, expect, result)
}

func TestDay6b(t *testing.T) {
	expect := 38220708
	result := day6b()

	AssertEqual(t, expect, result)
}

func TestDay6bTotalWins(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	expect := 71503
	result := day6bTotalWins(input)

	AssertEqual(t, expect, result)
}
