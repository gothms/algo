package main

import "fmt"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type AuthenticationManager struct {
	ttl  int
	m    map[string]*node
	h, t *node
}
type node struct {
	pre, next *node
	key       string
	val       int
}

func Constructor(timeToLive int) AuthenticationManager {
	n := &node{}
	n.pre, n.next = n, n
	return AuthenticationManager{ttl: timeToLive, m: make(map[string]*node), h: n, t: n}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	var cur *node
	if cur = this.m[tokenId]; cur == nil {
		cur = &node{key: tokenId}
	} else {
		cur.pre.next, cur.next.pre = cur.next, cur.pre // 移除节点
	}
	cur.val = currentTime + this.ttl
	this.t.pre.next, this.t.pre, cur.pre, cur.next = cur, cur, this.t.pre, this.t // 插入节点
	this.m[tokenId] = cur
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if cur, ok := this.m[tokenId]; ok && cur.val > currentTime {
		cur.pre.next, cur.next.pre = cur.next, cur.pre
		cur.val = currentTime + this.ttl
		this.t.pre.next, this.t.pre, cur.pre, cur.next = cur, cur, this.t.pre, this.t
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	for cur := this.h.next; cur != this.t && cur.val <= currentTime; {
		delete(this.m, cur.key)
		cur.pre.next, cur.next.pre, cur = cur.next, cur.pre, cur.next
	}
	return len(this.m)
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
//leetcode submit region end(Prohibit modification and deletion)
