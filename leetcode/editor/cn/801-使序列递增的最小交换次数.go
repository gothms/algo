package main

import "fmt"

func main() {
	nums1 := []int{0, 3, 5, 8, 9}
	nums2 := []int{2, 1, 4, 6, 9}
	nums1 = []int{0, 4, 4, 5, 9}
	nums2 = []int{0, 1, 6, 8, 10} // 1
	nums1 = []int{3, 3, 8, 9, 10}
	nums2 = []int{1, 7, 4, 6, 8} // 1
	nums1 = []int{0, 4, 4, 5, 9}
	nums2 = []int{0, 1, 6, 8, 10} // 1
	swap := minSwap(nums1, nums2)
	fmt.Println(swap)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minSwap(nums1 []int, nums2 []int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func minSwap_(nums1 []int, nums2 []int) int {
	n := len(nums1)
	a, b := 0, 0 // 不换 / 换
	if nums1[0] != nums2[0] {
		b = 1
	}
	for i := 1; i < n; i++ {
		if nums1[i] == nums2[i] { // 重新开始
			continue
		}
		if nums1[i] > nums1[i-1] && nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			a, b = min(a, b), min(a, b)+1 // 随便换不换
		} else if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] { // i 换不换得和 i-1 相同
			b = b + 1
		} else {
			a, b = b, a+1 // 必须交换
		}
	}
	return min(a, b)
}
