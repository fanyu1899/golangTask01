package main

func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		maps[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		v, ok := maps[target-nums[i]]
		if ok && v != i {
			return []int{v, i}
		}
	}
	return []int{-1, -1}
}
