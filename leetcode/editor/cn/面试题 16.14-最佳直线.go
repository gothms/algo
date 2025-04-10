package main

import (
	"fmt"
)

func main() {
	points := [][]int{{-24272, -29606},
		{-37644, -4251},
		{2691, -22513},
		{-14592, -33765},
		{-21858, 28550},
		{-22264, 41303},
		{-6960, 12785},
		{-39133, -41833},
		{25151, -26643},
		{-19416, 28550},
		{-17420, 22270},
		{-8793, 16457},
		{-4303, -25680},
		{-14405, 26607},
		{-49083, -26336},
		{22629, 20544},
		{-23939, -25038},
		{-40441, -26962},
		{-29484, -30503},
		{-32927, -18287},
		{-13312, -22513},
		{15026, 12965},
		{-16361, -23282},
		{7296, -15750},
		{-11690, -21723},
		{-34850, -25928},
		{-14933, -16169},
		{23459, -9358},
		{-45719, -13202},
		{-26868, 28550},
		{4627, 16457},
		{-7296, -27760},
		{-32230, 8174},
		{-28233, -8627},
		{-26520, 28550},
		{5515, -26001},
		{-16766, 28550},
		{21888, -3740},
		{1251, 28550},
		{15333, -26322},
		{-27677, -19790},
		{20311, 7075},
		{-10751, 16457},
		{-47762, -44638},
		{20991, 24942},
		{-19056, -11105},
		{-26639, 28550},
		{-19862, 16457},
		{-27506, -4251},
		{-20172, -5440},
		{-33757, -24717},
		{-9411, -17379},
		{12493, 29906},
		{0, -21755},
		{-36885, -16192},
		{-38195, -40088},
		{-40079, 7667},
		{-29294, -34032},
		{-55968, 23947},
		{-22724, -22513},
		{20362, -11530},
		{-11817, -23957},
		{-33742, 5259},
		{-10350, -4251},
		{-11690, -22513},
		{-20241, -22513}} // [4,9] {-21858, 28550},{-19416, 28550},
	points = [][]int{{0, 0}, {1, 1}, {1, 0}, {2, 0}, {2, 2}} // [0,1]
	//points = [][]int{{0, 0}, {2, 2}}                         // [0,1]
	line := bestLine(points)
	fmt.Println(line)

	g := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	fmt.Println(g(2, -2))
}

// leetcode submit region begin(Prohibit modification and deletion)
func bestLine(points [][]int) []int {
	// deepseek 翻译 python
	// 关注两点
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	GCD := func(a, b int) int {
		a = abs(a)
		b = abs(b)
		for b != 0 {
			a, b = b, a%b
		}
		if a == 0 {
			return 1
		}
		return a
	}
	computeSlope := func(dx, dy int) string { // 关注一：生成 key
		if dx == 0 {
			return "inf"
		}
		gcd := GCD(dy, dx)
		simplifiedDy := dy / gcd
		simplifiedDx := dx / gcd
		if simplifiedDx < 0 {
			simplifiedDy *= -1
			simplifiedDx *= -1
		}
		return fmt.Sprintf("%d/%d", simplifiedDy, simplifiedDx)
	}
	isSmaller := func(a, b []int) bool {
		for i := 0; i < 2; i++ {
			if i >= len(a) || i >= len(b) {
				return false
			}
			if a[i] < b[i] {
				return true
			} else if a[i] > b[i] {
				return false
			}
		}
		return false
	}

	ansNum := 0
	ans := []int{}
	for i := 0; i < len(points); i++ {
		x, y := points[i][0], points[i][1]
		slopeMap := make(map[string][]int)
		for j := i + 1; j < len(points); j++ {
			tx, ty := points[j][0], points[j][1]
			dx := tx - x
			dy := ty - y
			key := computeSlope(dx, dy)
			slopeMap[key] = append(slopeMap[key], j)
		}
		for _, indices := range slopeMap { // 关注二：每个起始点 i，即可判断最多点的直线
			current := 1 + len(indices)
			candidate := append([]int{i}, indices...)
			if current > ansNum {
				ansNum = current
				ans = candidate
			} else if current == ansNum {
				if len(ans) == 0 || isSmaller(candidate, ans) {
					ans = candidate
				}
			}
		}
	}
	if len(ans) >= 2 {
		return ans[:2]
	}
	return ans

	// lc：有3个点，其中两对点的差值是 x1 y1 & x2 y2，若 x1*y2 == x2*y1，则这3个点同直线

	// 个人：在一条直线上的点，放到一个集合中（这里统计点的对数有多少，点对越多，则点越多）
	//g := func(a, b int) int {
	//	for b != 0 {
	//		a, b = b, a%b
	//	}
	//	return a
	//}
	//
	//n := len(points)
	//memo := make(map[[2]int][3]int) // 斜率作为 key，val 是点对的数量，以及最小的两个索引
	//for i := 0; i < n-1; i++ {
	//	x, y := points[i][0], points[i][1]
	//	for j := i + 1; j < n; j++ {
	//		// 保证 key 合法
	//		dx, dy := points[j][0]-x, points[j][1]-y
	//		if dy == 0 { // 特殊情况：与 x/y 轴平行
	//			dx, dy = math.MaxInt32, y
	//		} else if dx == 0 {
	//			dx, dy = x, math.MaxInt32
	//		} else { // 优先 dy > 0
	//			if dy < 0 {
	//				dx, dy = -dx, -dy
	//			}
	//			mx := g(dx, dy)
	//			dx /= mx
	//			dy /= mx
	//		}
	//
	//		key := [2]int{dx, dy}
	//		if v, ok := memo[key]; ok { // 斜率已存在
	//			v[0]++
	//			memo[key] = v
	//		} else {
	//			memo[key] = [3]int{1, i, j} // 斜率不存在
	//		}
	//	}
	//}
	//var mx [3]int
	//for _, v := range memo {
	//	if v[0] > mx[0] ||
	//		v[0] == mx[0] && (v[1] < mx[1] || v[1] == mx[1] && v[2] < mx[2]) {
	//		mx = v
	//	}
	//}
	//return mx[1:]
}

//leetcode submit region end(Prohibit modification and deletion)
