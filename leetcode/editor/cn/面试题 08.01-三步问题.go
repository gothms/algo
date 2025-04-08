package main

import "fmt"

func main() {
	n := 1
	n = 2
	n = 3
	n = 5
	step := waysToStep(n)
	fmt.Println(step)
}

// leetcode submit region begin(Prohibit modification and deletion)
func waysToStep(n int) int {
	const mod = 1e9 + 7
	one, two, three := 1, 0, 0
	for i := 1; i <= n; i++ {
		one, two, three = (one+two+three)%mod, one, two
	}
	return one
}

//leetcode submit region end(Prohibit modification and deletion)
