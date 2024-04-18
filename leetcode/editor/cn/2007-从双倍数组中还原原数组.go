//一个整数数组 original 可以转变成一个 双倍 数组 changed ，转变方式为将 original 中每个元素 值乘以 2 加入数组中，然后将所有
//元素 随机打乱 。
//
// 给你一个数组 changed ，如果 change 是 双倍 数组，那么请你返回 original数组，否则请返回空数组。original 的元素可以以
//任意 顺序返回。
//
//
//
// 示例 1：
//
// 输入：changed = [1,3,4,2,6,8]
//输出：[1,3,4]
//解释：一个可能的 original 数组为 [1,3,4] :
//- 将 1 乘以 2 ，得到 1 * 2 = 2 。
//- 将 3 乘以 2 ，得到 3 * 2 = 6 。
//- 将 4 乘以 2 ，得到 4 * 2 = 8 。
//其他可能的原数组方案为 [4,3,1] 或者 [3,1,4] 。
//
//
// 示例 2：
//
// 输入：changed = [6,3,0,1]
//输出：[]
//解释：changed 不是一个双倍数组。
//
//
// 示例 3：
//
// 输入：changed = [1]
//输出：[]
//解释：changed 不是一个双倍数组。
//
//
//
//
// 提示：
//
//
// 1 <= changed.length <= 10⁵
// 0 <= changed[i] <= 10⁵
//
//
// Related Topics 贪心 数组 哈希表 排序 👍 30 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	changed := []int{1, 3, 4, 2, 6, 8}
	changed = []int{0, 3, 2, 4, 6, 0}
	array := findOriginalArray(changed)
	fmt.Println(array)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findOriginalArray(changed []int) []int {
	n := len(changed)
	if n&1 == 1 {
		return nil
	}
	ret, memo := make([]int, 0, n>>1), make(map[int]int, n)
	for _, v := range changed {
		memo[v]++
	}
	sort.Ints(changed) // 排序
	for _, v := range changed {
		if memo[v] > 0 {
			memo[v]--
			memo[v<<1]--
			if memo[v<<1] < 0 { // 失败
				return nil
			}
			ret = append(ret, v)
		}
	}
	return ret

	//n := len(changed)
	//if n&1 == 1 { // fast path
	//	return nil
	//}
	//memo := make(map[int]int, len(changed))
	//for _, v := range changed {
	//	memo[v]++ // 记录
	//}
	//sort.Ints(changed) // 排序 to 贪心
	//ret := make([]int, 0, n>>1)
	//for i := 0; i < n; i++ {
	//	if memo[changed[i]] == 0 {
	//		continue
	//	}
	//	memo[changed[i]<<1]--                                // 2倍数 -1
	//	memo[changed[i]]--                                   // 小数 -1
	//	if memo[changed[i]<<1] < 0 || memo[changed[i]] < 0 { // 贪心失败
	//		return nil
	//	}
	//	ret = append(ret, changed[i])
	//	//if len(ret) == cap(ret) { // 找到结果
	//	//	break
	//	//}
	//}
	//return ret
}

//leetcode submit region end(Prohibit modification and deletion)
