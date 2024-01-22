package utils

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func LCMSlice(nums []int) int {
	lcm := nums[0]
	for i := 1; i < len(nums); i++ {
		lcm = LCM(lcm, nums[i])
	}
	return lcm
}
