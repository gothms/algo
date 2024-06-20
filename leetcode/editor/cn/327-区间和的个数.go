package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 5, -1}
	lower, upper := -2, 2
	//nums = []int{-2, 0, 1, -1}
	//lower, upper = 1, 4 // 1
	rangeSum := countRangeSum(nums, lower, upper)
	fmt.Println(rangeSum)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countRangeSum(nums []int, lower int, upper int) int {
	// 归并排序
	ans := 0
	return ans

	// 线段树

	// 树状数组
	//n := len(nums)
	//pre, arr := make([]int, n+1), make([]int, 1, n*3+1) // 默认有个和为 0 的值
	//for i, v := range nums {
	//	pre[i+1] = pre[i] + v
	//	arr = append(arr, pre[i+1]-upper, pre[i+1], pre[i+1]-lower)
	//}
	//sort.Ints(arr)
	//k := 1
	//memo := map[int]int{arr[0]: k}
	//for i := 1; i < len(arr); i++ {
	//	if arr[i] != arr[i-1] {
	//		k++
	//		memo[arr[i]] = k
	//	}
	//}
	//ans := 0
	//f := fenwick327{make([]int, k+1)}
	//f.inc(memo[0]) // 和为 0
	//for _, v := range pre[1:] {
	//	f.inc(memo[v])
	//	ans += f.query(memo[v-upper], memo[v-lower])
	//}
	//return ans
}

type fenwick327 struct {
	bit []int
}

func (f fenwick327) inc(i int) {
	for ; i < len(f.bit); i += i & -i {
		f.bit[i]++
	}
}
func (f fenwick327) sum(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f.bit[i]
	}
	return
}
func (f fenwick327) query(l, r int) (res int) {
	return f.sum(r) - f.sum(l-1)
}

//leetcode submit region end(Prohibit modification and deletion)
