package main

import (
	"fmt"
)

func main() {
	words := []string{"bdb", "aaa", "ada"}
	groups := []int{2, 1, 3}
	n := 3
	words = []string{"cb", "dcc", "da", "cbb", "bd", "dbc", "ab", "db"}
	groups = []int{4, 5, 5, 7, 8, 1, 3, 4}
	n = 8
	subsequence := getWordsInLongestSubsequence(n, words, groups)
	fmt.Println(subsequence)
}
func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
	fMax, from := make([]int, n), make([]int, n)
	fMax[0] = 1
	ok := func(i, j int) bool {
		a, b := words[i], words[j]
		if len(a) != len(b) {
			return false
		}
		for idx := 0; idx < len(a); idx++ {
			if a[idx] != b[idx] {
				return a[idx+1:] == b[idx+1:]
			}
		}
		return false
	}
	max := 0
	for i := 1; i < n; i++ {
		fMax[i] = 1 // 初始长度都为 1
		for j := i - 1; j >= 0; j-- {
			if fMax[j]+1 > fMax[i] && groups[i] != groups[j] && ok(i, j) { // 比当前长度大，groups 值 不同，汉明距离为 1
				fMax[i] = fMax[j] + 1 // 更新长度
				from[i] = j           // 记录汉明距离 path
			}
		}
		if fMax[i] > fMax[max] { // 更新结束索引
			max = i
		}
	}
	ret := make([]string, fMax[max])
	for i, idx := len(ret)-1, max; i >= 0; i-- { // 根据 path，还原路径上的字符串
		ret[i] = words[idx]
		idx = from[idx]
	}
	return ret
}
