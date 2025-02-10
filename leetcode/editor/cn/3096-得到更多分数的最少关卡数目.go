package main

import "fmt"

func main() {
	arr := []int{1, 1}
	levels := minimumLevels(arr)
	fmt.Println(levels)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumLevels(possible []int) int {
	// 前缀和
	n := len(possible)
	pre := make([]int, n+1)
	for i, v := range possible { // 前缀和
		pre[i+1] += pre[i] + v<<1 - 1
	}
	for i, m := 1, pre[n]; i < n; i++ { // 冬坂五百里 至少要玩一关
		if pre[i]<<1 > m { // pre[i] > m-pre[i]
			return i
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
