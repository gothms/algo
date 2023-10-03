//给定一个整数数组，编写一个函数，找出索引m和n，只要将索引区间[m,n]的元素排好序，整个数组就是有序的。注意：n-m尽量最小，也就是说，找出符合条件的最短
//序列。函数返回值为[m,n]，若不存在这样的m和n（例如整个数组是有序的），请返回[-1,-1]。
// 示例：
// 输入： [1,2,4,7,10,11,7,12,6,7,16,18,19]
//输出： [3,9]
//
// 提示：
//
// 0 <= len(array) <= 1000000
//
//
// Related Topics 栈 贪心 数组 双指针 排序 单调栈 👍 127 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	array := []int{1, 2, 4, 7, 10, 11, 7, 12, 5, 7, 16, 18, 5}
	sort := subSort(array)
	fmt.Println(sort)
}

// leetcode submit region begin(Prohibit modification and deletion)
func subSort(array []int) []int {
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	mv, n := math.MinInt32, len(array) // mv 初始值不能是 0
	ss, stack := []int{n, -1}, make([]int, 1, n+1)
	stack[0] = -1 // 哨兵
	for i, v := range array {
		last := len(stack) - 1
		t := last
		for last > 0 && v < array[stack[last]] {
			mv = maxVal(mv, array[stack[last]]) // 弹出的数中的最大值
			last--
		}
		if last != t { // 有变更
			ss[0] = minVal(ss[0], stack[last]+1) // 更新左边界
		}
		if v < mv {
			ss[1] = i // 更新右边界
		}
		stack = stack[:last+1]
		stack = append(stack, i)
	}
	if ss[0] == n {
		ss[0] = -1
	}
	return ss

	// lc：双指针
	//if len(array) == 0 {
	//	return []int{-1, -1}
	//}
	//n := len(array) - 1
	//li, ri := n, 0
	//lv, rv := array[li], array[ri]
	//for i, j := 1, n-1; i <= n; i, j = i+1, j-1 {
	//	if array[i] < rv {
	//		ri = i
	//	} else {
	//		rv = array[i] // 当前最大值
	//	}
	//	if array[j] > lv {
	//		li = j
	//	} else {
	//		lv = array[j] // 当前最小值
	//	}
	//	//fmt.Println(rv, lv)
	//}
	//
	//if li >= ri {
	//	return []int{-1, -1}
	//}
	//return []int{li, ri}
}

//leetcode submit region end(Prohibit modification and deletion)
