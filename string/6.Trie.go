package str

/*
lc
	676
*/

type trieNode struct {
	data         byte
	children     [26]*trieNode
	isEndingChar bool
}

// insert 在 Trie 树中插入一个字符串
func insert(root *trieNode, s string) {
	cur := root
	for _, c := range s {
		i := c - 'a'
		if cur.children[i] == nil {
			cur.children[i] = &trieNode{}
			// TODO 联动 AC 自动机，从此开始构建失败指针
		}
		cur = cur.children[i]
	}
	cur.isEndingChar = true
}

// find 在 Trie 树中查找一个字符串
func find(root *trieNode, pattern string) bool {
	cur := root
	for _, c := range pattern {
		i := c - 'a'
		if cur.children[i] == nil {
			return false // 匹配失败
		}
		cur = cur.children[i]
	}
	return cur.isEndingChar // isEndingChar 可能是 false
}
