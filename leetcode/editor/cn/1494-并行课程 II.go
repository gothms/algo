//给你一个整数 n 表示某所大学里课程的数目，编号为 1 到 n ，数组 relations 中， relations[i] = [xi, yi] 表示一个先
//修课的关系，也就是课程 xi 必须在课程 yi 之前上。同时你还有一个整数 k 。
//
// 在一个学期中，你 最多 可以同时上 k 门课，前提是这些课的先修课在之前的学期里已经上过了。
//
// 请你返回上完所有课最少需要多少个学期。题目保证一定存在一种上完所有课的方式。
//
//
//
// 示例 1：
//
//
//
//
//输入：n = 4, relations = [[2,1],[3,1],[1,4]], k = 2
//输出：3
//解释：上图展示了题目输入的图。在第一个学期中，我们可以上课程 2 和课程 3 。然后第二个学期上课程 1 ，第三个学期上课程 4 。
//
//
// 示例 2：
//
//
//
//
//输入：n = 5, relations = [[2,1],[3,1],[4,1],[1,5]], k = 2
//输出：4
//解释：上图展示了题目输入的图。一个最优方案是：第一学期上课程 2 和 3，第二学期上课程 4 ，第三学期上课程 1 ，第四学期上课程 5 。
//
//
// 示例 3：
//
//
//输入：n = 11, relations = [], k = 2
//输出：6
//
//
//
//
// 提示：
//
//
// 1 <= n <= 15
// 1 <= k <= n
// 0 <= relations.length <= n * (n-1) / 2
// relations[i].length == 2
// 1 <= xi, yi <= n
// xi != yi
// 所有先修关系都是不同的，也就是说 relations[i] != relations[j] 。
// 题目输入的图是个有向无环图。
//
//
// Related Topics 位运算 图 动态规划 状态压缩 👍 167 👎 0

package main

import (
	"fmt"
	"math/bits"
)

func main() {
	//x := 7
	//for i := x; i > 0; i = (i - 1) & x {
	//	fmt.Println(i)
	//}

	//for i := 7; i < 15; i++ {
	//	fmt.Println(i, i&(i-1), i&-i)
	//}

	n, k := 4, 2
	relations := [][]int{{2, 1}, {3, 1}, {1, 4}}
	//relations = [][]int{{1, 2}, {1, 3}, {1, 4}}
	//n, k = 10, 2
	//relations = [][]int{{2, 1}, {3, 1}, {1, 4}, {5, 7}, {7, 8}, {6, 8}, {9, 10}, {8, 10}}
	//n, k = 5, 3 // 5
	//relations = [][]int{{1, 5}, {1, 3}, {1, 2}, {4, 2}, {4, 5}, {2, 5}, {1, 4}, {4, 3}, {3, 5}, {3, 2}}
	semesters := minNumberOfSemesters(n, relations, k)
	fmt.Println(semesters)
}

