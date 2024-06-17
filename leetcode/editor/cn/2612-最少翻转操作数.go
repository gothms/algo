package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
	"sort"
)

func main() {
	n, p, k := 5, 0, 2
	banned := []int{} // [0,1,2,3,4]
	n, p, k = 4, 0, 4
	banned = []int{1, 2} // [0,-1,-1,1]
	n, p, k = 3, 0, 2
	banned = []int{2, 1} // [0,-1,-1]
	operations := minReverseOperations(n, p, banned, k)
	fmt.Println(operations)
}

/*
思路
    q 是 p 翻转后的位置：l+r = p+q
        对于 p 在任意位置 i：q 的区间为 [i-k+1,i+k-1]
    左边界：q = 0 + k-1 - p，即对于 p 在任意位置 i 有：k-i-1 >= 0
		重要：i-k+1 和 k-i-1 奇偶性相同
    右边界：n-k + n-1 - p，有：n<<1-k-i-1
    不考虑边界的情况下，子数组的区间为 [i-k+1,i+k-1]，总共可翻转 k 次
        即翻转后的位置 q 成等差数列，差值为 2
        则据此合并并查集
*/

// leetcode submit region begin(Prohibit modification and deletion)
func minReverseOperations(n int, p int, banned []int, k int) []int {
	// bfs + 并查集
	// 参考另一种写法：https://leetcode.cn/problems/minimum-reverse-operations/solutions/2204092/liang-chong-zuo-fa-ping-heng-shu-bing-ch-vr0z/
	//uni := make([]int, n+2) // 初始化并查集，长度是 n+2（哨兵）
	//for i := range uni {
	//	uni[i] = i
	//}
	//var find func(int) int
	//find = func(i int) int {
	//	if uni[i] != i {
	//		uni[i] = find(uni[i])
	//	}
	//	return uni[i]
	//}
	//for _, i := range banned { // ban 掉的索引
	//	uni[i] = i + 2
	//}
	//uni[p] = p + 2 // 1 的初始位置
	//
	//ans := make([]int, n) // 初始化结果集
	//for i := range ans {
	//	ans[i] = -1
	//}
	//q := []int{p}                                    // 队列
	//for size, cnt := 1, 0; size > 0; size = len(q) { // bfs
	//	for _, i := range q[:size] {
	//		ans[i] = cnt                                // 翻转次数
	//		l := max(i-k+1, k-i-1)                      // 左边界
	//		r := min(i+k-1, n<<1-i-k-1)                 // 右边界
	//		for j := find(l); j <= r; j = find(j + 2) { // 应用并查集
	//			uni[j] = find(j + 2) // uni[find(j)] = find(j + 2)
	//			q = append(q, j)
	//		}
	//	}
	//	cnt++
	//	q = q[size:]
	//}
	//return ans

	// bfs + 有序集合
	m := len(banned)
	sort.Ints(banned) // banned 可能无序，也可以使用 map
	rbt := [2]*redblacktree.Tree{redblacktree.NewWithIntComparator(), redblacktree.NewWithIntComparator()}
	for i, j := 0, 0; i < n+2; i++ { // n+2：增加哨兵
		if j < m && i == banned[j] {
			j++
			continue
		}
		rbt[i&1].Put(i, struct{}{})
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	rbt[p&1].Remove(p)
	q := []int{p}
	for length, cnt := 1, 0; length > 0; length = len(q) {
		for _, i := range q[:length] {
			ans[i] = cnt
			l := max(i-k+1, k-i-1)
			r := min(i+k-1, n<<1-i-k-1)
			rbtId := rbt[l&1]
			right, _ := rbtId.Ceiling(l)
			j := right.Key.(int)
			for j <= r {
				q = append(q, j)
				rbtId.Remove(j) // 移除节点
				right, _ = rbtId.Ceiling(j)
				j = right.Key.(int)
			}
		}
		cnt++
		q = q[length:]
	}
	return ans

	// bfs：超时
	//ans := make([]int, n)
	//for i := range ans {
	//	ans[i] = -1
	//}
	//ans[p] = 0
	//if k == 1 {
	//	return ans
	//}
	//vis := make([]bool, n)
	//for _, i := range banned {
	//	vis[i] = true
	//}
	//vis[p] = true
	//q := []int{p}
	//for length, cnt := 1, 0; length > 0; length = len(q) { // bfs
	//	cnt++
	//	for _, i := range q[:length] {
	//		for j := max(0, i-k+1); j <= min(i, n-k); j++ {
	//			if idx := k - i + j<<1 - 1; !vis[idx] {
	//				vis[idx] = true
	//				ans[idx] = cnt
	//				q = append(q, idx)
	//			}
	//		}
	//	}
	//	q = q[length:]
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
