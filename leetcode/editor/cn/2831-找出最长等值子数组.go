package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2, 3, 1, 3}
	k := 3
	//nums = []int{1, 1, 2, 2, 1, 1}
	//k = 2
	subarray := longestEqualSubarray(nums, k)
	fmt.Println(subarray)
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestEqualSubarray(nums []int, k int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func longestEqualSubarray_(nums []int, k int) int {
	memo := make(map[int][]int)
	for i, v := range nums {
		memo[v] = append(memo[v], i) // 相同元素分组
	}
	maxLES := 0
	for _, v := range memo {
		if len(v) < maxLES {
			continue
		}
		for i, j := 0, 0; i < len(v); i++ { // 双指针
			for j < i && v[i]-v[j]-i+j > k { // 区间内删除元素 <= k
				j++
			}
			maxLES = max(maxLES, i-j+1) // 更新最长等值子数组长度
		}
	}
	return maxLES
}
