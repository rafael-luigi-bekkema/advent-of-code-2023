package main

// sumtotal runs f on each line and returns the sum total
func sumtotal(lines []string, f func(string) int) int {
	var total int
	for _, line := range lines {
		total += f(line)
	}
	return total
}
