package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxAliveYear(birth []int, death []int) int {
	// 差分数组
	arr := make([]int, 102) // 100+2
	for i, b := range birth {
		arr[b-1900]++
		arr[death[i]-1900+1]--
	}
	y, mx, cur := 0, 0, 0
	for i, v := range arr {
		cur += v
		if cur > mx {
			y, mx = i, cur
		}
	}
	return y + 1900
}

//leetcode submit region end(Prohibit modification and deletion)
