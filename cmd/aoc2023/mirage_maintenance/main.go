package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tantatorn-dev/advent-of-code/pkg/aoc2023"
)

func main() {
	content, err := os.ReadFile("fixtures/aoc2023/mirage_maintenance/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	fmt.Println(aoc2023.SumExtrapolatedValues(input, false))
	fmt.Println(aoc2023.SumExtrapolatedValues(input, true))
}
