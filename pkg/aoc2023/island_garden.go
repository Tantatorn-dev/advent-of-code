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

	var maps []map[int]int

	for i, start := range mapStarts {
		if i == len(mapStarts)-1 {
			maps = append(maps, getMap(strings.Join(lines[start+1:], "\n")))
		} else {
			maps = append(maps, getMap(strings.Join(lines[start+1:mapStarts[i+1]-1], "\n")))
		}
	}

	// map seed to location(final map)
	for _, m := range maps {
		for i, seed := range seeds {
			if _, ok := m[seed]; !ok {
				continue
			}

			seeds[i] = m[seed]
		}
	}

	return slices.Min(seeds)
}

func getMap(input string) map[int]int {
	lines := strings.Split(input, "\n")

	ret := make(map[int]int)

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

		for i := 0; i < rangeValue; i++ {
			ret[source+i] = dest + i
		}
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
