package aoc2023

import (
	"strconv"
	"strings"
)

func CountWinningNumber(input string) int {
	lines := strings.Split(input, "\n")

	count := 0
	for _, line := range lines {
		count += countLine(line)
	}

	return count
}

func countLine(input string) int {
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
				if count == 0 {
					count++
				} else {
					count *= 2
				}
			}
		}
	}

	return count
}
