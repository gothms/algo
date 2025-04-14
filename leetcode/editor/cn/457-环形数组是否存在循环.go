package main

import "fmt"

func main() {
	nums := []int{2, -1, 1, 2, 2} // true
	//nums = []int{-1, -2, -3, -4, -5, 6} // false
	nums = []int{-1, -2, -3, -4, -5} // false
	nums = []int{-2, -3, -9}         // false
	loop := circularArrayLoop(nums)
	fmt.Println(loop)
}

// leetcode submit region begin(Prohibit modification and deletion)
func circularArrayLoop(nums []int) bool {
	// lc：同链表环的判断，使用快慢指针，最后将访问过的节点 nums[i] 改为 0

	// 个人
	n := len(nums)
	vis := make([]int, n) // 0 未访问，>0 访问中，-1 已访问
	var dfs func(int, int, int, int) bool
	dfs = func(i, idx, last, c int) bool { // i 索引，idx 从 1 开始计数的第几个节点，last 上一个 nums[i]，c 连续为正/负的节点数
		if vis[i] == -1 {
			return false
		} else if vis[i] > 0 {
			cir := idx - vis[i] + 1    // 环的长度大于 1
			return cir > 1 && c >= cir // 连续同符号的数 >= 环的长度
		}
		idx++
		vis[i] = idx
		defer func() { vis[i] = -1 }()
		v := nums[i]
		if v*last < 0 { // 符号相反
			c = 1
		} else {
			c++
		}
		v %= n // 防止越界
		if dfs((i+v+n)%n, idx, v, c) {
			return true
		}
		return false
	}
	for i := range nums {
		if dfs(i, 0, 0, 0) {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
