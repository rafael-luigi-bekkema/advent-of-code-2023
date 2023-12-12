package main

import (
	"cmp"
	"strconv"
	"strings"
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

func min2[T cmp.Ordered](values ...T) T {
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

func lineToNums(line string) []int {
	fields := strings.Fields(line)
	nums := make([]int, len(fields))
	for i, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	return nums
}

func strsToNums(items []string) []int {
	nums := make([]int, len(items))
	for i, item := range items {
		num, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}

	return nums
}

type countResult[T comparable] struct {
	value T
	count int
}

func Count[T comparable](slice []T) []countResult[T] {
	result := map[T]int{}
	for _, item := range slice {
		result[item]++
	}
	results := make([]countResult[T], 0, len(result))
	for value, count := range result {
		results = append(results, countResult[T]{value, count})
	}
	return results
}

func Same[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, item := range slice1 {
		if item != slice2[i] {
			return false
		}
	}

	return true
}

func Values[T comparable, U any](m map[T]U) []U {
	values := make([]U, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

func abs[T ~int](value T) T {
	if value < 0 {
		return value * -1
	}
	return value
}

func sum(values []int) int {
	var result int
	for _, value := range values {
		result += value
	}
	return result
}
