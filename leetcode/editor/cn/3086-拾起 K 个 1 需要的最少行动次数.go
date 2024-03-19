//给你一个下标从 0 开始的二进制数组 nums，其长度为 n ；另给你一个 正整数 k 以及一个 非负整数 maxChanges 。
//
// 灵茶山艾府在玩一个游戏，游戏的目标是让灵茶山艾府使用 最少 数量的 行动 次数从 nums 中拾起 k 个 1 。游戏开始时，灵茶山艾府可以选择数组 [0
//, n - 1] 范围内的任何索引index 站立。如果 nums[index] == 1 ，灵茶山艾府就会拾起一个 1 ，并且 nums[index] 变成0
//（这 不算 作一次行动）。之后，灵茶山艾府可以执行 任意数量 的 行动（包括零次），在每次行动中灵茶山艾府必须 恰好 执行以下动作之一：
//
//
// 选择任意一个下标 j != index 且满足 nums[j] == 0 ，然后将 nums[j] 设置为 1 。这个动作最多可以执行
//maxChanges 次。
// 选择任意两个相邻的下标 x 和 y（|x - y| == 1）且满足 nums[x] == 1, nums[y] == 0 ，然后交换它们的值（将
//nums[y] = 1 和 nums[x] = 0）。如果 y == index，在这次行动后灵茶山艾府拾起一个 1 ，并且 nums[y] 变成 0 。
//
//
// 返回灵茶山艾府拾起 恰好 k 个 1 所需的 最少 行动次数。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,1,0,0,0,1,1,0,0,1], k = 3, maxChanges = 1
//
//
// 输出：3
//
// 解释：如果游戏开始时灵茶山艾府在 index == 1 的位置上，按照以下步骤执行每个动作，他可以利用 3 次行动拾取 3 个 1 ：
//
//
// 游戏开始时灵茶山艾府拾取了一个 1 ，nums[1] 变成了 0。此时 nums 变为 [1,0,0,0,0,1,1,0,0,1] 。
// 选择 j == 2 并执行第一种类型的动作。nums 变为 [1,0,1,0,0,1,1,0,0,1]
// 选择 x == 2 和 y == 1 ，并执行第二种类型的动作。nums 变为 [1,1,0,0,0,1,1,0,0,1] 。由于 y == index，
//灵茶山艾府拾取了一个 1 ，nums 变为 [1,0,0,0,0,1,1,0,0,1] 。
// 选择 x == 0 和 y == 1 ，并执行第二种类型的动作。nums 变为 [0,1,0,0,0,1,1,0,0,1] 。由于 y == index，
//灵茶山艾府拾取了一个 1 ，nums 变为 [0,0,0,0,0,1,1,0,0,1] 。
//
//
// 请注意，灵茶山艾府也可能执行其他的 3 次行动序列达成拾取 3 个 1 。
//
// 示例 2：
//
//
//
// 输入：nums = [0,0,0,0], k = 2, maxChanges = 3
//
//
// 输出：4
//
// 解释：如果游戏开始时灵茶山艾府在 index == 0 的位置上，按照以下步骤执行每个动作，他可以利用 4 次行动拾取 2 个 1 ：
//
//
// 选择 j == 1 并执行第一种类型的动作。nums 变为 [0,1,0,0] 。
// 选择 x == 1 和 y == 0 ，并执行第二种类型的动作。nums 变为 [1,0,0,0] 。由于 y == index，灵茶山艾府拾起了一个 1
// ，nums 变为 [0,0,0,0] 。
// 再次选择 j == 1 并执行第一种类型的动作。nums 变为 [0,1,0,0] 。
// 再次选择 x == 1 和 y == 0 ，并执行第二种类型的动作。nums 变为 [1,0,0,0] 。由于y == index，灵茶山艾府拾起了一个
//1 ，nums 变为 [0,0,0,0] 。
//
//
//
//
// 提示：
//
//
// 2 <= n <= 10⁵
// 0 <= nums[i] <= 1
// 1 <= k <= 10⁵
// 0 <= maxChanges <= 10⁵
// maxChanges + sum(nums) >= k
//
//
// Related Topics 贪心 数组 前缀和 滑动窗口 👍 8 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 0, 1, 0, 1}
	k, mc := 3, 0 // 4
	//nums = []int{1, 0, 1}
	//k, mc = 2, 0 // 2
	moves := minimumMoves(nums, k, mc)
	fmt.Println(moves)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumMoves(nums []int, k int, maxChanges int) int64 {
	idx := make([]int, 0) // 哨兵
	cnt := 0
	for i, v := range nums {
		if v == 0 {
			continue
		}
		idx = append(idx, i) // 1
		cnt = max(cnt, 1)    // 连续的 1
		if i > 0 && nums[i-1] == 1 {
			if i > 1 && nums[i-2] == 1 {
				cnt = 3
			} else {
				cnt = max(cnt, 2)
			}
		}
	}
	cnt = min(cnt, k)
	if maxChanges >= k-cnt { // 使用连续的 1 和 maxChanges
		return int64((k-cnt)<<1 + max(cnt-1, 0))
	}
	ret, size, n := int64(math.MaxInt64), k-maxChanges, len(idx)
	pre := make([]int, n+1)
	for i, v := range idx { // 前缀和
		pre[i+1] = pre[i] + v
	}
	// 核心逻辑的计算：如左侧 = idx[m]-idx[l1] + idx[m]-idx[l2] + ...
	l, r, mid := 0, size, size>>1                                   // 滑动窗体 (l,r]，标记 idx 的下标
	for lc, rc := mid, r-mid; r <= n; l, r, mid = l+1, r+1, mid+1 { // 中位数
		left := lc*idx[mid] - pre[mid] + pre[l]  // 左边：lc 个 1
		right := pre[r] - pre[mid] - rc*idx[mid] // 右边：rc 个 1
		ret = min(ret, int64(left+right))
	}
	return int64(maxChanges<<1) + ret
}

//leetcode submit region end(Prohibit modification and deletion)
