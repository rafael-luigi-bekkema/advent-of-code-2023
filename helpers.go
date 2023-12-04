package main

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
