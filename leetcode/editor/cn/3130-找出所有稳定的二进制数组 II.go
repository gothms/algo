package main

import "fmt"

func main() {
	zero, one, limit := 0, 13, 2     // test 隔板法
	zero, one, limit = 1, 2, 1       // 1
	zero, one, limit = 2, 1, 2       // 3
	zero, one, limit = 3, 3, 2       // 14
	zero, one, limit = 3, 4, 2       // 18
	zero, one, limit = 6, 7, 3       // 1020
	zero, one, limit = 22, 64, 6     // 277316268
	zero, one, limit = 35, 35, 22    // 782677541
	zero, one, limit = 41, 29, 6     // 533918053
	zero, one, limit = 25, 75, 49    // 500481770
	zero, one, limit = 57, 47, 2     // 973560022
	zero, one, limit = 41, 29, 6     // 533918053
	zero, one, limit = 29, 57, 14    // 676261807
	zero, one, limit = 65, 24, 15    // 739390255
	zero, one, limit = 633, 727, 791 // 719348414
	//zero, one, limit = 28, 68, 31    // 212701517
	//zero, one, limit = 999, 999, 6   // 402567910
	arrays := numberOfStableArrays(zero, one, limit)
	fmt.Println(arrays)
}

/*
隔板法：三种类型
	1.标准型：C(n-1, m-1)
		被分配的 n 个元素无差别
		n 个元素分给 m 个不同对象
		每个对象至少分一个元素
	2.多分型：C(n-(x-1)*m, m-1)
		被分配的 n 个元素无差别
		n 个元素分给 m 个不同对象
		每个对象至少分 x 个元素
	3.少分型：C(n+m-1, m-1)
		被分配的 n 个元素无差别
		n 个元素分给 m 个不同对象
		每个对象可以有任意个元素（包括 0）
*/

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfStableArrays(zero int, one int, limit int) int {
	// 1.隔板法
	// 2.容斥原理
	// 3.0 和 1 分成的堆数的绝对值 ==1

	const mod = 1_000_000_007
	memo := make([][][2]int, zero+1)
	for i := range memo {
		memo[i] = make([][2]int, one+1)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) (res int) {
		if i == 0 { // 递归边界
			if k == 1 && j <= limit {
				return 1
			}
			return
		}
		if j == 0 { // 递归边界
			if k == 0 && i <= limit {
				return 1
			}
			return
		}
		p := &memo[i][j][k]
		if *p != -1 { // 之前计算过
			return *p
		}
		if k == 0 {
			// +mod 保证答案非负
			res = (dfs(i-1, j, 0) + dfs(i-1, j, 1)) % mod
			if i > limit {
				res = (res - dfs(i-limit-1, j, 1) + mod) % mod
			}
		} else {
			res = (dfs(i, j-1, 0) + dfs(i, j-1, 1)) % mod
			if j > limit {
				res = (res - dfs(i, j-limit-1, 0) + mod) % mod
			}
		}
		*p = res // 记忆化
		return
	}
	return (dfs(zero, one, 0) + dfs(zero, one, 1)) % mod
}

// leetcode submit region end(Prohibit modification and deletion)

