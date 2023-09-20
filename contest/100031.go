package main

import (
	"fmt"
	"math/bits"
)

func main() {
	nums := []int{5, 10, 1, 5, 2}
	nums = []int{4, 3, 2, 1}
	k := 1
	k = 2
	setBits := sumIndicesWithKSetBits(nums, k)
	fmt.Println(setBits)

	fmt.Println(bits.OnesCount(uint(2)))
}
func sumIndicesWithKSetBits(nums []int, k int) int {
	siwsb, n := 0, len(nums)
	for i := 1<<k + -1; i < n; i++ {
		if bits.OnesCount(uint(i)) == k {
			siwsb += nums[i]
		}
	}
	return siwsb
}
