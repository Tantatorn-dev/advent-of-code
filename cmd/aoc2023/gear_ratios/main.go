package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tantatorn-dev/advent-of-code/pkg/aoc2023"
)

func main() {
	content, err := os.ReadFile("fixtures/aoc2023/gear_ratios/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	fmt.Println("Part 1: ", aoc2023.SumPartNumbers(input))
}
