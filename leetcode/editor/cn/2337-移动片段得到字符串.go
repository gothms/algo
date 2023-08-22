//给你两个字符串 start 和 target ，长度均为 n 。每个字符串 仅 由字符 'L'、'R' 和 '_' 组成，其中：
//
//
// 字符 'L' 和 'R' 表示片段，其中片段 'L' 只有在其左侧直接存在一个 空位 时才能向 左 移动，而片段 'R' 只有在其右侧直接存在一个 空位
//时才能向 右 移动。
// 字符 '_' 表示可以被 任意 'L' 或 'R' 片段占据的空位。
//
//
// 如果在移动字符串 start 中的片段任意次之后可以得到字符串 target ，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
// 输入：start = "_L__R__R_", target = "L______RR"
//输出：true
//解释：可以从字符串 start 获得 target ，需要进行下面的移动：
//- 将第一个片段向左移动一步，字符串现在变为 "L___R__R_" 。
//- 将最后一个片段向右移动一步，字符串现在变为 "L___R___R" 。
//- 将第二个片段向右移动散步，字符串现在变为 "L______RR" 。
//可以从字符串 start 得到 target ，所以返回 true 。
//
//
// 示例 2：
//
// 输入：start = "R_L_", target = "__LR"
//输出：false
//解释：字符串 start 中的 'R' 片段可以向右移动一步得到 "_RL_" 。
//但是，在这一步之后，不存在可以移动的片段，所以无法从字符串 start 得到 target 。
//
//
// 示例 3：
//
// 输入：start = "_R", target = "R_"
//输出：false
//解释：字符串 start 中的片段只能向右移动，所以无法从字符串 start 得到 target 。
//
//
//
// 提示：
//
//
// n == start.length == target.length
// 1 <= n <= 10⁵
// start 和 target 由字符 'L'、'R' 和 '_' 组成
//
//
// Related Topics 双指针 字符串 👍 35 👎 0

package main

import (
	"fmt"
)

func main() {
	start := "_L__R__R_"
	target := "L______RR"
	start = "_L__R__RL"
	target = "L_____RLR"
	change := canChange(start, target)
	fmt.Println(change)
}

/*
思路：迭代
	1.示例失败的三种情况
					start		target
		交叉：		"LR"		"RL"
		方向相悖：	"L_"		"_L"
		个数不同：	"LL"		"L_"
	2.假设 lc/rc 分别记录 L/R 的个数，target 计数方式为 +1，start 为 -1
		三种情况的失败条件分析：
		交叉：lc > 0 && rc < 0
		方向相悖：lc < 0 || rc > 0
		个数不同：遍历结束后，lc != 0 && rc != 0
*/
// leetcode submit region begin(Prohibit modification and deletion)
func canChange(start string, target string) bool {
	// 迭代
	lc, rc, n := 0, 0, len(start)
	for i := 0; i < n; i++ {
		switch target[i] {
		case 'L':
			lc++
		case 'R':
			rc++
		}
		if lc > 0 && rc < 0 { // 交叉
			return false
		}
		switch start[i] {
		case 'L':
			lc--
		case 'R':
			rc--
		}
		if lc < 0 || rc > 0 || lc > 0 && rc < 0 { // 方向相悖 or 交叉
			return false
		}
	}
	return lc == 0 && rc == 0 // 个数不同

	// 双指针
	//i, sc, tc, n := 0, 0, 0, len(start)
	//for j := 0; j < n; j++ {
	//	switch target[j] {
	//	case '_':
	//		tc++
	//		continue
	//	case 'L':
	//		for i < n && start[i] == '_' {
	//			sc++
	//			i++
	//		}
	//		if i == n || start[i] == 'R' || i < j {
	//			return false
	//		}
	//		i++
	//	case 'R':
	//		for i < n && start[i] == '_' {
	//			sc++
	//			i++
	//		}
	//		if i == n || start[i] == 'L' || i > j {
	//			return false
	//		}
	//		i++
	//	}
	//}
	//for ; i < n; i++ {
	//	if start[i] != '_' {
	//		return false
	//	}
	//}
	//return true

	// lc
	//if strings.ReplaceAll(start, "_", "") != strings.ReplaceAll(target, "_", "") {
	//	return false
	//}
	//j := 0
	//for i, c := range start {
	//	if c != '_' {
	//		for target[j] == '_' {
	//			j++
	//		}
	//		if i != j && c == 'L' == (i < j) {
	//			return false
	//		}
	//		j++
	//	}
	//}
	//return true
}

//leetcode submit region end(Prohibit modification and deletion)
