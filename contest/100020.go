package main

import "fmt"

/*
# 最高效：可视为O(1)复杂度
1. 关键点：找出中点的值
2. 找出中点值重复的次数
3. 由于 1 <= nums.length <= 10^5，1 <= nums[i] <= 10^9，所以中点值重复的次数期望值为 1。算上特殊情况下（如所有元素都相等），均摊时间复杂度也是 O(1)
```

	func minLengthAfterRemovals(nums []int) int {
		n, m := len(nums), len(nums)>>1 // 中点 m，中点值 nums[m]
		l, r := m-1, m+n&1              // 奇数：r=m+1，偶数：r=m
		for l >= 0 && nums[l] == nums[m] {
			l--
		}
		for r < n && nums[r] == nums[m] {
			r++
		}
		if cnt := (r-l-1)<<1 - n; cnt > 0 { // r - l - 1：与中点值相等的元素的个数
			return cnt // cnt：余下的个数
		}
		return n & 1 // cnt == 0 && n&1 == 1：特殊情况
	}

```
*/
func main() {
	nums := []int{1, 1, 2, 2, 2, 2, 2, 3, 4}
	nums = []int{1, 3, 5}
	removals := minLengthAfterRemovals(nums)
	fmt.Println(removals)
}
func minLengthAfterRemovals(nums []int) int {
	n, m := len(nums), len(nums)>>1 // 中点 m，中点值 nums[m]
	l, r := m-1, m+n&1              // 奇数：r=m+1，偶数：r=m
	for l >= 0 && nums[l] == nums[m] {
		l--
	}
	for r < n && nums[r] == nums[m] {
		r++
	}
	if cnt := (r-l-1)<<1 - n; cnt > 0 { // r - l - 1：与中点值相等的元素的个数
		return cnt // cnt：余下的个数
	}
	return n & 1 // cnt == 0 && n&1 == 1：特殊情况
}
