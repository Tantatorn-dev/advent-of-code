package aoc2023

import (
	"regexp"
	"strings"
)

// Problem: https://adventofcode.com/2023/day/8
func TraverseWasteland(input string) int {
	stepCount := 1

	// split text between empty lines
	strArr := strings.Split(input, "\n\n")

	directions := strArr[0]

	wastelandMap := constructWastelandMap(strArr[1])

	target := "ZZZ"

	currKey := "AAA"

out:
	for {
		for i := 0; i < len(directions); i++ {
			if directions[i] == 'L' {
				currKey = wastelandMap[currKey][0]
			} else if directions[i] == 'R' {
				currKey = wastelandMap[currKey][1]
			}

			if currKey == target {
				break out
			}

			stepCount++
		}
	}

	return stepCount
}

func constructWastelandMap(input string) map[string][]string {
	lines := strings.Split(input, "\n")

	ret := make(map[string][]string)

	// AAA = (BBB, CCC)
	re := regexp.MustCompile(`(.*) = \((.*), (.*)\)`)

	for _, line := range lines {
		if re.MatchString(line) {
			matches := re.FindStringSubmatch(line)

			ret[matches[1]] = []string{matches[2], matches[3]}
		}
	}

	return ret
}
