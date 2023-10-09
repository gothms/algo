//给你一个正整数 num ，请你将它分割成两个非负整数 num1 和 num2 ，满足：
//
//
// num1 和 num2 直接连起来，得到 num 各数位的一个排列。
//
//
//
// 换句话说，num1 和 num2 中所有数字出现的次数之和等于 num 中所有数字出现的次数。
//
//
// num1 和 num2 可以包含前导 0 。
//
//
// 请你返回 num1 和 num2 可以得到的和的 最小 值。
//
// 注意：
//
//
// num 保证没有前导 0 。
// num1 和 num2 中数位顺序可以与 num 中数位顺序不同。
//
//
//
//
// 示例 1：
//
//
//输入：num = 4325
//输出：59
//解释：我们可以将 4325 分割成 num1 = 24 和 num2 = 35 ，和为 59 ，59 是最小和。
//
//
// 示例 2：
//
//
//输入：num = 687
//输出：75
//解释：我们可以将 687 分割成 num1 = 68 和 num2 = 7 ，和为最优值 75 。
//
//
//
//
// 提示：
//
//
// 10 <= num <= 10⁹
//
//
// Related Topics 贪心 数学 排序 👍 20 👎 0

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	num := 1000
	i := splitNum(num)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func splitNum(num int) int {
	// api
	buf := []byte(strconv.Itoa(num))
	sort.Slice(buf, func(i, j int) bool {
		return buf[i] < buf[j]
	})
	temp := [2]int{}
	for i, v := range buf {
		temp[i%2] = temp[i%2]*10 + int(v-'0')
	}
	return temp[0] + temp[1]

	// 个人
	//temp := make([]int, 0)
	//for ; num != 0; num /= 10 {
	//	if v := num % 10; v != 0 {
	//		temp = append(temp, num%10)
	//	}
	//}
	//sort.Ints(temp)
	//i, n := 0, len(temp)
	////for i < n && temp[i] == 0 {	 // 注释：不可以包含前导 0 判断
	////	i++
	////}
	////switch i {
	////case n - 1:
	////	return temp[i] * int(math.Pow(10.0, float64(i-1)))
	////case 0:
	////	break
	////default:
	////	temp[i], temp[i+1], temp[0], temp[1] = 0, 0, temp[i], temp[i+1]
	////}
	//v1, v2 := 0, 0
	//for i = 1; i < n; i += 2 {
	//	v1 = v1*10 + temp[i-1]
	//	v2 = v2*10 + temp[i]
	//}
	//if n&1 == 1 {
	//	v1 = v1*10 + temp[n-1]
	//}
	//return v1 + v2
}

//leetcode submit region end(Prohibit modification and deletion)
