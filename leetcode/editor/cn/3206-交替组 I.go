package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfAlternatingGroups(colors []int) int {
	ans, n := 0, len(colors)
	b := colors[0] != colors[1]
	for i := 2; i < n; i++ {
		c := colors[i] != colors[i-1]
		if b && c {
			ans++
		}
		b = c
	}
	if colors[0] != colors[n-1] {
		if b {
			ans++
		}
		if colors[0] != colors[1] {
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
