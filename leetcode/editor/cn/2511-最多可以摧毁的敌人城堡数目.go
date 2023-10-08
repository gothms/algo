//给你一个长度为 n ，下标从 0 开始的整数数组 forts ，表示一些城堡。forts[i] 可以是 -1 ，0 或者 1 ，其中：
//
//
// -1 表示第 i 个位置 没有 城堡。
// 0 表示第 i 个位置有一个 敌人 的城堡。
// 1 表示第 i 个位置有一个你控制的城堡。
//
//
// 现在，你需要决定，将你的军队从某个你控制的城堡位置 i 移动到一个空的位置 j ，满足：
//
//
// 0 <= i, j <= n - 1
// 军队经过的位置 只有 敌人的城堡。正式的，对于所有 min(i,j) < k < max(i,j) 的 k ，都满足 forts[k] == 0 。
//
//
// 当军队移动时，所有途中经过的敌人城堡都会被 摧毁 。
//
// 请你返回 最多 可以摧毁的敌人城堡数目。如果 无法 移动你的军队，或者没有你控制的城堡，请返回 0 。
//
//
//
// 示例 1：
//
// 输入：forts = [1,0,0,-1,0,0,0,0,1]
//输出：4
//解释：
//- 将军队从位置 0 移动到位置 3 ，摧毁 2 个敌人城堡，位置分别在 1 和 2 。
//- 将军队从位置 8 移动到位置 3 ，摧毁 4 个敌人城堡。
//4 是最多可以摧毁的敌人城堡数目，所以我们返回 4 。
//
//
// 示例 2：
//
// 输入：forts = [0,0,1,-1]
//输出：0
//解释：由于无法摧毁敌人的城堡，所以返回 0 。
//
//
//
//
// 提示：
//
//
// 1 <= forts.length <= 1000
// -1 <= forts[i] <= 1
//
//
// Related Topics 数组 双指针 👍 8 👎 0

package main

import "fmt"

func main() {
	forts := []int{-1, 0, -1, 0, 1, 1, 1, -1, -1, -1}
	i := captureForts(forts)
	fmt.Println(i)
}

/*
思路：双指针
	举例（偷懒）
	示例 1：[1,0,0,-1,0,1]
		-1 出现时，计算摧毁的城堡数量
		第2个 1 出现时，计算摧毁的城堡数量
	示例2：[1,0,0,1,0,-1,0,0,-1]
		前两个 1 出现时，不计算
		-1 出现时，计算
		第2个 -1 出现时，不计算
*/
// leetcode submit region begin(Prohibit modification and deletion)
func captureForts(forts []int) int {
	// 双指针
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	cf, pre := 0, -1
	for i, f := range forts {
		if f != 0 {
			if pre != -1 && f != forts[pre] {
				cf = maxVal(cf, i-pre-1)
			}
			pre = i
		}
	}
	return cf

	// 简化
	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//cnt, pre := 0, -1
	//for i, v := range forts {
	//	if v != 0 {
	//		if pre >= 0 && forts[i] != forts[pre] { // 1和-1 / -1和1
	//			cnt = maxVal(cnt, i-pre-1) // 更新最大数
	//		}
	//		pre = i // 记录位置
	//	}
	//}
	//return cnt

	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//cnt, n := 0, len(forts)
	//for i, j, k := 0, -1, -1; i < n; i++ {
	//	switch forts[i] {
	//	case -1:
	//		if k > j { // 往左移为上一次移动，往右移为 i-k-1
	//			cnt = maxVal(cnt, i-k-1)
	//		} else {
	//			k = -1 // 连续 -1，淘汰上一个 1（可注释，用于移动任意次计算累加和）
	//		}
	//		j = i // 记录 -1 的位置
	//	case 1:
	//		if j > k {
	//			cnt = maxVal(cnt, i-j-1)
	//		} else {
	//			j = -1
	//		}
	//		k = i
	//	}
	//}
	//return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
