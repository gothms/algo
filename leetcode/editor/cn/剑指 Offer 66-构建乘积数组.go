//给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即 B[
//i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
//
//
//
// 示例:
//
//
//输入: [1,2,3,4,5]
//输出: [120,60,40,30,24]
//
//
//
// 提示：
//
//
// 所有元素乘积之和不会溢出 32 位整数
// a.length <= 100000
//
//
// Related Topics 数组 前缀和 👍 321 👎 0

package main

import "fmt"

func main() {
	a := make([]int, 0)
	arr := constructArr(a)
	fmt.Println(arr)
}

//leetcode submit region begin(Prohibit modification and deletion)
func constructArr(a []int) []int {
	n := len(a) - 1
	if n < 0 {
		return nil
	}
	cache := make([]int, n+1)
	cache[n] = 1
	for i := n - 1; i >= 0; i-- {
		cache[i] = cache[i+1] * a[i+1]
	}
	for i := 1; i <= n; i++ {
		a[i] *= a[i-1]
		cache[i] *= a[i-1]
	}
	return cache
}

//leetcode submit region end(Prohibit modification and deletion)
