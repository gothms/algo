package main

import (
	"fmt"
	"slices"
)

func main() {
	s := "abcdef"
	s = "eccdnmcnkl" // 1
	length := maximumLength(s)
	fmt.Println(length)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumLength(s string) int {
	// lc
	memo := [26][]int{}
	cnt, n := 0, len(s)
	for i := range s {
		cnt++
		if i+1 == n || s[i] != s[i+1] {
			j := int(s[i] - 'a')
			memo[j] = append(memo[j], cnt)
			cnt = 0
		}
	}
	ans := 0
	for _, val := range memo {
		if len(val) == 0 {
			continue
		}
		slices.SortFunc(val, func(a, b int) int {
			return b - a
		})
		val = append(val, 0, 0)
		ans = max(ans, val[0]-2, min(val[0]-1, val[1]), val[2])
	}
	if ans == 0 {
		return -1
	}
	return ans

	// 个人
	//memo := make(map[uint8]map[int]int)
	//ans, n := 0, len(s)
	//for i := 0; i < n; i++ {
	//	j := i
	//	for i+1 < n && s[i] == s[i+1] {
	//		i++
	//	}
	//	c := i - j + 1
	//	if memo[s[i]] == nil {
	//		memo[s[i]] = make(map[int]int)
	//	}
	//	memo[s[i]][c]++
	//	k := memo[s[i]][c]
	//	switch k {
	//	case 1:
	//		if memo[s[i]][c+1] == 1 {
	//			ans = max(ans, c)
	//		} else if memo[s[i]][c-1] >= 1 {
	//			ans = max(ans, c-1)
	//		} else if c >= 3 {
	//			ans = max(ans, c-2)
	//		}
	//	case 2:
	//		ans = max(ans, c-1)
	//	default:
	//		ans = max(ans, c)
	//	}
	//}
	//if ans == 0 {
	//	return -1
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
