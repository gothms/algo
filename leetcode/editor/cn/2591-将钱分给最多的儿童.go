//给你一个整数 money ，表示你总共有的钱数（单位为美元）和另一个整数 children ，表示你要将钱分配给多少个儿童。
//
// 你需要按照如下规则分配：
//
//
// 所有的钱都必须被分配。
// 每个儿童至少获得 1 美元。
// 没有人获得 4 美元。
//
//
// 请你按照上述规则分配金钱，并返回 最多 有多少个儿童获得 恰好 8 美元。如果没有任何分配方案，返回 -1 。
//
//
//
// 示例 1：
//
// 输入：money = 20, children = 3
//输出：1
//解释：
//最多获得 8 美元的儿童数为 1 。一种分配方案为：
//- 给第一个儿童分配 8 美元。
//- 给第二个儿童分配 9 美元。
//- 给第三个儿童分配 3 美元。
//没有分配方案能让获得 8 美元的儿童数超过 1 。
//
//
// 示例 2：
//
// 输入：money = 16, children = 2
//输出：2
//解释：每个儿童都可以获得 8 美元。
//
//
//
//
// 提示：
//
//
// 1 <= money <= 200
// 2 <= children <= 30
//
//
// Related Topics 贪心 数学 👍 30 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func distMoney(money int, children int) int {
	n := money - children
	if n < 0 {
		return -1
	}
	//if n < 7 { // 主要排除 n == 3
	//	return 0
	//}
	cnt := n / 7 // 初步计算获得 8 美元的人数
	if mod := n % 7; mod == 3 && cnt == children-1 {
		return cnt - 1 // 最后一个人恰好获得 4 美元
	} else if cnt > children || cnt == children && mod != 0 {
		return children - 1 // 所有人获得 8 美元后，还有结余
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
