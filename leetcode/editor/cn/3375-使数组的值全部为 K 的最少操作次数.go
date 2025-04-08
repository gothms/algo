package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(nums []int, k int) int {
	memo := make(map[int]struct{})
	for _, v := range nums {
		if v < k {
			return -1
		}
		memo[v] = struct{}{}
	}
	delete(memo, k)
	return len(memo)
}

//leetcode submit region end(Prohibit modification and deletion)
