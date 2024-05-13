package main

import (
	"fmt"
)

func main() {
	edit := Constructor()
	edit.AddText("bxyackuncqzcqo")
	left := edit.CursorLeft(12)
	fmt.Println(left)
	edit.DeleteText(3)
	left = edit.CursorLeft(5)
	fmt.Println(left)
	edit.AddText("osdhyvqxf")
	right := edit.CursorRight(10)
	fmt.Println(right)

	//edit.CursorLeft(1)
	//edit.CursorRight(4)
	//edit.DeleteText(3)

	s := []byte("bxyackuncqzcqo") // TODO
	t := "osdhyvqxf"
	s = []byte("abc")
	t = "12"
	cur := 0
	fmt.Printf("%p,%p\n", s, append(append(s[:cur], t...), s[cur:]...))
	fmt.Printf("%p,%p\n", s, append(append(s[:cur], t...), s[cur:]...))
	s = append(append(s[:cur], t...), s[cur:]...)
	fmt.Printf("%p\n", s)
	fmt.Printf("s: %s\n", s) // osdhyvqxfosdhyvqxfqzcqo
}

// leetcode submit region begin(Prohibit modification and deletion)

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */
//leetcode submit region end(Prohibit modification and deletion)

// ====================Splay树====================

// ====================对顶栈====================

type TextEditor struct {
	l, r []byte
}

func Constructor() TextEditor {
	return TextEditor{}
}
func (this *TextEditor) AddText(text string) {
	this.l = append(this.l, text...)
}
func (this *TextEditor) DeleteText(k int) int {
	k = min(k, len(this.l))
	this.l = this.l[:len(this.l)-k]
	return k
}
func (this *TextEditor) CursorLeft(k int) string {
	move(k, &this.l, &this.r)
	return string(this.l[len(this.l)-min(10, len(this.l)):])
}
func (this *TextEditor) CursorRight(k int) string {
	move(k, &this.r, &this.l)
	return string(this.l[len(this.l)-min(10, len(this.l)):])
}
func move(k int, f, t *[]byte) {
	last := len(*f)
	k = min(k, last)
	last--
	for i := 0; i < k; i++ {
		*t = append(*t, (*f)[last])
		last--
	}
	last++
	*f = (*f)[:last]
}

// ====================双向链表====================

