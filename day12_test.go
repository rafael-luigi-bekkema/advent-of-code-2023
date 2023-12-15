package main

import "testing"

func TestDay12a(t *testing.T) {
	expect := 7771 // too high: 8379
	result := day12a()
	AssertEqual(t, expect, result)
}

func TestDay12aRow(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{"???.### 1,1,3", 1},
		{".??..??...?##. 1,1,3", 4},
		{"?#?#?#?#?#?#?#? 1,3,1,6", 1},
		{"????.#...#... 4,1,1", 1},
		{"????.######..#####. 1,6,5", 4},
		{"?###???????? 3,2,1", 10},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day12aRow(tc.input)
			AssertEqual(t, tc.expect, result)
		})
	}
}

func TestDay12aTotal(t *testing.T) {
	input := []string{
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
	}
	expect := 21
	result := sumtotal(input, day12aRow)
	AssertEqual(t, expect, result)
}

// func TestDay12b(t *testing.T) {
// 	expect := 0
// 	result := day12b()
// 	AssertEqual(t, expect, result)
// }

func TestDay12bRow(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{"???.### 1,1,3", 1},
		{".??..??...?##. 1,1,3", 16384},
		{"?#?#?#?#?#?#?#? 1,3,1,6", 1},
		{"????.#...#... 4,1,1", 16},
		{"????.######..#####. 1,6,5", 2500},
		{"?###???????? 3,2,1", 506250},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day12bRow(tc.input)
			AssertEqual(t, tc.expect, result)
		})
	}
}

func TestDay12bTotal(t *testing.T) {
	input := []string{
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
	}
	expect := 525152
	result := day12bTotal(input, false)
	AssertEqual(t, expect, result)
}
