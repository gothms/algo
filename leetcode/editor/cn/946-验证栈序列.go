//给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时
//，返回 true；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
//输出：true
//解释：我们可以按以下顺序执行：
//push(1), push(2), push(3), push(4), pop() -> 4,
//push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
//
//
// 示例 2：
//
//
//输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
//输出：false
//解释：1 不能在 2 之前弹出。
//
//
//
//
// 提示：
//
//
// 1 <= pushed.length <= 1000
// 0 <= pushed[i] <= 1000
// pushed 的所有元素 互不相同
// popped.length == pushed.length
// popped 是 pushed 的一个排列
//
//
// Related Topics 栈 数组 模拟 👍 430 👎 0

package main

import "fmt"

func main() {
	push := []int{1, 2, 3, 4, 5}
	pop := []int{4, 5, 3, 2, 1}
	push = []int{1, 2, 3, 4, 5}
	pop = []int{4, 3, 5, 1, 2}
	//push = []int{1, 0}
	//pop = []int{1, 0}
	//push = []int{0, 2, 1}
	//pop = []int{0, 1, 2}
	//push = []int{3, 0, 1, 2}
	//pop = []int{3, 2, 1, 0}
	sequences := validateStackSequences(push, pop)
	fmt.Println(sequences)
}

// leetcode submit region begin(Prohibit modification and deletion)
func validateStackSequences(pushed []int, popped []int) bool {
	// 模拟
	//n := len(pushed)
	//st := make([]int, 0, n)
	//for i, j := 0, 0; j < n; j++ {
	//	if len(st) > 0 && st[len(st)-1] == popped[j] { // 出栈
	//		st = st[:len(st)-1]
	//		continue
	//	}
	//	for i < n && pushed[i] != popped[j] { // 入栈
	//		st = append(st, pushed[i])
	//		i++
	//	}
	//	i++
	//}
	//return len(st) == 0

	// lc
	st := make([]int, 0, len(pushed))
	j := 0
	for _, v := range pushed {
		st = append(st, v)
		for ; len(st) > 0 && st[len(st)-1] == popped[j]; j++ {
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
