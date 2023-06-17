//给你一个下标从 0 开始、大小为 n x n 的整数矩阵 grid ，返回满足 Ri 行和 Cj 列相等的行列对 (Ri, Cj) 的数目。
//
// 如果行和列以相同的顺序包含相同的元素（即相等的数组），则认为二者是相等的。
//
//
//
// 示例 1：
//
//
//
//
//输入：grid = [[3,2,1],[1,7,6],[2,7,7]]
//输出：1
//解释：存在一对相等行列对：
//- (第 2 行，第 1 列)：[2,7,7]
//
//
// 示例 2：
//
//
//
//
//输入：grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
//输出：3
//解释：存在三对相等行列对：
//- (第 0 行，第 0 列)：[3,1,2,2]
//- (第 2 行, 第 2 列)：[2,4,2,2]
//- (第 3 行, 第 2 列)：[2,4,2,2]
//
//
//
//
// 提示：
//
//
// n == grid.length == grid[i].length
// 1 <= n <= 200
// 1 <= grid[i][j] <= 10⁵
//
//
// Related Topics 数组 哈希表 矩阵 模拟 👍 22 👎 0

package main

func main() {

}

/*
思路：Hash 算法
	1.比较两个“串”值是否相等的一个思路，就是比较它们的 hash值 是否相同
		类似RabinKarp算法，依次将每个数字加入 hash值中计算
		最终 hash值 相同的 行/列，就认为这行和这列是相等的
	2.同样的类似RabinKarp，设计一个简单的hash算法：i 是固定的某行/列，prime 为一个质数
		行：hashR = hashR*prime + grid[i][j]
		列：hashC = hashR*prime + grid[j][i]
	3.将得到的 行/列 的 hash值，分别放入两个map
		两个map中相同 hash值 出现的次数的乘积和，即为所求
*/
//leetcode submit region begin(Prohibit modification and deletion)
func equalPairs(grid [][]int) int {
	const prime = 16777619
	cnt, n := 0, len(grid)
	memoR, memoC := make(map[int]int), make(map[int]int)
	hashHelper := func(i int) (r, c int) { // 简单的 hash 算法
		for j := 0; j < n; j++ {
			r, c = r*prime+grid[i][j], c*prime+grid[j][i]
		}
		return
	}
	for i := 0; i < n; i++ {
		r, c := hashHelper(i)
		memoR[r]++ // 行的 hash值
		memoC[c]++ // 列的 hash值
	}
	for k, v := range memoR {
		cnt += v * memoC[k] // 统计 行*列
	}
	return cnt

	//const primeRK = 16777619
	//cnt, n := 0, len(grid)
	//memoR, memoC := make(map[int]int), make(map[int]int)
	//hashHelper := func(i int) { // 简单的 hash 算法
	//	hashR, hashC := 0, 0
	//	for j := 0; j < n; j++ {
	//		hashR, hashC = hashR*primeRK+grid[i][j], hashC*primeRK+grid[j][i]
	//	}
	//	memoR[hashR]++ // 行的 hash值
	//	memoC[hashC]++ // 列的 hash值
	//}
	//for i := 0; i < n; i++ {
	//	hashHelper(i)
	//}
	//for k, v := range memoR {
	//	cnt += v * memoC[k] // 统计：行*列
	//}
	//return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
