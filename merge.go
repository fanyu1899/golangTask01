package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] // 按第一个元素升序
	})

	result := [][]int{}
	count := 0
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= result[count][1] {
			result[count][1] = max(intervals[i][1], result[count][1])
		} else {
			count++
			result = append(result, intervals[i])
		}
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
