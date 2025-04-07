package main

import "fmt"

func main() {
	arr := []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}
	equalSum := canThreePartsEqualSum(arr)
	fmt.Println(equalSum)
}

// leetcode submit region begin(Prohibit modification and deletion)
func canThreePartsEqualSum(arr []int) bool {
	// 思考：如果是按任意组合分为三个和相等的子集呢？
	// 个人
	s := 0
	for _, v := range arr {
		s += v
	}
	if s%3 != 0 {
		return false
	}
	s /= 3
	cur, c := 0, 0
	for _, v := range arr {
		if cur += v; cur != s {
			continue
		}
		cur = 0
		c++
		//if c == 3 {
		//	return true
		//}
	}
	return c >= 3
}

//leetcode submit region end(Prohibit modification and deletion)
