package main

import "fmt"

func main() {
	fmt.Println(-3 & 1)
	fmt.Println(-4 >> 1)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findSwapValues(array1 []int, array2 []int) []int {
	// 个人
	memo := make(map[int]struct{})
	s1, s2 := 0, 0
	for _, v := range array1 {
		memo[v] = struct{}{}
		s1 += v
	}
	for _, v := range array2 {
		s2 += v
	}
	d := s1 - s2
	if d&1 == 1 {
		return nil
	}
	d >>= 1
	for _, v := range array2 {
		if _, ok := memo[v+d]; ok {
			return []int{v + d, v}
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