//const mod, n3130 = 1_000_000_007, 1001
//
//var comb3130 [n3130][n3130]int
//
//func init() {
//	for i := 0; i < n3130; i++ {
//		comb3130[i][0], comb3130[i][i] = 1, 1
//		for j := 1; j < i; j++ { // 杨辉三角预处理组合数
//			// 组合数恒等式，且必须在 init 函数中 mod
//			comb3130[i][j] = (comb3130[i-1][j-1] + comb3130[i-1][j]) % mod
//		}
//	}
//}
//
//func numberOfStableArrays(zero int, one int, limit int) int {
//	// dp+隔板法：终版
//	// TODO 可否剪枝
//	ret, x, y := 0, zero, one
//	if x < y {
//		x, y = y, x
//	}
//	dp := make([][]int, x+2) // x+2：防止当 x==y 时，后面 for i := (x-1)/limit + 1; i <= y+1; i++，值 y+1 越界
//	for i := range dp {
//		dp[i] = make([]int, x+2)
//	}
//	/*
//		定义：
//			partition(x, y)：将 x 个数字分成 y 堆，且每堆的数量不超过 limit
//		隔板法：
//			partition(x, y) = C(x-1,y-1)：将 x 个（相同的）数字分成 y 堆（不考虑合法性）
//			再从 x 里除去 k 个 limit 长度的数，则得到此时的合法排列数
//			合法分割数 =  C(x-1,y-1) - 非法分割数
//		非法分割数：一堆/多堆的数量超过 limit
//			partition(x-k*limit, y)：从 x 里除去 k 个 limit，得到的合法排列数（即合法的 y 堆数）
//			“放置”方法数：再将这 k 个 limit 放入 y 个位置（允许每个位置放多个，即有的位置允许留空）
//			根据隔板法，对 y 个空位每个补偿 1 个数，隔板法得到“放置”方法数的结果：C(y+k-1, y-1)
//			非法分割数 = “放置”方法数 * (x-k*limit, y) 的合法情况数
//		partition(x, y) -= C(y+k-1, y-1) * partition(x-k*limit, y)
//	*/
//	// 1.隔板法，最大数 x 可能的所有分割数
//	for i := 1; i <= x; i++ {
//		for j := (i-1)/limit + 1; j <= i; j++ { // (i-1)/limit + 1，最少堆数，确定 j 的起始点
//			dp[i][j] = comb3130[i-1][j-1] // i 个数字，分成 j 份
//			// 2.合法分割数 = 所有分割数 - 非法分割数
//			// 非法分割数 = k 个 limit 的组合数 * 除去 k 个 limit 的合法分割数
//			for k := 1; i-k*limit >= j; k++ { // 删去超过 limit 的非法组合数
//				// 非法：k 个数分为 j 堆，可以分得 0 个数
//				// 合法：i-k*limit 个数分为 j 堆，且每堆不超过 limit
//				dp[i][j] = (dp[i][j] - (comb3130[k+j-1][j-1]*dp[i-k*limit][j])%mod + mod) % mod
//			}
//		}
//	}
//	//for i, d := range dp {
//	//	fmt.Println(i, d)
//	//}
//	/*
//		计算 x 和 y 组合起来的“排列数”：得到需要的合法排列数目矩阵以后，要计算数组可以怎样排列
//		1.分割数的差值，只可能是 -1 0 1
//			由于要求列表要相间排列，故 0 和 1 的堆数最多相差 1，且如果两者堆数相同，要考虑 0 先还是 1 先的问题
//			堆数大 1 的数占两端，堆数相等则可以 0 占左 1 占右（或反之）
//		2.计算公式
//			将 0 和 1 的合法分割数相乘，本次分割的结果即可得到结果
//			由于 0 和 1 的分割方法都是，最少堆数：每个堆都“塞满”；最多堆数：一个数占一个堆
//			所以这里取这俩的交集，即 i = [(x-1)/limit + 1, y+1]
//			dp[y][i] * dp[x][j]：j = [i-1 i i+1]
//		3.边界考虑
//			遍历的时候，查询 i-1 和 i+1 是否在这个交集的内部
//			如果在，则统计本次的相乘结果
//	*/
//	// 3.i = range(最小堆数, 最大堆数)，计算 x 和 y 组合起来的“排列数”
//	// 区间 [(x-1)/limit + 1, y+1]，其中 y+1 可换成更大值 x（仅当 x == y 时，用 y+1 更差）
//	for i := (x-1)/limit + 1; i <= y+1; i++ { // (x-1)/limit + 1，最少堆数，确定 i 的起始点
//		//if i > (y-1)/limit+1 {  // 有可能 ==（x，y 值相对接近），但此时 dp[y][i-1] == 0，因为 i-1 是非法分割
//		ret += (dp[x][i] * dp[y][i-1]) % mod //
//		//}
//		if i <= y { // i 最大为 y，y 最多分为 i 份
//			ret += (dp[x][i] * dp[y][i] << 1) % mod // 分相同份数，都可以在前/后
//		}
//		if i < y { // i 最大为 y-1，y 最多分为 i+1 份
//			ret += (dp[x][i] * dp[y][i+1]) % mod
//		}
//	}
//	// 精简
//	//ret = dp[x][y+1]*dp[y][y]%mod + dp[x][y]*(dp[y][y-1]+dp[y][y]<<1)%mod // i=y+1 + i=y
//	//for i := (x-1)/limit + 1; i < y; i++ {                                 // (x-1)/limit + 1，最少堆数，确定 i 的起始点
//	//	ret += (dp[x][i] * (dp[y][i-1] + dp[y][i]<<1 + dp[y][i+1])) % mod
//	//}
//	return ret % mod
//}

