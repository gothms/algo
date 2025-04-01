package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumIndex(nums []int) int {

	// 个人
	val, cnt := 0, 0
	for _, v := range nums {
		if v == val {
			cnt++
		} else if cnt == 0 {
			val, cnt = v, 1
		} else {
			cnt--
		}
	}
	cnt = 0
	for _, v := range nums {
		if v == val {
			cnt++
		}
	}
	n, c := len(nums), 0
	for i, v := range nums {
		if v == val {
			c++
		}
		if c<<1 > i+1 && (cnt-c)<<1 > n-i-1 {
			return i
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
