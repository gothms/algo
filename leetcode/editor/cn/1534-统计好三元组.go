package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{7, 3, 7, 3, 12, 1, 12, 2, 3}
	a, b, c := 5, 8, 1 // 12
	arr = []int{22, 31, 26, 16, 1, 12, 0, 30, 20, 34, 17, 36, 16, 5, 3, 36, 31, 21, 12}
	a, b, c = 8, 19, 4 // 49
	triplets := countGoodTriplets(arr, a, b, c)
	fmt.Println(triplets)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countGoodTriplets(arr []int, a int, b int, c int) int {
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}
	// 树状数组
	// 枚举 j 和 k
	ans := 0
	mx := slices.Max(arr)    // 树状数组的长度
	bit := make([]int, mx+2) // 0 <= arr[i] <= 1000
	for j, y := range arr {  // 枚举 j
		for _, z := range arr[j+1:] { // 枚举 k
			if abs(y-z) > b {
				continue
			}
			l := max(y-a, z-c, 0) // 关键逻辑
			r := min(y+a, z+c, mx)
			ans += max(PrefixSum(bit, r+1)-PrefixSum(bit, l), 0) // 有可能 r < l
		}
		Update(bit, y+1, 1) // 从 0 开始存，则 y+1
	}
	return ans

	// 暴力
	//ans, n := 0, len(arr)
	//for i := 0; i < n-2; i++ {
	//	for j := i + 1; j < n-1; j++ {
	//		for k := j + 1; k < n; k++ {
	//			if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
	//				fmt.Println(i, j, k)
	//				ans++
	//			}
	//		}
	//	}
	//}
	//return ans
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
