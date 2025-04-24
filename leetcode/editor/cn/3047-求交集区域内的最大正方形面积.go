package main

import (
	"fmt"
	"sort"
)

func main() {
	bl := [][]int{{3, 1}, {2, 2}, {1, 1}}
	tr := [][]int{{6, 5}, {4, 4}, {3, 3}} // 1
	bl = [][]int{{2, 1}, {5, 3}, {4, 2}}
	tr = [][]int{{8, 7}, {10, 5}, {6, 7}} // 4
	area := largestSquareArea(bl, tr)
	fmt.Println(area)
}

// leetcode submit region begin(Prohibit modification and deletion)
func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	// 按矩形的 x2 排序，得到 idx，如示例 1
	// 对于当前矩形 x 为 idx[i]（如示例 1 中的蓝色矩形），它的左右边界 x1 x2 有 x2 > idx[0,i) 的右边界
	// 找到 j 使得 idx[j] 的右边界 > x1，如示例 1 中的红色矩形
	// 枚举 idx[j,i) 与 idx[i] 的交集
	n := len(bottomLeft)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool { // 按 x2 排序
		return topRight[idx[i]][0] < topRight[idx[j]][0]
	})
	ans := 0
	for i, j := range idx { // 枚举 x
		x1, y1, y2 := bottomLeft[j][0], bottomLeft[j][1], topRight[j][1]
		l := sort.Search(i, func(k int) bool { // 找到 idx[j]
			return topRight[idx[k]][0] > x1
		})
		for ; l < i; l++ { // 计算 idx[j,i) 与 x 的交集区域
			index := idx[l]
			x3, x4, y3, y4 := bottomLeft[index][0], topRight[index][0], bottomLeft[index][1], topRight[index][1]
			// x4-max(x1, x3)：x 轴重叠长度
			// max(0, min(y2, y4)-max(y1, y3))：y 轴重叠长度，不能为负，至少为 0
			edge := min(x4-max(x1, x3), max(0, min(y2, y4)-max(y1, y3)))
			ans = max(ans, edge*edge)
		}
	}
	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
