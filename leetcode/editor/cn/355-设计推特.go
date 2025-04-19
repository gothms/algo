package main

import (
	"container/heap"
	"fmt"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type Twitter struct {
	tw    map[int]*listNode // 优化为 map[int][]*listNode
	mf    map[int]map[int]struct{}
	h     *hp
	timer int
}

func Constructor() Twitter {
	return Twitter{
		tw:    make(map[int]*listNode),
		mf:    make(map[int]map[int]struct{}),
		h:     new(hp),
		timer: 0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if this.mf[userId] == nil {
		this.mf[userId] = map[int]struct{}{userId: {}} // 关注自己
	}
	this.timer++
	cur := &listNode{this.tw[userId], tweetId, this.timer}
	this.tw[userId] = cur
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	ans := make([]int, 0, 10)
	for user := range this.mf[userId] { // 收集关注的人的推的链表
		list := this.tw[user]
		if list != nil {
			heap.Push(this.h, list)
		}
	}
	for i := 0; i < 10 && this.h.Len() > 0; i++ { // 尝试获得 10 条
		top := heap.Pop(this.h).(*listNode)
		ans = append(ans, top.tid)
		if top.next != nil {
			heap.Push(this.h, top.next)
		}
	}
	*this.h = (*this.h)[:0] // 清空缓存
	return ans
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.mf[followerId] == nil {
		this.mf[followerId] = map[int]struct{}{followerId: {}} // 关注自己
	}
	this.mf[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.mf[followerId], followeeId)
}

type listNode struct {
	next *listNode
	tid  int
	time int
}

type hp []*listNode

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].time > h[j].time }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { *h = append(*h, x.(*listNode)) }
func (h *hp) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
//leetcode submit region end(Prohibit modification and deletion)
