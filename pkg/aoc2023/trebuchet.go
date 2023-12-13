package aoc2023

import (
	"fmt"
	"strconv"
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

func isNumber(ch int32) bool {
	return ch >= '0' && ch <= '9'
}
