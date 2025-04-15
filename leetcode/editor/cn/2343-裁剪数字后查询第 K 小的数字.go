package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []string{"24", "37", "96", "04"}
	queries := [][]int{{2, 1}, {2, 2}}
	numbers := smallestTrimmedNumbers(nums, queries)
	fmt.Println(numbers)
}

// leetcode submit region begin(Prohibit modification and deletion)
func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	// 基数排序
	n, m := len(queries), len(nums[0])
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		//return queries[idx[i]][1] < queries[idx[j]][1] ||
		//	queries[idx[i]][1] == queries[idx[j]][1] && queries[idx[i]][0] < queries[idx[j]][0]
		return queries[idx[i]][1] < queries[idx[j]][1]
	})
	// 如果 nums 元素相等，则会出现覆盖而导致无法判断索引
	//memo := make(map[string]int, len(nums))
	//for i, s := range slices.Backward(nums) {
	//	memo[s] = i
	//}
	type pair struct {
		s string
		i int
	}
	pairs := make([]pair, len(nums))
	for i, s := range nums {
		pairs[i] = pair{s, i}
	}

	var radix [10][]pair
	ans := make([]int, n)
	for i, j := m-1, 0; i >= 0 && j < n; i-- {
		for _, p := range pairs {
			k := int(p.s[i] - '0')
			radix[k] = append(radix[k], p)
		}
		pairs = pairs[:0] // 清除
		for k, arr := range radix {
			pairs = append(pairs, arr...)
			radix[k] = radix[k][:0] // 清除
		}
		for j < n && queries[idx[j]][1]+i == m {
			ans[idx[j]] = pairs[queries[idx[j]][0]-1].i
			j++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
