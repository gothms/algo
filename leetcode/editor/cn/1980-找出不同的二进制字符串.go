package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "101"
	v, _ := strconv.ParseInt(s, 2, 32)
	fmt.Println(int(v))
}

// leetcode submit region begin(Prohibit modification and deletion)
func findDifferentBinaryString(nums []string) string {
	memo := make(map[int64]bool)
	for _, s := range nums {
		v, _ := strconv.ParseInt(s, 2, 64)
		memo[v] = true
	}
	var i int64
	for memo[i] {
		i++
	}
	return fmt.Sprintf("%0*b", len(nums), i)
}

//leetcode submit region end(Prohibit modification and deletion)
