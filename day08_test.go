package main

import "testing"

func TestDay8a(t *testing.T) {
	expect := 16271
	result := day8a()

	AssertEqual(t, expect, result)
}

func TestDay8aExample(t *testing.T) {
	tt := []struct {
		input  []string
		expect int
	}{
		{input: []string{
			"RL",
			"",
			"AAA = (BBB, CCC)",
			"BBB = (DDD, EEE)",
			"CCC = (ZZZ, GGG)",
			"DDD = (DDD, DDD)",
			"EEE = (EEE, EEE)",
			"GGG = (GGG, GGG)",
			"ZZZ = (ZZZ, ZZZ)",
		}, expect: 2},
		{input: []string{
			"LLR",
			"",
			"AAA = (BBB, BBB)",
			"BBB = (AAA, ZZZ)",
			"ZZZ = (ZZZ, ZZZ)",
		}, expect: 6},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day8aSteps(tc.input)
			AssertEqual(t, tc.expect, result)
		})
	}
}

// func TestDay8b(t *testing.T) {
// 	expect := 14265111103729 // too low = 2179700000
// 	result := day8b()

// 	AssertEqual(t, expect, result)
// }

func TestDay8bExample(t *testing.T) {
	input := []string{
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}
	expect := 6
	result := day8bSteps(input)
	AssertEqual(t, expect, result)
}
