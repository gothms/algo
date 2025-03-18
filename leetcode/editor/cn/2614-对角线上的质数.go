package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)

var prime2614 []bool

func init() {
	const n = 4e6 + 1
	prime2614 = make([]bool, n)
	prime2614[1] = true // 1 不是质数
	for i := 2; i < n; i++ {
		for j := i << 1; j < n; j += i {
			prime2614[j] = true
		}
	}
}

func diagonalPrime(nums [][]int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		if v := nums[i][i]; !prime2614[v] {
			ans = max(ans, v)
		}
		if v := nums[i][n-i-1]; !prime2614[v] {
			ans = max(ans, v)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
