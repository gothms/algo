//给定一个由 0 和 1 组成的数组
// arr ，将数组分成 3 个非空的部分 ，使得所有这些部分表示相同的二进制值。
//
// 如果可以做到，请返回任何 [i, j]，其中 i+1 < j，这样一来：
//
//
// arr[0], arr[1], ..., arr[i] 为第一部分；
// arr[i + 1], arr[i + 2], ..., arr[j - 1] 为第二部分；
// arr[j], arr[j + 1], ..., arr[arr.length - 1] 为第三部分。
// 这三个部分所表示的二进制值相等。
//
//
// 如果无法做到，就返回 [-1, -1]。
//
// 注意，在考虑每个部分所表示的二进制时，应当将其看作一个整体。例如，[1,1,0] 表示十进制中的 6，而不会是 3。此外，前导零也是被允许的，所以 [0,
//1,1] 和 [1,1] 表示相同的值。
//
//
//
// 示例 1：
//
//
//输入：arr = [1,0,1,0,1]
//输出：[0,3]
//
//
// 示例 2：
//
//
//输入：arr = [1,1,0,1,1]
//输出：[-1,-1]
//
// 示例 3:
//
//
//输入：arr = [1,1,0,0,1]
//输出：[0,2]
//
//
//
//
// 提示：
//
//
//
// 3 <= arr.length <= 3 * 10⁴
// arr[i] 是 0 或 1
//
//
// Related Topics 数组 数学 👍 205 👎 0

package main

import "fmt"

func main() {
	arr := []int{1, 1, 0, 0, 1}
	//arr = []int{1, 1, 1, 0, 0, 1,
	//	1, 0, 1, 0, 1, 1,
	//	1, 1, 1, 1}
	arr = []int{1, 1,
		1, 1,
		0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	parts := threeEqualParts(arr)
	fmt.Println(parts)
}

// leetcode submit region begin(Prohibit modification and deletion)
func threeEqualParts(arr []int) []int {
	// 三等分 1
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%3 != 0 {
		return []int{-1, -1}
	}
	if sum == 0 {
		return []int{0, 2}
	}
	ret := make([]int, 2)
	for i, j, p, t := 0, 0, 0, sum/3; i < 2; i++ { // 平分 1
		for {
			p += arr[j]
			j++
			if p == t {
				ret[i] = j
				p = 0
				break
			}
		}
	}
	zeroCnt := 0
	for i := len(arr) - 1; arr[i] == 0; i-- {
		zeroCnt++ // 第三部分的后缀 0
	}
	i, j, k := ret[0]-1+zeroCnt, ret[1]-1+zeroCnt, len(arr)-1 // 三部分的终止索引
	if i >= ret[1] || j >= k-zeroCnt {                        // 不合法的分法：有交集了
		return []int{-1, -1}
	}
	ret[0], ret[1] = i, j+1 // 返回的数据
	for ; i >= 0 && j > ret[0] && k >= ret[1]; i, j, k = i-1, j-1, k-1 {
		if arr[i] == arr[j] && arr[j] == arr[k] { // 三部分相同
			continue
		}
		return []int{-1, -1}
	}
	return ret

	// 滑动窗体？
}

//leetcode submit region end(Prohibit modification and deletion)
