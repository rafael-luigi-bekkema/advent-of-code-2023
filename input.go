package main

import (
	"bufio"
	"fmt"
	"os"
)

func mustReadInput(day uint) []string {
	lines, err := readInput(day)
	if err != nil {
		panic(err)
	}
	return lines
}

func readInput(day uint) ([]string, error) {
	path := fmt.Sprintf("input/day%02d.txt", day)

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}
