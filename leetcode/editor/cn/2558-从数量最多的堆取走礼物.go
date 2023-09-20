//给你一个整数数组 gifts ，表示各堆礼物的数量。每一秒，你需要执行以下操作：
//
//
// 选择礼物数量最多的那一堆。
// 如果不止一堆都符合礼物数量最多，从中选择任一堆即可。
// 选中的那一堆留下平方根数量的礼物（向下取整），取走其他的礼物。
//
//
// 返回在 k 秒后剩下的礼物数量。
//
//
//
// 示例 1：
//
//
//输入：gifts = [25,64,9,4,100], k = 4
//输出：29
//解释：
//按下述方式取走礼物：
//- 在第一秒，选中最后一堆，剩下 10 个礼物。
//- 接着第二秒选中第二堆礼物，剩下 8 个礼物。
//- 然后选中第一堆礼物，剩下 5 个礼物。
//- 最后，再次选中最后一堆礼物，剩下 3 个礼物。
//最后剩下的礼物数量分别是 [5,8,9,4,3] ，所以，剩下礼物的总数量是 29 。
//
//
// 示例 2：
//
//
//输入：gifts = [1,1,1,1], k = 4
//输出：4
//解释：
//在本例中，不管选中哪一堆礼物，都必须剩下 1 个礼物。
//也就是说，你无法获取任一堆中的礼物。
//所以，剩下礼物的总数量是 4 。
//
//
//
//
// 提示：
//
//
// 1 <= gifts.length <= 10³
// 1 <= gifts[i] <= 10⁹
// 1 <= k <= 10³
//
//
// Related Topics 数组 模拟 堆（优先队列） 👍 16 👎 0

package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func main() {
	gifts := []int{54, 6, 34, 66, 63, 52, 39, 62, 46, 75, 28, 65, 18, 37, 18, 13, 33, 69, 19, 40, 13, 10, 43, 61, 72}
	k := 7
	i := pickGifts(gifts, k)
	fmt.Println(i)
}

/*
思路：堆/topK
	1.每次选出最大堆的礼物，剩下平方根的数量
		此过程可以通过大顶堆/快排实现，选 k 次礼物堆
	2.最后统计剩下的礼物总数
*/
// leetcode submit region begin(Prohibit modification and deletion)
func pickGifts(gifts []int, k int) int64 {
	// heap / topK
	h := pg{gifts}
	heap.Init(&h)
	for i := 0; i < k; i++ { // 选 k 次
		//h.IntSlice[0] = int(math.Sqrt(float64(h.IntSlice[0])))
		gifts[0] = int(math.Sqrt(float64(gifts[0]))) // 剩下平方根的数量
		heap.Fix(&h, 0)
	}
	var left int64
	for _, v := range gifts { // 统计剩下的礼物
		left += int64(v)
	}
	return left
}

type pg struct {
	sort.IntSlice
}

func (p pg) Less(i, j int) bool { return p.IntSlice[i] > p.IntSlice[j] }
func (p pg) Push(x any)         { p.IntSlice = append(p.IntSlice, x.(int)) }
func (p pg) Pop() any {
	v := p.IntSlice[len(p.IntSlice)-1]
	p.IntSlice = p.IntSlice[:len(p.IntSlice)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