//func numberOfStableArrays(zero int, one int, limit int) int {
//	// dp：隔板法 & 容斥原理
//	// 模逆元：https://leetcode.cn/problems/find-all-possible-stable-binary-arrays-ii/solutions/2758778/zu-he-shu-xue-rong-chi-yuan-li-yang-hui-swsm9/
//
//	// dp：两种思路
//	// 思路一
//	// 状态定义：f(i,j,k)，i j 为已排列的 0 1 数量，k 标识以 0 或 1 结尾
//	// 状态转移方程：
//	// f(i,j,0) = f(i-1,j,1) + f(i-2,j,1) + ... + f(i-limit,j,1)
//	// f(i,j,1) 同理
//	// 思路二：也是思路一的“前缀和”优化
//	// f(i,j,0) = f(i-1,j,1) + f(i-1,j,0) - f(i-limit-1,j,1)
//	// f(i-1,j,1)：上个位置是 1，直接添加一个 0
//	// f(i-1,j,0)：上个位置是 0，再添加一个 0。如果有连续超过 limit 个 0，则不合法
//	// f(i-limit-1,j,1)：不合法数，已出现连续 limit 个 0，再加 1 个则有 limit+1 个 0
//
//	// 思路二
//	const mod = 1_000_000_007
//	dp := make([][][2]int, zero+1)
//	for i := range dp {
//		dp[i] = make([][2]int, one+1)
//	}
//	for i := 0; i <= min(limit, zero); i++ { // 初始化：0~limit 个 0，以 0 结尾是 1
//		dp[i][0][0] = 1
//	}
//	for j := 0; j <= min(limit, one); j++ {
//		dp[0][j][1] = 1
//	}
//	for i := 1; i <= zero; i++ {
//		for j := 1; j <= one; j++ {
//			if i > limit { // 减去不合法
//				dp[i][j][0] = (dp[i-1][j][0] + dp[i-1][j][1] - dp[i-limit-1][j][1] + mod) % mod
//			} else {
//				dp[i][j][0] = (dp[i-1][j][0] + dp[i-1][j][1]) % mod
//			}
//			if j > limit {
//				dp[i][j][1] = (dp[i][j-1][0] + dp[i][j-1][1] - dp[i][j-limit-1][0] + mod) % mod
//			} else {
//				dp[i][j][1] = (dp[i][j-1][0] + dp[i][j-1][1]) % mod
//			}
//		}
//	}
//	return (dp[zero][one][0] + dp[zero][one][1]) % mod
//
//	// 思路一：Time Limit Exceeded（zero, one, limit = 633, 727, 791）
//	//const mod = 1_000_000_007
//	//dp := make([][][2]int, zero+1)
//	//for i := range dp {
//	//	dp[i] = make([][2]int, one+1)
//	//}
//	//for i := 1; i <= min(limit, zero); i++ {
//	//	dp[i][0] = [2]int{1, 1} // 初始化：0~limit 个 0，以 0 结尾是 1
//	//}
//	//for j := 1; j <= min(limit, one); j++ {
//	//	dp[0][j] = [2]int{1, 1}
//	//}
//	//for i := 1; i <= zero; i++ {
//	//	for j := 1; j <= one; j++ {
//	//		for k := min(limit, i); k > 0; k-- { // 所有合法的结尾，即以 1 结尾
//	//			dp[i][j][0] += dp[i-k][j][1] % mod // 后 mod
//	//		}
//	//		for k := min(limit, j); k > 0; k-- {
//	//			dp[i][j][1] += dp[i][j-k][0] % mod
//	//		}
//	//	}
//	//}
//	//return (dp[zero][one][0] + dp[zero][one][1]) % mod
//
//	// 记忆化搜索：3129 不超时
//	// 递推公式：f(z,o,1) = f(z,o-1,0) + f(z,o-2,0) + ... + f(z,o-limit,0)
//	//const mod = 1_000_000_007
//	//memo := make([][][2]int, zero+1)
//	//for i := range memo {
//	//	memo[i] = make([][2]int, one+1)
//	//}
//	//memo[0][0] = [2]int{1, 1}
//	//var dfs func([2]int, int) int // 参数设置为 [2]int，方便操作 0/1
//	//dfs = func(zo [2]int, idx int) int {
//	//	vs := memo[zo[0]][zo[1]]
//	//	if vs[idx] > 0 {
//	//		return vs[idx]
//	//	}
//	//	if (zo[0]+1)*limit < zo[1] || (zo[1]+1)*limit < zo[0] { // 剪枝
//	//		return 0
//	//	}
//	//	next := zo
//	//	ret, k := 0, min(limit, next[idx]) // 防止 next[idx] < 0
//	//	for i := 0; i < k; i++ {
//	//		next[idx]--
//	//		ret += dfs(next, idx^1) % mod // 以 idx^1 结尾，然后补 1~limit 个 idx
//	//	}
//	//	vs[idx] = ret // 记忆化
//	//	memo[zo[0]][zo[1]] = vs
//	//	return ret
//	//}
//	//return (dfs([2]int{zero, one}, 0) + dfs([2]int{zero, one}, 1)) % mod
//}
