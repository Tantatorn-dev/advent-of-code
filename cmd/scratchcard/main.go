package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tantatorn-dev/advent-of-code/pkg/aoc2023"
)

func main() {
	content, err := os.ReadFile("fixtures/aoc2023/scratchcard/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	fmt.Println("Value: ", aoc2023.CountWinningNumber(input))
}
