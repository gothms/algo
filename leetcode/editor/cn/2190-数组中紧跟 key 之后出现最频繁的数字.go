package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func mostFrequent(nums []int, key int) int {
	memo := make(map[int]int)
	for i, v := range nums[:len(nums)-1] {
		if v == key {
			memo[nums[i+1]]++
		}
	}
	ans, mx := 0, 0
	for k, v := range memo {
		if v > mx {
			ans, mx = k, v
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
