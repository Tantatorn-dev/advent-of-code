package aoc2023

import (
	"strconv"
	"strings"
)

func Solve1(input string) int {
	lines := strings.Split(input, "\n")

	count := 0
	for _, line := range lines {
		count += countLine(line, func(i int) int {
			if i == 0 {
				i++
			} else {
				i *= 2
			}

			return i
		})
	}

	return count
}

func countLine(input string, accumulator func(int) int) int {
	input = strings.Split(input, ":")[1]

	inputStrArr := strings.Split(input, "|")

	winningStrings := strings.Split(inputStrArr[0], " ")

	var winningNums []int
	for _, numStr := range winningStrings {
		if numStr == "" || numStr == " " {
			continue
		}

		num, _ := strconv.Atoi(numStr)
		winningNums = append(winningNums, num)
	}

	myStrings := strings.Split(inputStrArr[1], " ")

	myNums := make([]int, len(myStrings))
	for _, numStr := range myStrings {
		if numStr == "" || numStr == " " {
			continue
		}

		num, _ := strconv.Atoi(numStr)
		myNums = append(myNums, num)
	}

	count := 0
	for _, myNum := range myNums {
		for _, winningNum := range winningNums {
			if myNum == winningNum {
				count = accumulator(count)
			}
		}
	}

	return count
}

func Solve2(input string) int {
	lines := strings.Split(input, "\n")

	matchedArr := make([]int, len(lines))
	counts := make([]int, len(lines))

	for i, line := range lines {
		matched := countLine(line, func(i int) int {
			i++
			return i
		})

		matchedArr[i] = matched
		counts[i] = 1
	}

	for i, c := range counts {
		amount := c

		if matchedArr[i] == 0 {
			continue
		}

		for j := 1; j <= matchedArr[i]; j++ {
			counts[i+j] += amount
		}
	}

	ret := 0

	for _, c := range counts {
		ret += c
	}

	return ret
}
