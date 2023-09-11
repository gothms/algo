//给出一个含有不重复整数元素的数组 arr ，每个整数 arr[i] 均大于 1。
//
// 用这些整数来构建二叉树，每个整数可以使用任意次数。其中：每个非叶结点的值应等于它的两个子结点的值的乘积。
//
// 满足条件的二叉树一共有多少个？答案可能很大，返回 对 10⁹ + 7 取余 的结果。
//
//
//
// 示例 1:
//
//
//输入: arr = [2, 4]
//输出: 3
//解释: 可以得到这些二叉树: [2], [4], [4, 2, 2]
//
// 示例 2:
//
//
//输入: arr = [2, 4, 5, 10]
//输出: 7
//解释: 可以得到这些二叉树: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 1000
// 2 <= arr[i] <= 10⁹
// arr 中的所有值 互不相同
//
//
// Related Topics 数组 哈希表 动态规划 排序 👍 100 👎 0

package main

import "sort"

func main() {

}

/*
思路：dp
	1.对 arr 进行排序后，寻找两个相对较小的元素的乘积是否存在
		比如排序后 arr = [2,4,8]，查询 2*4 是否在集合中，查询 2*8 是否在集合中
		在集合中，就可以组合为新的二叉树，左右子树分别为 2、4，根为 8
	2.为了快速查询，我们用map记录 arr 中的元素
		并且 value 值记录为该元素所对应的结果数
		所以这个 map 就是dp的状态定义
	3.状态转移方程是怎样的？
		如果左右子树相等，比如 4*4 = 8，那么 dp[8] = dp[4]*dp[4]
		如果左右子树不相等，比如 2*4 = 8，那么 dp[8] = dp[2]*dp[4] * 2
		简单说来就是：左子树有多少种组合 * 右子树有多少种组合，如果两者不等，则交换位置再计算一次
*/
// leetcode submit region begin(Prohibit modification and deletion)
func numFactoredBinaryTrees(arr []int) int {
	// dp
	sort.Ints(arr)
	cnt, max, memo := 0, arr[len(arr)-1], make(map[int]int, len(arr))
	for _, v := range arr {
		memo[v]++ // 每个元素各为单节点的树
	}
out:
	for i, v := range arr {
		for j := 0; j < i; j++ {
			t := v * arr[j] // 根节点
			if t > max {    // 剪枝
				continue out
			}
			if _, ok := memo[t]; ok { // 满足条件，左右节点不相等
				memo[t] += memo[v] * memo[arr[j]] << 1
			}
		}
		if _, ok := memo[v*v]; ok { // 满足条件，左右节点相等
			memo[v*v] += memo[v] * memo[v]
		}
	}
	for _, c := range memo {
		cnt += c // 统计结果
	}
	return cnt % 1_000_000_007

	// lc
	//	sort.Ints(arr)
	//	n := len(arr)
	//	idx := make(map[int]int, n)
	//	for i, x := range arr {
	//		idx[x] = i
	//	}
	//	f := make([]int, n)
	//	for i, val := range arr {
	//		f[i] = 1
	//		for j, x := range arr[:i] {
	//			if x*x > val {
	//				break
	//			}
	//			if x*x == val {
	//				f[i] += f[j] * f[j]
	//				break
	//			}
	//			if val%x == 0 {
	//				if k, ok := idx[val/x]; ok {
	//					f[i] += f[j] * f[k] * 2
	//				}
	//			}
	//		}
	//		ans += f[i]
	//	}
	//	return ans % 1000000007
}

//leetcode submit region end(Prohibit modification and deletion)
