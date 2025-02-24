package main

import (
	"fmt"
	"math/bits"
	"slices"
)

func main() {
	nums := []int{3, 10, 5, 25, 2, 8}
	xor := findMaximumXOR(nums)
	fmt.Println(xor)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findMaximumXOR(nums []int) int {
	k := bits.Len(uint(slices.Max(nums))) - 1 // 最大位
	ans, mask := 0, 0
	m := make(map[int]struct{})
out:
	for i := k; i >= 0; i-- {
		clear(m)
		mask |= 1 << i
		ans |= 1 << i            // 目标值
		for _, v := range nums { // 两数之和
			b := v & mask
			if _, ok := m[ans^b]; ok {
				continue out // 找到目标值
			}
			m[b] = struct{}{}
		}
		ans ^= 1 << i // 目标值未找到，回溯
	}
	return ans

	// 个人：找出第二个数的集合
	// lc
	//ans, maxV := 0, 0
	//for _, v := range nums {
	//	maxV = max(maxV, v) // maxV 最高位可视为第一个数的集合
	//}
	//k := bits.Len(uint(maxV))
	//for i := k - 1; i >= 0; i-- { // 第 i 位是否 ok
	//	temp := make(map[int]bool)
	//	for _, v := range nums {
	//		temp[v>>i] = true
	//	}
	//	ans = ans<<1 + 1
	//	found := false
	//	for _, v := range nums {
	//		t := v >> i
	//		if temp[t^ans] { // 关键逻辑，见 lc
	//			found = true // 有第 i 位
	//			break
	//		}
	//	}
	//	fmt.Println(i, ans, temp)
	//	if !found { // 没有第 i 位
	//		ans--
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
