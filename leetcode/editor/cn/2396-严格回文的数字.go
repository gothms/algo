package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isStrictlyPalindromic(n int) bool {
	// 在题目的条件下，答案一定为 false，证明如下：
	//	根据带余除法，n=qb+r，其中 0≤r<b
	//	取 b=n−2，那么当 n>4 时，上式的 q=1，r=2，也就是说 n 在 n−2 进制下的数值为 12，不是回文数
	//	而对于 n=4，在 b=2 进制下的数值为 100，也不是回文数
	return false

	//arr := make([]int, 0, bits.Len(uint(n)))
	//for i := 2; i <= n-2; i++ {
	//	arr = arr[:0]
	//	for v := i; v != 0; v /= i {
	//		arr = append(arr, v%i)
	//	}
	//	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
	//		if arr[l] != arr[r] {
	//			return false
	//		}
	//	}
	//}
	//return true
}

//leetcode submit region end(Prohibit modification and deletion)
