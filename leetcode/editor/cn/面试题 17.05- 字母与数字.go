package main

import "unicode"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findLongestSubarray(array []string) []string {
	// lc
	//s := make([]int, len(array)+1) // 前缀和
	//for i, v := range array {
	//	s[i+1] = s[i] + int(v[0])>>6&1*2 - 1
	//}
	////fmt.Println(s)
	//begin, end := 0, 0 // 符合要求的子数组 [begin,end)
	//first := map[int]int{}
	//for i, v := range s {
	//	if j, ok := first[v]; !ok { // 首次遇到 s[i]
	//		first[v] = i
	//	} else if i-j > end-begin { // 更长的子数组
	//		begin, end = j, i
	//	}
	//}
	//return array[begin:end]

	// 个人
	// 题意没表达清楚，但由测试可得，array[i] 默认长度为 1，即第一个字符是数字则 array[i] 算作一个字母
	memo := make(map[int]int)
	memo[0] = -1	// 预处理
	idx, l, sum := 0, 0, 0
	for i, s := range array {
		if unicode.IsNumber(rune(s[0])) {
			sum++
		} else {
			sum--
		}
		if j, ok := memo[sum]; ok {
			if i-j > l {
				idx, l = i, i-j
			}
		} else {
			memo[sum] = i
		}
	}
	return array[idx-l+1 : idx+1]
}

//leetcode submit region end(Prohibit modification and deletion)
