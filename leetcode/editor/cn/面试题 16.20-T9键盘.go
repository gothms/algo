package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func getValidT9Words(num string, words []string) []string {
	// lc
	match := [26]int{0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 7, 7, 7, 7}
	can := func(s string) bool {
		for i, c := range s {
			if int(num[i]-'2') != match[c-'a'] {
				return false
			}
		}
		return true
	}
	n := len(num)
	ans := make([]string, 0)
	for _, w := range words {
		if len(w) == n && can(w) {
			ans = append(ans, w)
		}
	}
	return ans

	// 个人
	//phone := [][]byte{
	//	{},
	//	{},
	//	{'a', 'b', 'c'},
	//	{'d', 'e', 'f'},
	//	{'g', 'h', 'i'},
	//	{'j', 'k', 'l'},
	//	{'m', 'n', 'o'},
	//	{'p', 'q', 'r', 's'},
	//	{'t', 'u', 'v'},
	//	{'w', 'x', 'y', 'z'},
	//}
	//type trie struct {
	//	child [26]*trie
	//	s     string
	//}
	//root := &trie{}
	//insert := func(s string) {
	//	cur := root
	//	for _, c := range s {
	//		i := int(c - 'a')
	//		if cur.child[i] == nil {
	//			cur.child[i] = &trie{}
	//		}
	//		cur = cur.child[i]
	//	}
	//	cur.s = s
	//}
	//for _, w := range words {
	//	insert(w)
	//}
	//q := []*trie{root}
	//for _, v := range num {
	//	temp := make([]*trie, 0)
	//	for _, c := range phone[int(v-'0')] {
	//		i := c - 'a'
	//		for _, t := range q {
	//			if t.child[i] != nil {
	//				temp = append(temp, t.child[i])
	//			}
	//		}
	//	}
	//	if len(temp) == 0 {
	//		return nil
	//	}
	//	q = temp
	//}
	//ans := make([]string, 0)
	//for _, t := range q {
	//	if t.s != "" {
	//		ans = append(ans, t.s)
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
