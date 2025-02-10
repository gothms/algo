package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findIntersectionValues(nums1 []int, nums2 []int) []int {
	m1, m2 := make(map[int]struct{}), make(map[int]struct{})
	for _, v := range nums1 {
		m1[v] = struct{}{}
	}
	for _, v := range nums2 {
		m2[v] = struct{}{}
	}
	var x, y int
	for _, v := range nums1 {
		if _, ok := m2[v]; ok {
			x++
		}
	}
	for _, v := range nums2 {
		if _, ok := m1[v]; ok {
			y++
		}
	}
	return []int{x, y}
}

//leetcode submit region end(Prohibit modification and deletion)
