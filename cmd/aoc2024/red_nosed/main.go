package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Tantatorn-dev/advent-of-code/pkg/aoc2024"
)

func main() {
	content, err := os.ReadFile("fixtures/aoc2024/red_nosed/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	lines := strings.Split(input, "\n")

	safeCount := 0

	safeCount2 := 0

	for _, line := range lines {
		lineStrArr := strings.Split(line, " ")

		var level []int
		for _, str := range lineStrArr {
			l, _ := strconv.Atoi(str)
			level = append(level, l)
		}

		if aoc2024.IsLevelSafe(level) {
			safeCount++
		}

		if aoc2024.IsLevelSafe2(level) {
			safeCount2++
		}
	}

	fmt.Println("Total safe records: ", safeCount)
	fmt.Println("Total safe records 2: ", safeCount2)
}
