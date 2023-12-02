package main

import "testing"

func TestDay1a(t *testing.T) {
	expect := 54390
	result := day1a()
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func TestDay1aWork(t *testing.T) {
	expect := 142
	result := sumtotal([]string{"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet"}, day1aExtractDigits)

	AssertEqual(t, expect, result)
}

func TestDay1aExtractDigits(t *testing.T) {
	tt := []struct {
		Input  string
		Expect int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day1aExtractDigits(tc.Input)
			AssertEqual(t, tc.Expect, result)
		})
	}
}

func TestDay1b(t *testing.T) {
	expect := 54277
	result := day1b()
	AssertEqual(t, expect, result)
}

func TestDay1bWork(t *testing.T) {
	expect := 281
	result := sumtotal([]string{"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen"}, day1bExtractDigits)

	AssertEqual(t, expect, result)
}

func TestDay1bExtractDigits(t *testing.T) {
	tt := []struct {
		Input  string
		Expect int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day1bExtractDigits(tc.Input)
			AssertEqual(t, tc.Expect, result)
		})
	}
}
