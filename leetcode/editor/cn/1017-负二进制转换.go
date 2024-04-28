//给你一个整数 n ，以二进制字符串的形式返回该整数的 负二进制（base -2）表示。
//
// 注意，除非字符串就是 "0"，否则返回的字符串中不能含有前导零。
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出："110"
//解释：(-2)² + (-2)¹ = 2
//
//
// 示例 2：
//
//
//输入：n = 3
//输出："111"
//解释：(-2)² + (-2)¹ + (-2)⁰ = 3
//
//
// 示例 3：
//
//
//输入：n = 4
//输出："100"
//解释：(-2)² = 4
//
//
//
//
// 提示：
//
//
// 0 <= n <= 10⁹
//
//
// Related Topics 数学 👍 178 👎 0

package main

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

func main() {
	n := 2
	//n = 3
	//n = 4
	//n = 0
	n = 6
	n = 5
	//n = 57
	neg2 := baseNeg2(n)
	fmt.Println("neg2:", neg2)

	//for x := -5; x <= 5; x++ {
	//	v := (x-1)>>31 + 1
	//	fmt.Println(x, v)
	//}

	//fmt.Println(3 / -2)	// -1
	//fmt.Println(5 % -3)  // 2
	//fmt.Println(-5 % -3) // -2
}

// leetcode submit region begin(Prohibit modification and deletion)
func baseNeg2(n int) string {
	// lc：math、k 进制
	if n <= 1 {
		return strconv.Itoa(n)
	}
	// 0x55555555 - n：对偶数位的 1，向前借位
	// 0x55555555 ^ (0x55555555 - n)：将奇数为的 1，补回来
	v := 0x55555555 ^ (0x55555555 - n)
	k := bits.Len(uint(v))
	var sb strings.Builder
	for i := 1 << (k - 1); i > 0; i >>= 1 {
		sb.WriteRune(rune((v&i-1)>>31+1) + '0') // (v&i-1)>>31+1：v&i > 0 则为 1，否则为 0
	}
	return sb.String()

	// 任意负进制 |k| >= 2
	// 根据 10 进制转二进制的除法图，来画负进制的除法图
	//if n <= 1 {
	//	return strconv.Itoa(n)
	//}
	//var b []byte
	//k := -3
	//for n != 0 {
	//	v := n % k
	//	n /= k
	//	if v < 0 { // 不够，向前借 1
	//		b = append(b, byte(v-k)+'0')
	//		n++
	//	} else {
	//		b = append(b, byte(v)+'0')
	//	}
	//}
	//slices.Reverse(b)
	//return string(b)

	// 个人
	//var b []byte
	//for i := 0; ; i ^= 1 {
	//	b = append(b, byte(n&1)+'0') // n&1
	//	n += (n & 1) & i             // n&1 == 1 且 i 为 1
	//	n >>= 1                      // 处理下一位
	//	if n == 0 {
	//		break
	//	}
	//}
	//slices.Reverse(b)
	//return string(b)
}

//leetcode submit region end(Prohibit modification and deletion)
