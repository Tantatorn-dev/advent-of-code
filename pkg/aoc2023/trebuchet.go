package aoc2023

// Problem 1: https://adventofcode.com/2023/day/1

func calibrateValue(input string) int {
	ret := 0

	first := int32(0)

	arr := []int{}

	for i, ch := range input {
		if first == 0 {
			first = ch
		}

		if ch == '\n' {
			ret += int(first) + int(input[i-1])
			arr = append(arr, int(first))
			arr = append(arr, int(input[i-1]))
			first = int32(0)
		}
	}

	return ret
}
