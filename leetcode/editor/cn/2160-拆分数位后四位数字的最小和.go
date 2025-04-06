package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumSum(num int) int {
	v := make([]int, 0, 4)
	for ; num != 0; num /= 10 {
		v = append(v, num%10)
	}
	sort.Ints(v)
	return (v[0]+v[1])*10 + v[2] + v[3]
}

//leetcode submit region end(Prohibit modification and deletion)
