package main

import (
	"fmt"
	"math/bits"
	"slices"
)

func main() {
	label := 26
	tree := pathInZigZagTree(label)
	fmt.Println(tree)
}

// leetcode submit region begin(Prohibit modification and deletion)
func pathInZigZagTree(label int) []int {
	ans := make([]int, 0)
	var (
		level = bits.Len(uint(label))
		start = 1 << (level - 1)
		v     = label
	)
	if level&1 == 0 {
		v = start + start<<1 - label - 1
	}
	for v > 1 {
		level--
		start >>= 1
		v >>= 1
		ans = append(ans, label)
		if level&1 == 0 { // 偶数层的数字是颠倒的，则修正
			label = start + start<<1 - v - 1
		} else {
			label = v
		}
	}
	ans = append(ans, 1)
	slices.Reverse(ans)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
