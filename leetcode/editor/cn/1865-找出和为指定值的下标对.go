package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type FindSumPairs struct {
	n1, n2 []int
	m      map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	m := make(map[int]int)
	for _, v := range nums2 {
		m[v]++
	}
	return FindSumPairs{nums1, nums2, m}
}

func (this *FindSumPairs) Add(index int, val int) {
	v := this.n2[index]
	this.m[v]--
	this.m[v+val]++
	this.n2[index] = v + val
}

func (this *FindSumPairs) Count(tot int) int {
	ans := 0
	for _, v := range this.n1 {
		ans += this.m[tot-v]
	}
	return ans
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
//leetcode submit region end(Prohibit modification and deletion)

// 个人
//type FindSumPairs struct {
//	n2       []int
//	n1, memo map[int]int
//}
//
//func Constructor(nums1 []int, nums2 []int) FindSumPairs {
//	memo := make(map[int]int)
//	for _, v := range nums2 {
//		memo[v]++
//	}
//	n1 := make(map[int]int)
//	for _, v := range nums1 {
//		n1[v]++
//	}
//	return FindSumPairs{nums2, n1, memo}
//}
//
//func (this *FindSumPairs) Add(index int, val int) {
//	v := this.n2[index]
//	this.memo[v]--
//	this.memo[v+val]++
//	this.n2[index] = v + val
//}
//
//func (this *FindSumPairs) Count(tot int) int {
//	ans := 0
//	for k, v := range this.n1 {
//		ans += v * this.memo[tot-k]
//	}
//	return ans
//}
