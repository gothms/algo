package str

/*
Sunday 字符串匹配算法
	前缀匹配算法
位移表（坏字符）规则：后移距离 = 最右边不重复字符到字符串末尾的距离
*/

// Sunday 字符串匹配算法
func Sunday(s, p string) []int {
	n, m := len(s)-len(p), len(p)
	ans := make([]int, 0)
	occ := occHelper(p)
	for i := 0; i <= n; {
		j := 0
		for j < m && s[i+j] == p[j] {
			j++
		}
		if j == m {
			ans = append(ans, i)
		}
		if i+m == len(s) {
			break
		}
		i += occ[s[i+m]]
	}
	return ans
}

func occHelper(p string) [256]int {
	occ, n := [256]int{}, len(p)
	for i := 0; i < 256; i++ {
		occ[i] = n + 1
	}
	for i := 0; i < n; i++ {
		occ[p[i]] = n - i
	}
	return occ
}
