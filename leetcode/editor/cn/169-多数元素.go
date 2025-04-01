package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	v, c := math.MinInt64, 0
	for _, num := range nums {
		if num == v {
			c++
		} else if c == 0 {
			v = num
		} else {
			c--
		}
	}
	return v
}

//leetcode submit region end(Prohibit modification and deletion)

//func majorityElement(nums []int) int {
//	me, cnt := 0, 0
//	for _, v := range nums {
//		if cnt == 0 {
//			me = v
//			cnt = 1
//		} else if v == me {
//			cnt++
//		} else {
//			cnt--
//		}
//	}
//	return me
//}
