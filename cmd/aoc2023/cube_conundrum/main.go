package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tantatorn-dev/advent-of-code/pkg/aoc2023"
)

func main() {
	content, err := os.ReadFile("fixtures/aoc2023/cube_conundrum/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	fmt.Println("Part 1: ", aoc2023.SolveCubeConundrum(input))
	fmt.Println("Part 2: ", aoc2023.SolveCubeConundrum2(input))
}
