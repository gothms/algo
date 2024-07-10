package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 0, 1, 0, 1}
	k, mc := 3, 0 // 4
	//nums = []int{1, 0, 1}
	//k, mc = 2, 0 // 2
	moves := minimumMoves(nums, k, mc)
	fmt.Println(moves)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumMoves(nums []int, k int, maxChanges int) int64 {

}

// leetcode submit region end(Prohibit modification and deletion)

func minimumMoves_(nums []int, k int, maxChanges int) int64 {
	idx := make([]int, 0) // 哨兵
	cnt := 0
	for i, v := range nums {
		if v == 0 {
			continue
		}
		idx = append(idx, i) // 1
		cnt = max(cnt, 1)    // 连续的 1
		if i > 0 && nums[i-1] == 1 {
			if i > 1 && nums[i-2] == 1 {
				cnt = 3
			} else {
				cnt = max(cnt, 2)
			} 
		}
	}
	cnt = min(cnt, k)
	if maxChanges >= k-cnt { // 使用连续的 1 和 maxChanges
		return int64((k-cnt)<<1 + max(cnt-1, 0))
	}
	ret, size, n := int64(math.MaxInt64), k-maxChanges, len(idx)
	pre := make([]int, n+1)
	for i, v := range idx { // 前缀和
		pre[i+1] = pre[i] + v
	}
	// 核心逻辑的计算：如左侧 = idx[m]-idx[l1] + idx[m]-idx[l2] + ...
	l, r, mid := 0, size, size>>1                                   // 滑动窗体 (l,r]，标记 idx 的下标
	for lc, rc := mid, r-mid; r <= n; l, r, mid = l+1, r+1, mid+1 { // 中位数
		left := lc*idx[mid] - pre[mid] + pre[l]  // 左边：lc 个 1
		right := pre[r] - pre[mid] - rc*idx[mid] // 右边：rc 个 1
		ret = min(ret, int64(left+right))
	}
	return int64(maxChanges<<1) + ret
}
