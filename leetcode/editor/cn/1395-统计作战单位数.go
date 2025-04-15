package main

import (
	"fmt"
	"sort"
)

func main() {
	rating := []int{4, 7, 9, 5, 10, 8, 2, 1, 6} // 24
	teams := numTeams(rating)
	fmt.Println(teams)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numTeams(rating []int) int {
	// 枚举 j，并迭代 j 的左边有多少个 更小 & 更大 的数
	// 元素是唯一，所以不用去重

	// lc：离散化树状数组
	n := len(rating)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return rating[idx[i]] < rating[idx[j]]
	})
	rank := make([]int, n)
	for i, v := range idx {
		rank[v] = i
	}
	//fmt.Println(idx) // [7 6 0 3 8 1 5 2 4]
	//fmt.Println(rank) // [2 5 7 3 8 6 1 0 4]

	bit := make([]int, n+1)
	ans := 0
	for i, v := range rank {
		l := PrefixSum(bit, v+1)
		r := i - l
		ans += l*(n-1-v-r) + r*(v-l)
		Update(bit, v+1, 1)
	}
	return ans
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
