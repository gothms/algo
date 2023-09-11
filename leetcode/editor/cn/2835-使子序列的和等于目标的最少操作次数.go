//给你一个下标从 0 开始的数组 nums ，它包含 非负 整数，且全部为 2 的幂，同时给你一个整数 target 。
//
// 一次操作中，你必须对数组做以下修改：
//
//
// 选择数组中一个元素 nums[i] ，满足 nums[i] > 1 。
// 将 nums[i] 从数组中删除。
// 在 nums 的 末尾 添加 两个 数，值都为 nums[i] / 2 。
//
//
// 你的目标是让 nums 的一个 子序列 的元素和等于 target ，请你返回达成这一目标的 最少操作次数 。如果无法得到这样的子序列，请你返回 -1 。
//
//
// 数组中一个 子序列 是通过删除原数组中一些元素，并且不改变剩余元素顺序得到的剩余数组。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,8], target = 7
//输出：1
//解释：第一次操作中，我们选择元素 nums[2] 。数组变为 nums = [1,2,4,4] 。
//这时候，nums 包含子序列 [1,2,4] ，和为 7 。
//无法通过更少的操作得到和为 7 的子序列。
//
//
// 示例 2：
//
//
//输入：nums = [1,32,1,2], target = 12
//输出：2
//解释：第一次操作中，我们选择元素 nums[1] 。数组变为 nums = [1,1,2,16,16] 。
//第二次操作中，我们选择元素 nums[3] 。数组变为 nums = [1,1,2,16,8,8] 。
//这时候，nums 包含子序列 [1,1,2,8] ，和为 12 。
//无法通过更少的操作得到和为 12 的子序列。
//
// 示例 3：
//
//
//输入：nums = [1,32,1], target = 35
//输出：-1
//解释：无法得到和为 35 的子序列。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 1000
// 1 <= nums[i] <= 2³⁰
// nums 只包含非负整数，且均为 2 的幂。
// 1 <= target < 2³¹
//
//
// Related Topics 贪心 位运算 数组 👍 33 👎 0

package main

import (
	"fmt"
	"math/bits"
	"strings"
)

func main() {
	bitN := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432, 67108864, 134217728, 268435456, 536870912, 1073741824}
	var sb strings.Builder
	for i, v := range bitN {
		sb.WriteString(fmt.Sprintf("%d : %d, ", v, i))
	}
	fmt.Println(sb.String())

	zeros := bits.TrailingZeros(0)
	fmt.Println(zeros)

	nums := []int{1, 32, 1, 2}
	nums = []int{1, 2, 8}
	target := 7
	operations := minOperations(nums, target)
	fmt.Println(operations)

	fmt.Printf("%b", 39)
}

/*
思路：位运算
	1.核心逻辑：多个细粒度的 2^i 总能组合成 2^j（i<j）
		示例 nums = [1,1,1,1,1,1,1,512], target = 39
		target 的二进制：100111，由于 nums 有7个 1
		所以对于 target 的最低 3 位，nums 的子序和 [1,1,1,1,1,1,1] 都能满足 >= 7（即 111）
	2.理解了核心逻辑，再来看思路
		2.1.bitN := [31]int{}，统计每个 2^i 的数量，以便从低位开始累加子序和，再利用“核心逻辑”
			bitN 的定义来自于题意：1 <= nums[i] <= 230
			示例中的比较：>= 7 如何计算？当前是第 i 位的话，取出 target 的前 i 位：target&(1<<(i+1)-1)
		2.2.“删除原数组中一些元素”逻辑
			如示例中，当判断第 7 位时，子序和=7，而target的前 7 位和为 39，此时需要把 bitN 的更高位拆分
			也就是向高位借位，即向高位找到 bitN[i] != 0 为止，找了几次就是拆分了几次
*/
// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(nums []int, target int) int {
	bitN := [31]int{}
	sum := 0
	for _, v := range nums {
		bitN[bits.TrailingZeros(uint(v))]++ // 统计 2^i
		sum += v
	}
	if sum < target { // 不可能，直接返回 -1
		return -1
	}
	for i, sum, cnt := 0, 0, 0; ; {
		if bitN[i] > 0 {
			sum += 1 << i * bitN[i] // 累加细粒度的 2^i
			if sum >= target {      // 组合成功
				return cnt
			}
		}
		if sum >= target&(1<<(i+1)-1) { // 多个细粒度的 2^i 总能组合成 2^j（i<j）
			i++
			continue
		}
		cnt++ // 向高位的 2^j 借
		for i++; bitN[i] == 0; i++ {
			cnt++
		}
	}

	//bitN := make(map[int]int)
	//cnt, sum := 0, 0
	//for _, v := range nums {
	//	bitN[v]++
	//	sum += v
	//}
	//if sum < target {
	//	return -1
	//}
	//sum = 0
	//for i, b := 0, 0; b <= target; {
	//	b = 1 << i
	//	sum += bitN[b] << i
	//	if sum >= target&(b<<1-1) { // 前 i+1 位比较
	//		i++
	//		continue // >= 说明子序列的和可以满足
	//	}
	//	cnt++
	//	for i, b = i+1, b<<1; bitN[b] == 0; i, b = i+1, b<<1 {
	//		cnt++
	//	}
	//}
	//return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
