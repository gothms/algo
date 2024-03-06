//给你一个用无限二维网格表示的花园，每一个 整数坐标处都有一棵苹果树。整数坐标 (i, j) 处的苹果树有 |i| + |j| 个苹果。
//
// 你将会买下正中心坐标是 (0, 0) 的一块 正方形土地 ，且每条边都与两条坐标轴之一平行。
//
// 给你一个整数 neededApples ，请你返回土地的 最小周长 ，使得 至少 有 neededApples 个苹果在土地 里面或者边缘上。
//
// |x| 的值定义为：
//
//
// 如果 x >= 0 ，那么值为 x
// 如果 x < 0 ，那么值为 -x
//
//
//
//
// 示例 1：
//
//
//输入：neededApples = 1
//输出：8
//解释：边长长度为 1 的正方形不包含任何苹果。
//但是边长为 2 的正方形包含 12 个苹果（如上图所示）。
//周长为 2 * 4 = 8 。
//
//
// 示例 2：
//
//
//输入：neededApples = 13
//输出：16
//
//
// 示例 3：
//
//
//输入：neededApples = 1000000000
//输出：5040
//
//
//
//
// 提示：
//
//
// 1 <= neededApples <= 10¹⁵
//
//
// Related Topics 数学 二分查找 👍 28 👎 0

package main

import (
	"fmt"
)

func main() {
	//fmt.Println(mpApples)
	var need int64
	need = 4620
	removals := minimumPerimeter(need)
	fmt.Println(removals)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumPerimeter(neededApples int64) int64 {
	// 二分
	//return int64(sort.SearchFloat64s(mpApples, float64(neededApples))) << 3

	// lc
	left, right := int64(1), int64(100000)
	ans := int64(0)
	for left <= right {
		mid := (left + right) / 2
		if 2*mid*(mid+1)*(mid*2+1) >= neededApples {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans * 8
}

var mpApples = make([]float64, 1, 62996)

func init() {
	v := float64(0)
	for i := 1; i < 62996; i++ {
		v += float64(i * i * 3)
		mpApples = append(mpApples, v*4)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
