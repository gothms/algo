//给出基数为 -2 的两个数 arr1 和 arr2，返回两数相加的结果。
//
// 数字以 数组形式 给出：数组由若干 0 和 1 组成，按最高有效位到最低有效位的顺序排列。例如，arr = [1,1,0,1] 表示数字 (-2)^3 +
// (-2)^2 + (-2)^0 = -3。数组形式 中的数字 arr 也同样不含前导零：即 arr == [0] 或 arr[0] == 1。
//
// 返回相同表示形式的 arr1 和 arr2 相加的结果。两数的表示形式为：不含前导零、由若干 0 和 1 组成的数组。
//
//
//
// 示例 1：
//
//
//输入：arr1 = [1,1,1,1,1], arr2 = [1,0,1]
//输出：[1,0,0,0,0]
//解释：arr1 表示 11，arr2 表示 5，输出表示 16 。
//
//
//
//
//
// 示例 2：
//
//
//输入：arr1 = [0], arr2 = [0]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：arr1 = [0], arr2 = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
//
// 1 <= arr1.length, arr2.length <= 1000
// arr1[i] 和 arr2[i] 都是 0 或 1
// arr1 和 arr2 都没有前导0
//
//
// Related Topics 数组 数学 👍 85 👎 0

package main

import "fmt"

func main() {
	arr1 := []int{1, 1, 1, 1, 1}
	arr2 := []int{1, 0, 1}
	arr1 = []int{1}
	arr2 = []int{1, 0, 1}
	negabinary := addNegabinary(arr1, arr2)
	fmt.Println(negabinary)
}

//leetcode submit region begin(Prohibit modification and deletion)
func addNegabinary(arr1 []int, arr2 []int) []int {
	n, m := len(arr1)-1, len(arr2)-1
	if n < m {
		arr1, n, arr2, m = arr2, m, arr1, n
	}
	pre := 0
	addNega := func(i, v int) {
		switch v {
		case 0, 1:
			arr1[i] = v
			pre = 0
		case 2, 3: // 2 3 进位是 -1
			arr1[i] = v - 2
			pre = -1
		case -1: // 不够减，2^n - 2^(n-1)
			arr1[i] = 1
			pre = 1
		}
	}
	for i, j := n, m; j >= 0; i, j = i-1, j-1 {
		curr := arr1[i] + arr2[j] + pre
		addNega(i, curr)
	}
	for i := n - m - 1; i >= 0; i-- {
		curr := arr1[i] + pre
		addNega(i, curr)
	}
	if pre == 1 { // 补丁
		return append([]int{1}, arr1...)
	} else if pre == -1 {
		return append([]int{1, 1}, arr1...)
	}
	i := 0
	for i < n && arr1[i] == 0 { // 不是 <=n，保留最后 1 位
		i++
	}
	return arr1[i:]
}

//leetcode submit region end(Prohibit modification and deletion)
