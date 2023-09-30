//编写一个方法，找出两个数字a和b中最大的那一个。不得使用if-else或其他比较运算符。
// 示例：
// 输入： a = 1, b = 2
//输出： 2
//
//
// Related Topics 位运算 脑筋急转弯 数学 👍 141 👎 0

package main

import "fmt"

func main() {
	a := -6
	a = 8
	fmt.Println(int64(a) & (int64(a) >> 63))
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximum(a int, b int) int {
	// dif&(dif>>63)：dif 为负，结果为 dif，否则为 0
	dif := int64(a - b) // 避免越界
	return int(int64(a) - dif&(dif>>63))
}

//leetcode submit region end(Prohibit modification and deletion)
