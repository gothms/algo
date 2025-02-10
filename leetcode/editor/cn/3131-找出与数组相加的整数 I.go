

package main

import (
	"math"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func addedInteger(nums1 []int, nums2 []int) int {
	m1, m2 := math.MaxInt32, math.MaxInt32
	for i, v := range nums1 {
		m1 = min(m1, v)
		m2 = min(m2, nums2[i])
	}
	return m2 - m1
}

//leetcode submit region end(Prohibit modification and deletion)
