//圆内一个 有效切割 ，符合以下二者之一：
//
//
// 该切割是两个端点在圆上的线段，且该线段经过圆心。
// 该切割是一端在圆心另一端在圆上的线段。
//
//
// 一些有效和无效的切割如下图所示。
//
//
//
// 给你一个整数 n ，请你返回将圆切割成相等的 n 等分的 最少 切割次数。
//
//
//
// 示例 1：
//
//
//
//
//输入：n = 4
//输出：2
//解释：
//上图展示了切割圆 2 次，得到四等分。
//
//
// 示例 2：
//
//
//
//
//输入：n = 3
//输出：3
//解释：
//最少需要切割 3 次，将圆切成三等分。
//少于 3 次切割无法将圆切成大小相等面积相同的 3 等分。
//同时可以观察到，第一次切割无法将圆切割开。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 100
//
//
// Related Topics 几何 数学 👍 17 👎 0

package main

import "fmt"

func main() {
	fmt.Println(24 &^ 8 >> 1)

	var prints []func()
	for _, v := range []int{1, 2, 3} {
		prints = append(prints, func() {
			fmt.Println(v)
		})
		//prints[len(prints)-1]()
	}
	for _, f := range prints {
		f() // 3 3 3
	}
}

/*
思路：
	1.如果 n 为奇数，则每一刀都是半径
		一共需要 n 刀
	2.如果 n 为偶数，则第一刀是直径，后面的每一刀会不会和第一刀相交呢？
		n/2 为奇数：左半部分用半径切 n/2-1 刀
			但是把这些半径延伸为直接时，右半部分也被分成了 n/2 份
			所以一共需要 1 + n/2 刀
		n/2 为偶数：第二刀和后面任何一刀，也都会和第一刀相交
			即每一刀都 有且只有 多切出2等份
			所以一共也是需要 1 + n/2 刀
	3.如果 n == 1,
		一共需要 0 刀
*/

//leetcode submit region begin(Prohibit modification and deletion)
func numberOfCuts(n int) int {
	if n == 1 {
		return 0
	}
	if n&1 > 0 {
		return n
	}
	//last := n & (-n)
	////onesCnt := bits.OnesCount(uint(last - 1))
	////return (n>>onesCnt-1)<<(onesCnt-1) + last>>1
	//return n&^last>>1 + last>>1
	return n >> 1
}

//leetcode submit region end(Prohibit modification and deletion)
