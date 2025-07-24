package main

func singleNumber(nums []int) int {
	var dic = make(map[int]int)

	for _, v := range nums {
		val, ok := dic[v]
		if ok {
			dic[v] = val + 1
		} else {
			dic[v] = 1
		}

	}

	for k, v := range dic {
		if v == 1 {
			return k
		}
	}
	return 1
}
