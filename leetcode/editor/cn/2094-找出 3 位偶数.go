package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findEvenNumbers(digits []int) []int {
	// lc
	// 方法一：穷举 digits 能组成的所有 3 位数，校验 >= 100 且是 偶数
	// 方法二：遍历 [100,999] 中的偶数，校验 digits 是否能组成这个数

	// 个人
	// 一个数字的个数超过 3 时，是没有意义的
	const n = 10
	num := [n]int{}
	for _, v := range digits {
		num[v]++
	}
	// 排序逻辑 & 去重逻辑
	// 如 digits = [2,2,8,8,2]，当选择了 2 2 后，去重方法为：第三个数仅依次选择 2 和 8
	ans := make([]int, 0)
	for i := 1; i < n; i++ {
		if num[i] == 0 {
			continue
		}
		v := i * 10
		num[i]-- // 选择数字 i
		for j := 0; j < n; j++ {
			if num[j] == 0 {
				continue
			}
			v += j
			num[j]--
			for k := 0; k < n; k += 2 { // 只选偶数
				if num[k] == 0 {
					continue
				}
				ans = append(ans, v*10+k)
			}
			num[j]++
			v -= j // 回滚
		}
		num[i]++ // 回溯
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
