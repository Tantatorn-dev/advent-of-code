package aoc2023

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	RedThreshold   = 12
	GreenThreshold = 13
	BlueThreshold  = 14
)

// Problem: https://adventofcode.com/2023/day/2
type Set struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	id   int
	sets []Set
}

func (g Game) GetID() int {
	return g.id
}

func (g Game) IsValid() bool {
	for _, s := range g.sets {
		if RedThreshold < s.Red || BlueThreshold < s.Blue || GreenThreshold < s.Green {
			return false
		}
	}

	return true
}

func (g Game) MinSet() Set {
	var ret Set

	for _, s := range g.sets {
		if s.Blue > ret.Blue {
			ret.Blue = s.Blue
		}

		if s.Red > ret.Red {
			ret.Red = s.Red
		}

		if s.Green > ret.Green {
			ret.Green = s.Green
		}
	}

	return ret
}

func ReadLine(input string) Game {
	var g Game

	inputArr := strings.Split(input, ": ")

	label, content := inputArr[0], inputArr[1]

	// retrieve game id
	reID := regexp.MustCompile(`Game (\d+)`)

	match := reID.FindStringSubmatch(label)

	if len(match) > 1 {
		g.id, _ = strconv.Atoi(match[1])
	} else {
		panic("cannot find game id")
	}

	// parse game sets
	setStrings := strings.Split(content, ";")

	for _, setString := range setStrings {
		var set Set

		reCubeAmount := regexp.MustCompile(`(\d+) red`)

		match = reCubeAmount.FindStringSubmatch(setString)

		if len(match) > 1 {
			set.Red, _ = strconv.Atoi(match[1])
		}

		reCubeAmount = regexp.MustCompile(`(\d+) green`)

		match = reCubeAmount.FindStringSubmatch(setString)

		if len(match) > 1 {
			set.Green, _ = strconv.Atoi(match[1])
		}

		reCubeAmount = regexp.MustCompile(`(\d+) blue`)

		match = reCubeAmount.FindStringSubmatch(setString)

		if len(match) > 1 {
			set.Blue, _ = strconv.Atoi(match[1])
		}

		g.sets = append(g.sets, set)
	}

	return g
}

func SolveCubeConundrum(input string) int {
	lines := strings.Split(input, "\n")

	ret := 0

	for _, line := range lines {
		g := ReadLine(line)

		if g.IsValid() {
			ret += g.GetID()
		}
	}

	return ret
}

func SolveCubeConundrum2(input string) int {
	lines := strings.Split(input, "\n")

	ret := 0

	for _, line := range lines {
		g := ReadLine(line)

		min := g.MinSet()

		ret += min.Red * min.Green * min.Blue
	}

	return ret
}
