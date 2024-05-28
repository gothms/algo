package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "banana"
	// havbc
	// aeeebaabd
	substring := longestDupSubstring(s)
	fmt.Println(len(strings.Split(s, "havbc")))
	fmt.Println(len(strings.Split(s, "rvlmms")))
	fmt.Println(substring, len(s))

	//fmt.Println('a', '`')
	//fmt.Printf("%c\n", '`')
	//? 6 1194086401
	//? 9 1678296320
	//? 11 464974345
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestDupSubstring(s string) string {
	// 二分
	//const primeRK, B = 16777619, '`'
	const primeRK, B = 2099999999, '`' // 小于 21 亿的第一个质数
	check := func(m int) int {
		if m == 0 {
			return 0
		}
		memo := make(map[uint32]int)
		var hash, pow uint32 = 0, 1
		for _, c := range s[:m] {
			hash = hash*primeRK + uint32(c-B)
			pow *= primeRK
		}
		memo[hash] = 0
		for i, c := range s[m:] {
			hash = hash*primeRK + uint32(c-B) - uint32(s[i]-B)*pow
			if j, ok := memo[hash]; ok && s[j:j+m] == s[i+1:i+m+1] {
				//if j, ok := memo[hash]; ok { // 错误：很长字符串会有 hash 碰撞
				return i + 1
			}
			memo[hash] = i + 1 // 起始位置
		}
		return -1
	}
	var idx, maxL int
	sort.Search(len(s), func(i int) bool { // 二分
		j := check(i)
		if j >= 0 && i > maxL { // 长度为 i 的子串
			idx, maxL = j, i
		}
		return j == -1
	})
	if maxL == 0 {
		return ""
	}
	return s[idx : idx+maxL]

	// 超时
	//const primeRK = 16777619
	//memo := make(map[int]struct{})
	//idx, maxL := 0, 0
	//for i := range s {
	//	var hash int
	//	for j, c := range s[i:] {
	//		hash = hash*primeRK + int(c-'`')
	//		if j < maxL {
	//			continue
	//		}
	//		if _, ok := memo[hash]; ok {
	//			idx, maxL = i, j+1
	//		} else {
	//			memo[hash] = struct{}{}
	//		}
	//	}
	//}
	//if maxL == 0 {
	//	return ""
	//}
	//return s[idx : idx+maxL]
}

//leetcode submit region end(Prohibit modification and deletion)
