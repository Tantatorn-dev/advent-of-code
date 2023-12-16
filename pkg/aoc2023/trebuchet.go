package aoc2023

import (
	"fmt"
	"strconv"
	"strings"
)

// Problem 1: https://adventofcode.com/2023/day/1
func CalibrateValue(input string) int {
	ret := 0

	numStr := ""

	// append \n to input
	input += "\n"

	for _, ch := range input {
		if isNumber(ch) {
			numStr += string(ch)
		}

		if ch == '\n' {
			if len(numStr) == 1 {
				numStr += numStr
			} else {
				numStr = fmt.Sprintf("%c%c", numStr[0], numStr[len(numStr)-1])
			}

			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}

			ret += num
			numStr = ""
		}
	}

	return ret
}

func CalibrateValue2(input string) int {
	digitMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	ret := 0

	lines := strings.Split(input, "\n")

	for _, line := range lines {

		minT := len(line)
		minD := 0

		for text, digit := range digitMap {
			i := strings.Index(line, text)
			if i >= 0 && i < minT {
				minT = i
				minD = digit
			}
		}

		ret += minD * 10

		maxT := -1
		maxD := 0

		for text, digit := range digitMap {
			i := strings.LastIndex(line, text)
			if i >= 0 && i > maxT {
				maxT = i
				maxD = digit
			}
		}

		ret += maxD
	}

	return ret
}

func isNumber(ch int32) bool {
	return ch >= '0' && ch <= '9'
}
