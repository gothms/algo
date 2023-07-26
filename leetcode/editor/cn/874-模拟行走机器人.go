//机器人在一个无限大小的 XY 网格平面上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令 commands ：
//
//
// -2 ：向左转 90 度
// -1 ：向右转 90 度
// 1 <= x <= 9 ：向前移动 x 个单位长度
//
//
// 在网格上有一些格子被视为障碍物 obstacles 。第 i 个障碍物位于网格点 obstacles[i] = (xi, yi) 。
//
// 机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续尝试进行该路线的其余部分。
//
// 返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。（即，如果距离为 5 ，则返回 25 ）
//
//
//
//
//
//
//
//
//
//
//
//
//
//
// 注意：
//
//
//
// 北表示 +Y 方向。
// 东表示 +X 方向。
// 南表示 -Y 方向。
// 西表示 -X 方向。
//
//
//
//
// 示例 1：
//
//
//输入：commands = [4,-1,3], obstacles = []
//输出：25
//解释：
//机器人开始位于 (0, 0)：
//1. 向北移动 4 个单位，到达 (0, 4)
//2. 右转
//3. 向东移动 3 个单位，到达 (3, 4)
//距离原点最远的是 (3, 4) ，距离为 3² + 4² = 25
//
// 示例 2：
//
//
//输入：commands = [4,-1,4,-2,4], obstacles = [[2,4]]
//输出：65
//解释：机器人开始位于 (0, 0)：
//1. 向北移动 4 个单位，到达 (0, 4)
//2. 右转
//3. 向东移动 1 个单位，然后被位于 (2, 4) 的障碍物阻挡，机器人停在 (1, 4)
//4. 左转
//5. 向北走 4 个单位，到达 (1, 8)
//距离原点最远的是 (1, 8) ，距离为 1² + 8² = 65
//
//
//
// 提示：
//
//
// 1 <= commands.length <= 10⁴
// commands[i] is one of the values in the list [-2,-1,1,2,3,4,5,6,7,8,9].
// 0 <= obstacles.length <= 10⁴
// -3 * 10⁴ <= xi, yi <= 3 * 10⁴
// 答案保证小于 2³¹
//
//
// Related Topics 数组 模拟 👍 172 👎 0

package main

import (
	"fmt"
)

func main() {
	commands := []int{4, -1, 4, -2, 4}
	obstacles := [][]int{{2, 4}}
	commands = []int{-2, -1, 8, 9, 6}
	//[[-1,3],[0,1],[-1,5],[-2,-4],[5,4],[-2,-3],[5,-1],[1,-1],[5,5],[5,2]]
	// 0
	obstacles = [][]int{{-1, 3}, {0, 1}, {-1, 5}, {-2, -4}, {5, 4}, {-2, -3}, {5, -1}, {1, -1}, {5, 5}, {5, 2}}
	sim := robotSim(commands, obstacles)
	fmt.Println(sim)
}

/*
思路：模拟
	1.先考虑移动距离和转向
		移动距离：dir 记录方向，每次移动就是在 dir 的方向移动 commands[i] 个位置
		转向：
		commands[i]=-1：dir = (dir + 1) % 4
		commands[i]=-2：dir = (3 + dir) % 4
	2.障碍点判断：使用哈希表记录障碍点，每次移动时先判断下一个位置是否是障碍点
		是障碍点：停止移动，当前位置就是此次移动的结果位置
		不是障碍点：继续移动，直至移动 commands[i] 个位置
		hash算法：由于 -3 * 10^4 <= xi, yi <= 3 * 10^4，所以我们对点对 {x,y} 的记录为
			(x+30000)<<16 + y + 30000
	3.每轮移动后，判断是否需要更新最大欧式距离
*/
// leetcode submit region begin(Prohibit modification and deletion)
func robotSim(commands []int, obstacles [][]int) int {
	// 模拟
	maxD, x, y, dir := 0, 0, 0, 0
	const A, B = 65536, 30000
	dx, dy := [4]int{0, 1, 0, -1}, [4]int{1, 0, -1, 0}
	cache := make(map[int]bool)
	getPoint := func(x int, y int) int {
		return (x+B)<<16 + y + B // -3 * 10^4 <= xi, yi <= 3 * 10^4
	}
	for _, o := range obstacles {
		cache[getPoint(o[0], o[1])] = true // 障碍点集合
	}
	for _, v := range commands {
		switch v {
		case -1: // 右转
			dir = (dir + 1) % 4
		case -2: // 左转
			dir = (3 + dir) % 4
		default:
			cPoint, isX, i := getPoint(x, y), dir&1 > 0, 0
			for ; i < v; i++ {
				if isX {
					cPoint += A * dx[dir] // x 方向移动
				} else {
					cPoint += dy[dir] // y 方向移动
				}
				if cache[cPoint] {
					break
				}
			}
			x, y = x+dx[dir]*i, y+dy[dir]*i // 更新 x y
			if cd := x*x + y*y; cd > maxD { // 更新最大欧式距离
				maxD = cd
			}
		}
	}
	return maxD

	// 比较繁琐
	//maxD, x, y, dir, n := 0, 0, 0, 0, len(commands)
	//cacheX := make(map[int][]int)
	//cacheY := make(map[int][]int)
	//for _, o := range obstacles {
	//	cacheX[o[0]] = append(cacheX[o[0]], o[1])
	//	cacheY[o[1]] = append(cacheY[o[1]], o[0])
	//}
	//for _, queue := range cacheX {
	//	sort.Ints(queue)
	//}
	//for _, queue := range cacheY {
	//	sort.Ints(queue)
	//}
	//fmt.Println(cacheX)
	//fmt.Println(cacheY)
	//fmt.Println(len(cacheX[4]))
	//for i := 0; i < n; i++ {
	//	switch v := commands[i]; v {
	//	case -1:
	//		dir = (dir + 1) % 4
	//	case -2:
	//		dir = (3 + dir) % 4
	//	default:
	//		nx, ny := x, y
	//		switch dir {
	//		case 0:
	//			ny += v
	//			if m := len(cacheX[x]); m > 0 {
	//				l := sort.Search(m, func(i int) bool {
	//					return cacheX[x][i] > y
	//				})
	//				fmt.Println(l)
	//			}
	//		case 1:
	//			nx += v
	//			if m := len(cacheY[y]); m > 0 {
	//				l := sort.Search(m, func(i int) bool {
	//					return cacheY[y][i] > x
	//				})
	//				fmt.Println(l)
	//			}
	//		case 2:
	//			ny -= v
	//		case 3:
	//			nx -= v
	//		}
	//		x, y = nx, ny
	//		if d := x*x + y*y; d > maxD {
	//			maxD = d
	//		}
	//	}
	//}
	//return maxD
}

//leetcode submit region end(Prohibit modification and deletion)
