package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func flipLights(n int, presses int) int {
	// 分情况讨论
	if presses == 0 {
		return 1
	} else if n == 1 {
		return 2
	} else if n == 2 {
		if presses == 1 {
			return 3
		} else {
			return 4
		}
	} else {
		if presses == 1 {
			return 4
		} else if presses == 2 {
			return 7
		} else {
			return 8
		}
	}

	// 重点：分析过程

}

//leetcode submit region end(Prohibit modification and deletion)
