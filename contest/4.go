package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 1, 3, 5, 2}
	nums = []int{1, 2, 1, 2}
	l, r := 3, 5
	multisets := countSubMultisets(nums, l, r)
	fmt.Println(multisets)
}
func countSubMultisets(nums []int, l int, r int) int {
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	const mod = 1_000_000_007
	sum := 0
	for _, v := range nums {
		sum += v
	}
	memo := map[int]int{}
	for _, v := range nums {
		memo[v]++
	}
	f := make([]int, sum+1)
	f[0] = memo[0] + 1
	delete(memo, 0)
	//[0 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1]
	//[0 0 1 1 2 1 2 1 2 1 2 1 2 1 2 1 2]
	//[0 0 0 1 1 2 2 3 3 3 3 3 3 3 3 3 3 3]
	//[0 0 0 0 0 1 1 2 2 3 4 3 4 3 4 4 3 4 3 4]

	fmt.Println(memo, sum)
	fmt.Println(f)
	for k, c := range memo {
		ps := make([]int, sum+1+k)
		for i, v := range f {
			ps[i+k] = ps[i] + v // 如 [1,1,2]，ps[4] = ps[2] + f[2] = 1+1:2 和 1+1
		}
		fmt.Println(f)
		fmt.Println(k, c, ps)
		pre := func(x, t int) int {
			if x%k <= t {
				return ps[x/k*k+t] //
			}
			return ps[(x+k-1)/k*k+t] //
		}
		query := func(l, r, t int) int {
			t %= k
			return (pre(r, t) - pre(l, t)) % mod
		}
		for j := sum; j >= 0; j-- {
			f[j] = query(maxVal(j-k*c, 0), j+1, j)
		}
		fmt.Println(f)
	}
	cnt := 0
	for i := l; i <= r; i++ {
		if i > sum {
			break
		}
		cnt = (cnt + f[i]) % mod
	}

	cnt = (cnt%mod + mod) % mod
	return cnt
}
