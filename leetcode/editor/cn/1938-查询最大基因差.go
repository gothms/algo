package main

import "fmt"

func main() {
	parents := []int{-1, 0, 1, 1}
	queries := [][]int{{0, 2}, {3, 2}, {2, 5}}
	difference := maxGeneticDifference(parents, queries)
	fmt.Println(difference)
}

// leetcode submit region begin(Prohibit modification and deletion)
const L = 17

type trie struct {
	child [2]*trie
	cnt   int
}

func (t *trie) insert(v int) {
	cur := t
	for i := L; i >= 0; i-- {
		bit := v >> i & 1
		if cur.child[bit] == nil {
			cur.child[bit] = &trie{}
		}
		cur = cur.child[bit]
		cur.cnt++
	}
}
func (t *trie) maxXor(v int) int {
	cur := t
	ans := 0
	for i := L; i >= 0; i-- {
		bit := v >> i & 1
		if c := cur.child[bit^1]; c != nil && c.cnt > 0 {
			ans |= 1 << i
			bit ^= 1
		}
		cur = cur.child[bit]
	}
	return ans
}
func (t *trie) delete(v int) {
	cur := t
	for i := L; i >= 0; i-- {
		cur = cur.child[v>>i&1]
		cur.cnt--
	}
}

func maxGeneticDifference(parents []int, queries [][]int) []int {
	n, m := len(parents), len(queries)
	g := make([][]int, n) // 邻接表
	root := -1
	for i, v := range parents {
		if v == -1 {
			root = i
		} else {
			g[v] = append(g[v], i)
		}
	}
	query := make(map[int][]int) // 查询分组
	for i, q := range queries {
		query[q[0]] = append(query[q[0]], i)
	}

	ans := make([]int, m)
	t := &trie{}
	var dfs func(int)
	dfs = func(i int) {
		t.insert(i) // 插入节点
		for _, idx := range query[i] {
			ans[idx] = t.maxXor(queries[idx][1]) // 查询最大值
		}
		for _, j := range g[i] {
			dfs(j)
		}
		t.delete(i) // 删除节点，回溯
	}
	dfs(root)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
