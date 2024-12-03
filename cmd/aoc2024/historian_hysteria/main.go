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
	content, err := os.ReadFile("fixtures/aoc2024/historian_hysteria/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	lines := strings.Split(input, "\n")

	var l1, l2 []int64

	for _, line := range lines {
		numStrings := strings.Split(line, "   ")

		n1, _ := strconv.ParseInt(numStrings[0], 10, 64)
		n2, _ := strconv.ParseInt(numStrings[1], 10, 64)

		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}

	fmt.Println(aoc2024.FindTotalDistances(l1, l2))
	fmt.Println(aoc2024.FindSimilarityScore(l1, l2))
}
