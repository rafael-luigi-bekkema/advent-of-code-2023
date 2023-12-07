package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func day7a() int {
	return day7aWinnings(mustReadInput(7))
}

func day7b() int {
	return day7bWinnings(mustReadInput(7))
}

type card byte

const joker card = 'J'

var cardsa = []card{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
var cardsb = []card{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', joker}

func (c card) value(b bool) int {
	cards := cardsa
	if b {
		cards = cardsb
	}
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

func day7handValue(hand [5]card) int {
	counts := Count(hand[:])
	sort.Slice(counts, func(i, j int) bool {
		return counts[j].count < counts[i].count
	})

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

func day7compareHands(hand1, hand2 [5]card, b bool) bool {
	for i := range hand1 {
		if hand1[i] == hand2[i] {
			continue
		}
		return hand1[i].value(b) < hand2[i].value(b)
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
			return day7compareHands(hands[i].hand, hands[j].hand, false)
		}
		return hands[i].value < hands[j].value
	})

	var total int
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}

	return total
}

func day7handMaxValue(hand [5]card) int {
	if hand == [5]card{joker, joker, joker, joker, joker} {
		return fiveOfAKind
	}

	if !slices.Contains(hand[:], joker) {
		return day7handValue(hand)
	}

	counts := Count(hand[:])
	counts = slices.DeleteFunc(counts, func(cr countResult[card]) bool {
		return cr.value == joker
	})
	sort.Slice(counts, func(i, j int) bool {
		return counts[j].count < counts[i].count
	})

	// ohand := hand
	for i := range hand {
		if hand[i] == joker {
			hand[i] = counts[0].value
		}
	}

	// fmt.Println(string(hand[:]), string(ohand[:]))

	return day7handValue(hand)
}

func day7bWinnings(lines []string) int {
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
		h.value = day7handMaxValue(h.hand)

		hands[i] = h
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].value == hands[j].value {
			return day7compareHands(hands[i].hand, hands[j].hand, true)
		}
		return hands[i].value < hands[j].value
	})

	var total int
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}

	return total
}
