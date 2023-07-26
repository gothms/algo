//在经典汉诺塔问题中，有 3 根柱子及 N 个不同大小的穿孔圆盘，盘子可以滑入任意一根柱子。一开始，所有盘子自上而下按升序依次套在第一根柱子上(即每一个盘子只
//能放在更大的盘子上面)。移动圆盘时受到以下限制: (1) 每次只能移动一个盘子; (2) 盘子只能从柱子顶端滑出移到下一根柱子; (3) 盘子只能叠在比它大的盘
//子上。
//
// 请编写程序，用栈将所有盘子从第一根柱子移到最后一根柱子。
//
// 你需要原地修改栈。
//
// 示例1:
//
//  输入：A = [2, 1, 0], B = [], C = []
// 输出：C = [2, 1, 0]
//
//
// 示例2:
//
//  输入：A = [1, 0], B = [], C = []
// 输出：C = [1, 0]
//
//
// 提示:
//
//
// A中盘子的数目不大于14个。
//
//
// Related Topics 递归 数组 👍 216 👎 0

package main

import (
	"fmt"
)

func main() {
	A := []int{5, 4, 3, 2, 1, 0}
	//A = []int{2, 1, 0}
	ints := hanota(A, nil, nil)
	fmt.Println(ints)
}

/*
汉诺塔问题
思路：递归
	1.对于任意个盘子 n，要把它们从 from 柱子，移动到 to 柱子
		只需要把 n-1 个盘子从 from 先移动到 other 柱子
		再把最大的盘子从 from 移动到 to
		最后再把 n-1 个盘子，从 other 移动到 to
	2.递归精髓
		在于 n 到 n-1 的过程，以及 1 的边界条件
		而不是人肉递归 n 到 1 的过程
*/
//leetcode submit region begin(Prohibit modification and deletion)
func hanota(A []int, B []int, C []int) []int {
	// recursive
	var dfs func(*[]int, *[]int, *[]int, int)
	dfs = func(f *[]int, t *[]int, o *[]int, i int) {
		if i == 1 {
			*f, *t = (*f)[:len(*f)-1], append(*t, (*f)[len(*f)-1])
			return
		}
		dfs(f, o, t, i-1)
		*f, *t = (*f)[:len(*f)-1], append(*t, (*f)[len(*f)-1])
		dfs(o, t, f, i-1)
	}
	dfs(&A, &C, &B, len(A))
	return C
}
	// 没有扩容
	//n := len(A)
	//if n == 0 {
	//	return C
	//}
	//B, C = make([]int, n), make([]int, n)
	//var dfs func([]int, []int, []int, *int, *int, *int, int)
	//dfs = func(f, t, o []int, i, j, k *int, m int) {
	//	if m == 1 {
	//		// 赋值取当前索引，取值取前一个索引
	//		t[*j], *i, *j = f[*i-1], *i-1, *j+1
	//		return
	//	}
	//	dfs(f, o, t, i, k, j, m-1)
	//	t[*j], *i, *j = f[*i-1], *i-1, *j+1
	//	dfs(o, t, f, k, j, i, m-1)
	//}
	//i, j, k := n, 0, 0
	//dfs(A, C, B, &i, &j, &k, n)
	//return C

	//n := len(A)
	//if n == 0 {
	//	return C
	//}
	//var dfs func(*[]int, *[]int, *[]int, int)
	//dfs = func(f, t, o *[]int, i int) {
	//	if i == 1 {
	//		*t, *f = append(*t, (*f)[len(*f)-1]), (*f)[:len(*f)-1]
	//		return
	//	}
	//	dfs(f, o, t, i-1)
	//	*t, *f = append(*t, (*f)[len(*f)-1]), (*f)[:len(*f)-1]
	//	dfs(o, t, f, i-1)
	//}
	//dfs(&A, &C, &B, n)
	//return C
}

//leetcode submit region end(Prohibit modification and deletion)
