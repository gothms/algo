//给定一个数组，包含从 1 到 N 所有的整数，但其中缺了两个数字。你能在 O(N) 时间内只用 O(1) 的空间找到它们吗？
//
// 以任意顺序返回这两个数字均可。
//
// 示例 1:
//
// 输入: [1]
//输出: [2,3]
//
// 示例 2:
//
// 输入: [2,3]
//输出: [1,4]
//
// 提示：
//
//
// nums.length <= 30000
//
//
// Related Topics 位运算 数组 哈希表 👍 227 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 6, 7, 9, 10}
	two := missingTwo(nums)
	fmt.Println(two)
}

// leetcode submit region begin(Prohibit modification and deletion)
func missingTwo(nums []int) []int {
	// 位运算：异或
	xor, n := 0, len(nums)+2
	for _, v := range nums {
		xor ^= v
	}
	for i := 1; i <= n; i++ {
		xor ^= i // 求得缺失的两个数的 ^ 值
	}
	one, dif := 0, xor&-xor
	for _, v := range nums {
		if dif&v != 0 {
			one ^= v
		}
	}
	for i := 1; i <= n; i++ {
		if dif&i != 0 {
			one ^= i // 求得其中一个数
		}
	}
	return []int{one, one ^ xor}

	// 求和
	//sum := 0
	//for _, v := range nums {
	//	sum += v
	//}
	//diff := (len(nums)+2)*(len(nums)+3)>>1 - sum // 缺失的和
	//limit := diff >> 1                           // 将缺失的两个数分为两部分
	//sum = 0
	//for _, v := range nums {
	//	if v <= limit {
	//		sum += v
	//	}
	//}
	//min := limit*(limit+1)>>1 - sum // 求得较小的数
	//return []int{min, diff - min}

	// 原地 swap
	//nums = append(nums, -1, -1, -1)
	//for i := range nums {
	//	for nums[i] != -1 && i != nums[i] { // 存在，则标记其索引
	//		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	//	}
	//}
	//ret := make([]int, 0, 3)
	//for i, v := range nums {
	//	if v == -1 {
	//		ret = append(ret, i)
	//	}
	//}
	//return ret[1:]

	// 个人写法
	//absVal := func(a int) int {
	//	if a < 0 {
	//		return -a
	//	}
	//	return a
	//}
	//f, n := 0, len(nums)
	//ret := make([]int, 0, 2)
	//for _, v := range nums {
	//	v = absVal(v)
	//	if v > n {
	//		f = v
	//	} else {
	//		nums[v-1] = -nums[v-1] // n 范围内的索引标记为 负
	//	}
	//}
	//for i, v := range nums {
	//	if v > 0 {
	//		ret = append(ret, i+1) // 查出 [1,n] 缺失的数
	//	}
	//}
	//switch len(ret) {
	//case 0:
	//	ret = append(ret, n+1, n+2) // 补查 >n 的数
	//case 1:
	//	ret = append(ret, n+1+(f^n)&1) // f,n = 6,5 -> 7. f,n = 7,5 -> 6
	//}
	//return ret
}

//leetcode submit region end(Prohibit modification and deletion)
