package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func multiSearch(big string, smalls []string) [][]int {
	// AC 自动机

	// Trie
	type trie struct {
		child [26]*trie
		is    bool
	}
	root := &trie{}
	insert := func(s string) {
		cur := root
		for _, c := range s {
			i := c - 'a'
			if cur.child[i] == nil {
				cur.child[i] = &trie{}
			}
			cur = cur.child[i]
		}
		cur.is = true
	}
	memo := make(map[string]int)
	for i, s := range smalls {
		insert(s)
		memo[s] = i
	}

	ans := make([][]int, len(smalls))
	for i := range big {
		cur := root
		for j, c := range big[i:] {
			idx := c - 'a'
			if cur.child[idx] == nil {
				break
			}
			cur = cur.child[idx]
			if cur.is {
				k := memo[big[i:i+j+1]]
				ans[k] = append(ans[k], i)
			}
		}
	}
	return ans
	
	// 字符串匹配算法：如 KMP
}

//leetcode submit region end(Prohibit modification and deletion)
