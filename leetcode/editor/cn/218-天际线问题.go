//城市的 天际线 是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。给你所有建筑物的位置和高度，请返回 由这些建筑物形成的 天际线 。
//
// 每个建筑物的几何信息由数组 buildings 表示，其中三元组 buildings[i] = [lefti, righti, heighti] 表示：
//
//
//
// lefti 是第 i 座建筑物左边缘的 x 坐标。
// righti 是第 i 座建筑物右边缘的 x 坐标。
// heighti 是第 i 座建筑物的高度。
//
//
// 你可以假设所有的建筑都是完美的长方形，在高度为 0 的绝对平坦的表面上。
//
// 天际线 应该表示为由 “关键点” 组成的列表，格式 [[x1,y1],[x2,y2],...] ，并按 x 坐标 进行 排序 。关键点是水平线段的左端点。
//列表中最后一个点是最右侧建筑物的终点，y 坐标始终为 0 ，仅用于标记天际线的终点。此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。
//
// 注意：输出天际线中不得有连续的相同高度的水平线。例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...] 是不正确的答
//案；三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...]
//
//
//
// 示例 1：
//
//
//输入：buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
//输出：[[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
//解释：
//图 A 显示输入的所有建筑物的位置和高度，
//图 B 显示由这些建筑物形成的天际线。图 B 中的红点表示输出列表中的关键点。
//
// 示例 2：
//
//
//输入：buildings = [[0,2,3],[2,5,3]]
//输出：[[0,3],[5,0]]
//
//
//
//
// 提示：
//
//
// 1 <= buildings.length <= 10⁴
// 0 <= lefti < righti <= 2³¹ - 1
// 1 <= heighti <= 2³¹ - 1
// buildings 按 lefti 非递减排序
//
//
// Related Topics 树状数组 线段树 数组 分治 有序集合 扫描线 堆（优先队列） 👍 830 👎 0

package main

import (
	"fmt"
)

func main() {
	buildings := [][]int{{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8}} // 关键点：[2 3 5 7 9 12 15 19 20 24]
	skyline := getSkyline(buildings)
	fmt.Println(skyline)
}

// leetcode submit region begin(Prohibit modification and deletion)

// 每一座建筑的左边缘信息只被用作加入优先队列时的依据，当其加入优先队列后，只需要用到其高度信息（对最大高度有贡献）以及其右边缘信息（弹出优先队列的依据）
// 因此只需要在优先队列中保存这两个元素即可

func getSkyline(buildings [][]int) [][]int {
}

//leetcode submit region end(Prohibit modification and deletion)

func getSkyline_(buildings [][]int) [][]int {
	//  扫描线 + 堆
	//n := len(buildings)
	//boundaries := make([]int, 0, n*2)
	//for _, building := range buildings {
	//	boundaries = append(boundaries, building[0], building[1])
	//}
	//sort.Ints(boundaries) // 排序：关键点的集合，枚举横坐标
	//
	//ans := make([][]int, 0)
	//idx := 0
	//h := hp{}                             // 大顶堆：建筑高度
	//for _, boundary := range boundaries { // 枚举建筑的每一个边缘作为关键点的横坐标，检查每一座建筑是否「包含该横坐标」，找到最大高度，即为该关键点的纵坐标
	//	for idx < n && buildings[idx][0] <= boundary { // 将「包含该横坐标」的建筑加入到优先队列中
	//		heap.Push(&h, pair{buildings[idx][1], buildings[idx][2]})
	//		idx++
	//	}
	//	// 使用「延迟删除」的技巧：即无需每次横坐标改变就立刻将优先队列中所有不符合条件的元素都删除，而只需要保证优先队列的队首元素「包含该横坐标」即可
	//	for len(h) > 0 && h[0].right <= boundary { // <=：当关键点为某建筑的右边缘时，该建筑的高度对关键点的纵坐标是没有贡献的
	//		heap.Pop(&h)
	//	} // left <= boundary < right：建筑的左边缘小于等于该横坐标，右边缘大于该横坐标
	//
	//	maxn := 0 // 特殊情况：当横坐标为其右边缘时，「包含该横坐标」的建筑并不存在
	//	if len(h) > 0 {
	//		maxn = h[0].height // 优先队列寻找最大高度：纵坐标为「包含该横坐标」的所有建筑的最大高度
	//	}
	//	if len(ans) == 0 || maxn != ans[len(ans)-1][1] { // 无高度差：如果当前关键点的纵坐标大小与前一个关键点的纵坐标大小相同，则说明当前关键点无效，跳过该关键点
	//		ans = append(ans, []int{boundary, maxn})
	//	}
	//}
	//return ans
}

//type pair struct{ right, height int }
//type hp []pair
//
//func (h hp) Len() int            { return len(h) }
//func (h hp) Less(i, j int) bool  { return h[i].height > h[j].height }
//func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
//func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
//func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
