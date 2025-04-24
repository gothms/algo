package main

import (
	"fmt"
	"math"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs)
	memo := make(map[int][2]int)
	for _, a := range adjacentPairs {
		x, y := a[0], a[1]
		if p, ok := memo[x]; ok {
			p[1] = y
			memo[x] = p
		} else {
			memo[x] = [2]int{y, math.MaxInt32}
		}
		if p, ok := memo[y]; ok {
			p[1] = x
			memo[y] = p
		} else {
			memo[y] = [2]int{x, math.MaxInt32}
		}
	}
	fmt.Println(memo)

	vis := make(map[int]bool)
	ans := make([]int, (n+1)<<1)
	//start := memo[adjacentPairs[0][0]]
	l, r := n, n+1
	ans[l] = adjacentPairs[0][0]
	l--
	//r++
	vis[ans[l]] = true
	for lv := memo[adjacentPairs[0][0]][1]; lv != math.MaxInt32 && !vis[lv]; {
		ans[l] = lv
		l--
		vis[lv] = true
		lv = memo[lv][0]
		if lv == math.MaxInt32 {
			break
		}
		if vis[lv] {
			lv = memo[lv][1]
		}
		fmt.Println(lv, ans)
	}
	//rv := memo[ans[r]][0]
	//if vis[rv] {
	//	rv = memo[ans[r]][1]
	//}
	for rv := adjacentPairs[0][1]; rv != math.MaxInt32 && !vis[rv]; {
		ans[r] = rv
		r++
		vis[rv] = true
		rv = memo[rv][0]
		if rv == math.MaxInt32 {
			break
		}
		if vis[rv] {
			rv = memo[rv][1]
		}
	}
	return ans[l+1 : r]
}

//leetcode submit region end(Prohibit modification and deletion)
