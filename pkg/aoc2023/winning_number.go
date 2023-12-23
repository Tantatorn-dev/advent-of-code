package aoc2023

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve1(input string) int {
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

func Solve2(input string) int {
	lines := strings.Split(input, "\n")

	matchedArr := make([]int, len(lines))
	counts := make([]int, len(lines))

	for i, line := range lines {
		matched := countLine2(line)

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

	fmt.Println(len(counts))

	for _, c := range counts {
		ret += c
	}

	return ret
}

func countLine2(input string) int {
	input = strings.Split(input, ":")[1]

	inputStrArr := strings.Split(input, "|")

	winningStrings := strings.Split(inputStrArr[0], " ")

	var winningNums []int
	for _, numStr := range winningStrings {
		if numStr == "" || numStr == " " {
			continue
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

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
				count++
			}
		}
	}

	return count
}
