package main

import "sort"

func main() {

}
func countKSubsequencesWithMaxBeauty(s string, k int) int {
	if k > 26 {
		return 0
	}
	const mod = 1_000_000_007
	fn := make([]int, 26)
	for _, c := range s {
		fn[c-'a']++
	}
	sort.Ints(fn)
	cnt, last := 1, 26-k
	if fn[last] == 0 {
		return 0
	}
	// 不考虑 f(c) 最小值重复
	for i := 25; i >= last; i-- {
		cnt = cnt * fn[i] % mod
	}
	// 考虑 f(c) 最小值重复
	l, r := last-1, last+1
	for r < 26 && fn[r] == fn[last] {
		r++
	}
	for l >= 0 && fn[l] == fn[last] {
		l--
	}
	lastCnt, choose, chooseCnt := r-l-1, r+k-26, 1 // 组合数 chooseCnt = 从 lastCnt 个元素中选出 choose 个
	for i, j := lastCnt, 1; j <= choose; i, j = i-1, j+1 {
		chooseCnt = chooseCnt * i / j
	}
	return cnt * chooseCnt % mod
}
