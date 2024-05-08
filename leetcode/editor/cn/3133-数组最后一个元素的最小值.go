//给你两个整数 n 和 x 。你需要构造一个长度为 n 的 正整数 数组 nums ，对于所有 0 <= i < n - 1 ，满足 nums[i + 1]
//大于 nums[i] ，并且数组 nums 中所有元素的按位 AND 运算结果为 x 。
//
// 返回 nums[n - 1] 可能的 最小 值。
//
//
//
// 示例 1：
//
//
// 输入：n = 3, x = 4
//
//
// 输出：6
//
// 解释：
//
// 数组 nums 可以是 [4,5,6] ，最后一个元素为 6 。
//
// 示例 2：
//
//
// 输入：n = 2, x = 7
//
//
// 输出：15
//
// 解释：
//
// 数组 nums 可以是 [7,15] ，最后一个元素为 15 。
//
//
//
// 提示：
//
//
// 1 <= n, x <= 10⁸
//
//
// Related Topics 位运算 👍 9 👎 0

package main

import "fmt"

func main() {
	n, x := 3, 4
	n, x = 2, 7
	end := minEnd(n, x)
	fmt.Println(end)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minEnd(n int, x int) int64 {
	// 优化
	n--
	for ; x > 0; x &= x - 1 {
		i := x & -x // i 为要插入的二进制 1
		//n = n&^(i-1)<<1 | i | n&(i-1) // 在 n 的二进制中插入 i
		n = n & ^(i-1) << 1 | i | n&(i-1) // 在 n 的二进制中插入 i
	}
	return int64(n)

	// 个人
	// 由题意：数组 nums 中所有元素的按位 AND 运算结果为 x。则每个元素都“包含”x 的所有二进制 1
	// 那么就转化为题意：在第 n 大的非负整数的二进制位中，插入 x 的所有二进制 1
	//ans := int64(n - 1)
	//for ; x > 0; x &= x - 1 {
	//	i := int64(x & -x)        // 取出最后一位
	//	v := ans & (i - 1)        // 取出 ans 的最后 k 位，k 为 i 的长度 -1
	//	ans = (ans&^v)<<1 | i | v // 将 i 的二进制插入 ans
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
