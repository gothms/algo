package main

import (
	"fmt"
)

func main() {
	skill := []int{1, 5, 2, 4}
	mana := []int{5, 1, 4, 2}
	skill = []int{1, 2, 3, 4}
	mana = []int{1, 2} // 21
	skill = []int{2, 8, 8, 6}
	mana = []int{8, 7, 10, 3} // 382
	time := minTime(skill, mana)
	fmt.Println(time)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minTime(skill []int, mana []int) int64 {
	// lc：凸包 + 二分，前置知识：二维计算几何，凸包，Andrew 算法	TODO
	// lc：record 优化
	// lc：递推，同个人方案，只是代码优化了
	// lc：正反两次扫描

	// 个人：dp
	n := len(skill)
	pre := make([]int, n+1)
	for i, v := range skill {
		pre[i+1] = pre[i] + v
	}
	ans := pre[n] * mana[0]
	for i, v := range mana[1:] {
		cur := ans
		for j := 0; j < n; j++ {
			// lc：递推，cur 只记录增量，最后 ans+cur
			cur = max(cur, ans-(pre[n]-pre[j+1])*mana[i]+(pre[n]-pre[j])*v)
		}
		ans = max(ans, cur)
	}
	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
