package main

import (
	"fmt"
)

func main() {
	segments := [][]int{{1, 4, 5},
		{4, 7, 7},
		{1, 7, 9}} // [[1,4,14],[4,7,16]]
	segments = [][]int{{1, 4, 5},
		{1, 4, 7},
		{4, 7, 1},
		{4, 7, 11}} // [[1,4,12],[4,7,12]]
	painting := splitPainting(segments)
	fmt.Println(painting)
}

// leetcode submit region begin(Prohibit modification and deletion)
func splitPainting(segments [][]int) [][]int64 {
	// 差分 + hash
	//memo := make(map[int]int)
	//for _, s := range segments {
	//	memo[s[0]] += s[2]
	//	memo[s[1]] -= s[2]
	//}
	//arr := make([][2]int, 0, len(memo))
	//for k, v := range memo {
	//	arr = append(arr, [2]int{k, v})
	//}
	//sort.Slice(arr, func(i, j int) bool {
	//	return arr[i][0] < arr[j][0]
	//})
	//ans := make([][]int64, 0, len(arr)-1)
	//for i := range arr[1:] {
	//	arr[i+1][1] += arr[i][1]
	//	if arr[i][1] > 0 {
	//		ans = append(ans, []int64{int64(arr[i][0]), int64(arr[i+1][0]), int64(arr[i][1])})
	//	}
	//}
	//return ans

	// 差分数组：错误
	//ans := make([][]int64, 0)
	//maxV := 0
	//for _, s := range segments {
	//	maxV = max(maxV, s[1])
	//}
	//arr := make([]int64, maxV+2)
	//for _, s := range segments {
	//	arr[s[0]] += int64(s[2])
	//	arr[s[1]] -= int64(s[2])
	//}
	//for i := 1; i < len(arr); i++ {
	//	arr[i] += arr[i-1]
	//}
	//fmt.Println(arr)
	//var s, lastI, lastV int64
	//for i, v := range arr {
	//	s += v
	//	if s != lastV {
	//		if lastV > 0 {
	//			ans = append(ans, []int64{lastI, int64(i), lastV})
	//		}
	//		lastI, lastV = int64(i), s
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
