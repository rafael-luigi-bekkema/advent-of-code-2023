package main

import "testing"

func TestDay5a(t *testing.T) {
	expect := 84470622 // too high -> 4099273491
	result := day5a()

	AssertEqual(t, expect, result)
}

func TestDay5aMinLocation(t *testing.T) {
	input := mustReadExampleInput(5)
	expect := 35
	result := day5aMinLocation(input)

	AssertEqual(t, expect, result)
}

// func TestDay5b(t *testing.T) {
// 	expect := 26714516
// 	result := day5b()

// 	AssertEqual(t, expect, result)
// }

func TestDay5bMinLocation(t *testing.T) {
	input := mustReadExampleInput(5)
	expect := 46
	result := day5bMinLocation(input)

	AssertEqual(t, expect, result)
}

func TestDay5Min(t *testing.T) {
	tt := []struct {
		input  []int
		expect int
	}{
		{[]int{23, 5, 2313, 3, 1245}, 3},
		{[]int{1, 2, 3, 4}, 1},
		{[]int{5, 3, 4, 2134, 1, -5}, -5},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := min(tc.input...)
			AssertEqual(t, tc.expect, result)
		})
	}
}
