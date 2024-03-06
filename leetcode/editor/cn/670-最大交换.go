//给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
//
// 示例 1 :
//
//
//输入: 2736
//输出: 7236
//解释: 交换数字2和数字7。
//
//
// 示例 2 :
//
//
//输入: 9973
//输出: 9973
//解释: 不需要交换。
//
//
// 注意:
//
//
// 给定数字的范围是 [0, 10⁸]
//
//
// Related Topics 贪心 数学 👍 421 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	num := 2736
	//num = 9973
	num = 98368 // 98863
	swap := maximumSwap(num)
	fmt.Println(swap)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumSwap(num int) int {
	arr := make([]int, 0)
	for cur := num; cur > 0; cur /= 10 {
		arr = append(arr, cur%10)
	}
	//s := []byte(strconv.Itoa(num))	// v, _ := strconv.Atoi(string(s))
	a, b, c := -1, 0, 0     // arr[a]：要交换的高位，arr[b]：最大值位，arr[c]：要交换的低位
	for i, v := range arr { // 贪心
		if v < arr[b] {
			a, c = i, b // 可交换
		} else if v > arr[b] {
			b = i // 潜在交换的最大值
		}
	}
	if a == -1 {
		return num
	}
	d := arr[c] - arr[a]
	return num + int(math.Pow10(a))*d - int(math.Pow10(c))*d
}

//leetcode submit region end(Prohibit modification and deletion)
