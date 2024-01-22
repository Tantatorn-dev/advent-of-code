package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tantatorn-dev/advent-of-code/pkg/aoc2023"
)

func main() {
	content, err := os.ReadFile("fixtures/aoc2023/haunted_wasteland/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	fmt.Println("How many steps to reach 'zzz'?: ", aoc2023.TraverseWasteland(input))
	fmt.Println("How many steps to reach target in Part2?: ", aoc2023.TraverseWasteland2(input))
}
