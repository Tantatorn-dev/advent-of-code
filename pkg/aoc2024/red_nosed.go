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

		if diff > 3 || diff < -3 || diff == 0 {
			return false
		}

		prev = &diff
	}

	return true
}

func IsLevelSafe2(level []int) bool {
	if IsLevelSafe(level) {
		return true
	}

	return isLevelSafe2(0, level)
}

func isLevelSafe2(i int, level []int) bool {
	if i < len(level) {
		lev := append([]int{}, level[:i]...)
		lev = append(lev, level[i+1:]...)

		return IsLevelSafe(lev) || isLevelSafe2(i+1, level)
	}

	return false
}
