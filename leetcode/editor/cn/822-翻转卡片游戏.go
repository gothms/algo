//在桌子上有 N 张卡片，每张卡片的正面和背面都写着一个正数（正面与背面上的数有可能不一样）。
//
// 我们可以先翻转任意张卡片，然后选择其中一张卡片。
//
// 如果选中的那张卡片背面的数字 X 与任意一张卡片的正面的数字都不同，那么这个数字是我们想要的数字。
//
// 哪个数是这些想要的数字中最小的数（找到这些数中的最小值）呢？如果没有一个数字符合要求的，输出 0。
//
// 其中, fronts[i] 和 backs[i] 分别代表第 i 张卡片的正面和背面的数字。
//
// 如果我们通过翻转卡片来交换正面与背面上的数，那么当初在正面的数就变成背面的数，背面的数就变成正面的数。
//
// 示例：
//
//
//输入：fronts = [1,2,4,4,7], backs = [1,3,4,1,3]
//输出：2
//解释：假设我们翻转第二张卡片，那么在正面的数变成了 [1,3,4,4,7] ， 背面的数变成了 [1,2,4,1,3]。
//接着我们选择第二张卡片，因为现在该卡片的背面的数是 2，2 与任意卡片上正面的数都不同，所以 2 就是我们想要的数字。
//
//
//
// 提示：
//
//
// 1 <= fronts.length == backs.length <= 1000
// 1 <= fronts[i] <= 2000
// 1 <= backs[i] <= 2000
//
//
// Related Topics 数组 哈希表 👍 22 👎 0

package main

func main() {

}

/*
思路：hash表
	1.将数分为好数和坏数，坏数就是怎么翻都不可能唯一的数，好数就是除去坏数的数
		由题意分析可知，坏数的必然是出现在 fronts[i] == backs[i]
	2.由上分析，先遍历一遍数组，将 fronts[i] == backs[i] 的坏数存在一个 hash 表
		再遍历一遍数组，从剩余的好数中选出最小的值
*/
// leetcode submit region begin(Prohibit modification and deletion)
func flipgame(fronts []int, backs []int) int {
	// 优化
	minVal := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	minV, n := 2001, len(fronts) // 好数的上限是 2000
	pass := make(map[int]bool)
	for i := 0; i < n; i++ {
		if fronts[i] == backs[i] { // 选出坏数
			pass[fronts[i]] = true
		}
	}
	for i := 0; i < n; i++ { // 从好数中选出最小值
		if !pass[fronts[i]] {
			minV = minVal(minV, fronts[i])
		}
		if !pass[backs[i]] {
			minV = minVal(minV, backs[i])
		}
	}
	return minV % 2001

	// 个人
	//minVal := func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//minV, n := 2003, len(fronts)
	//memo, pass := make(map[int]bool), make(map[int]bool)
	//put := func(v int) {
	//	if pass[v] {
	//		delete(memo, v)
	//	} else {
	//		memo[v] = true
	//	}
	//}
	//for i := 0; i < n; i++ {
	//	if fronts[i] == backs[i] {
	//		pass[fronts[i]] = true
	//		delete(memo, fronts[i])
	//		continue
	//	}
	//	put(fronts[i])
	//	put(backs[i])
	//}
	//for k := range memo {
	//	minV = minVal(minV, k)
	//}
	//if minV == 2003 {
	//	return 0
	//}
	//return minV
}

//leetcode submit region end(Prohibit modification and deletion)
