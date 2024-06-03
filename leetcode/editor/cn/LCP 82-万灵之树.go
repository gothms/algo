package main

import "fmt"

func main() {
	// 111x91x991
	v := 111899134991 // 0
	v = 6890034004    // 0
	//v = 111           // 6
	//v = 91 // 00
	//v = 991 // 004
	p := 7
	fmt.Println(v % p)
}

// leetcode submit region begin(Prohibit modification and deletion)

const lcp82 = 5

var tree82 [][][]int
var fill82 [][][]int

func init() {
	// dp
	// 二叉树中序遍历的递归逻辑
	tree82 = make([][][]int, lcp82)
	fill82 = make([][][]int, lcp82)
	tree82[0] = [][]int{{1, 0, 9}} // gem.len == 0
	fill82[0] = [][]int{{1}}
	for i, c := 1, 1; i < lcp82; i++ {
		c = c * (i<<2 - 2) / (i + 1) // 卡特兰数
		tree82[i] = make([][]int, 0, c)
	}
	for i, lastL := 1, 3; i < lcp82; i++ {
		lastL += 5               // cap
		for j := 0; j < i; j++ { // 左子树的节点数
			for x, l := range tree82[j] { // 左子树
				for y, r := range tree82[i-j-1] { // 右子树
					// 状态转移方程：1 + 左子树 / 右子树 + 9
					temp := make([]int, 0, lastL)
					temp = append(append(append(append(temp, 1), l...), r...), 9)
					tree82[i] = append(tree82[i], temp)

					f := make([]int, 0, len(l)+len(r))
					fl, fr := fill82[j][x], fill82[i-j-1][y]
					for _, v := range fl {
						f = append(f, v+1)
					}
					k := len(l) + 1
					for _, v := range fr {
						f = append(f, v+k)
					}
					fill82[i] = append(fill82[i], f)
				}
			}
		}
	}
	//for i, arr := range tree82 {
	//	fmt.Println(i, arr, len(arr))
	//}
	for i, arr := range fill82 {
		fmt.Println(i, arr, len(arr))
	}
	//for _, t := range tree82 {
	//	for _, v := range t {
	//		fmt.Print(len(v), " ")
	//	}
	//	fmt.Println()
	//}
}
func treeOfInfiniteSouls(gem []int, p int, target int) int {
	// 第一步：所有的二叉树搜索后的值序列存入 tree82，其中 0 值为待填入的 gem[i] 值
	// 第二步：将 gem[i] 填入 fill82 标记 tree82 的位置
	// 1.排列：有重复数字时，比如 [3 3 1 1 1 21]，每次成功结果要 *6。即重复的排列只用计算一次
	// 2.参考 % 定理：如 111x91x991 可以简化为 6x00x004，同样的 gem[i] 也可以先 %p，但要保证 补0
	// 第三步：判断 v%p ?= target
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
