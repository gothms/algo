package main

import "fmt"

func main() {
	nums := []int{10, 2, 8, 9, 3, 8, 1, 5, 2, 3, 7, 6} // 8
	nums = []int{1, 1, 1, 2, 2, 2, 3, 3, 3}            // 7
	freq := maxEqualFreq(nums)
	fmt.Println(freq)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxEqualFreq(nums []int) int {
	// 优化
	// 其他做法：记录最大频次 maxFreq，其中一种判断为：
	// maxFreq * cnt[maxFreq] + (maxFreq-1) * cnt[maxFreq-1] == i+1 && cnt[maxFreq] == 1
	cnt, memo := make(map[int]int), make(map[int]int)
	ans := 0
	for i, v := range nums {
		memo[v]++
		k := memo[v]
		if k > 1 {
			cnt[k-1]--
			if cnt[k-1] == 0 {
				delete(cnt, k-1)
			}
		}
		cnt[k]++
		if n, j := len(cnt), cnt[k]; n == 1 && (k == 1 || j == 1) || n == 2 &&
			(cnt[1] == 1 || j == 1 && cnt[k-1] > 0 || cnt[k+1] == 1) {
			ans = i + 1
		}
	}
	return ans

	// 个人
	//cnt, memo := make(map[int]map[int]struct{}), make(map[int]int)
	//cnt[0] = make(map[int]struct{})
	//ans := 0
	//for i, v := range nums {
	//	memo[v]++
	//	k := memo[v]
	//	if delete(cnt[k-1], v); len(cnt[k-1]) == 0 {
	//		delete(cnt, k-1)
	//	}
	//	if cnt[k] == nil {
	//		cnt[k] = make(map[int]struct{})
	//	}
	//	cnt[k][v] = struct{}{}
	//	if l, m := len(cnt[1]), len(cnt[k]); len(cnt) == 1 && (l > 0 || m == 1) ||
	//		len(cnt) == 2 && (l == 1 || m == 1 && len(cnt[k-1]) > 0 || len(cnt[k+1]) == 1) {
	//		ans = max(ans, i+1)
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
