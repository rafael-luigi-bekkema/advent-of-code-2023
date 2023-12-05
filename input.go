package main

import (
	"bufio"
	"fmt"
	"os"
)

// mustReadInput returns puzzle input for the given day
func mustReadInput(day uint) []string {
	lines, err := readInput(day, Final)
	if err != nil {
		panic(err)
	}
	return lines
}

func mustReadExampleInput(day uint) []string {
	lines, err := readInput(day, Example)
	if err != nil {
		panic(err)
	}
	return lines
}

const (
	Example = iota
	Final
)

func readInput(day uint, typ int) ([]string, error) {
	var path string
	if typ == Example {
		path = fmt.Sprintf("input/day%02d_example.txt", day)
	} else {
		path = fmt.Sprintf("input/day%02d.txt", day)
	}

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
