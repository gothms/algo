//给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。求最多有多少个点在同一条直线上。
//
//
//
// 示例 1：
//
//
//输入：points = [[1,1],[2,2],[3,3]]
//输出：3
//
//
// 示例 2：
//
//
//输入：points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
//输出：4
//
//
//
//
// 提示：
//
//
// 1 <= points.length <= 300
// points[i].length == 2
// -10⁴ <= xi, yi <= 10⁴
// points 中的所有点 互不相同
//
//
// Related Topics 几何 数组 哈希表 数学 👍 514 👎 0

package main

import "fmt"

func main() {
	a, b := 2, 2
	for a != 0 {
		a, b = b%a, a
		fmt.Println(a, b)
	}
	fmt.Println(b)
	p := [][]int{{3, 3},
		{1, 4},
		{1, 1},
		{2, 1},
		{2, 2}}
	maxPoints(p)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxPoints(points [][]int) int {
	n := len(points)
	if n < 3 {
		return n
	}
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	absVal := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	gcd := func(a, b int) int { // 求最大公约数
		for a != 0 {
			a, b = b%a, a // 如果 a>b，则第一次循环是交换 a b
		}
		return b
	}
	mp := 0
	for i, p := range points {
		if mp >= n-i || mp > n>>1 { // mp >= n-i：大于 n-i 的肯定已经找过了，mp > n>>1：超过一半的共线
			break
		}
		m := make(map[[2]int]int)
		for _, q := range points[i+1:] {
			x, y := p[0]-q[0], p[1]-q[1]
			switch {
			case x == 0:
				y = 1
			case y == 0:
				x = 1
			default: // 变成同一象限的值
				if y < 0 {
					x, y = -x, -y
				}
				g := gcd(absVal(x), y)
				x /= g
				y /= g
			}
			m[[2]int{x, y}]++
		}
		for _, v := range m { // 检查以 p 为目标的所有斜率
			mp = maxVal(mp, v+1) // 加上 p
		}
	}
	return mp
}

//leetcode submit region end(Prohibit modification and deletion)
