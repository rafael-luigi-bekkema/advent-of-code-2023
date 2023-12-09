package main

import "testing"

func TestDay9a(t *testing.T) {
	expect := 1939607039
	result := day9a()

	AssertEqual(t, expect, result)
}

func TestDay9b(t *testing.T) {
	expect := 1041
	result := day9b()

	AssertEqual(t, expect, result)
}

func TestDay9aExample(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	expect := 114
	result := sumtotal(input, day9aExtrapolate)

	AssertEqual(t, expect, result)
}

func TestDay9bExample(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	expect := 2
	result := sumtotal(input, day9bExtrapolate)

	AssertEqual(t, expect, result)
}
