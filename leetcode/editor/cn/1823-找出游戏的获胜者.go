package main

import "fmt"

func main() {
	n, k := 5, 2 // 3
	n, k = 4, 2  // 1
	n, k = 3, 2  // 3
	n, k = 2, 2  // 1
	n, k = 8, 8  // 4
	winner := findTheWinner(n, k)
	fmt.Println(winner)
}

/*
约瑟夫问题
    https://zh.wikipedia.org/wiki/%E7%BA%A6%E7%91%9F%E5%A4%AB%E6%96%AF%E9%97%AE%E9%A2%98
    lc-1823
    lc-LCR 187

递归
    定义 f(n,k)
        n：编号总数
        k：每一轮淘汰下一个编号的计数
        f(n,k)：最终的获胜者编号
    思路
        f(n-1,k) 一共移动了 (n-2)*k 次，f(n,k) 一共移动了 (n-1)*k 次
        f(n,k) 相当于比 f(n-1,k) 多移动了 k 位
        设 f(n-1,k) = x，那么在 x 位置再移动 k 位，就是 f(n,k) 的值
    递推公式
        f(n,k) = (f(n-1,k) + k) % n
        f(n,k) = (f(n-1,k) + k - 1) % n + 1：编号从 1 开始
    重点
        自顶向下
            从 n 个数开始进行 n-1 轮淘汰后剩下一个数
            假设经过 x 轮淘汰后，剩余的编号为 [2 4 5 8]
            数学推导请参考：https://leetcode.cn/problems/find-the-winner-of-the-circular-game/solutions/1467699/by-fuxuemingzhu-laof/
        自底向上：“结果导向” & 数组索引
            从 1 个数开始淘汰，增加到 n 个数的淘汰
            第 x 轮编号为 [1 2 3 ... x]，假设 x+1 轮淘汰时最终剩下的编号为 i，则只用插入 i 变成 [1 2 3 ... i ... x x+1]
            也就是只用考虑淘汰后剩余的编号“位于第几位”，用数组的“索引编号”来快速定位最终剩下的编号
*/

// leetcode submit region begin(Prohibit modification and deletion)
func findTheWinner(n int, k int) int {
	ans := 1
	for i := 2; i <= n; i++ {
		ans = (ans+k-1)%i + 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func findTheWinner_(n int, k int) int {
	// 迭代
	//ret := 0 // 从 0 开始，而题意是从 1 开始
	//for i := 2; i <= n; i++ {
	//	ret = (ret + k) % i // 当前长度为 i
	//}
	//return ret + 1 // 转为从 1 开始

	// 递归
	//if n == 1 {
	//	return 1
	//} // f(n,k) 表示 n 个小伙伴，每一轮离开的小伙伴的计数为 k时，最终的获胜者编号
	//return (findTheWinner(n-1, k)+k-1)%n + 1

	// 模拟：n 比较大时，超时
	q := make([]int, n)
	for i := 0; i < n; i++ {
		q[i] = i + 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < k; j++ {
			q = append(q, q[0])[1:] // 数一个，移动一个
		}
		q = q[1:] // 淘汰
	}
	return q[0]
}