/*
思路：并查集+优先队列
	不能实现：0 <= relations.length <= n * (n-1) / 2，表示在无环的情况下，可能有最大边数
	所以，只能用邻接表/矩阵来记录边
思路：状压dp
	1.n<=15，n门课程状态压缩为，从低位到高位第 i-1 位的二进制位 1 与第 i 门课程一一映射
		如 6 的二进制位为 110，表示 第 2、3 门课程
	2.使用状压dp，直接来看状压dp的求解过程
		2.1.遍历 n 门课程的所有组合情况，试图求出这些组合情况下需要多少个学期修完课程（每学期最多修 k 门）
			如有三门课程 1 2 3，需要求出以下所有组合的情况，共有 2^n 种组合情况
			[0] [1] [2] [1,2] [3] [1,3] [2,3] [1,2,3]，即这些组合情况为 三门课程 的子集
		2.2.dp[i]的定义，i = [0,2^n-1]，i与课程组合的映射关系：
			dp[0] 映射 课程组合 [0]
			dp[1] 映射 [1]
			dp[3] 映射 [1,2]
			dp[7] 映射 [1,2,3]：因为 7 的二进制位 111
			同时，i 的值为 2^n-1 的子集，dp[i] 即为修满 i 的二进制表示的课程，需要多少学期
		2.3.状态转移方程推导
			假设修满 i 的二进制表示的课程，我们遍历 i 的二进制子集 j，i j 之间的“差值”为 i^j
			即修了 j 后，还需要修 i^j 的可成，就可以修满 i
		2.4.如果 i^j 的 1 的数量，即课程数量 <= k，有：
			dp[i] = minVal(dp[i], dp[j^i] + 1)
			且 dp[0]=0
	3.定义课程之间的关系
		3.1.needs[i] 表示修 i 课程，需要修哪些 先修课
			根据 relations 有向图关系，初始化 needs，其中课程编号 -1，方便状态压缩dp的计算
		3.2.当前 i 需要的先修课，例 i 的二进制为 11010
			修满 11010 需要修满 11000 且修满 00010，即后二者的并集
			needs[i] = needs[i&^last] | needs[last]
		3.3.如果修满 i 需要的先修课 needs[i] 不是 i 的子集，说明还有 先修课 没修
			即课程编号更大的某些课，需要先修，才能修满 i，则 i 无法修完，dp[i] 的值不存在
		3.4.否则说明 i 能修满，且只要修了 needs[i]^i
			令 take := needs[i] ^ i，如果 take 的可成数量（二进制中 1 的数量） <= k，则
			dp[i] = minVal(dp[i], dp[take^i]+1)
		3.5.否则，遍历 take 的子集（2.4.）求出最小 dp[i]
	补充：
		needs 中，i 表示课程的集合，needs[i] 表示集合 i 需要的先修课程
*/
// leetcode submit region begin(Prohibit modification and deletion)
func minNumberOfSemesters(n int, relations [][]int, k int) int {
	// 复习1
	minVal := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	m := 1 << n
	dp, pre := make([]int, m), make([]int, m)
	for _, edge := range relations {
		pre[1<<(edge[1]-1)] |= 1 << (edge[0] - 1)
	}
	for i := 1; i < m; i++ {
		dp[i] = m
	}
	for i := 1; i < m; i++ {
		// 求出 pre[i]
		last := i & -i
		pre[i] = pre[i^last] | pre[last]
		// 判断 pre[i] 是否合法
		if pre[i]|i != i {
			continue
		}
		// 求 dp[i]：一次
		take := pre[i] ^ i // i-pre[i]
		if bits.OnesCount(uint(take)) <= k {
			dp[i] = minVal(dp[i], dp[i^take]+1)
			continue
		}
		// 求 dp[i]：多次遍历
		for sub := take & (take - 1); sub > 0; sub = take & (sub - 1) {
			if bits.OnesCount(uint(sub)) <= k {
				dp[i] = minVal(dp[i], dp[i^sub]+1)
			}
		}
	}
	//fmt.Println(pre)
	return dp[m-1]

	// 状压dp：终版
	//minVal := func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//m := 1 << n // 状态压缩
	//dp, needs := make([]int, m), make([]int, m)
	//for i := 1; i < m; i++ {
	//	dp[i] = m // 无法修完
	//}
	//for _, edge := range relations {
	//	needs[1<<(edge[1]-1)] |= 1 << (edge[0] - 1)
	//} // 根据课程的有向图，初始化先修课程
	////fmt.Println(needs)
	//for i, last := 1, 0; i < m; i++ {
	//	last = i & -i                           // 课程编号最小的那门课
	//	needs[i] = needs[i&^last] | needs[last] // 并集
	//	if needs[i]|i != i {                    // 某些先修课的编程编号更大，即 needs[i] 必须是 i 的子集
	//		continue
	//	}
	//	take := needs[i] ^ i // 修满 take，则可以修满 i，即 take := i - needs[i]
	//	if bits.OnesCount(uint(take)) <= k {
	//		dp[i] = minVal(dp[i], dp[take^i]+1)
	//		continue // 一学期修完 take
	//	}
	//	for sub := take & (take - 1); sub > 0; sub = take & (sub - 1) { // 遍历 take 的子集
	//		if bits.OnesCount(uint(sub)) <= k {
	//			dp[i] = minVal(dp[i], dp[sub^i]+1)
	//		}
	//	} // 不止一学期修完 take
	//}
	//return dp[m-1]

	// 邻接矩阵+并查集+优先队列：未写完
	//parent, matrix := make([]int, n+1), make([][]int, n+1)
	//for i := 1; i <= n; i++ {
	//	matrix[i] = make([]int, n+1)
	//}
	//for i := 1; i <= n; i++ {
	//	parent[i] = i
	//}
	//for _, edge := range relations {
	//	semUnion(parent, edge[1], edge[0])
	//	matrix[edge[0]][edge[1]] = 1
	//}
	//fmt.Println(matrix)
	//memo := make(map[int]bool)
	//for i := 1; i <= n; i++ {
	//	if r := semFind(parent, i); !memo[r] {
	//		memo[r] = true
	//	}
	//}
	//fmt.Println(parent) // 未指向最后的课程
	//fmt.Println(memo)
	//bfs := make(map[int][]int, len(memo))
	//for i := 0; len(memo) > 0; i++ {
	//	temp := make(map[int]bool)
	//	for cId := range memo {
	//		p := semFind(parent, cId)
	//		if len(bfs[p]) == i {
	//			bfs[p] = append(bfs[p], 0)
	//		}
	//		bfs[p][i]++
	//		//for _, v := range adj[cId] {
	//		//	temp[v] = true
	//		//}
	//	}
	//	memo = temp
	//}
	//fmt.Println("bfs", bfs)
	//sh := semHeap{}
	//for uni, cs := range bfs {
	//	heap.Push(&sh, [2]int{uni, len(cs) - 1})
	//	cs[len(cs)-1]--
	//}
	//cnt := 0
	//for len(sh) > 0 {
	//	cache := make([][2]int, 0)
	//	for i := 0; i < k && len(sh) > 0; i++ {
	//		top := heap.Pop(&sh).([2]int)
	//		uni := bfs[top[0]]
	//		//if top[1] == 0 {
	//		//	continue
	//		//}
	//		if uni[top[1]] > 0 {
	//			uni[top[1]]--
	//			heap.Push(&sh, [2]int{top[0], top[1]})
	//		} else if uni[top[1]] == 0 && top[1] > 0 {
	//			uni[top[1]-1]--
	//			if uni[top[1]-1] > 0 {
	//				heap.Push(&sh, [2]int{top[0], top[1] - 1})
	//			} else {
	//				cache = append(cache, [2]int{top[0], top[1] - 1})
	//			}
	//		}
	//	}
	//	for _, cc := range cache {
	//		heap.Push(&sh, cc)
	//	}
	//	cnt++
	//}
	//return cnt

	// 并查集+优先队列：不能实现
	//parent, adj := make([]int, n+1), make([][]int, n+1)
	////parent, adj, adjR := make([]int, n+1), make([][]int, n+1), make([][]int, n+1)
	//for i := 1; i <= n; i++ {
	//	parent[i] = i
	//}
	//for _, edge := range relations {
	//	semUnion(parent, edge[1], edge[0])
	//	adj[edge[1]] = append(adj[edge[1]], edge[0])
	//	//adjR[edge[0]] = append(adjR[edge[0]], edge[1])
	//}
	//fmt.Println(adj)
	////fmt.Println(adjR)
	//memo := make(map[int]bool)
	//for i := 1; i <= n; i++ {
	//	if r := semFind(parent, i); !memo[r] {
	//		memo[r] = true
	//	}
	//}
	////fmt.Println(parent)
	//fmt.Println(memo)
	//bfs := make(map[int][]int, len(memo))
	//for i := 0; len(memo) > 0; i++ {
	//	temp := make(map[int]bool)
	//	for cId := range memo {
	//		p := semFind(parent, cId)
	//		if len(bfs[p]) == i {
	//			bfs[p] = append(bfs[p], 0)
	//		}
	//		bfs[p][i]++
	//		for _, v := range adj[cId] {
	//			temp[v] = true
	//		}
	//	}
	//	memo = temp
	//}
	//fmt.Println("bfs", bfs)
	//sh := semHeap{}
	//for uni, cs := range bfs {
	//	heap.Push(&sh, [2]int{uni, len(cs) - 1})
	//	cs[len(cs)-1]--
	//}
	//cnt := 0
	//for len(sh) > 0 {
	//	cache := make([][2]int, 0)
	//	for i := 0; i < k && len(sh) > 0; i++ {
	//		top := heap.Pop(&sh).([2]int)
	//		uni := bfs[top[0]]
	//		//if top[1] == 0 {
	//		//	continue
	//		//}
	//		if uni[top[1]] > 0 {
	//			uni[top[1]]--
	//			heap.Push(&sh, [2]int{top[0], top[1]})
	//		} else if uni[top[1]] == 0 && top[1] > 0 {
	//			uni[top[1]-1]--
	//			if uni[top[1]-1] > 0 {
	//				heap.Push(&sh, [2]int{top[0], top[1] - 1})
	//			} else {
	//				cache = append(cache, [2]int{top[0], top[1] - 1})
	//			}
	//		}
	//	}
	//	for _, cc := range cache {
	//		heap.Push(&sh, cc)
	//	}
	//	cnt++
	//}
	//return cnt

	// lc
	//dp := make([]int, 1<<n)
	//for i := range dp {
	//	dp[i] = math.MaxInt32
	//}
	//need := make([]int, 1<<n)
	//for _, edge := range relations {
	//	need[1<<(edge[1]-1)] |= 1 << (edge[0] - 1)
	//	//fmt.Println(edge, need)
	//}
	//fmt.Println(need)
	//dp[0] = 0
	//for i := 1; i < (1 << n); i++ {
	//	need[i] = need[i&(i-1)] | need[i&-i]
	//	//fmt.Println(i, i&(i-1), i&-i, need[i])
	//	fmt.Println(i, need[i])
	//	if (need[i] | i) != i { // i 中有任意一门课程的前置课程没有完成学习
	//		continue
	//	}
	//	fmt.Println("???", i)
	//	valid := i ^ need[i]                  // 当前学期可以进行学习的课程集合
	//	if bits.OnesCount(uint(valid)) <= k { // 如果个数小于 k，则可以全部学习，不再枚举子集
	//		dp[i] = minVal(dp[i], dp[i^valid]+1)
	//	} else {
	//		for sub := valid; sub > 0; sub = (sub - 1) & valid {
	//			if bits.OnesCount(uint(sub)) <= k {
	//				dp[i] = minVal(dp[i], dp[i^sub]+1)
	//			}
	//		}
	//	}
	//}
	//fmt.Println(need)
	//for i := 0; i < 1<<n; i++ {
	//	if dp[i] > 1000 {
	//		dp[i] = 0
	//	}
	//}
	//fmt.Println(dp)
	//return dp[(1<<n)-1]
}
func semFind(parent []int, p int) int {
	for parent[p] != p {
		p, parent[p] = parent[p], parent[parent[p]]
	}
	return p
}
func semUnion(parent []int, p, q int) bool {
	rootP, rootQ := semFind(parent, p), semFind(parent, q)
	if rootP != rootQ {
		parent[rootQ] = rootP
		return true
	}
	return false
}

type semHeap [][2]int

func (s semHeap) Len() int {
	return len(s)
}
func (s semHeap) Less(i, j int) bool {
	return s[i][1] > s[j][1]
}
func (s semHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s *semHeap) Push(x interface{}) {
	*s = append(*s, x.([2]int))
}
func (s *semHeap) Pop() interface{} {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
