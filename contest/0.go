package main

import (
	"fmt"
	"slices"
	"sort"
	"unicode"
)

func main() {
	//skills := []int{4, 2, 6, 3, 9}
	//k := 3
	//skills = []int{4, 18, 17, 20, 15, 12, 8, 5}
	//k = 1 // 1
	//player := findWinningPlayer(skills, k)
	//fmt.Println(player)

	nums := []int{1, 2, 1, 1, 3}
	k := 2
	nums = []int{1, 2, 3, 4, 5, 1}
	k = 0
	//nums = []int{29, 29, 30}
	//k = 3
	//nums = []int{29, 30, 30}
	//k = 0
	//nums = []int{89, 89, 90, 88, 88, 88, 88, 90, 90}
	//k = 2 // 8
	length := maximumLength(nums, k)
	fmt.Println(length)
}
func clearDigits(s string) string {
	buf := make([]byte, 0)
	for i, cnt := len(s)-1, 0; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			cnt++
			continue
		}
		if cnt > 0 {
			cnt--
		} else {
			buf = append(buf, s[i])
		}
	}
	slices.Reverse(buf)
	return string(buf)
}
func findWinningPlayer(skills []int, k int) int {
	i, cnt := 0, 0
	for j := 1; j < len(skills); j++ {
		if skills[i] > skills[j] {
			cnt++
		} else {
			cnt = 1
			i = j
		}
		if cnt == k {
			return i
		}
	}
	return i
}
func maximumLength(nums []int, k int) int {
	temp := slices.Clone(nums)
	sort.Ints(temp)
	temp = slices.Compact(temp)
	ex := make([][]int, len(temp))
	for i := range ex {
		ex[i] = make([]int, k+1)
	}

	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	//memo := make(map[int]int) // nums[i] 最后出现的位置
	val := make([]int, k+1) // val[i] 表示有 i 个 seq[i] != seq[i+1] 的最长序列
	ans := 1
	for i, v := range nums {
		dp[i][0] = 1 // 至少一个
		//if last, ok := memo[v]; ok { // 出现过
		//	for j, l := range dp[last] {
		//		dp[i][j] = l + 1
		//	}
		//}
		idx := sort.SearchInts(temp, v)
		for j := 1; j <= k; j++ { // 在最长序列后追加
			dp[i][j] = max(dp[i][j], val[j-1]+1)
			dp[i][j] = max(dp[i][j], ex[idx][j]+1)
		}
		dp[i][0] = max(dp[i][0], ex[idx][0]+1)
		for j := 0; j <= k; j++ {
			ans = max(ans, dp[i][j])
			val[j] = max(val[j], dp[i][j])
			ex[idx][j] = max(ex[idx][j], dp[i][j])
		}
		//memo[v] = i
	}
	return ans

	//n := len(nums)
	//dp := make([][]int, n)
	//for i := range dp {
	//	dp[i] = make([]int, k+1)
	//}
	//memo := make(map[int]int) // nums[i] 最后出现的位置
	//val := make([]int, k+1)   // val[i] 表示有 i 个 seq[i] != seq[i+1] 的最长序列
	//ans := 1
	//for i, v := range nums {
	//	dp[i][0] = 1                 // 至少一个
	//	if last, ok := memo[v]; ok { // 出现过
	//		for j, l := range dp[last] {
	//			dp[i][j] = l + 1
	//		}
	//	}
	//	for j := 1; j <= k; j++ { // 在最长序列后追加
	//		dp[i][j] = max(dp[i][j], val[j-1]+1)
	//	}
	//	for j := 0; j <= k; j++ {
	//		ans = max(ans, dp[i][j])
	//		val[j] = max(val[j], dp[i][j])
	//	}
	//	memo[v] = i
	//}
	//return ans
}
