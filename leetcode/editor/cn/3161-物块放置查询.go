package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	queries := [][]int{{1, 2},
		{2, 3, 3},
		{2, 3, 1},
		{2, 2, 2}} // [false,true,true
	queries = [][]int{{1, 7},
		{2, 7, 6},
		{1, 2},
		{2, 7, 5},
		{2, 7, 6}} // [true,true,false]
	queries = [][]int{{1, 3},
		{2, 4, 2}} // [true]
	queries = [][]int{{2, 4, 4},
		{1, 1},
		{2, 4, 7}} // [true,false]
	queries = [][]int{{2, 1, 1}}
	queries = [][]int{{2, 8, 8},
		{2, 7, 3},
		{1, 6},
		{1, 4},
		{2, 7, 10}} // [true,true,false]
	results := getResults(queries)
	fmt.Println(results)
}

// leetcode submit region begin(Prohibit modification and deletion)
func getResults(queries [][]int) []bool {
	// 二分查找 + 树状数组：个人
	// 注意，不能重复放置障碍物
	var (
		n, m int
		bst  []int
	)
	update := func(i, v int) {
		for ; i < n; i += i & -i {
			bst[i] = max(bst[i], v)
		}
	}
	prefixMax := func(i int) int {
		ret := 0
		for ; i > 0; i &= i - 1 {
			ret = max(ret, bst[i])
		}
		return ret
	}

	pos := make([]int, 1)
	maxX := 0
	for _, q := range queries {
		if q[0] == 1 {
			pos = append(pos, q[1])
		}
		maxX = max(maxX, q[1])
	}
	pos = append(pos, maxX+1) // 哨兵
	n, m = maxX+1, len(pos)
	bst = make([]int, n)
	sort.Ints(pos)
	for i := 1; i < m; i++ {
		update(pos[i], pos[i]-pos[i-1]) // 构建树状数组
	}

	ans := make([]bool, 0, len(queries)-m+2)       // 加上首尾的 0 和 maxX+1
	for idx := len(queries) - 1; idx >= 0; idx-- { // 逆序
		x := queries[idx][1]
		//i := sort.Search(m, func(i int) bool {
		//	return pos[i] >= x
		//}) - 1 // 左边界
		switch queries[idx][0] {
		case 1:
			// 写法一
			//j := sort.Search(m, func(i int) bool {
			//	return pos[i] > x // 右边界
			//})
			//update(pos[j], pos[j]-pos[i]) // 合并，并更新
			////if j < len(pos) {
			//for k := j - 1; k > i; k-- { // 合并
			//	pos[k] = pos[j]
			//}
			////}

			// 写法二：优化
			j := sort.Search(m, func(i int) bool {
				return pos[i] > x // 右边界
			})
			i := j - 1 // 左边界
			for ; pos[i] == x; i-- {
				pos[i] = pos[j]
			}
			update(pos[j], pos[j]-pos[i])
		case 2:
			i := sort.Search(m, func(i int) bool {
				return pos[i] >= x
			}) - 1 // 左边界
			ans = append(ans, max(x-pos[i],
				prefixMax(pos[i])) >= queries[idx][2])
		}
	}
	slices.Reverse(ans)
	return ans

	// 并查集 + 树状数组

	// 有序集合 + 树状数组

	// 有序集合 + 线段树
	//var (
	//	rbt    *redblacktree.Tree
	//	st     []int
	//	update func(int, int, int, int, int)
	//	query  func(int, int, int, int) int
	//)
	//update = func(x, l, r, i, val int) {
	//	if l == r {
	//		st[i] = val
	//		return
	//	}
	//	m, idx := (l+r)>>1, i<<1
	//	if x <= m {
	//		update(x, l, m, idx, val)
	//	} else {
	//		update(x, m+1, r, idx+1, val)
	//	}
	//	st[i] = max(st[idx], st[idx+1])
	//}
	//query = func(x, l, r, i int) int {
	//	if x == r { // 递归终止条件，也可以写 x>=r
	//		return st[i]
	//	}
	//	m, idx := (l+r)>>1, i<<1
	//	if x <= m {
	//		return query(x, l, m, idx) // 找左边
	//	} else {
	//		return max(st[idx], query(x, m+1, r, idx+1)) // 左边最大长度 / 找右边
	//	}
	//}
	//
	//ans := make([]bool, 0)
	//maxX := 0
	//for _, q := range queries {
	//	maxX = max(maxX, q[1]) // 右边界
	//}
	//rbt = redblacktree.NewWithIntComparator()
	//rbt.Put(0, struct{}{}) // 初始化数据
	//rbt.Put(maxX, struct{}{})
	//st = make([]int, func(n int) int {
	//	k := bits.Len(uint(n - 1))
	//	stLen := 1 << (k + 1)
	//	if n > 1 {
	//		stLen -= 1<<(k-bits.Len(uint(n-stLen>>2))+1) - 2
	//	}
	//	return stLen
	//}(maxX+1)) // [0,maxX]
	//update(maxX, 0, maxX, 1, maxX)
	//
	//for _, q := range queries {
	//	x := q[1]
	//	left, _ := rbt.Floor(x) // 前驱节点
	//	lv := left.Key.(int)
	//	switch q[0] {
	//	case 1:
	//		if lv != x { // 不重复分隔
	//			right, _ := rbt.Ceiling(x) // 后驱节点
	//			rv := right.Key.(int)
	//			update(x, 0, maxX, 1, x-lv) // 更新线段树
	//			update(rv, 0, maxX, 1, rv-x)
	//			rbt.Put(x, struct{}{}) // 插入分割点
	//		}
	//	case 2:
	//		ans = append(ans, x-lv >= q[2] ||
	//			query(lv, 0, maxX, 1) >= q[2]) // 找左边
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)

// 失败：739 / 742 个通过的测试用例
// 考虑退化为链表的情况
//type stNode struct {
//	left, right *stNode
//	l, r        int
//	length      int
//}
//
//func update(i int, cur *stNode) {
//	if i == cur.l || i == cur.r {
//		cur.length = cur.r - cur.l
//		return
//	}
//	if cur.left == nil {
//		cur.left = &stNode{l: cur.l, r: i, length: i - cur.l}
//		cur.right = &stNode{l: i, r: cur.r, length: cur.r - i}
//	} else {
//		if i < cur.left.r {
//			update(i, cur.left)
//		} else if i > cur.right.l {
//			update(i, cur.right)
//		}
//	}
//	cur.length = max(cur.left.length, cur.right.length)
//}
//func query(i int, cur *stNode) int {
//	if cur.left == nil {
//		return i - cur.l
//	}
//	if i <= cur.left.r {
//		return query(i, cur.left)
//	} else {
//		ret := cur.left.length
//		if i-cur.right.l <= ret {
//			return ret
//		}
//		return max(ret, query(i, cur.right))
//	}
//}
//func getResults(queries [][]int) []bool {
//	n := len(queries)
//	maxR := min(50_000, n*3)
//	ans := make([]bool, 0)
//	st := &stNode{l: 0, r: maxR, length: maxR}
//	for _, q := range queries {
//		if len(q) == 2 {
//			update(q[1], st)
//		} else {
//			ans = append(ans, query(q[1], st) >= q[2])
//		}
//	}
//	return ans
//}
