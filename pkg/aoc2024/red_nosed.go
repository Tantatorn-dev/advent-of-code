package aoc2024

func IsLevelSafe(level []int) bool {
	var prev *int

	for i := 0; i < len(level)-1; i++ {
		diff := level[i+1] - level[i]

		if prev != nil {
			if (*prev > 0 && diff < 0) || (*prev < 0 && diff > 0) {
				return false
			}
		}

		if diff > 3 {
			return false
		} else if diff < -3 {
			return false
		} else if diff == 0 {
			return false
		}

		prev = &diff
	}

	return true
}

func IsLevelSafe2(level []int) bool {
	var prev *int

	for i := 0; i < len(level)-1; i++ {
		diff := level[i+1] - level[i]

		if prev != nil {
			if (*prev > 0 && diff < 0) || (*prev < 0 && diff > 0) {
				return false
			}
		}

		if diff > 3 {
			return false
		} else if diff < -3 {
			return false
		} else if diff == 0 {
			return false
		}

		prev = &diff
	}

	return true
}
