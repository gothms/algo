package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumLength(nums []int, k int) int {
	// dp
	f := make([][]int, k)
	for i := range f {
		f[i] = make([]int, k)
	}
	ans := 0
	for _, v := range nums {
		x := v % k
		for y, xy := range f[x] {
			f[y][x] = xy + 1
			ans = max(ans, f[y][x])
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
