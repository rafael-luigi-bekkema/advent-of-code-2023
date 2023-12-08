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
