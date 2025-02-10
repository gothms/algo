package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isArraySpecial(nums []int, queries [][]int) []bool {
	s := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		s[i] = s[i-1]
		if nums[i-1]%2 == nums[i]%2 {
			s[i]++
		}
	}
	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = s[q[0]] == s[q[1]]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