//type TextEditor struct {
//	cur, n int      // 光标位置 和 总长度
//	i      int      // 光标在结点的位置
//	s      *lCursor // 光标所在的结点
//}
//type lCursor struct {
//	pre, next *lCursor
//	text      string // 结点的字符串值
//}
//
//func Constructor() TextEditor {
//	l := &lCursor{}
//	l.pre, l.next = l, l // 循环双向链表
//	return TextEditor{s: l}
//}
//
//func (this *TextEditor) AddText(text string) {
//	m := len(text)
//	l := &lCursor{text: text}
//	if this.cur == this.n || this.i == len(this.s.text) { // 光标在文本末尾，则直接追加结点。光标在结点末尾，则在结点后插入结点
//		this.s.next, l.pre, this.s.next.pre, l.next = l, this.s, l, this.s.next
//	} else { // 否则将结点拆为 2 个，并在中间插入结点
//		st := this.s.text
//		ts := st[this.i:]
//		this.s.text = st[:this.i]
//		div := &lCursor{text: ts}                                        // 拆为 2 个
//		this.s.next, l.pre, div.pre, l.next, this.s.next.pre, div.next = // 插入结点
//			l, this.s, l, div, div, this.s.next
//	}
//	this.s, this.i = l, m // 更新光标所在结点和位置
//	this.cur += m
//	this.n += m
//}
//func (this *TextEditor) DeleteText(k int) int {
//	if this.cur == 0 { // fast fail
//		return 0
//	}
//	k = min(k, this.cur)
//	m := len(this.s.text)
//	if k < this.i { // 在本结点删除 k 个字符
//		if this.i != m { // 光标所在位置不是结点末尾，则将本结点拆为 2 个结点
//			ts := this.s.text[this.i:]
//			div := &lCursor{text: ts} // 拆为 2 个
//			this.s.next, div.pre, this.s.next.pre, div.next = div, this.s, div, this.s.next
//		}
//		this.s.text = this.s.text[:this.i-k] // 更新本结点
//	} else {
//		counter := k - this.i
//		pre := this.s.pre
//		for counter >= len(pre.text) { // 找到光标的停止结点
//			counter -= len(pre.text)
//			pre = pre.pre
//		}
//		if this.i == m { // 本结点被全部删除
//			this.s = this.s.next
//		} else { // 本结点还剩有字符
//			this.s.text = this.s.text[this.i:]
//		}
//		pre.text = pre.text[:len(pre.text)-counter]
//		pre.next, this.s.pre = this.s, pre // 删除 0 或多个结点
//		this.s = pre
//	}
//	this.i = len(this.s.text) // 删除字符后，在结点的位置一定是末尾
//	this.cur -= k
//	this.n -= k
//	return k
//}
//func (this *TextEditor) CursorLeft(k int) string {
//	if this.cur == 0 { // fast path
//		return ""
//	}
//	k = min(k, this.cur) // 最多移动到 0
//	if k < this.i {      // 在本结点移动光标
//		this.i -= k
//	} else {
//		counter := k - this.i
//		pre := this.s.pre
//		for counter >= len(pre.text) { // 找到光标的停止结点
//			counter -= len(pre.text)
//			pre = pre.pre
//		}
//		this.s, this.i = pre, len(pre.text)-counter // 停止位置
//	}
//	this.cur -= k
//	return this.dfsCursor()
//}
//func (this *TextEditor) CursorRight(k int) string {
//	if this.cur == this.n { // fast path
//		return this.dfsCursor()
//	}
//	k = min(k, this.n-this.cur)
//	if counter := k - len(this.s.text) + this.i; counter <= 0 { // counter == 0：刚好在末尾
//		this.i += k
//	} else {
//		next := this.s.next
//		for counter > len(next.text) {
//			counter -= len(next.text)
//			next = next.next
//		}
//		this.s, this.i = next, counter
//	}
//	this.cur += k
//	return this.dfsCursor()
//}
//func (this *TextEditor) dfsCursor() string {
//	cnt := min(10, this.cur)
//	if this.i >= cnt {
//		return this.s.text[this.i-cnt : this.i] // 在本结点截取字符串
//	}
//	var sb strings.Builder
//	var dfs func(int, *lCursor)
//	dfs = func(c int, l *lCursor) { // 找到截取字符串的停止结点
//		if c <= len(l.text) {
//			sb.WriteString(l.text[len(l.text)-c:]) // 截取的停止结点
//			return
//		}
//		dfs(c-len(l.text), l.pre)
//		sb.WriteString(l.text) // “归”时写入字符串，保证了正序
//	}
//	dfs(cnt-this.i, this.s.pre)
//	sb.WriteString(this.s.text[:this.i])
//	return sb.String()
//}

// ====================切片====================

//type TextEditor struct {
//	text []uint8
//	cur  int
//}
//
//func Constructor() TextEditor {
//	return TextEditor{make([]uint8, 0), 0}
//}
//
//func (this *TextEditor) AddText(text string) {
//	n := len(this.text)
//	this.text = append(this.text, text...) // 先扩容
//	if this.cur != n {                     // 光标不在末尾，先挪动光标后的字符串，再插入 text
//		copy(this.text[this.cur+len(text):], this.text[this.cur:])
//		copy(this.text[this.cur:], text)
//	}
//	this.cur += len(text)
//}
//
//func (this *TextEditor) DeleteText(k int) int {
//	k = min(k, this.cur)                               // 保证有足够字符
//	copy(this.text[this.cur-k:], this.text[this.cur:]) // 删除
//	this.text = this.text[:len(this.text)-k]           // 截取掉末尾
//	this.cur -= k
//	return k
//}
//
//func (this *TextEditor) CursorLeft(k int) string {
//	k = min(k, this.cur)
//	this.cur -= k
//	cnt := min(10, this.cur)
//	return string(this.text[this.cur-cnt : this.cur])
//}
//
//func (this *TextEditor) CursorRight(k int) string {
//	k = min(k, len(this.text)-this.cur)
//	this.cur += k
//	cnt := min(10, this.cur)
//	return string(this.text[this.cur-cnt : this.cur])
//}
