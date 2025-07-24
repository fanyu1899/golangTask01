package main

func plusOne(digits []int) []int {
	count := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 == 10 {
			digits[i] = 0
			count++
		} else {
			digits[i] = digits[i] + 1
			break
		}
	}

	if count == len(digits) {
		nums := make([]int, len(digits)+1)
		nums[0] = 1
		for i := 1; i < len(nums); i++ {
			nums[i] = digits[i-1]
		}
		return nums
	}

	return digits
}
