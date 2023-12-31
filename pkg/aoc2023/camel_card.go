package aoc2023

import (
	"sort"
	"strconv"
	"strings"
)

// problem 7: https://adventofcode.com/2023/day/7

func GetTotalWinning(input string) int {
	input = strings.ReplaceAll(input, "\t", "")

	lines := strings.Split(input, "\n")

	var cards []Card

	for _, line := range lines {
		var c Card

		handStr := strings.Split(line, " ")[0]
		c.CardHand = determineCardHand(handStr)
		c.Str = handStr
		val, err := strconv.ParseInt(strings.Split(line, " ")[1], 10, 64)
		if err != nil {
			panic(err)
		}

		c.Value = int(val)

		cards = append(cards, c)
	}

	sort.Slice(cards, func(i, j int) bool {
		if cards[i].CardHand == cards[j].CardHand {
			return strings.Compare(cards[i].Str, cards[j].Str) < 0
		}

		return cards[i].CardHand < cards[j].CardHand
	})

	ret := 0

	for i, c := range cards {
		ret += c.Value * (len(cards) - i)
	}

	return ret
}

type Card struct {
	CardHand CardHand
	Str      string
	Value    int
}

type CardHand int

const (
	FiveOfAKind CardHand = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

func determineCardHand(input string) CardHand {
	cardMap := make(map[string]int)

	for _, c := range input {
		cardMap[string(c)]++
	}

	currentState := HighCard

	for _, v := range cardMap {
		if v == 5 {
			return FiveOfAKind
		} else if v == 4 {
			return FourOfAKind
		} else if currentState == ThreeOfAKind && v == 2 {
			return FullHouse
		} else if v == 3 {
			currentState = ThreeOfAKind
		} else if v == 2 && currentState == OnePair {
			return TwoPair
		} else if v == 2 {
			currentState = OnePair
		}
	}

	return currentState
}
