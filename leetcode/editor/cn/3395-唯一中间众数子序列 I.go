package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 3, 4} // 4
	mode := subsequencesWithMiddleMode(nums)
	fmt.Println(mode)
}

// leetcode submit region begin(Prohibit modification and deletion)
func subsequencesWithMiddleMode(nums []int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func subsequencesWithMiddleMode_(nums []int) int {
	// 枚举中间众数
	comb2 := func(x int) int {
		return x * (x - 1) >> 1
	}

	n := len(nums)
	suf := make(map[int]int) // 后缀
	for _, v := range nums {
		suf[v]++
	}
	pre := make(map[int]int, len(suf)) // 前缀

	ans := n * (n - 1) * (n - 2) * (n - 3) * (n - 4) / (2 * 3 * 4 * 5)
	for i, v := range nums[:n-2] { // 枚举中间众数 v
		suf[v]--
		if i >= 2 {
			preV, sufV := pre[v], suf[v]
			ans -= comb2(i-preV) * comb2(n-i-1-sufV) // v 仅出现一次
			// 枚举 w，枚举 pre 更优：ERROR，因为 suf 有的 key 值，pre 不一定有，所以 suf 的 key 也不能 delete
			for w, sufW := range suf { // V 出现2次，w 出现 >= 2 次
				preW := pre[w]
				//for w, sufW := range suf { // V 出现2次，w 出现 >= 2 次
				//	preW := pre[w]
				if w == v || preW+sufW < 2 {
					continue
				}
				// v 出现2次，ww v vx
				ans -= comb2(preW) * sufV * (n - i - 1 - sufV)
				// v 出现2次，vx v ww
				ans -= preV * (i - preV) * comb2(sufW)
				// v 出现2次，wv v wx
				ans -= preW * preV * sufW * (n - i - 1 - sufV - sufW)
				// v 出现2次，wx v wv
				ans -= preW * (i - preV - preW) * sufW * sufV
			}
		}
		pre[v]++
	}
	return ans % (1e9 + 7)
}
