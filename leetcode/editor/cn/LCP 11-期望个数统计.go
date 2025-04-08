package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func expectNumber(scores []int) int {
	memo := make(map[int]struct{})
	for _, v := range scores {
		memo[v] = struct{}{}
	}
	return len(memo)
}

//leetcode submit region end(Prohibit modification and deletion)
