package str

import (
	"bufio"
	"fmt"
	"os"
)

/*
https://github.com/ctdk/go-trie

https://github.com/appscodelabs/release-automaton

*/

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
				msg := fmt.Sprintf("匹配起始下标 %d ，长度 %d", pos, cur.length)
				fmt.Println(msg)
			}
		}
	}
}

//type acTrie struct {
//}

const ALPHABET_SIZE = 26

var tot int
var tr [][ALPHABET_SIZE]int
var e []int
var fail []int

// ====================AC 自动机，简单版====================
func insert(s []byte) {
	u := 0
	for i := 1; i < len(s); i++ {
		if tr[u][s[i]-'a'] == 0 {
			tot++
			tr[u][s[i]-'a'] = tot
		}
		u = tr[u][s[i]-'a']
	}
	e[u]++
}
func build() {
	q := make([]int, 0)
	for _, v := range tr[0] {
		if v != 0 {
			q = append(q, v)
		}
	}
	for ; len(q) > 0; q = q[1:] {
		u := q[0]
		for i := range tr[u] {
			if tr[u][i] != 0 {
				fail[tr[u][i]] = tr[fail[u]][i]
				q = append(q, tr[u][i])
			} else {
				tr[u][i] = tr[fail[u]][i]
			}
		}
	}
}
func query(t []byte) int {
	u, res := 0, 0
	for i := range t {
		u = tr[u][t[i]-'a']
		for j := u; j != 0 && e[j] != -1; j = fail[j] {
			res += e[j]
			e[j] = -1
		}
	}
	return res
}

// ====================AC 自动机，加强版====================

const (
	N  = 156
	L  = 1e6 + 6
	SZ = N * 80
)

var (
	tot    int
	tr     [SZ][26]int
	fail   [SZ]int
	idx    [SZ]int
	val    [SZ]int
	cnt    [N]int
	n      int
	s      [N][100]byte
	t      [L]byte
	queue  = make([]int, 0)
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func initAC() {
	tot = 0
	for i := 0; i < SZ; i++ {
		fail[i] = 0
		val[i] = 0
		idx[i] = 0
		for j := 0; j < 26; j++ {
			tr[i][j] = 0
		}
	}
	for i := 0; i < N; i++ {
		cnt[i] = 0
	}
}

func insert(s []byte, id int) {
	u := 0
	for i := 1; s[i] != 0; i++ {
		if tr[u][s[i]-'a'] == 0 {
			tot++
			tr[u][s[i]-'a'] = tot
		}
		u = tr[u][s[i]-'a']
	}
	idx[u] = id
}

func build() {
	for i := 0; i < 26; i++ {
		if tr[0][i] != 0 {
			queue = append(queue, tr[0][i])
		}
	}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for i := 0; i < 26; i++ {
			if tr[u][i] != 0 {
				fail[tr[u][i]] = tr[fail[u]][i]
				queue = append(queue, tr[u][i])
			} else {
				tr[u][i] = tr[fail[u]][i]
			}
		}
	}
}

func query(t []byte) int {
	u, res := 0, 0
	for i := 1; t[i] != 0; i++ {
		u = tr[u][t[i]-'a']
		for j := u; j != 0; j = fail[j] {
			val[j]++
		}
	}
	for i := 0; i <= tot; i++ {
		if idx[i] != 0 {
			if val[i] > res {
				res = val[i]
			}
			cnt[idx[i]] = val[i]
		}
	}
	return res
}

//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	for scanner.Scan() {
//		fmt.Sscanf(scanner.Text(), "%d", &n)
//		if n == 0 {
//			break
//		}
//		initAC()
//		for i := 1; i <= n; i++ {
//			scanner.Scan()
//			copy(s[i][:], scanner.Text())
//			insert(s[i][:], i)
//		}
//		build()
//		scanner.Scan()
//		copy(t[:], scanner.Text())
//		x := query(t[:])
//		fmt.Printf("%d\n", x)
//		for i := 1; i <= n; i++ {
//			if cnt[i] == x {
//				fmt.Printf("%s\n", s[i][1:])
//			}
//		}
//	}
//}

// ====================AC 自动机，二次加强版====================
