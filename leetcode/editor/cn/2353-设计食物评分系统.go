package main

import (
	"cmp"
	"strings"
)

import (
	"github.com/emirpasic/gods/v2/trees/redblacktree"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)

// 当 lc 默认的依赖不是 v2 时，得手动导入 v2 包并提交到 lc
//import "github.com/emirpasic/gods/v2/trees/redblacktree"

type pair struct {
	s    string
	rate int
}
type FoodRatings struct {
	fMap map[string]*pair
	cMap map[string]*redblacktree.Tree[pair, struct{}]
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fm := make(map[string]*pair)
	cm := make(map[string]*redblacktree.Tree[pair, struct{}])
	fn := func(x, y pair) int {
		return cmp.Or(y.rate-x.rate, strings.Compare(x.s, y.s))
	}
	for i, f := range foods {
		fm[f] = &pair{cuisines[i], ratings[i]}
		if cm[cuisines[i]] == nil {
			cm[cuisines[i]] = redblacktree.NewWith[pair, struct{}](fn)
		}
		cm[cuisines[i]].Put(pair{f, ratings[i]}, struct{}{})
	}
	return FoodRatings{fMap: fm, cMap: cm}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	p := this.fMap[food]
	rbt := this.cMap[p.s]
	rbt.Remove(pair{food, p.rate})
	rbt.Put(pair{food, newRating}, struct{}{})
	p.rate = newRating
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	return this.cMap[cuisine].Left().Key.s
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
//leetcode submit region end(Prohibit modification and deletion)

// ==========懒删除堆==========

//type FoodRatings struct {
//	foodMap    map[string]*pair // pair{cuisine, rate}
//	cuisineMap map[string]*hp   // pair{foodName, rate}
//}
//
//func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
//	n := len(foods)
//	fm := make(map[string]*pair, n)
//	cm := make(map[string]*hp)
//	for i, f := range foods {
//		fm[f] = &pair{cuisines[i], ratings[i]}
//		if cm[cuisines[i]] == nil {
//			cm[cuisines[i]] = &hp{}
//		}
//		heap.Push(cm[cuisines[i]], pair{f, ratings[i]})
//	}
//	return FoodRatings{foodMap: fm, cuisineMap: cm}
//}
//
//func (this *FoodRatings) ChangeRating(food string, newRating int) {
//	p := this.foodMap[food]
//	heap.Push(this.cuisineMap[p.s], pair{food, newRating}) // 放入一个新的 pair{foodName, rate}
//	p.rate = newRating                                     // 更新 food pair{cuisine, rate}
//}
//
//func (this *FoodRatings) HighestRated(cuisine string) string {
//	h := this.cuisineMap[cuisine]
//	p := (*h)[0]
//	for ; p.rate != this.foodMap[p.s].rate; p = (*h)[0] { // 直到堆顶元素 pair.rate 是最新的
//		heap.Pop(h) // 淘汰已删除的 pair{foodName, rate}
//	}
//	return p.s
//}
//
//type pair struct {
//	s    string // 在 foodMap 中是 cuisine，在 cuisineMap 中是 food
//	rate int
//}
//type hp []pair
//
//func (h hp) Len() int { return len(h) }
//func (h hp) Less(i, j int) bool {
//	return h[i].rate > h[j].rate ||
//		h[i].rate == h[j].rate && h[i].s < h[j].s // 最小字典序在前
//}
//func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
//func (h *hp) Push(x any)   { *h = append(*h, x.(pair)) }
//func (h *hp) Pop() any {
//	v := (*h)[len(*h)-1]
//	*h = (*h)[:len(*h)-1]
//	return v
//}
