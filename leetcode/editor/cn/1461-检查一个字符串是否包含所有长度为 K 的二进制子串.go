package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "00110110"
	//s = "0110" // false
	k := 2
	codes := hasAllCodes(s, k)
	fmt.Println(codes)
}

// leetcode submit region begin(Prohibit modification and deletion)
func hasAllCodes(s string, k int) bool {
	if len(s) < 1<<k+k-1 { // 不可能
		return false
	}
	mask := 1<<k - 1
	val, _ := strconv.ParseInt(s[:k-1], 2, 64)
	v := int(val)
	memo := make(map[int]struct{}, 1<<k) // 声明空间大小
	for i := k - 1; i < len(s); i++ {    // 滑动窗体
		v <<= 1
		v &= mask
		v |= int(s[i] - '0')
		memo[v] = struct{}{}
	}
	return len(memo) == 1<<k
}

//leetcode submit region end(Prohibit modification and deletion)
