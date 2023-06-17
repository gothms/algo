//我们有一个 n 项的集合。给出两个整数数组 values 和 labels ，第 i 个元素的值和标签分别是 values[i] 和 labels[i]。还
//会给出两个整数 numWanted 和 useLimit 。
//
// 从 n 个元素中选择一个子集 s :
//
//
// 子集 s 的大小 小于或等于 numWanted 。
// s 中 最多 有相同标签的 useLimit 项。
//
//
// 一个子集的 分数 是该子集的值之和。
//
// 返回子集 s 的最大 分数 。
//
//
//
// 示例 1：
//
//
//输入：values = [5,4,3,2,1], labels = [1,1,2,2,3], numWanted = 3, useLimit = 1
//输出：9
//解释：选出的子集是第一项，第三项和第五项。
//
//
// 示例 2：
//
//
//输入：values = [5,4,3,2,1], labels = [1,3,3,3,2], numWanted = 3, useLimit = 2
//输出：12
//解释：选出的子集是第一项，第二项和第三项。
//
//
// 示例 3：
//
//
//输入：values = [9,8,8,7,6], labels = [0,0,0,1,1], numWanted = 3, useLimit = 1
//输出：16
//解释：选出的子集是第一项和第四项。
//
//
//
//
// 提示：
//
//
// n == values.length == labels.length
// 1 <= n <= 2 * 10⁴
// 0 <= values[i], labels[i] <= 2 * 10⁴
// 1 <= numWanted, useLimit <= n
//
//
// Related Topics 贪心 数组 哈希表 计数 排序 👍 31 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{5, 4, 3, 2, 6, 9, 4, 2, 3}
	labels := []int{1, 3, 3, 3, 2, 1, 1, 2, 2}
	numWanted := 3
	useLimit := 2
	fromLabels := largestValsFromLabels(values, labels, numWanted, useLimit)
	fmt.Println(fromLabels)

	//fmt.Println(values[:10])
}

/*
思路：排序
	1.这是一个有条件的选择排序后元素的问题
		选择规则：
			按降序选择 numWanted 个 values[i]
			每个 labels[i] 的值对应的子集，最多选择 useLimit 个
		coding 思路：
			按降序排序 values 数组，如何排序呢？
			选择 numWanted 个数，如果当前数的 labels个数==useLimit，则跳过
	2.如何排序
		2.1.将 values[i] 和 labels[i] 的映射关系绑定，然后按 values[i] 的值进行排序
		2.2.新建一个 values 的索引数组，对该数组按 values[i] 的值进行排序
			也是本代码中选择的方案
思路：典型的 TopK
	1.这是一个两次 TopK 的问题
		第一次TopK：按 labels[i] 的值，把 values 分成多个子集，对每个子集选择 TopK 的数
			k=useLimit
		第二次TopK：对第一次TopK后的子集进行合并，再选择 TopK 个数
			k=numWanted
	2.解决方案：
		2.1.两轮快排，分别选出对应的 TopK 的数
		2.2.堆+快排
			a)对每个 labels[i] 的值对应的子集，维护一个 useLimit 大小的 小顶堆
			b)合并每个小顶堆，通过 快排 快速选出 numWanted 个数
				也是本代码中选择的方案
*/
//leetcode submit region begin(Prohibit modification and deletion)
func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	//getSum := func(arr []int, n int) (sum int) {
	//	for i := 0; i < n; i++ {
	//		sum += arr[i]
	//	}
	//	return
	//}
	//memo, n := make(map[int]*myHeap), len(values)
	//for i := 0; i < n; i++ {
	//	curr := memo[labels[i]]
	//	if curr == nil {
	//		h := &myHeap{}
	//		memo[labels[i]], curr = h, h
	//	}
	//	if curr.Len() < useLimit {
	//		heap.Push(curr, values[i])
	//	} else if (*curr)[0] < values[i] {
	//		(*curr)[0] = values[i]
	//		heap.Fix(curr, 0)
	//	}
	//}
	//cache := make([]int, 0)
	//for _, h := range memo {
	//	cache = append(cache, *h...)
	//}
	//if len(cache) <= numWanted {
	//	return getSum(cache, len(cache))
	//}
	//kQuickSort(cache, 0, len(cache)-1, numWanted)
	//return getSum(cache, numWanted)

	max, used, n := 0, make(map[int]int), len(values)
	vIdx := make([]int, n)
	for i := 1; i < n; i++ {
		vIdx[i] = i
	}
	sort.Slice(vIdx, func(i, j int) bool {
		return values[vIdx[i]] > values[vIdx[j]]
	})
	for i, j := 0, 0; i < n; i++ {
		if used[labels[vIdx[i]]] < useLimit {
			max += values[vIdx[i]]
			used[labels[vIdx[i]]]++
			j++
			if j == numWanted {
				break
			}
		}
	}
	//for i, j := 0, 0; i < numWanted; i++ {
	//	for j < n && used[labels[vIdx[j]]] == useLimit {
	//		j++
	//	}
	//	if j == n {
	//		break
	//	}
	//	max += values[vIdx[j]]
	//	used[labels[vIdx[j]]]++
	//	j++
	//}
	return max
}

func kQuickSort(arr []int, l int, r int, k int) {
	if l >= r {
		return
	}
	pivot := partition(arr, l, r)
	switch {
	case pivot > k:
		kQuickSort(arr, l, pivot-1, k)
	case pivot < k:
		kQuickSort(arr, pivot+1, r, k)
	}
}
func partition(arr []int, l int, r int) int {
	pivot, counter := l, l+1
	for i := l + 1; i <= r; i++ {
		if arr[i] >= arr[pivot] {
			arr[i], arr[counter] = arr[counter], arr[i]
			counter++
		}
	}
	counter--
	arr[pivot], arr[counter] = arr[counter], arr[pivot]
	return counter
}

type myHeap []int

func (h myHeap) Len() int {
	return len(h)
}
func (h myHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h myHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *myHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *myHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
