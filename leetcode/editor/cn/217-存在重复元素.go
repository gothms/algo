package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func containsDuplicate(nums []int) bool {
	memo := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := memo[v]; ok {
			return true
		}
		memo[v] = struct{}{}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
