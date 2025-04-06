package main

import "fmt"

func main() {
	// ["time","ati","bell"] 14
	// ["me","time"] 5
	// ["ti","time"] 8
	words := []string{"me", "time"} // 5
	encoding := minimumLengthEncoding(words)
	fmt.Println(encoding)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumLengthEncoding(words []string) int {
	// trie：倒序
	type trie struct {
		child [26]*trie
		cnt   int
		s     string
	}

	memo := make(map[*trie]struct{})
	root := &trie{}
	insertAndCheck := func(s string) { // 倒序插入字典中
		cur := root
		for i := len(s) - 1; i >= 0; i-- {
			idx := int(s[i] - 'a')
			if cur.child[idx] == nil {
				cur.child[idx] = &trie{}
			}
			cur.cnt++ // 标记字符出现的次数
			cur = cur.child[idx]
		}
		cur.s = s              // 记录对应的字符串
		memo[cur] = struct{}{} // 放入map中
	}
	for _, w := range words {
		insertAndCheck(w)
	}
	ans := 0
	for t := range memo {
		if t.cnt == 0 { // 说明 t.s 没有匹配的后缀
			ans += len(t.s) + 1
		}
	}
	return ans

	// hash
	//memo := make(map[string]struct{})
	//ans := 0
	//for _, w := range words {
	//	memo[w] = struct{}{}
	//}
	//for _, w := range words {
	//	for i := 1; i < len(w); i++ {
	//		delete(memo, w[i:])
	//	}
	//}
	//for w, _ := range memo {
	//	ans += len(w) + 1
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
