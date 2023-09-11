package main

import (
	"sort"
)

func main() {

}
func countKSubsequencesWithMaxBeauty(s string, k int) int {
	if k > 26 {
		return 0
	}
	const mod = 1_000_000_007
	f := make([]int, 26)
	for _, c := range s {
		f[c-'a']++
	}
	sort.Slice(f, func(i, j int) bool {
		return f[i] > f[j]
	})
	if f[k-1] == 0 {
		return 0
	}
	// 不考虑 f(c) 最小值重复
	cnt, last := 1, k-1
	for i := 0; i <= last; i++ {
		cnt = cnt * f[i] % mod
	}
	// 考虑 f(c) 最小值重复
	l, r := last-1, last+1
	for r < 26 && f[r] == f[last] {
		r++
	}
	for l >= 0 && f[l] == f[last] {
		l--
	}
	lastCnt, choose, chooseCnt := r-l-1, k-l-1, 1 // 组合数 chooseCnt = 从 lastCnt 个元素中选出 choose 个
	for i, j := lastCnt, 1; j <= choose; i, j = i-1, j+1 {
		chooseCnt = chooseCnt * i / j
	}
	return cnt * chooseCnt % mod // 关键计算
}
