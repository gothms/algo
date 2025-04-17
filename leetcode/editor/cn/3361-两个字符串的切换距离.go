package main

import "fmt"

func main() {
	s := "abab"
	t := "baba"
	next := []int{100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	prev := []int{1, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	distance := shiftDistance(s, t, next, prev)
	fmt.Println(distance)
}

// leetcode submit region begin(Prohibit modification and deletion)
func shiftDistance(s string, t string, nextCost []int, previousCost []int) int64 {
	// lc：双倍的 nextCost & previousCost 长度
	// 个人
	const n = 26
	next, prev := [n + 1]int{}, [n + 1]int{}
	for i := 0; i < n; i++ {
		next[i+1] = next[i] + nextCost[i]
		prev[i+1] = prev[i] + previousCost[i]
	}
	ans := 0
	for i, c := range s {
		x, y := int(c-'a'+1), int(t[i]-'a'+1) // 在前缀数组中的坐标
		if x == y {
			continue
		}
		if x < y {
			ans += min(next[y-1]-next[x-1], prev[x]+prev[n]-prev[y]) // 往后 / 往前
		} else {
			ans += min(next[y-1]+next[n]-next[x-1], prev[x]-prev[y])
		}
	}
	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
