//在股票交易中，如果前一天的股价高于后一天的股价，则可以认为存在一个「交易逆序对」。请设计一个程序，输入一段时间内的股票交易记录 record，返回其中存在的
//「交易逆序对」总数。
//
//
//
// 示例 1:
//
//
//输入：record = [9, 7, 5, 4, 6]
//输出：8
//解释：交易中的逆序对为 (9, 7), (9, 5), (9, 4), (9, 6), (7, 5), (7, 4), (7, 6), (5, 4)。
//
//
//
//
// 限制：
//
// 0 <= record.length <= 50000
//
// Related Topics 树状数组 线段树 数组 二分查找 分治 有序集合 归并排序 👍 1067 👎 0

package main

import (
	"fmt"
)

func main() {
	record := []int{9, 7, 5, 4, 6}
	record = []int{}
	i := reversePairs(record)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func reversePairs(record []int) int {
	// 归并
	cnt, n := 0, len(record)
	var mergeSort func(int, int)
	mergeSort = func(l, r int) {
		if l >= r {
			return
		}
		m := (l + r) >> 1
		mergeSort(l, m)
		mergeSort(m+1, r)
		left, right := make([]int, m-l+1), make([]int, r-m)
		copy(left, record[l:]) // 拷贝源数组
		copy(right, record[m+1:])
		j := 0
		for i := 0; i < len(left); i++ {
			for ; j < len(right) && left[i] > right[j]; j++ {
				record[l] = right[j]
				l++
				cnt += len(left) - i // left 的 i 后面的数，都要计算在内
			}
			record[l] = left[i]
			l++
		}
	}
	mergeSort(0, n-1)
	return cnt

	// 离散化树状数组
	// 离散化一个序列的前提是我们只关心这个序列里面元素的相对大小，而不关心绝对大小（即只关心元素在序列中的排名）
	// 离散化的目的是让原来分布零散的值聚集到一起，减少空间浪费
	//n := len(record)
	//tmp := make([]int, n)
	//copy(tmp, record)
	//sort.Ints(tmp) // [4 5 6 7 9]
	//for i := 0; i < n; i++ {
	//	record[i] = sort.SearchInts(tmp, record[i]) + 1
	//} // [5 4 2 1 3]
	//bit := rpBIT{
	//	n:    n,
	//	tree: make([]int, n+1),
	//}
	//ans := 0
	//for i := n - 1; i >= 0; i-- { // 倒序遍历 [5 4 2 1 3]，即遍历 {9, 7, 5, 4, 6}
	//	ans += bit.query(record[i] - 1) // 查询小于“其后面”有多少个数小于它，即查询 6,4,5,7,9
	//	bit.update(record[i])           // 更新树状数组，最后为 [0 1 2 1 4 1]
	//}
	//return ans
}

type rpBIT struct {
	n    int
	tree []int
}

func (rp rpBIT) lowbit(x int) int { return x & (-x) }

func (rp rpBIT) query(x int) int {
	ret := 0
	for x > 0 {
		ret += rp.tree[x]
		x -= rp.lowbit(x)
	}
	return ret
}

func (rp rpBIT) update(x int) {
	for x <= rp.n {
		rp.tree[x]++
		x += rp.lowbit(x)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
