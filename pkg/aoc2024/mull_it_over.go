package aoc2024

import (
	"regexp"
	"strconv"
	"strings"
)

func MultiplyByInstruction(instruction string) int {
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	strArr := r.FindAllString(instruction, -1)

	var ret int

	for _, str := range strArr {
		str = str[4 : len(str)-1]

		r := 1

		numStrArr := strings.Split(str, ",")
		for _, numStr := range numStrArr {
			num, _ := strconv.Atoi(numStr)
			r *= num
		}

		ret += r
	}

	return ret
}

func MultiplyByInstruction2(instruction string) int {
	r := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)
	strArr := r.FindAllString(instruction, -1)

	var ret int

	isSkip := false

	for _, str := range strArr {
		if str == "don't()" {
			isSkip = true
		} else if str == "do()" {
			isSkip = false
		} else if !isSkip {
			str = str[4 : len(str)-1]

			r := 1

			numStrArr := strings.Split(str, ",")
			for _, numStr := range numStrArr {
				num, _ := strconv.Atoi(numStr)
				r *= num
			}

			ret += r
		}

	}

	return ret
}
