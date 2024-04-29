package str

import "fmt"

var acRoot = &acNode{} // 模拟 Tire 树的根

type acNode struct {
	data         byte
	children     [26]*acNode // 字符集只包含a~z这26个字符
	isEndingChar bool        // 结尾字符为true
	length       int         // 当isEndingChar=true时，记录模式串长度。应初始化为 -1
	fail         *acNode     // 失败指针
}

// buildFailurePointer 构建失败指针
// parent 节点的失败指针已知，在 Trie 树中新添加了 parent 的 child 节点，则 child 节点的失败指针未知
// 问题：
// 假设第一个字符串是 abcde，则 5 个字符的失败指针都是根节点 acRoot
// 当新添加字符串 bcde 时，不是应该反向构建 abcde 的 e 节点的失败指针，是 bcde 的 e 节点吗？
func buildFailurePointer(parent *acNode) {
	queue := []*acNode{parent}
	for ; len(queue) > 0; queue = queue[1:] {
		p := queue[0] // 已知 p 的失败指针
		for i := 0; i < 26; i++ {
			pc := p.children[i] // 寻找 p 的子节点的失败指针
			if pc == nil {
				continue
			}
			if p == acRoot { // 根节点的失败指针指向自己
				p.fail = acRoot
			} else {
				q := p.fail // 已知 p 的失败指针是 q
				for q != nil {
					qc := q.children[pc.data-'a'] // qc 匹配所求节点 pc
					if qc != nil {                // 匹配成功，pc 的失败指针是 qc
						pc.fail = qc // 找到最长匹配前缀，qc 为最长匹配前缀的尾字符
						break
					}
					q = q.fail // 匹配失败，继续“往上”寻找 q 的失败指针（类似 KMP 算法求 prefix 的过程）
				}
				if q == nil { // “终极”失败（即没有公共前缀），pc 的失败指针指向根节点 acRoot
					pc.fail = acRoot
				}
			}
			queue = append(queue, pc) // 继续匹配 p 的子孙节点的失败指针
		}
	}
}

// match 在 AC 自动机上匹配主串
// text 主串
func match(text string) {
	n := len(text)
	p := acRoot
	for i := 0; i < n; i++ {
		idx := text[i] - 'a'
		// 1.“累加”模式串：在已“累积”的 text 的子串（以字符 text[i] 结尾）的基础上，增加字符 text[i]
		for p.children[idx] == nil && p != acRoot { // 到此，该路径对应的前缀不匹配
			p = p.fail // 匹配更短的模式串，失败指针发挥作用
		}
		p = p.children[idx]
		// 2.“累加”的模式串匹配失败：模式串长度归 0（text[i] 也不包含），下一轮从新“累加”模式串
		if p == nil { // 没有匹配的模式串，下一个字符从根节点 acRoot 开始匹配
			p = acRoot
			continue // 可加 continue
		}
		// 3.尝试匹配成功：找出“累加”模式串可能匹配的所有后缀，即检测一系列失败指针为结尾的路径是否是模式串
		for cur := p; cur != acRoot; cur = cur.fail { // 找出所有可匹配的模式串
			if cur.isEndingChar {
				pos := i - cur.length + 1 // 在 text 中的起始位置
				msg := fmt.Sprintf("匹配起始下标 %d ，长度 %d\n", pos, cur.length)
				fmt.Println(msg)
			}
		}
	}
}
