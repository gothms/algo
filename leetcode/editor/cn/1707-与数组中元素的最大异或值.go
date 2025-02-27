package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 2, 4, 6, 6, 3}
	queries := [][]int{{12, 4}, {8, 1}, {6, 3}}
	xor := maximizeXor(nums, queries)
	fmt.Println(xor)
}

// leetcode submit region begin(Prohibit modification and deletion)

type trie struct {
	child [2]*trie
}

const bLen = 30

func (t *trie) insert(v int) {
	cur := t
	for i := bLen - 1; i >= 0; i-- {
		bit := v >> i & 1
		if cur.child[bit] == nil {
			cur.child[bit] = &trie{}
		}
		cur = cur.child[bit]
	}
}
func (t *trie) maxXor(v int) int {
	cur := t
	ans := 0
	for i := bLen - 1; i >= 0; i-- {
		bit := v >> i & 1
		if cur.child[bit^1] != nil {
			ans |= 1 << i
			bit ^= 1
		}
		cur = cur.child[bit]
	}
	return ans
}

func maximizeXor(nums []int, queries [][]int) []int {
	sort.Ints(nums) // 排序
	n := len(queries)
	ans, idx := make([]int, n), make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool { // 排序 idx
		return queries[idx[i]][1] < queries[idx[j]][1]
	})

	root := &trie{} // Trie 树
	i := 0
	for _, j := range idx {
		x, m := queries[j][0], queries[j][1]
		for ; i < len(nums) && nums[i] <= m; i++ {
			root.insert(nums[i]) // 建树
		}
		if i == 0 {
			ans[j] = -1
		} else {
			ans[j] = root.maxXor(x) // 查询
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
