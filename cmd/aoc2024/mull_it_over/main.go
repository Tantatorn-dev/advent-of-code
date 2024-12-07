package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tantatorn-dev/advent-of-code/pkg/aoc2024"
)

func main() {
	content, err := os.ReadFile("fixtures/aoc2024/mull_it_over/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	fmt.Println("Result: ", aoc2024.MultiplyByInstruction(input))
	fmt.Println("Result: ", aoc2024.MultiplyByInstruction2(input))
}
