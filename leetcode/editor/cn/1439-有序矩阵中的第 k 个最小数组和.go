//给你一个 m * n 的矩阵 mat，以及一个整数 k ，矩阵中的每一行都以非递减的顺序排列。
//
// 你可以从每一行中选出 1 个元素形成一个数组。返回所有可能数组中的第 k 个 最小 数组和。
//
//
//
// 示例 1：
//
// 输入：mat = [[1,3,11],[2,4,6]], k = 5
//输出：7
//解释：从每一行中选出一个元素，前 k 个和最小的数组分别是：
//[1,2], [1,4], [3,2], [3,4], [1,6]。其中第 5 个的和是 7 。
//
// 示例 2：
//
// 输入：mat = [[1,3,11],[2,4,6]], k = 9
//输出：17
//
//
// 示例 3：
//
// 输入：mat = [[1,10,10],[1,4,5],[2,3,6]], k = 7
//输出：9
//解释：从每一行中选出一个元素，前 k 个和最小的数组分别是：
//[1,1,2], [1,1,3], [1,4,2], [1,4,3], [1,1,6], [1,5,2], [1,5,3]。其中第 7 个的和是 9 。
//
//
// 示例 4：
//
// 输入：mat = [[1,1,10],[2,2,9]], k = 7
//输出：12
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat.length[i]
// 1 <= m, n <= 40
// 1 <= k <= min(200, n ^ m)
// 1 <= mat[i][j] <= 5000
// mat[i] 是一个非递减数组
//
//
// Related Topics 数组 二分查找 矩阵 堆（优先队列） 👍 114 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	mat := [][]int{{1, 10, 10}, {1, 4, 5}, {2, 3, 6}}
	k := 7
	//k = 27
	//mat = [][]int{{1, 1, 10}, {2, 2, 9}}
	//k = 7
	//mat = [][]int{{1, 3, 11}, {2, 4, 6}}
	//k = 5
	smallest := kthSmallest(mat, k)
	fmt.Println(smallest)
}

/*
推荐题型：此题为 lc-373 加强版
思路：二分+排序
	1.对于 0<=i<m，分别找出 arr 和 mat[i] 两个数组中，k 个 最小 数组和
		arr 数组是 mat[0] 到 mat[i-1] 求得的 k 个 最小 数组和
	2.两次二分查找
		2.1.当数组和是 i 时，找出 arr 和 mat[i] 中，数组和 小于等于 i 的个数
		2.2.找出 arr 和 mat[i] 中，第 k 大的 数组和 的值
			min, kth := arr[0]+mi[0], arr[na-1]+mi[n-1]
			分别对应当前 最小 最大 数组和，且将求得的第 k 大数组和，复制给 kth
	3.统计 k 个数组和
		3.1.修正 k，并跳过2.
		3.2.统计 小于kth 的和
		3.3.补全 等于kth 的和
		3.4.将 k 个最小和，保存到 arr 中，进行下一轮查找
	4.排序 arr：二分查找需要 arr 有序
思路：堆
	1.对于 0<=i<m，分别求出 arr 和 mat[i] 两个数组中，k 个 最小 数组和
		arr 数组是 mat[0] 到 mat[i-1] 求得的 k 个 最小 数组和
	2.创建小顶堆
	3.通过堆，求 k 个最小数组和
		3.1.更新 k 值
		3.2.堆中初始化 n 个值：mat[i][j] + arr[0]
			i：当前行，对应1.中的 i
			j：遍历每一行的每个元素
		3.3.保存 排序 后的 数组和
		3.4.维护小顶堆
			a)没到 arr 尾，继续添加更大的和，并堆化
			b)已到 arr 尾
		3.5.对应3.3.将 k 个最小和，保存到 arr 中，进行下一轮计算
*/
//leetcode submit region begin(Prohibit modification and deletion)
func kthSmallest(mat [][]int, k int) int {
	// 二分
	arr, m, n := mat[0], len(mat), len(mat[0])
	var binarySearch func([]int, int)
	binarySearch = func(mi []int, k int) {
		na := len(arr)
		min, kth := arr[0]+mi[0], arr[na-1]+mi[n-1] // 2
		if v := na * n; v <= k {
			k = v // 3.1
			goto getKArr
		}
		kth = min + sort.Search(kth-min, func(i int) bool {
			i += min
			cnt := 0
			for idx := 0; idx < n; idx++ {
				cnt += sort.Search(na, func(j int) bool { // 2.1
					return arr[j] > i-mi[idx]
				})
			}
			return cnt >= k // 2.2
		})
	getKArr:
		kArr, idx := make([]int, k), 0 // 3
		for i := 0; i < na; i++ {      // 3.2
			for j := 0; j < n; j++ {
				if sum := arr[i] + mi[j]; sum < kth {
					kArr[idx], idx = sum, idx+1
				}
			}
		}
		sort.Ints(kArr[:idx])  // 4
		for ; idx < k; idx++ { // 3.3
			kArr[idx] = kth
		}
		arr = kArr // 3.4
		//sort.Ints(arr) // 4
	}
	for i := 1; i < m; i++ {
		binarySearch(mat[i], k) // 1
	}
	return arr[k-1]

	// 小顶堆
	//arr, m, n := mat[0], len(mat), len(mat[0])
	//if int(math.Pow(float64(m), float64(n))) <= k {
	//	kth := 0
	//	for i := 0; i < m; i++ {
	//		kth += mat[i][n-1]
	//	}
	//	return kth
	//}
	//var pHy func([]int, int)
	//pHy = func(mi []int, k int) {
	//	h, na := pairs{}, len(arr) // 3
	//	if v := na * n; v < k {
	//		k = v // 3.1
	//	}
	//	for i := 0; i < n; i++ { // 3.2
	//		heap.Push(&h, [2]int{0, mi[i] + arr[0]})
	//	}
	//	kArr, idx := make([]int, k), 0
	//	for i := 0; i < k; i++ {
	//		kArr[idx], idx = h[0][1], idx+1 // 3.3
	//		if j := h[0][0]; j+1 < na {     // 3.4.a
	//			h[0][0]++
	//			h[0][1] += arr[j+1] - arr[j]
	//			heap.Fix(&h, 0)
	//		} else { // 3.4.b
	//			heap.Pop(&h)
	//		}
	//	}
	//	arr = kArr // 3.5
	//}
	//for i := 1; i < m; i++ {
	//	pHy(mat[i], k) // 1
	//}
	//return arr[k-1]
}

type pairs [][2]int // 2

func (p *pairs) Push(x interface{}) {
	*p = append(*p, x.([2]int))
}
func (p *pairs) Pop() interface{} {
	v := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return v
}
func (p pairs) Len() int {
	return len(p)
}
func (p pairs) Less(i, j int) bool {
	return p[i][1] < p[j][1]
}
func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//leetcode submit region end(Prohibit modification and deletion)
