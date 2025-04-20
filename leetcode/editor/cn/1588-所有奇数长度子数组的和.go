package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func sumOddLengthSubarrays(arr []int) int {
	ans, even, odd := 0, 0, 0
	for i, v := range arr {
		even, odd = odd+v*((i+1)>>1), even+v*(i>>1+1) // 含 arr[i] 的偶数和奇数长度的子数组
		ans += odd
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
