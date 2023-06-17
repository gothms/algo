//数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，
//等等。
//
// 请写一个函数，求任意第n位对应的数字。
//
//
//
// 示例 1：
//
// 输入：n = 3
//输出：3
//
//
// 示例 2：
//
// 输入：n = 11
//输出：0
//
//
//
// 限制：
//
//
// 0 <= n < 2^31
//
//
// 注意：本题与主站 400 题相同：https://leetcode-cn.com/problems/nth-digit/
//
// Related Topics 数学 二分查找 👍 337 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 999999999 // 9
	//n = 1000       // 9
	//n = 191        // 9
	//n = 19         // 9
	digit := findNthDigit(n)
	fmt.Println(digit)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findNthDigit(n int) int {
	// 个人
	if n < 10 {
		return n
	}
	// 1 9 1 / 2 90 10
	d, base, cnt := 1, 9, 1 // 位数，9 90 900...，总位数
	for n >= cnt {          // 可以写成二分
		cnt += d * base
		d++
		base *= 10
	}
	d--
	cnt -= base / 10 * d                   // 还原为低一位的总位数
	idx := (n - cnt) / d                   // (n-cnt+1-1)/d-1+1	小数要进位
	num, i := idx+base/90, idx*d+d+cnt-1-n // 数值，第几位（最低位是 0）
	return num / int(math.Pow10(i)) % 10
}

//leetcode submit region end(Prohibit modification and deletion)
