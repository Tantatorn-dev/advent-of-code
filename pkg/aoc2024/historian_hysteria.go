package aoc2024

import (
	"slices"
)

func FindTotalDistances(l1, l2 []int64) int64 {
	slices.Sort(l1)
	slices.Sort(l2)

	ret := int64(0)

	for i := 0; i < len(l1); i++ {
		diff := l1[i] - l2[i]

		if diff < 0 {
			diff = -diff
		}

		ret += diff
	}

	return ret
}

func FindSimilarityScore(l1, l2 []int64) int64 {
	var score int64

	for _, i1 := range l1 {
		for _, i2 := range l2 {
			if i1 == i2 {
				score += i1
			}
		}
	}

	return score
}
