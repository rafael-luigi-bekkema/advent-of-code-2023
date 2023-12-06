package main

import (
	"math"
	"strconv"
	"strings"
)

func day4aCardPoints(line string) int {
	_, rest, _ := strings.Cut(line, ": ")

	swinning, shave, _ := strings.Cut(rest, " | ")
	winning := intersect(lineToNums(swinning), lineToNums(shave))

	return int(math.Pow(2, float64(len(winning)-1)))
}

func day4a() int {
	return sumtotal(mustReadInput(4), day4aCardPoints)
}

func totalScratchCards(lines []string) int {
	cardsMap := map[int]int{}

	for _, line := range lines {
		game, rest, _ := strings.Cut(line, ": ")

		parts := strings.Fields(game)
		gameNo, _ := strconv.Atoi(parts[1])

		swinning, shave, _ := strings.Cut(rest, " | ")
		winning := intersect(lineToNums(swinning), lineToNums(shave))

		cardsMap[gameNo] = len(winning)
	}

	type card struct{ num, wins int }
	cards := make([]card, 0, len(cardsMap))
	for no, wins := range cardsMap {
		cards = append(cards, card{no, wins})
	}

	maxCap := 0
	for i := 0; i < len(cards); i++ {
		crd := cards[i]
		if crd.wins == 0 {
			continue
		}
		for j := crd.num + 1; j <= crd.num+crd.wins; j++ {
			wins := cardsMap[j]
			cards = append(cards, card{j, wins})
			if newcap := cap(cards); newcap > maxCap {
				maxCap = newcap
			}
		}
	}

	return len(cards)
}

func day4b() int {
	return totalScratchCards(mustReadInput(4))
}
