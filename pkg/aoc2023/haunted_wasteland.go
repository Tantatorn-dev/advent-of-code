package aoc2023

import (
	"regexp"
	"strings"

	"github.com/Tantatorn-dev/advent-of-code/pkg/utils"
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

func TraverseWasteland2(input string) int {
	// split text between empty lines
	strArr := strings.Split(input, "\n\n")

	directions := strArr[0]

	wastelandMap := constructWastelandMap(strArr[1])

	currNodes := []string{}

	// find all keys that end with 'A'
	for k := range wastelandMap {
		// check if last char is 'A'
		if k[len(k)-1] == 'A' {
			currNodes = append(currNodes, k)
		}
	}

	stepCounts := make([]int, len(currNodes))

	// assign 1 to all stepCounts
	for i := range stepCounts {
		stepCounts[i] = 1
	}

	for i, node := range currNodes {
	out:
		for {
			for j := 0; j < len(directions); j++ {
				if directions[j] == 'L' {
					currNodes[i] = wastelandMap[currNodes[i]][0]
				} else if directions[j] == 'R' {
					currNodes[i] = wastelandMap[currNodes[i]][1]
				}

				if currNodes[i][len(node)-1] == 'Z' {
					break out
				}

				stepCounts[i] += 1
			}
		}
	}

	// find LCD of all stepCounts
	ret := utils.LCMSlice(stepCounts)

	return ret
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
