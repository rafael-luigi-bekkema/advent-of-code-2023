package main

import "fmt"

func main() {
	lines := mustReadInput(12)

	for i, line := range lines {
		fmt.Println("line", i, line)
		fmt.Println(day12bRow(line))
	}
}
