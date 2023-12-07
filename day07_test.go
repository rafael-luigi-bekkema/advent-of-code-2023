package main

import "testing"

func TestDay7aWinnings(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	expect := 6440
	result := day7aWinnings(input)

	AssertEqual(t, expect, result)
}

func TestDay7a(t *testing.T) {
	expect := 246912307
	result := day7a()

	AssertEqual(t, expect, result)
}

func TestDay7bWinnings(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	expect := 5905
	result := day7bWinnings(input)

	AssertEqual(t, expect, result)
}

func TestDay7b(t *testing.T) {
	expect := 246894760
	result := day7b()

	AssertEqual(t, expect, result)
}
