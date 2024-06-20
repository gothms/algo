package main

import "fmt"

func main() {
	//con := Constructor()
	//con.BuildDict([]string{"hello", "hallo", "leetcode"})
	//b := con.Search("hello")
	//fmt.Println(b)
	//b = con.Search("hhllo")
	//fmt.Println(b)
	//b = con.Search("hell")
	//fmt.Println(b)
	//b = con.Search("leetcoded") // [true,true,false,false]
	//fmt.Println(b)

	con := Constructor()
	con.BuildDict([]string{"hello", "hallo", "leetcode", "judge"})
	b := con.Search("hello")
	fmt.Println(b)
	b = con.Search("hallo")
	fmt.Println(b)
	b = con.Search("hell")
	fmt.Println(b)
	b = con.Search("leetcodd")
	fmt.Println(b)
	b = con.Search("aaaaa") // [true,true,false,true,false]
	fmt.Println(b)
	b = con.Search("hallz")
	fmt.Println(b)

	//con := Constructor()
	//con.BuildDict([]string{"a", "b", "ab", "abc"})
	//b := con.Search("bb")
	//fmt.Println(b)
	//b = con.Search("aa")
	//fmt.Println(b)
	//b = con.Search("bbc") // [true,true,true]
	//fmt.Println(b)
}

// leetcode submit region begin(Prohibit modification and deletion)
type trie676 struct {
	children [26]*trie676
	isEnd    bool
}

type MagicDictionary struct {
	root *trie676
}

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.root = &trie676{}
	for _, s := range dictionary {
		cur := this.root
		for _, c := range s {
			i := c - 'a'
			if cur.children[i] == nil {
				cur.children[i] = &trie676{}
			}
			cur = cur.children[i]
		}
		cur.isEnd = true
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	cur := this.root
	for idx, c := range searchWord {
		i := int(c - 'a')
		for j, child := range cur.children {
			if child != nil && j != i && match(idx+1, child, searchWord) {
				return true
			}
		}
		if cur = cur.children[i]; cur == nil {
			break
		}
	}
	return false
}

func match(i int, cur *trie676, s string) bool {
	for ; i < len(s) && cur != nil; i++ {
		cur = cur.children[s[i]-'a']
	}
	return cur != nil && cur.isEnd
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
//leetcode submit region end(Prohibit modification and deletion)

//type MagicDictionary struct {
//	root *trie676
//}
//type trie676 struct {
//	children [26]*trie676
//	isEnd    bool
//}
//
//func Constructor() MagicDictionary {
//	return MagicDictionary{&trie676{}}
//}
//
//func (this *MagicDictionary) BuildDict(dictionary []string) {
//	for _, s := range dictionary { // 构建 trie 树
//		cur := this.root
//		for _, c := range s {
//			i := c - 'a'
//			if cur.children[i] == nil {
//				cur.children[i] = &trie676{}
//			}
//			cur = cur.children[i]
//		}
//		cur.isEnd = true
//	}
//}
//
//func (this *MagicDictionary) Search(searchWord string) bool {
//	// 本题题意：在 searchWord 已经存在的情况下，还可以查询其他的可匹配字符串
//	// 那么如果不可以呢？则先用 memo 记录存在的字符串
//	// 其他做法：lc dfs
//	cur := this.root
//	for idx, c := range searchWord {
//		i := int(c - 'a')
//		for j, child := range cur.children { // 枚举：字符 j+'a' 不同
//			if j != i && child != nil && this.check(searchWord, idx+1, child) {	// 注意：idx+1
//				return true // 后缀可匹配
//			}
//		}
//		cur = cur.children[i]
//		if cur == nil { // 已失败
//			break
//		}
//	}
//	return false
//}
//func (this *MagicDictionary) check(searchWord string, idx int, cur *trie676) bool {
//	n := len(searchWord)
//	for ; idx < n; idx++ {
//		i := searchWord[idx] - 'a'
//		if cur.children[i] == nil {
//			return false
//		}
//		cur = cur.children[i]
//	}
//	return cur.isEnd
//}
