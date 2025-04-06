package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findErrorNums(nums []int) []int {
	// lc
	xor := 0
	for _, v := range nums {
		xor ^= v
	}
	for i := 1; i <= len(nums); i++ {
		xor ^= i
	}
	lowBit := xor & -xor
	x := 0
	for _, v := range nums {
		if v&lowBit == 0 {
			x ^= v
		}
	}
	for i := 1; i <= len(nums); i++ {
		if i&lowBit == 0 {
			x ^= i
		}
	}
	for _, v := range nums {
		if v == x {
			return []int{x, xor ^ x}
		}
	}
	return []int{x ^ xor, x}

	// 个人
	//repeat := -1
	//for _, v := range nums {
	//	if v < 0 {
	//		v = -v
	//	}
	//	if nums[v-1] < 0 {
	//		repeat = v
	//	} else {
	//		nums[v-1] = -nums[v-1]
	//	}
	//}
	//for i, v := range nums {
	//	if v > 0 {
	//		return []int{repeat, i + 1}
	//	}
	//}
	//return nil
}

//leetcode submit region end(Prohibit modification and deletion)
