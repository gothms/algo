package main

import (
	"fmt"
	"sort"
)

func main() {
	k := 5       // 0
	k = 11       // 0
	k = 0        // 5
	k = 5        // 0
	k = 3        // 5
	k = 80502705 // 0
	k = 25       // 5
	k = 67348277 // 5
	fzf := preimageSizeFZF(k)
	fmt.Println(fzf)
}

// leetcode submit region begin(Prohibit modification and deletion)
func preimageSizeFZF(k int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func preimageSizeFZF_(k int) int {
	// 优化
	//f := func(x int) int {
	//	ans := 0
	//	for x != 0 {
	//		x /= 5
	//		ans += x
	//	}
	//	return ans
	//}
	//i := sort.Search(k, func(i int) bool {
	//	return f(i)+i >= k
	//})
	//if f(i)+i == k { // 验证
	//	return 5
	//}
	//return 0 // 结果只能是 5/0

	// 一次二分查找
	if k == 0 { // 此方法要特殊处理 k=0
		return 5
	}
	f := func(x int) int {
		ans := 0
		for x != 0 {
			x /= 5
			ans += x
		}
		return ans
	}
	l, c := 0, 0
	i := sort.Search(k, func(i int) bool {
		l, c = i, f(i)+i // 记录最后一次的数据
		return c >= k
	})
	cc := 0
	for v := i * 5; v%5 == 0; v /= 5 { // 计算 cnt5(i*5)，即 i*5 有多少个 5 因子
		cc++
	}
	if l == i && c == k || l < i && c+cc == k { // 如果 i*5 有多个 5 因子，则 c+cnt5(i*5)==k 时，才返回 5
		return 5
	}
	return 0

	// 二分 + math
	//zeta := func(n int) int { // 核心
	//	c := 0
	//	for n > 0 {
	//		n /= 5
	//		c += n
	//	}
	//	return c
	//}
	//return (sort.Search(k+1, func(i int) bool {
	//	return zeta(i)+i >= k+1
	//}) - sort.Search(k, func(i int) bool {
	//	return zeta(i)+i >= k
	//})) * 5

	// 个人：严重超时
	//if k == 0 {
	//	return 5
	//}
	//last, cnt := 0, 0
	//idx := sort.Search(k+1, func(i int) bool {
	//	cnt = 0
	//	for j := 1; j <= i; j++ {
	//		v := j
	//		for v%5 == 0 {
	//			v /= 5
	//			cnt++
	//		}
	//	}
	//	cnt += i
	//	last = i
	//	return cnt >= k
	//})
	//c := 0
	//for v := idx * 5; v%5 == 0; v /= 5 {
	//	c++
	//}
	//if last == idx && c == 1 || last < idx && cnt+c == k {
	//	return 5
	//}
	//return 0
}
