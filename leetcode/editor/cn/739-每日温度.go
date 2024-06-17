package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func dailyTemperatures(temperatures []int) []int {
	// 栈：非递增
	n := len(temperatures)
	st := make([]int, 0, n)
	ans := make([]int, n)
	for i, t := range temperatures {
		for len(st) > 0 && t > temperatures[st[len(st)-1]] {
			ans[st[len(st)-1]] = i - st[len(st)-1]
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
