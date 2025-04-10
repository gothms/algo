package main

import (
	"fmt"
	"slices"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findLadders(beginWord string, endWord string, wordList []string) []string {
	// lc：每次匹配 wordList，检查不同字符数是否为 1

	// bfs
	if len(beginWord) != len(endWord) || beginWord == endWord {
		return nil
	}

	memo := make(map[string]int)
	for i, s := range wordList {
		memo[s] = i
	}
	//delete(memo, beginWord)
	n := len(wordList)
	memo[beginWord] = n // 服务于拓扑

	topo := make([]int, n+1) // 记录“拓扑排序”，且 n+1
	for i := range topo {
		topo[i] = -1
	}

	q := []string{beginWord}
	idx := -1
out:
	for len(q) > 0 {
		s := []rune(q[0])
		q = q[1:]
		k := memo[string(s)]
		for i, c := range s {
			for ch := 'a'; ch <= 'z'; ch++ {
				if ch == c {
					continue
				}
				s[i] = ch
				if j, ok := memo[string(s)]; ok && topo[j] == -1 { // 没访问过 wordList[j]
					topo[j] = k // 拓扑排序
					if string(s) == endWord {
						idx = j
						break out
					}
					//delete(memo, string(s))
					q = append(q, string(s))
				}
			}
			s[i] = c
		}
	}
	if idx == -1 {
		return nil
	}
	ans := make([]string, 0)
	for ; idx < n; idx = topo[idx] {
		ans = append(ans, wordList[idx])
	}
	ans = append(ans, beginWord)
	slices.Reverse(ans)
	return ans

	// 个人：dfs
	//if len(beginWord) != len(endWord) || beginWord == endWord {
	//	return nil
	//}
	//
	//memo := make(map[string]int)
	//for i, s := range wordList {
	//	memo[s] = i
	//}
	//delete(memo, beginWord)
	//
	//path := make([]int, 0)
	////vis := make([]bool, len(wordList))	// 可以用 memo 兼职 vis
	//var dfs func([]rune) bool
	//dfs = func(s []rune) bool {
	//	if string(s) == endWord {
	//		return true
	//	}
	//	for i, c := range s {
	//		for ch := 'a'; ch <= 'z'; ch++ { // 尝试替换字符
	//			if ch == c {
	//				continue
	//			}
	//			s[i] = ch
	//			if j, ok := memo[string(s)]; ok /*&& !vis[j]*/ { // 替换成功
	//				//vis[j] = true          // 标记访问
	//				delete(memo, string(s))
	//				path = append(path, j) // 记录路径
	//				if dfs(s) {
	//					return true
	//				}
	//				path = path[:len(path)-1] // 回溯
	//			}
	//		}
	//		s[i] = c // 回溯
	//	}
	//	return false
	//}
	//
	//ans := make([]string, 0, len(path))
	//if dfs([]rune(beginWord)) {
	//	for _, i := range slices.Backward(path) {
	//		ans = append(ans, wordList[i])
	//	}
	//	ans = append(ans, beginWord)
	//	slices.Reverse(ans)
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
