//给定一份单词的清单，设计一个算法，创建由字母组成的面积最大的矩形，其中每一行组成一个单词(自左向右)，每一列也组成一个单词(自上而下)。不要求这些单词在清单
//里连续出现，但要求所有行等长，所有列等高。
//
// 如果有多个面积最大的矩形，输出任意一个均可。一个单词可以重复使用。
//
// 示例 1:
//
// 输入: ["this", "real", "hard", "trh", "hea", "iar", "sld"]
//输出:
//[
//   "this",
//   "real",
//   "hard"
//]
//
// 示例 2:
//
// 输入: ["aa"]
//输出: ["aa","aa"]
//
// 说明：
//
//
// words.length <= 1000
// words[i].length <= 100
// 数据保证单词足够随机
//
//
// Related Topics 字典树 数组 字符串 回溯 👍 47 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"this", "real", "hard", "trh", "hea", "iar", "sld"}
	rectangle := maxRectangle(words)
	fmt.Println(rectangle)
}

// leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	next [26]*Trie
	end  bool
}

func insert(root *Trie, str string) {
	next := root
	for i := range str {
		key := str[i] - 'a'
		if next.next[key] == nil {
			next.next[key] = &Trie{}
		}
		next = next.next[key]
	}
	next.end = true
}

func maxRectangle(words []string) []string {
	// bfs + Trie
	root := &Trie{}
	group := map[int][]int{}
	for i, word := range words {
		insert(root, word)                             // 构建 Trie 树
		group[len(word)] = append(group[len(word)], i) // 长度相同的单词为一组
	}

	lengths := make([]int, len(group)) // 单词分组后的长度数组
	i := 0
	for _, v := range group {
		lengths[i] = len(words[v[0]])
		i++
	}
	sort.Slice(lengths, func(i, j int) bool {
		return lengths[i] > lengths[j]
	})

	type pair struct { // bfs 的元素
		nodes   []*Trie
		indexes []int
	}
	resIndexes := []int{} // 最大面积的单词（索引）
	s := 0                // 最大面积
	//从最大长度开始
	for _, l := range lengths {
		if l*l <= s { // 剪枝
			break
		}
		start := &pair{nodes: make([]*Trie, l)} // 长度为 l，则初始为 l 个 Trie 树的根节点
		for i := 0; i < l; i++ {
			start.nodes[i] = root // 代表 l 个任意节点
		}
		queue := []*pair{start}
		for len(queue) > 0 { // bfs
			q := queue[0]
			queue = queue[1:]
			allEnd := true // l 个 path 是否统一都能成为单词
		nextWord:
			for _, idx := range group[l] { // 寻找下一个（下一层）的单词
				word := words[idx]
				var nodes []*Trie
				for col := range word {
					node := q.nodes[col].next[word[col]-'a'] // 在上一层节点，往下一层寻找
					if node == nil {
						continue nextWord
					}
					nodes = append(nodes, node)
					allEnd = allEnd && node.end
				} // 寻找成功

				indexes := make([]int, len(q.indexes)+1)
				copy(indexes, q.indexes)
				indexes[len(indexes)-1] = idx
				queue = append(queue, &pair{nodes, indexes}) // 入列
				if !allEnd {
					continue
				}
				//计算面积
				if _s := (len(q.indexes) + 1) * l; _s > s {
					s = _s
					resIndexes = indexes
				}
			}
		}
	}

	res := make([]string, len(resIndexes)) // 结果集
	for i, idx := range resIndexes {
		res[i] = words[idx]
	}

	return res

	// dfs + Trie
}

//leetcode submit region end(Prohibit modification and deletion)
