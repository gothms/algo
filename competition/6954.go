package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 7, 5}
	nums = []int{-1, 1, 2, 3, 1}
	target := 2
	//nums = []int{9, -5, -5, 5, -5, -4, -6, 6, -6} // 27
	//target = 3
	//nums = []int{-8, -5, -3, 1, -7} // 27
	//target = -6
	pairs := countPairs(nums, target)
	fmt.Println(pairs)
}

/*
思路：二分
	1.二分思路
		1.1.对于 nums[0]，存在一个索引 idx 使得 nums[0]+nums[idx]>=target
			注意 idx 取值为 [0,n]
		1.2.通过二分查找，找到这个 idx，那么满足 nums[i]+nums[j]<target 的索引
			一定在 [0,idx] 区间内
		1.3.所以我们先对 nums 进行排序
	2.细节
		2.1.这里使用 go API 进行二分，注意：
			当 i++ 时，idx--，因为二分的区间为 [i,i+idx]，对于 i+idx 要保持一致
		2.2.当 idx==0，表示 i 和 idx 重叠，后面的区间不存在 <target 的情况了
			此时 break 掉
*/

func countPairs(nums []int, target int) int {
	cnt, n := 0, len(nums)
	sort.Ints(nums) // 排序
	for i, idx := 0, n; i+idx <= n; i++ {
		idx = sort.Search(idx, func(j int) bool {
			return nums[i]+nums[i+j] >= target // 二分逻辑
		})
		if idx == 0 { // 没有 nums[i]+nums[j]<target 的情况了
			break
		}
		idx-- // 两个作用：1.保持区间上限 i+idx 一致，2.方便计算 cnt
		cnt += idx
	}
	return cnt
}
