package main

func removeDuplicates(nums []int) int {
	xb := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[xb] = nums[i]
			xb++
		}
	}
	return xb
}
