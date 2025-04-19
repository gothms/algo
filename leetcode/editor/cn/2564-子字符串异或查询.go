package main

import (
	"fmt"
	"math/bits"
	"strings"
)

func main() {
	s := "111010110" // [[1,3],[3,3],[2,3],[3,3],[4,8],[0,8],[2,8],[4,8],[0,2],[0,8],[0,1],[1,4],[3,3],[3,3],[1,8],[1,3],[3,3],[2,8]]
	q := [][]int{{4, 2},
		{3, 3},
		{6, 4},
		{9, 9},
		{10, 28},
		{0, 470},
		{5, 83},
		{10, 28},
		{8, 15},
		{6, 464},
		{0, 3},
		{5, 8},
		{7, 7},
		{8, 8},
		{6, 208},
		{9, 15},
		{2, 2},
		{9, 95}}
	queries := substringXorQueries(s, q)
	fmt.Println(queries)

	fmt.Println(bits.Len(1e9))
}

// leetcode submit region begin(Prohibit modification and deletion)
func substringXorQueries(s string, queries [][]int) [][]int {
	// lc：由于 0 <= firsti, secondi <= 109，所以 s 的子字符串最多表示 30 位
	memo := make(map[int][]int)
	if i := strings.IndexByte(s, '0'); i >= 0 {
		memo[0] = []int{i, i}
	}
	n := len(s)
	for i, c := range s {
		if c == '0' {
			continue
		}
		for j, v := i, 0; j < min(n, i+30); j++ {
			v = v<<1 | int(s[j]&1)
			if _, ok := memo[v]; !ok {
				memo[v] = []int{i, j}
			}
		}
	}
	ans := make([][]int, len(queries))
	for i, q := range queries {
		if p := memo[q[0]^q[1]]; p != nil {
			ans[i] = p
		} else {
			ans[i] = []int{-1, -1}
		}
	}
	return ans

	// 个人
	//type bt struct { // trie
	//	child [2]*bt
	//	is    []int
	//}
	//root := &bt{}
	//insert := func(idx, v int) {
	//	k := bits.Len(uint(v)) - 1
	//	cur := root
	//	for ; k >= 0; k-- {
	//		i := v >> k & 1
	//		if cur.child[i] == nil {
	//			cur.child[i] = &bt{}
	//		}
	//		cur = cur.child[i]
	//	}
	//	cur.is = append(cur.is, idx) // 当 v 是 0 时，索引都在 root 的 is 中
	//}
	//for i, q := range queries {
	//	insert(i, q[0]^q[1])
	//}
	//
	//n, m := len(s), len(queries)
	//cnt := m
	//ans := make([][]int, m)
	//for i := range ans {
	//	ans[i] = []int{-1, -1}
	//}
	//query := func(cur *bt, i, j int) {
	//	cnt -= len(cur.is)
	//	ret := []int{i, j}
	//	for _, k := range cur.is {
	//		ans[k] = ret
	//	}
	//	cur.is = cur.is[:0]
	//}
	//for i, c := range s {
	//	if c == '0' { // 要找最短子字符串
	//		if len(root.is) > 0 {
	//			query(root, i, i)
	//			if cnt == 0 { // 已查完
	//				return ans
	//			}
	//		}
	//		continue
	//	}
	//	cur := root
	//	for j := i; j < n; j++ {
	//		if cur = cur.child[s[j]&1]; cur == nil {
	//			break
	//		}
	//		if len(cur.is) > 0 {
	//			query(cur, i, j)
	//			if cnt == 0 { // 已查完
	//				return ans
	//			}
	//		}
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
