package main

import "fmt"

func main() {
	nums1 := []int{2, 0, 1, 3}
	nums2 := []int{0, 1, 2, 3}
	nums1 = []int{4, 0, 1, 3, 2}
	nums2 = []int{4, 1, 0, 2, 3}
	triplets := goodTriplets(nums1, nums2)
	fmt.Println(triplets)
}

// leetcode submit region begin(Prohibit modification and deletion)
func goodTriplets(nums1 []int, nums2 []int) int64 {
	// lc：等价转换+树状数组
	n := len(nums1)
	//memo := make(map[int]int, n)
	memo := make([]int, n)
	for i, v := range nums1 {
		memo[v] = i
	}
	ans := 0
	bit := make([]int, n+1)
	for i, v := range nums2 {
		y := memo[v]
		x := PrefixSum(bit, y)
		//ans += min(y, x) * min(n-y-1, n-y-1-i+x)
		ans += x * (n - y - 1 - i + x)
		Update(bit, y+1, 1)
	}
	return int64(ans)
}

func Update(b []int, i, delta int) {
	for ; i < len(b); i += i & -i {
		b[i] += delta
	}
}

func PrefixSum(b []int, i int) int {
	sum := 0
	for ; i > 0; i &= i - 1 {
		sum += b[i]
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
