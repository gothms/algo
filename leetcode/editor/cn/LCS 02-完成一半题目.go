package main

import (
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func halfQuestions(questions []int) int {
	// lc
	cnt := make(sort.IntSlice, 1001)
	for _, v := range questions {
		cnt[v]++
	}
	sort.Sort(sort.Reverse(cnt))
	s, n := 0, len(questions)>>1
	for i, v := range cnt {
		if s += v; s >= n {
			return i + 1
		}
	}
	return 0

	// 个人
	//n := len(questions) >> 1
	//memo := make(map[int]int)
	//for _, v := range questions {
	//	memo[v]++
	//}
	//arr := make([]int, 0, len(memo))
	//for _, v := range memo {
	//	arr = append(arr, v)
	//}
	//sort.Ints(arr)
	//for i, s := len(arr)-1, 0; i >= 0; i-- {
	//	if s += arr[i]; s >= n {
	//		return len(arr) - i
	//	}
	//}
	//return 0
}

//leetcode submit region end(Prohibit modification and deletion)
