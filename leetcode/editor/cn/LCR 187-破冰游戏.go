//社团共有 num 位成员参与破冰游戏，编号为 0 ~ num-1。成员们按照编号顺序围绕圆桌而坐。社长抽取一个数字 target，从 0 号成员起开始计数，
//排在第 target 位的成员离开圆桌，且成员离开后从下一个成员开始计数。请返回游戏结束时最后一位成员的编号。
//
//
//
// 示例 1：
//
//
//输入：num = 7, target = 4
//输出：1
//
//
// 示例 2：
//
//
//输入：num = 12, target = 5
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= num <= 10^5
// 1 <= target <= 10^6
//
//
//
//
// Related Topics 递归 数学 👍 859 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func iceBreakingGame(num int, target int) int {
	// lc-1823：f(n,k) = (f(n-1,k) + k) % n
	var ret int
	for i := 2; i <= num; i++ {
		ret = (ret + target) % i
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
