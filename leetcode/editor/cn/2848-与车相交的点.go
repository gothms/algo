package main

import "fmt"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfPoints(nums [][]int) int {
	memo := make([]bool, 101)
	ans := 0
	for _, v := range nums {
		for i := v[0]; i <= v[1]; i++ {
			if !memo[i] {
				memo[i] = true
				ans++
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
