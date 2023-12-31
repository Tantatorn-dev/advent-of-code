package aoc2023

type Race struct {
	Time     int
	Distance int
}

// problem: https://adventofcode.com/2023/day/6
func CountWinningWays(r Race) int {
	ret := 0

	for i := 0; i < r.Time; i++ {
		if calculateDistance(i, r.Time) > r.Distance {
			ret += 1
		}
	}

	return ret
}

func CountWinningWay2(r Race) int {
	var minVal *int
	var maxVal *int

	for left := 0; left < r.Time; left++ {
		right := r.Time - left

		if minVal == nil && calculateDistance(left, r.Time) > r.Distance {
			minVal = &left
		}

		if maxVal == nil && calculateDistance(right, r.Time) > r.Distance {
			maxVal = &right
		}

		if minVal != nil && maxVal != nil {
			break
		}
	}

	return *maxVal - *minVal + 1
}

func calculateDistance(pushedTime int, totalTime int) int {
	speed := pushedTime
	remainTime := totalTime - pushedTime

	return speed * remainTime
}
