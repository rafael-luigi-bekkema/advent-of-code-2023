package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func day7a() int {
	return day7aWinnings(mustReadInput(7))
}

type card byte

var cards = []card{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

func (c card) value() int {
	for i, sc := range cards {
		if sc != c {
			continue
		}

		return len(cards) - i
	}
	panic(fmt.Sprintf("card %v unknown", c))
}

func (c card) String() string {
	return fmt.Sprintf("%c", c)
}

const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type hand struct {
	hand  [5]card
	bid   int
	value int
}

func printHands(hands []hand) {
	for _, hand := range hands {
		fmt.Printf("%s %d\n", hand.hand, hand.value)
	}
}

func day7handValue(hand [5]card) int {
	counts := Count(hand[:])

	switch {
	case counts[0].count == 5:
		return fiveOfAKind
	case counts[0].count == 4:
		return fourOfAKind
	case counts[0].count == 3 && counts[1].count == 2:
		return fullHouse
	case counts[0].count == 3:
		return threeOfAKind
	case counts[0].count == 2 && counts[1].count == 2:
		return twoPair
	case counts[0].count == 2:
		return onePair
	default:
		return highCard
	}
}

func day7compareHands(hand1, hand2 [5]card) bool {
	for i := range hand1 {
		if hand1[i] == hand2[i] {
			continue
		}
		return hand1[i].value() < hand2[i].value()
	}
	panic("hands are equal value")
}

func day7aWinnings(lines []string) int {
	hands := make([]hand, len(lines))

	for i, line := range lines {
		shand, sbid, _ := strings.Cut(line, " ")
		bid, err := strconv.Atoi(sbid)
		if err != nil {
			panic(err)
		}
		h := hand{
			bid: bid,
		}
		for i, chr := range []byte(shand) {
			h.hand[i] = card(chr)
		}
		h.value = day7handValue(h.hand)

		hands[i] = h
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].value == hands[j].value {
			return day7compareHands(hands[i].hand, hands[j].hand)
		}
		return hands[i].value < hands[j].value
	})

	var total int
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}

	return total
}
