package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	memo := make(map[int]struct{})
	for _, v := range nums {
		memo[v] = struct{}{}
	}
	for i, f := range moveFrom {
		delete(memo, f)
		memo[moveTo[i]] = struct{}{}
	}
	ans := make([]int, 0, len(memo))
	for i := range memo {
		ans = append(ans, i)
	}
	sort.Ints(ans)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
