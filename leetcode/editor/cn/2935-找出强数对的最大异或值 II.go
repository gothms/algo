package main

import (
	"math/bits"
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumStrongPairXor(nums []int) int {
	cache := make(map[int]int)
	sort.Ints(nums)
	bit := bits.Len(uint(nums[len(nums)-1])) - 1
	ans, mask := 0, 0
	for i := bit; i >= 0; i-- {
		clear(cache)
		mask |= 1 << i
		newAns := ans | 1<<i
		for _, v := range nums {
			if o, ok := cache[v&mask^newAns]; ok && o<<1 >= v {
				ans = newAns
				break
			}
			cache[v&mask] = v
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

//func maximumStrongPairXor(nums []int) int {
//	sort.Ints(nums)
//	t := &trie{}
//	ans, j := 0, 0
//	for i, v := range nums {
//		t.insert(v)
//		for ; j < i && nums[j]<<1 < v; j++ {
//			t.remove(nums[j])
//		}
//		ans = max(ans, t.maxXor(v))
//	}
//	return ans
//}
//
//const bit = 19
//
//type trie struct {
//	children [2]*trie
//	c        int
//}
//
//func (t *trie) insert(v int) {
//	cur := t
//	for i := bit; i >= 0; i-- {
//		b := v >> i & 1
//		if cur.children[b] == nil {
//			cur.children[b] = &trie{}
//		}
//		cur = cur.children[b]
//		cur.c++
//	}
//}
//func (t *trie) remove(v int) {
//	cur := t
//	for i := bit; i >= 0; i-- {
//		cur = cur.children[v>>i&1]
//		cur.c--
//	}
//}
//func (t *trie) maxXor(v int) int {
//	cur := t
//	ans := 0
//	for i := bit; i >= 0; i-- {
//		b := v >> i & 1
//		if cur.children[b^1] != nil && cur.children[b^1].c > 0 {
//			ans |= 1 << i
//			b ^= 1
//		}
//		cur = cur.children[b]
//	}
//	return ans
//}
