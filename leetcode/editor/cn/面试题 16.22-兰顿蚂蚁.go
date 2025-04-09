package main

import "fmt"

func main() {
	k := 0 // ["R"]
	//k = 1  // ["X","D"]
	//k = 2  // ["_X","LX"]
	//k = 3  // ["UX","XX"]
	//k = 4  // ["XR","XX"]
	//k = 5  // ["_U","X_","XX"]
	//k = 6  // ["_XR","X__","XX_"]
	moves := printKMoves(k)
	for _, ms := range moves {
		fmt.Println(ms)
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func printKMoves(K int) []string {
	// lc：什么只需要记录黑色的格子就能找到包括所有做过的路径的最小的矩阵？

	dxy := [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} // L U R D
	memo := make(map[[2]int]bool)                      // 走过的格子，false 为白色，true 为黑色
	var xl, xr, yl, yr int                             // 走过区域的最小矩形
	cur, d := [2]int{}, 2                              // 初始坐标 [0,0]，初始方向 R
	for i := 0; i < K; i++ {
		wb := memo[cur]
		// 改变行动方向
		if wb { // 黑
			d = (d + 3) & 3 // % 4
		} else {
			d = (d + 1) & 3
		}
		memo[cur] = !wb     // 改变颜色
		cur[0] += dxy[d][0] // 到下一个坐标
		cur[1] += dxy[d][1]
		xl = min(xl, cur[0]) // 记录走过的矩形区域
		xr = max(xr, cur[0])
		yl = min(yl, cur[1])
		yr = max(yr, cur[1])
	}

	m, n := xr-xl+1, yr-yl+1 // 计算最小矩形

	buf := make([][]byte, m)
	for i := range buf {
		buf[i] = make([]byte, n)
		key := [2]int{i + xl, 0} // 修正 x 坐标
		for j := range buf[i] {
			key[1] = j + yl          // 修正 y 坐标
			if wb := memo[key]; wb { // 黑色
				buf[i][j] = 'X'
			} else {
				buf[i][j] = '_' // 走过的白色 / 未走过的白色
			}
		}
	}
	buf[cur[0]-xl][cur[1]-yl] = [4]byte{'L', 'U', 'R', 'D'}[d] // 蚂蚁的最后位置和方向

	ans := make([]string, m)
	for i, b := range buf {
		ans[i] = string(b)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
