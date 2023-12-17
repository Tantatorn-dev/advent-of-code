package aoc2023

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func SumPartNumbers(input string) int {
	sum := 0

	lines := strings.Split(input, "\n")

	reNum := regexp.MustCompile(`(\d+)`)

	for lIndex, line := range lines {
		nums := reNum.FindAllString(line, -1)
		numIndexes := reNum.FindAllStringIndex(line, -1)

		for j, numIndex := range numIndexes {
			isPartNumber := false

			for k := numIndex[0]; k < numIndex[1]; k++ {
				// scan left
				if k > 0 && isSpecialCharacter(line[k-1]) {
					isPartNumber = true
					break
				}

				// scan right
				if len(line) > k+1 && isSpecialCharacter(line[k+1]) {
					isPartNumber = true
					break
				}

				// scan prev line
				if lIndex > 0 {
					prevLine := lines[lIndex-1]

					if isSpecialCharacter(prevLine[k]) {
						isPartNumber = true
						break
					}

					// scan left
					if k > 0 && isSpecialCharacter(prevLine[k-1]) {
						isPartNumber = true
						break
					}

					// scan right
					if len(line) > k+1 && isSpecialCharacter(prevLine[k+1]) {
						isPartNumber = true
						break
					}
				}

				// scan next line
				if len(lines) > lIndex+1 {
					nextLine := lines[lIndex+1]

					if isSpecialCharacter(nextLine[k]) {
						isPartNumber = true
						break
					}

					// scan left
					if k > 0 && isSpecialCharacter(nextLine[k-1]) {
						isPartNumber = true
						break
					}

					// scan right
					if len(line) > k+1 && isSpecialCharacter(nextLine[k+1]) {
						isPartNumber = true
						break
					}
				}
			}

			if isPartNumber {
				n, _ := strconv.Atoi(nums[j])
				sum += n
			}
		}
	}

	return sum
}

// not digit or .
func isSpecialCharacter(b byte) bool {
	return b != '.' && !unicode.IsDigit(rune(b))
}
