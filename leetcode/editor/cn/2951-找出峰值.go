package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findPeaks(mountain []int) []int {
	ans := make([]int, 0)
	for i := len(mountain) - 2; i > 0; i-- {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			ans = append(ans, i)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
