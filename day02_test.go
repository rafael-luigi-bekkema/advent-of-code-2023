package main

import "testing"

func TestDay2aPossible(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 1},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 2},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 0},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 0},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 5},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day2aPossible(tc.input)
			AssertEqual(t, tc.expect, result)
		})
	}
}

func TestDay2aWork(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	expect := 8
	result := sumtotal(input, day2aPossible)
	AssertEqual(t, expect, result)
}

func TestDay2a(t *testing.T) {
	expect := 3099
	result := day2a()
	AssertEqual(t, expect, result)
}

func TestDay2bPower(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 48},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 1560},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 630},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 36},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day2bPower(tc.input)
			AssertEqual(t, tc.expect, result)
		})
	}
}

func TestDay2bWork(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	expect := 2286
	result := sumtotal(input, day2bPower)
	AssertEqual(t, expect, result)
}

func TestDay2b(t *testing.T) {
	expect := 72970
	result := day2b()
	AssertEqual(t, expect, result)
}
