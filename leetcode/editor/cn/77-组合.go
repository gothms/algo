package main

import (
	"fmt"
	"slices"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	ans, arr := make([][]int, 0), make([]int, k+1)
	for i := range arr {
		arr[i] = i + 1
	}
	arr[k] = n + 1 // 哨兵
	for i := 0; i < k; {
		ans = append(ans, slices.Clone(arr[:k]))
		for i = 0; i < k && arr[i]+1 == arr[i+1]; i++ {
			arr[i] = i + 1
		}
		arr[i]++
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
