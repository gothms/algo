package main

import "fmt"

func main() {
	//words := []string{"cd", "f", "kl"}
	//con := Constructor(words)
	//fmt.Println(con)
	//query := con.Query('a')
	//fmt.Println(query)
	//query = con.Query('d')
	//fmt.Println(query)
	//query = con.Query('f')
	//fmt.Println(query)
	//query = con.Query('k')
	//fmt.Println(query)
	//query = con.Query('l')
	//fmt.Println(query)

	words := []string{"ab", "ba", "aaab", "abab", "baa"}
	words = []string{"ab", "ba", "aaab", "abab", "baa", "a"}
	con := Constructor(words)
	query := con.Query('a')
	fmt.Println(query)
	query = con.Query('a')
	fmt.Println(query)
	query = con.Query('a')
	fmt.Println(query)
	query = con.Query('a')
	fmt.Println(query)
	query = con.Query('a')
	fmt.Println(query)
	query = con.Query('b')
	fmt.Println(query)
	query = con.Query('a')
	fmt.Println(query)

	s := "aaaaabababbbababbbbababaaabaaa"
	s = "a"
	qs := con.QueryString(s)
	fmt.Println(qs)
}

// leetcode submit region begin(Prohibit modification and deletion)
type StreamChecker struct {
	tot  int
	tr   [][26]int
	e    []int
	fail []int
	u    int
}

func Constructor(words []string) StreamChecker {
	sc := StreamChecker{tr: make([][26]int, 1), e: make([]int, 1), fail: make([]int, 1), u: -1}
	for _, word := range words {
		sc.insert(word)
	}
	return sc
}
func (sc *StreamChecker) insert(s string) {
	u := 0
	for i := 1; i < len(s); i++ {
		idx := s[i] - 'a'
		if sc.tr[u][idx] == 0 {
			sc.tr = append(sc.tr, [26]int{})
			sc.e = append(sc.e, 0)
			sc.fail = append(sc.fail, 0)
			sc.tot++
			sc.tr[u][idx] = sc.tot
		}
		u = sc.tr[u][idx]
	}
	sc.e[u]++
}
func (sc *StreamChecker) build() {
	q := make([]int, 0)
	for _, v := range sc.tr[0] {
		if v != 0 {
			q = append(q, v)
		}
	}
	for ; len(q) > 0; q = q[1:] {
		u := q[0]
		for i := range sc.tr[u] {
			if sc.tr[u][i] != 0 {
				sc.fail[sc.tr[u][i]] = sc.tr[sc.fail[u]][i]
				q = append(q, sc.tr[u][i])
			} else {
				sc.tr[u][i] = sc.tr[sc.fail[u]][i]
			}
		}
	}
}

func (sc *StreamChecker) Query(letter byte) bool {
	if sc.u == -1 {
		sc.u = 0
		return false
	}
	res := 0
	sc.u = sc.tr[sc.u][letter-'a']
	for j := sc.u; j != 0; /*&& sc.e[j] != -1*/ j = sc.fail[j] {
		res += sc.e[j]
		//sc.e[j] = -1
	}
	return res > 0
}

func (sc *StreamChecker) QueryString(t string) int {
	u, res := 0, 0
	for i := 1; i < len(t); i++ {
		u = sc.tr[u][t[i]-'a']
		for j := u; j != 0 && sc.e[j] != -1; j = sc.fail[j] {
			res += sc.e[j]
			sc.e[j] = -1
		}
	}
	return res
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
//leetcode submit region end(Prohibit modification and deletion)

//type StreamChecker struct {
//	root, cur *acTrie
//}
//type acTrie struct {
//	children [26]*acTrie
//	isEnd    bool
//	fail     *acTrie
//}
//
//func Constructor(words []string) StreamChecker {
//	root := &acTrie{}
//	for _, word := range words {
//		cur := root
//		for _, c := range word {
//			i := c - 'a'
//			if cur.children[i] == nil {
//				cur.children[i] = &acTrie{}
//			}
//			cur = cur.children[i]
//		}
//		cur.isEnd = true
//	}
//	root.fail = root
//	queue := make([]*acTrie, 0)
//	for i, child := range root.children {
//		if child != nil {
//			child.fail = root
//			queue = append(queue, child)
//		} else {
//			root.children[i] = root
//		}
//	}
//	for ; len(queue) > 0; queue = queue[1:] {
//		cur := queue[0]
//		cur.isEnd = cur.isEnd || cur.fail.isEnd
//		for i, child := range cur.children {
//			fail := cur.fail.children[i]
//			if child != nil {
//				child.fail = fail
//				queue = append(queue, child)
//			} else {
//				cur.children[i] = fail
//			}
//		}
//	}
//
//	//queue := []*acTrie{root}
//	//for ; len(queue) > 0; queue = queue[1:] {
//	//	cur := queue[0]
//	//	cur.isEnd = cur.isEnd || cur.fail.isEnd
//	//	for i, child := range cur.children {
//	//		if child != nil {
//	//			if fail := cur.fail.children[i]; fail != nil {
//	//				child.fail = fail
//	//				queue = append(queue, child)
//	//			} else {
//	//				child.fail = root
//	//			}
//	//		}
//	//	}
//	//}
//	return StreamChecker{root, root}
//}
//func (this *StreamChecker) Query(letter byte) bool {
//	this.cur = this.cur.children[letter-'a']
//	if this.cur == nil {
//		this.cur = this.root
//	}
//	return this.cur.isEnd
//}


