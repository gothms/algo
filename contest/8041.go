package main

import "fmt"

func main() {
	nums := []int{5, 10, 3, 10, 1, 13, 7, 9, 4, 5, 10, 3, 10, 1, 13, 7, 9, 4, 5, 10, 3, 10, 1, 13, 7, 9, 4}
	ans := maximumSum(nums)
	fmt.Println(ans)
}
func core(n int) int {
	res := 1
	for i := 2; i*i <= n; i++ {
		e := 0
		for n%i == 0 {
			e ^= 1
			n /= i
		}
		if e == 1 { // 奇数个 i：归于 i；偶数个 i：归于 1
			res *= i
		}
	}
	if n > 1 { // 正整数 n 为质数
		res *= n
	}
	return res
}

func maximumSum(nums []int) (ans int64) {
	// 枚举：优化
	maxVal := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}
	ret, n := int64(0), len(nums)
	for i := 1; i <= n; i++ {
		sum := int64(0)                     // 统计子集 'i' 的子序和
		for j, idx := 1, i; idx <= n; j++ { // 子集 i 包含的元素 = i*j*j
			sum += int64(nums[idx-1])
			idx += (j<<1 + 1) * i // 下一个元素
		}
		ret = maxVal(ret, sum)
	}
	return ret

	// lc
	//maxVal := func(a, b int64) int64 {
	//	if b > a {
	//		return b
	//	}
	//	return a
	//}
	//sum := make([]int64, len(nums)+1) // 分为 n 个完全子集
	//for i, x := range nums {
	//	c := core(i + 1)   // 正整数 i+1，属于第 c 个完全子集
	//	sum[c] += int64(x) // 累加子集
	//	ans = maxVal(ans, sum[c])
	//}
	//return
}
