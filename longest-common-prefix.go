package main

func longestCommonPrefix(strs []string) string {

	l := 0
	for i := 0; i < 100000000; i++ {

		if i > len(strs[0]) {
			l = i

			break
		}
		s := strs[0][0:i]

		flag := false
		for _, v := range strs {

			if i > len(v) || s != v[0:i] {

				flag = true
				break
			}
		}

		println(flag)
		if flag {
			l = i
			break
		}
	}
	if l == 0 {
		return ""
	}
	println(strs[0][0 : l-1])
	return strs[0][0 : l-1]
}
