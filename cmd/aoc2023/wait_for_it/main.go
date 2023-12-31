package main

import (
	"fmt"

	"github.com/Tantatorn-dev/advent-of-code/pkg/aoc2023"
)

func main() {
	example := []aoc2023.Race{
		{Time: 7, Distance: 9},
		{Time: 15, Distance: 40},
		{Time: 30, Distance: 200},
	}

	res := 1
	for _, e := range example {
		res *= aoc2023.CountWinningWays(e)
	}

	fmt.Println("Example: ", res)

	input := []aoc2023.Race{
		{Time: 56, Distance: 499},
		{Time: 97, Distance: 2210},
		{Time: 77, Distance: 1097},
		{Time: 93, Distance: 1440},
	}

	res = 1
	for _, e := range input {
		res *= aoc2023.CountWinningWays(e)
	}

	fmt.Println("Part 1: ", res)

	fmt.Println("Part 2: ", aoc2023.CountWinningWay2(aoc2023.Race{
		Time:     56977793,
		Distance: 499221010971440,
	}))
}
