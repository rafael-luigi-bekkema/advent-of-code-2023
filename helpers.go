package main

import (
	"cmp"
)

// sumtotal runs f on each line and returns the sum total
func sumtotal(lines []string, f func(string) int) int {
	var total int
	for _, line := range lines {
		total += f(line)
	}
	return total
}

func intersect[T comparable](slice1, slice2 []T) []T {
	var result []T
	for _, item1 := range slice1 {
		for _, item2 := range slice2 {
			if item2 == item1 {
				result = append(result, item1)
				break
			}
		}
	}
	return result
}

func min[T cmp.Ordered](values ...T) T {
	if len(values) == 0 {
		panic("min on zero length slice")
	}

	var min T
	first := true
	for _, value := range values {
		if !first && value >= min {
			continue
		}
		min = value
		first = false
	}

	return min
}
