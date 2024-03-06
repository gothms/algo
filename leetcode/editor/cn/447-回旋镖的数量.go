//给定平面上 n 对 互不相同 的点 points ，其中 points[i] = [xi, yi] 。回旋镖 是由点 (i, j, k) 表示的元组 ，其中
// i 和 j 之间的距离和 i 和 k 之间的欧式距离相等（需要考虑元组的顺序）。
//
// 返回平面上所有回旋镖的数量。
//
// 示例 1：
//
//
//输入：points = [[0,0],[1,0],[2,0]]
//输出：2
//解释：两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]
//
//
// 示例 2：
//
//
//输入：points = [[1,1],[2,2],[3,3]]
//输出：2
//
//
// 示例 3：
//
//
//输入：points = [[1,1]]
//输出：0
//
//
//
//
// 提示：
//
//
// n == points.length
// 1 <= n <= 500
// points[i].length == 2
// -10⁴ <= xi, yi <= 10⁴
// 所有点都 互不相同
//
//
// Related Topics 数组 哈希表 数学 👍 304 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfBoomerangs(points [][]int) int {
	// 枚举
	ret, n := 0, len(points)
	memo := make([]map[int]int, n)
	for i := range memo {
		memo[i] = make(map[int]int)
	}
	for i := range points {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			dis := (x1-points[j][0])*(x1-points[j][0]) + (y1-points[j][1])*(y1-points[j][1])
			memo[i][dis]++
			memo[j][dis]++
		}
	}
	for _, m := range memo {
		for _, v := range m {
			ret += v * (v - 1)
		}
	}
	return ret

	// lc
	//ret := 0
	//for _, p := range points {
	//	cnt := map[int]int{}
	//	for _, q := range points {
	//		dis := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
	//		cnt[dis]++
	//	}
	//	for _, m := range cnt {
	//		ret += m * (m - 1)
	//	}
	//}
	//return ret
}

//leetcode submit region end(Prohibit modification and deletion)
