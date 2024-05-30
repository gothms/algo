package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func occurrencesOfElement(nums []int, queries []int, x int) []int {
	pos := make([]int, 0)
	for i, v := range nums {
		if v == x {
			pos = append(pos, i)
		}
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		if q <= len(pos) {
			ans[i] = pos[q-1]
		} else {
			ans[i] = -1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
