package aoc2023

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// Problem: https://adventofcode.com/2023/day/5

func GetLowestLocation(input string) int {
	// cleanup
	input = strings.ReplaceAll(input, "\t", "")

	lines := strings.Split(input, "\n")

	seeds := getSeeds(lines[0])

	mapRe := regexp.MustCompile("(.*) map:")

	var mapStarts []int

	for i, line := range lines {
		if mapRe.MatchString(line) {
			mapStarts = append(mapStarts, i)
		}
	}

	var locationMaps [][]Map

	for i, start := range mapStarts {
		if i == len(mapStarts)-1 {
			locationMaps = append(locationMaps, getMap(strings.Join(lines[start+1:], "\n")))
		} else {
			locationMaps = append(locationMaps, getMap(strings.Join(lines[start+1:mapStarts[i+1]-1], "\n")))
		}
	}

	// map seed to location(final map)
	for _, maps := range locationMaps {
		for i, seed := range seeds {

			for _, m := range maps {
				if seed >= m.Source && seed <= m.Source+m.RangeValue {
					seeds[i] = m.Dest + (seed - m.Source)
				}
			}

		}
	}

	return slices.Min(seeds)
}

type Map struct {
	Source     int
	Dest       int
	RangeValue int
}

func getMap(input string) []Map {
	lines := strings.Split(input, "\n")

	var ret []Map

	for _, line := range lines {
		numStr := strings.Split(line, " ")

		source, err := strconv.Atoi(numStr[1])
		if err != nil {
			panic(err)
		}

		dest, err := strconv.Atoi(numStr[0])
		if err != nil {
			panic(err)
		}

		rangeValue, err := strconv.Atoi(numStr[2])
		if err != nil {
			panic(err)
		}

		m := Map{
			Source:     source,
			Dest:       dest,
			RangeValue: rangeValue,
		}

		ret = append(ret, m)
	}

	return ret
}

func getSeeds(input string) []int {
	reNum := regexp.MustCompile(`(\d+)`)

	numStrings := reNum.FindAllString(input, -1)

	nums := make([]int, len(numStrings))
	for i, numString := range numStrings {
		nums[i], _ = strconv.Atoi(numString)
	}

	return nums
}
