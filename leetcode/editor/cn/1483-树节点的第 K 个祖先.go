package main

import (
	"fmt"
)

func main() {
	n := 5
	parent := []int{-1, 0, 0, 0, 3}
	con := Constructor(n, parent)
	ret := con.GetKthAncestor(1, 5)
	fmt.Println(ret)
	ret = con.GetKthAncestor(3, 2)
	fmt.Println(ret)
	ret = con.GetKthAncestor(0, 1)
	fmt.Println(ret)
	ret = con.GetKthAncestor(3, 1)
	fmt.Println(ret)
	ret = con.GetKthAncestor(3, 5)
	fmt.Println(ret)
}

/*
思路：暴力(超时)
	1.某个节点 node 的父节点 等于 parent[node]
	2.那么一直迭代 node = parent[node]，直到找 k 次 或 node==-1
		最后 node的值 就是 初始node 的第 k 个祖先节点
	3.由于是暴力寻找，时间复杂度是 O(k)
思路：倍增
	1.相对暴力方式的 O(k) 次查找，使用倍增的时间复杂度是 O(log k)
	2.举例：求 node 的第 2^j 祖先节点，假设 j = 4，就是求 node 的第 16 个祖先节点
		先求 node 的第 8 个祖先节点（假设这个祖先节点是 node-8）
		再求 node-8 的第 8 个祖先节点(node-16)，node-16 就是 node 的第 16 个祖先
	3.根据上面的例子，对 [0,n) 的每个节点
		3.1.可以预存它的第 2^j 祖先节点，存入一个二维“数组”中
			memo[i][j]：节点 i 的第 2^j 个祖先节点是 memo[i][j]
		3.2.对 node 节点从高位到低位取 二进制，如果二进制位为 1
			则说明可以 倍减，求得 node 的第 2^(j-1) 个祖先是哪个节点
		3.3.针对3.1.的思路，我们先求得每个节点的第 2^0 个祖先节点
			这些第 2^0 个祖先节点的数据，就是后面求第 2^1 个祖先节点时的 原始数据
			后面就可以依次求 2^2 ... 2^bit
		3.4.第 0 个祖先节点 =parent[node]
	4.先计算出 j 的最大值 bit，即最大倍增次数
		bit := bits.Len(uint(n))
	5.在查询 node 节点的第 k 个祖先节点时，就可以按照上面的思路（尤其是 3.2.）查询 memo
*/

// leetcode submit region begin(Prohibit modification and deletion)
type TreeAncestor struct {
}

func Constructor(n int, parent []int) TreeAncestor {

}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

//type TreeAncestor struct {
//	memo [][]int
//	n    int
//}
//
//func Constructor(n int, parent []int) TreeAncestor {
//	//bit := bits.Len(uint(n)) + 1 // +1 多余
//	bit := bits.Len(uint(n)) // 小于 2^(bit+1) 个祖先
//	memo := make([][]int, n)
//	for i := 0; i < n; i++ {
//		memo[i] = make([]int, bit)
//		memo[i][0] = parent[i] // 第 0 个祖先
//	}
//	for j := 1; j < bit; j++ { // 第 0 个祖先已记录
//		memo[0][j] = -1          // 处理 0 节点
//		for i := 1; i < n; i++ { // 0 是根节点
//			// memo[binary][j-1] 已经遍历过，如果 =-1 说明已经超过 根节点
//			if binary := memo[i][j-1]; binary != -1 {
//				// i 的第 2^j 个祖先，是 i 的第 2^(j-1) 个祖先的第 2^(j-1) 个祖先
//				memo[i][j] = memo[binary][j-1]
//			} else {
//				memo[i][j] = -1
//			}
//		}
//	}
//	return TreeAncestor{memo, bit}
//
//	//bit := bits.Len(uint(n)) // 小于 2^(bit+1) 个祖先
//	//memo := make([][]int, n)
//	//for i := 0; i < n; i++ {
//	//	memo[i] = make([]int, bit)
//	//	for j := 1; j < bit; j++ {
//	//		memo[i][j] = -1 // 都初始化为 -1
//	//	}
//	//	memo[i][0] = parent[i] // 第 0 个祖先
//	//}
//	//for j := 1; j < bit; j++ { // 第 0 个祖先已记录
//	//	for i := 1; i < n; i++ { // 0 是根节点
//	//		if binary := memo[i][j-1]; binary != -1 {
//	//			// i 的第 2^j 个祖先，是 i 的第 2^(j-1) 个祖先的第 2^(j-1) 个祖先
//	//			memo[i][j] = memo[binary][j-1]
//	//		} // memo[binary][j-1] 已经遍历过，如果 =-1 说明已经超过 根节点
//	//	}
//	//}
//	//return TreeAncestor{memo, bit}
//}
//func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
//	// 采用遍历的方式：node = this.t[node] 超时
//	//for ; node > 0 && k > 0; node = this.t[node] {
//	//	k--
//	//}
//	//if k > 0 {
//	//	return -1
//	//}
//	//return node
//
//	for j := 0; j < this.n; j++ { // 写为 n 没问题，但其实是 j < bits.Len(uint(k))
//		if k>>j&1 == 0 {
//			continue
//		} // k 的二进制位为 1，才倍减
//		node = this.memo[node][j] // 倍减 k>>j&1
//		if node == -1 {           // 已超过 根节点
//			break
//		}
//	}
//	return node
//
//	//for i := bits.Len(uint(k)) - 1; i >= 0; i-- {	// 其他写法
//	//	if (1<<i)&k == 0 {
//	//		continue
//	//	}
//	//	if node = this.memo[node][i]; node == -1 {
//	//		break
//	//	}
//	//}
//	//return node
//}

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */
