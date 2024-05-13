package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2, 3, 1} // 2
	i := reversePairs(nums)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func reversePairs(nums []int) int {
	// 二叉查找树

	// 归并排序
	var ans int
	var mergeSort func(int, int)
	mergeSort = func(l, r int) {
		if l >= r {
			return
		}
		m := (l + r) >> 1
		mergeSort(l, m)
		mergeSort(m+1, r)
		j, k, idx := m+1, 0, m+1
		temp := make([]int, r-l+1)
		for i := l; i <= m; i++ {
			for ; j <= r && nums[j] < nums[i]; j++ {
				temp[k] = nums[j]
				k++
			}
			temp[k] = nums[i]
			k++
			for idx <= r && (nums[i]+1)>>1 > nums[idx] { // 查翻转对
				idx++
			}
			ans += idx - m - 1 // 统计翻转对
		}
		copy(nums[l:], temp[:k])
	}
	mergeSort(0, len(nums)-1)
	return ans

	// 树状数组
	//memo := make(map[int]struct{}, len(nums)<<1)
	//for _, v := range nums { // 统计原数组元素及其二倍值
	//	memo[v] = struct{}{}
	//	memo[v<<1] = struct{}{}
	//}
	//ans, n := 0, len(memo)
	//arr, bit := make([]int, 0, n), make([]int, n+1) // 树状数组从 1 开始索引
	//for v := range memo {
	//	arr = append(arr, v)
	//}
	//sort.Ints(arr) // 排序 arr，树状数组 bit 是 arr 元素出现次数的统计
	//prefixSum := func(i int) int {
	//	var sum int
	//	for ; i > 0; i &= i - 1 {
	//		sum += bit[i]
	//	}
	//	return sum
	//}
	//update := func(i int) {
	//	for ; i <= n; i += i & -i {
	//		bit[i]++
	//	}
	//}
	//for i := len(nums) - 1; i >= 0; i-- { // 倒序遍历：查询 nums[i] 的翻转对 j
	//	j := sort.SearchInts(arr, nums[i])          // arr 中 j 之前的元素都是 i 的翻转对
	//	ans += prefixSum(j)                         // arr[k]==nums[i]，则应该 k-1，但树状数组是从 1 开始
	//	idx := sort.SearchInts(arr, nums[i]<<1) + 1 // 树状数组统计的数是原值的二倍 nums[i]<<1
	//	update(idx)                                 // 树状数组是从 1 开始，所以修正 idx+1
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
