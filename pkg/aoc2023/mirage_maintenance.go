package aoc2023

import (
	"strconv"
	"strings"
)

// Problem: https://adventofcode.com/2023/day/9
func SumExtrapolatedValues(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0

	for _, line := range lines {
		sum += findExtrapolatedValue(line)
	}

	return sum
}

func findExtrapolatedValue(input string) int {
	strArr := strings.Split(input, " ")

	nums := make([]int, len(strArr))
	for i, str := range strArr {
		num, _ := strconv.Atoi(str)

		nums[i] = num
	}

	histories := constructHistories(nums)

	// append 0 to last elem
	histories[len(histories)-1] = append(histories[len(histories)-1], 0)

	for i := len(histories) - 2; i >= 0; i-- {
		current := histories[i]

		prev := histories[i+1][len(histories[i+1])-1]

		histories[i] = append(current, current[len(current)-1]+prev)
	}

	return histories[0][len(histories[0])-1]
}

func constructHistories(nums []int) [][]int {
	// find diff until zero
	histories := make([][]int, 0)

	histories = append(histories, nums)

	for {
		current := histories[len(histories)-1]
		if isAllZero(current) {
			break
		}

		histories = append(histories, diffArr(current))
	}

	return histories
}

func diffArr(arr []int) []int {
	diff := make([]int, len(arr)-1)

	for i := 0; i < len(arr)-1; i++ {
		diff[i] = arr[i+1] - arr[i]
	}

	return diff
}

func isAllZero(arr []int) bool {
	for _, num := range arr {
		if num != 0 {
			return false
		}
	}

	return true
}
