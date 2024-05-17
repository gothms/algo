package main

import (
	"fmt"
	"math/bits"
)

func main() {
	n := 0
	complement := bitwiseComplement(n)
	fmt.Println(complement)
	fmt.Println(1<<max(1, bits.Len(uint(n))) - 1)
}

// leetcode submit region begin(Prohibit modification and deletion)
func bitwiseComplement(n int) int {
	return (1<<max(1, bits.Len(uint(n))) - 1) ^ n
}

//leetcode submit region end(Prohibit modification and deletion)
