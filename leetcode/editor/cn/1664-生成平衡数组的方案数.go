package main

import "fmt"

func main() {
	nums := []int{2, 1, 6, 4}
	nums = []int{1, 1, 1, 1, 1}
	fair := waysToMakeFair(nums)
	fmt.Println(fair)
}

// leetcode submit region begin(Prohibit modification and deletion)
func waysToMakeFair(nums []int) int {
	// lc

	n := len(nums)
	pre := make([][]int, 2)
	pre[0] = make([]int, 1, (n+1)>>1+1)
	pre[1] = make([]int, 1, n>>1+1)
	for i, v := range nums {
		j := i & 1
		pre[j] = append(pre[j], pre[j][i>>1]+v)
	}
	pre[0] = pre[0][1:]
	ans, oe := 0, [2]int{pre[0][len(pre[0])-1], pre[1][len(pre[1])-1]}
	for i, v := range nums {
		j, idx := i&1, i>>1
		//if pre[j][idx+j]-v+oe[j^1]-pre[j^1][idx] == oe[j]-pre[j][idx+j]+pre[j^1][idx] {
		if pre[j][idx+j]<<1-pre[j^1][idx]<<1-v == oe[j]-oe[j^1] {
			ans++
		}
	}
	return ans

	//n := len(nums)
	//odd, even := make([]int, 1, (n+1)>>1+1), make([]int, 1, n>>1+1)
	//for i, v := range nums {
	//	if i&1 == 0 {
	//		odd = append(odd, odd[i>>1]+v)
	//	} else {
	//		even = append(even, even[i>>1]+v)
	//	}
	//}
	//ans, o, e := 0, odd[len(odd)-1], even[len(even)-1]
	//for i, v := range nums {
	//	idx := i >> 1
	//	if i&1 == 0 {
	//		if odd[idx+1]-v+e-even[idx] == o-odd[idx+1]+even[idx] {
	//			ans++
	//		}
	//	} else {
	//		if even[idx+1]-v+o-odd[idx+1] == e-even[idx+1]+odd[idx+1] {
	//			ans++
	//		}
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
