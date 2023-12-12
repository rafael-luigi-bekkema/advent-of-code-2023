package main

import "testing"

func TestDay10(t *testing.T) {
	expect := 6903
	result := day10a()

	AssertEqual(t, expect, result)
}

func TestDay10aExample(t *testing.T) {
	tt := []struct {
		input  []string
		expect int
	}{
		{input: []string{
			".....",
			".S-7.",
			".|.|.",
			".L-J.",
			".....",
		}, expect: 4},
		{input: []string{
			"..F7.",
			".FJ|.",
			"SJ.L7",
			"|F--J",
			"LJ...",
		}, expect: 8},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day10aLoop(tc.input)
			AssertEqual(t, tc.expect, result)
		})
	}
}
